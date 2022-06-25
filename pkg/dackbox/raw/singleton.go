package raw

import (
	"log"

	"github.com/blevesearch/bleve"
	"github.com/stackrox/rox/pkg/concurrency"
	"github.com/stackrox/rox/pkg/dackbox"
	"github.com/stackrox/rox/pkg/dackbox/indexer"
	"github.com/stackrox/rox/pkg/dackbox/utils/queue"
	rocksdbInstance "github.com/stackrox/rox/pkg/rocksdb/instance"
	"github.com/stackrox/rox/pkg/sync"
)

var (
	// GraphBucket specifies the prefix for the id map DackBox tracks and stores in the DB.
	GraphBucket = []byte("dackbox_graph")
	// DirtyBucket specifies the prefix for the set of dirty keys (need re-indexing) to add to dackbox.
	DirtyBucket = []byte("dackbox_dirty")
	// ReindexIfMissingBucket is a bucket for all of the child buckets that do not need reindexing.
	ReindexIfMissingBucket = []byte("dackbox_reindex")

	toIndex  queue.WaitableQueue
	registry indexer.WrapperRegistry

	globalKeyLock concurrency.KeyFence

	duckBox *dackbox.DackBox

	once sync.Once
)

// GetGlobalDackBox returns the global dackbox.DackBox instance.
func GetGlobalDackBox() *dackbox.DackBox {
	initialize()
	return duckBox
}

// GetIndexQueue returns the queue of items waiting to be indexed.
func GetIndexQueue() queue.WaitableQueue {
	initialize()
	return toIndex
}

// GetKeyFence returns the global key fence.
func GetKeyFence() concurrency.KeyFence {
	initialize()
	return globalKeyLock
}

func initialize() {
	once.Do(func() {
		rocksdbInstance.RegisterBucket(GraphBucket, "Graph Keys")
		rocksdbInstance.RegisterBucket(DirtyBucket, "Dirty Keys")
		rocksdbInstance.RegisterBucket(ReindexIfMissingBucket, "Bucket for reindexed state")

		toIndex = queue.NewWaitableQueue()
		registry = indexer.NewWrapperRegistry()
		globalKeyLock = concurrency.NewKeyFence()

		var err error
		duckBox, err = dackbox.NewRocksDBDackBox(rocksdbInstance.GetRocksDB(), toIndex, GraphBucket, DirtyBucket, ReindexIfMissingBucket)
		if err != nil {
			log.Panicf("could not load stored indices: %v", err)
		}
	})
}

// StartIndexer starts lazy indexer
func StartIndexer(index bleve.Index) {
	initialize()
	lazy := indexer.NewLazy(toIndex, registry, index, duckBox.AckIndexed)
	lazy.Start()
}

// IndexRegister registers bucket for indexing
func IndexRegister(prefix []byte, wrapper indexer.Wrapper) {
	registry.RegisterWrapper(prefix, wrapper)
}
