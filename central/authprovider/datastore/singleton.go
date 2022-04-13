package datastore

import (
	"github.com/stackrox/stackrox/central/authprovider/datastore/internal/store"
	"github.com/stackrox/stackrox/central/globaldb"
	"github.com/stackrox/stackrox/pkg/auth/authproviders"
	"github.com/stackrox/stackrox/pkg/sync"
)

var (
	once         sync.Once
	soleInstance authproviders.Store
)

// Singleton returns the sole instance of the DataStore service.
func Singleton() authproviders.Store {
	once.Do(func() {
		soleInstance = New(store.New(globaldb.GetGlobalDB()))
	})
	return soleInstance
}
