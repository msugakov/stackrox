package store

import (
	"github.com/stackrox/stackrox/central/globaldb"
	"github.com/stackrox/stackrox/central/networkpolicies/datastore/internal/store"
	"github.com/stackrox/stackrox/central/networkpolicies/datastore/internal/undodeploymentstore/rocksdb"
	"github.com/stackrox/stackrox/central/networkpolicies/datastore/internal/undostore"
	"github.com/stackrox/stackrox/pkg/sync"
	"github.com/stackrox/stackrox/pkg/utils"
)

var (
	once sync.Once

	as DataStore
)

func initialize() {
	undoDeploymentStorage, err := rocksdb.New(globaldb.GetRocksDB())
	utils.CrashOnError(err)

	as = New(store.Singleton(), undostore.Singleton(), undoDeploymentStorage)
}

// Singleton provides the interface for non-service external interaction.
func Singleton() DataStore {
	once.Do(initialize)
	return as
}
