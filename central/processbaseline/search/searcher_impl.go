package search

import (
	"context"

	"github.com/stackrox/stackrox/central/processbaseline/index"
	"github.com/stackrox/stackrox/central/processbaseline/index/mappings"
	"github.com/stackrox/stackrox/central/processbaseline/store"
	"github.com/stackrox/stackrox/central/role/resources"
	v1 "github.com/stackrox/stackrox/generated/api/v1"
	"github.com/stackrox/stackrox/generated/storage"
	"github.com/stackrox/stackrox/pkg/debug"
	"github.com/stackrox/stackrox/pkg/sac"
	"github.com/stackrox/stackrox/pkg/search"
	"github.com/stackrox/stackrox/pkg/search/blevesearch"
	"github.com/stackrox/stackrox/pkg/search/paginated"
)

const (
	baselineBatchLimit = 10000
)

var (
	processBaselineSACSearchHelper = sac.ForResource(resources.ProcessWhitelist).MustCreateSearchHelper(mappings.OptionsMap)
)

type searcherImpl struct {
	storage           store.Store
	indexer           index.Indexer
	formattedSearcher search.Searcher
}

func (s *searcherImpl) buildIndex(ctx context.Context) error {
	defer debug.FreeOSMemory()
	log.Info("[STARTUP] Indexing process baselines")
	baselines := make([]*storage.ProcessBaseline, 0, baselineBatchLimit)
	if err := s.storage.Walk(ctx, func(baseline *storage.ProcessBaseline) error {
		baselines = append(baselines, baseline)
		if len(baselines) == baselineBatchLimit {
			if err := s.indexer.AddProcessBaselines(baselines); err != nil {
				return err
			}

			baselines = baselines[:0]
		}

		return nil
	}); err != nil {
		return err
	}

	if len(baselines) > 0 {
		return s.indexer.AddProcessBaselines(baselines)
	}
	log.Info("[STARTUP] Successfully indexed process baselines")
	return nil
}

func (s *searcherImpl) SearchRawProcessBaselines(ctx context.Context, q *v1.Query) ([]*storage.ProcessBaseline, error) {
	results, err := processBaselineSACSearchHelper.Apply(s.indexer.Search)(ctx, q)
	if err != nil || len(results) == 0 {
		return nil, err
	}
	ids := search.ResultsToIDs(results)
	baselines, _, err := s.storage.GetMany(ctx, ids)
	if err != nil {
		return nil, err
	}
	return baselines, nil
}

func (s *searcherImpl) Search(ctx context.Context, q *v1.Query) ([]search.Result, error) {
	return s.formattedSearcher.Search(ctx, q)
}

// Count returns the number of search results from the query
func (s *searcherImpl) Count(ctx context.Context, q *v1.Query) (int, error) {
	return s.formattedSearcher.Count(ctx, q)
}

// Helper functions which format our searching.
///////////////////////////////////////////////

func formatSearcher(unsafeSearcher blevesearch.UnsafeSearcher) search.Searcher {
	filteredSearcher := processBaselineSACSearchHelper.FilteredSearcher(unsafeSearcher) // Make the UnsafeSearcher safe.
	paginatedSearcher := paginated.Paginated(filteredSearcher)
	return paginatedSearcher
}
