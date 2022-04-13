package datastore

import (
	"context"

	"github.com/stackrox/stackrox/central/globaldb"
	"github.com/stackrox/stackrox/central/globalindex"
	"github.com/stackrox/stackrox/central/rbac/k8srolebinding/internal/index"
	"github.com/stackrox/stackrox/central/rbac/k8srolebinding/internal/store"
	"github.com/stackrox/stackrox/central/rbac/k8srolebinding/internal/store/postgres"
	"github.com/stackrox/stackrox/central/rbac/k8srolebinding/internal/store/rocksdb"
	"github.com/stackrox/stackrox/central/rbac/k8srolebinding/search"
	"github.com/stackrox/stackrox/pkg/features"
	"github.com/stackrox/stackrox/pkg/logging"
	"github.com/stackrox/stackrox/pkg/sync"
)

var (
	once sync.Once

	ad DataStore

	log = logging.LoggerForModule()
)

func initialize() {
	var storage store.Store
	var indexer index.Indexer
	if features.PostgresDatastore.Enabled() {
		storage = postgres.New(context.TODO(), globaldb.GetPostgres())
		indexer = postgres.NewIndexer(globaldb.GetPostgres())
	} else {
		storage = rocksdb.New(globaldb.GetRocksDB())
		indexer = index.New(globalindex.GetGlobalTmpIndex())
	}
	var err error
	ad, err = New(storage, indexer, search.New(storage, indexer))
	if err != nil {
		log.Panicf("Failed to initialize secrets datastore: %s", err)
	}
}

// Singleton provides the interface for non-service external interaction.
func Singleton() DataStore {
	once.Do(initialize)
	return ad
}
