// Code generated by blevebindings generator. DO NOT EDIT.

package index

import (
	bleve "github.com/blevesearch/bleve"
	v1 "github.com/stackrox/stackrox/generated/api/v1"
	storage "github.com/stackrox/stackrox/generated/storage"
	search "github.com/stackrox/stackrox/pkg/search"
	blevesearch "github.com/stackrox/stackrox/pkg/search/blevesearch"
)

type Indexer interface {
	AddRisk(risk *storage.Risk) error
	AddRisks(risks []*storage.Risk) error
	Count(q *v1.Query, opts ...blevesearch.SearchOption) (int, error)
	DeleteRisk(id string) error
	DeleteRisks(ids []string) error
	MarkInitialIndexingComplete() error
	NeedsInitialIndexing() (bool, error)
	Search(q *v1.Query, opts ...blevesearch.SearchOption) ([]search.Result, error)
}

func New(index bleve.Index) Indexer {
	return &indexerImpl{index: index}
}
