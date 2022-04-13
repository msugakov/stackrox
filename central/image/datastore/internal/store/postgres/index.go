package postgres

import (
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
	mappings "github.com/stackrox/stackrox/central/image/mappings"
	metrics "github.com/stackrox/stackrox/central/metrics"
	v1 "github.com/stackrox/stackrox/generated/api/v1"
	storage "github.com/stackrox/stackrox/generated/storage"
	ops "github.com/stackrox/stackrox/pkg/metrics"
	search "github.com/stackrox/stackrox/pkg/search"
	"github.com/stackrox/stackrox/pkg/search/blevesearch"
	"github.com/stackrox/stackrox/pkg/search/postgres"
	"github.com/stackrox/stackrox/pkg/search/postgres/mapping"
)

func init() {
	mapping.RegisterCategoryToTable(v1.SearchCategory_IMAGES, schema)
}

// NewIndexer returns a new image indexer.
func NewIndexer(db *pgxpool.Pool) *indexerImpl {
	return &indexerImpl{
		db: db,
	}
}

type indexerImpl struct {
	db *pgxpool.Pool
}

func (b *indexerImpl) Count(q *v1.Query, opts ...blevesearch.SearchOption) (int, error) {
	defer metrics.SetIndexOperationDurationTime(time.Now(), ops.Count, "Image")

	return postgres.RunCountRequest(v1.SearchCategory_IMAGES, q, b.db, mappings.OptionsMap)
}

func (b *indexerImpl) Search(q *v1.Query, opts ...blevesearch.SearchOption) ([]search.Result, error) {
	defer metrics.SetIndexOperationDurationTime(time.Now(), ops.Search, "Image")

	return postgres.RunSearchRequest(v1.SearchCategory_IMAGES, q, b.db, mappings.OptionsMap)
}

//// Stubs for satisfying interfaces

func (b *indexerImpl) AddImage(deployment *storage.Image) error {
	return nil
}

func (b *indexerImpl) AddImages(_ []*storage.Image) error {
	return nil
}

func (b *indexerImpl) DeleteImage(id string) error {
	return nil
}

func (b *indexerImpl) DeleteImages(_ []string) error {
	return nil
}

func (b *indexerImpl) MarkInitialIndexingComplete() error {
	return nil
}

func (b *indexerImpl) NeedsInitialIndexing() (bool, error) {
	return false, nil
}
