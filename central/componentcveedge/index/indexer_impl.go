// Code generated by blevebindings generator. DO NOT EDIT.

package index

import (
	"bytes"
	bleve "github.com/blevesearch/bleve"
	mappings "github.com/stackrox/stackrox/central/componentcveedge/mappings"
	metrics "github.com/stackrox/stackrox/central/metrics"
	v1 "github.com/stackrox/stackrox/generated/api/v1"
	storage "github.com/stackrox/stackrox/generated/storage"
	batcher "github.com/stackrox/stackrox/pkg/batcher"
	ops "github.com/stackrox/stackrox/pkg/metrics"
	search "github.com/stackrox/stackrox/pkg/search"
	blevesearch "github.com/stackrox/stackrox/pkg/search/blevesearch"
	"time"
)

const batchSize = 5000

const resourceName = "ComponentCVEEdge"

type indexerImpl struct {
	index bleve.Index
}

type componentCVEEdgeWrapper struct {
	*storage.ComponentCVEEdge `json:"component_c_v_e_edge"`
	Type                      string `json:"type"`
}

func (b *indexerImpl) AddComponentCVEEdge(componentcveedge *storage.ComponentCVEEdge) error {
	defer metrics.SetIndexOperationDurationTime(time.Now(), ops.Add, "ComponentCVEEdge")
	if err := b.index.Index(componentcveedge.GetId(), &componentCVEEdgeWrapper{
		ComponentCVEEdge: componentcveedge,
		Type:             v1.SearchCategory_COMPONENT_VULN_EDGE.String(),
	}); err != nil {
		return err
	}
	return nil
}

func (b *indexerImpl) AddComponentCVEEdges(componentcveedges []*storage.ComponentCVEEdge) error {
	defer metrics.SetIndexOperationDurationTime(time.Now(), ops.AddMany, "ComponentCVEEdge")
	batchManager := batcher.New(len(componentcveedges), batchSize)
	for {
		start, end, ok := batchManager.Next()
		if !ok {
			break
		}
		if err := b.processBatch(componentcveedges[start:end]); err != nil {
			return err
		}
	}
	return nil
}

func (b *indexerImpl) processBatch(componentcveedges []*storage.ComponentCVEEdge) error {
	batch := b.index.NewBatch()
	for _, componentcveedge := range componentcveedges {
		if err := batch.Index(componentcveedge.GetId(), &componentCVEEdgeWrapper{
			ComponentCVEEdge: componentcveedge,
			Type:             v1.SearchCategory_COMPONENT_VULN_EDGE.String(),
		}); err != nil {
			return err
		}
	}
	return b.index.Batch(batch)
}

func (b *indexerImpl) Count(q *v1.Query, opts ...blevesearch.SearchOption) (int, error) {
	defer metrics.SetIndexOperationDurationTime(time.Now(), ops.Count, "ComponentCVEEdge")
	return blevesearch.RunCountRequest(v1.SearchCategory_COMPONENT_VULN_EDGE, q, b.index, mappings.OptionsMap, opts...)
}

func (b *indexerImpl) DeleteComponentCVEEdge(id string) error {
	defer metrics.SetIndexOperationDurationTime(time.Now(), ops.Remove, "ComponentCVEEdge")
	if err := b.index.Delete(id); err != nil {
		return err
	}
	return nil
}

func (b *indexerImpl) DeleteComponentCVEEdges(ids []string) error {
	defer metrics.SetIndexOperationDurationTime(time.Now(), ops.RemoveMany, "ComponentCVEEdge")
	batch := b.index.NewBatch()
	for _, id := range ids {
		batch.Delete(id)
	}
	if err := b.index.Batch(batch); err != nil {
		return err
	}
	return nil
}

func (b *indexerImpl) MarkInitialIndexingComplete() error {
	return b.index.SetInternal([]byte(resourceName), []byte("old"))
}

func (b *indexerImpl) NeedsInitialIndexing() (bool, error) {
	data, err := b.index.GetInternal([]byte(resourceName))
	if err != nil {
		return false, err
	}
	return !bytes.Equal([]byte("old"), data), nil
}

func (b *indexerImpl) Search(q *v1.Query, opts ...blevesearch.SearchOption) ([]search.Result, error) {
	defer metrics.SetIndexOperationDurationTime(time.Now(), ops.Search, "ComponentCVEEdge")
	return blevesearch.RunSearchRequest(v1.SearchCategory_COMPONENT_VULN_EDGE, q, b.index, mappings.OptionsMap, opts...)
}
