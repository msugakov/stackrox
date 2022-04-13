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
	AddComponentCVEEdge(componentcveedge *storage.ComponentCVEEdge) error
	AddComponentCVEEdges(componentcveedges []*storage.ComponentCVEEdge) error
	Count(q *v1.Query, opts ...blevesearch.SearchOption) (int, error)
	DeleteComponentCVEEdge(id string) error
	DeleteComponentCVEEdges(ids []string) error
	MarkInitialIndexingComplete() error
	NeedsInitialIndexing() (bool, error)
	Search(q *v1.Query, opts ...blevesearch.SearchOption) ([]search.Result, error)
}

func New(index bleve.Index) Indexer {
	return &indexerImpl{index: index}
}
