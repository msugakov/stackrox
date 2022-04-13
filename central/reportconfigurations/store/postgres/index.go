// Code generated by pg-bindings generator. DO NOT EDIT.
package postgres

import (
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
	metrics "github.com/stackrox/stackrox/central/metrics"
	mappings "github.com/stackrox/stackrox/central/reportconfigurations/mappings"
	v1 "github.com/stackrox/stackrox/generated/api/v1"
	storage "github.com/stackrox/stackrox/generated/storage"
	ops "github.com/stackrox/stackrox/pkg/metrics"
	search "github.com/stackrox/stackrox/pkg/search"
	"github.com/stackrox/stackrox/pkg/search/blevesearch"
	"github.com/stackrox/stackrox/pkg/search/postgres"
	"github.com/stackrox/stackrox/pkg/search/postgres/mapping"
)

func init() {
	mapping.RegisterCategoryToTable(v1.SearchCategory_REPORT_CONFIGURATIONS, schema)
}

// NewIndexer returns new indexer for `storage.ReportConfiguration`.
func NewIndexer(db *pgxpool.Pool) *indexerImpl {
	return &indexerImpl{
		db: db,
	}
}

type indexerImpl struct {
	db *pgxpool.Pool
}

func (b *indexerImpl) Count(q *v1.Query, opts ...blevesearch.SearchOption) (int, error) {
	defer metrics.SetIndexOperationDurationTime(time.Now(), ops.Count, "ReportConfiguration")

	return postgres.RunCountRequest(v1.SearchCategory_REPORT_CONFIGURATIONS, q, b.db, mappings.OptionsMap)
}

func (b *indexerImpl) Search(q *v1.Query, opts ...blevesearch.SearchOption) ([]search.Result, error) {
	defer metrics.SetIndexOperationDurationTime(time.Now(), ops.Search, "ReportConfiguration")

	return postgres.RunSearchRequest(v1.SearchCategory_REPORT_CONFIGURATIONS, q, b.db, mappings.OptionsMap)
}

//// Stubs for satisfying interfaces

func (b *indexerImpl) AddReportConfiguration(deployment *storage.ReportConfiguration) error {
	return nil
}

func (b *indexerImpl) AddReportConfigurations(_ []*storage.ReportConfiguration) error {
	return nil
}

func (b *indexerImpl) DeleteReportConfiguration(id string) error {
	return nil
}

func (b *indexerImpl) DeleteReportConfigurations(_ []string) error {
	return nil
}

func (b *indexerImpl) MarkInitialIndexingComplete() error {
	return nil
}

func (b *indexerImpl) NeedsInitialIndexing() (bool, error) {
	return false, nil
}
