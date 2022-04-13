package service

import (
	"context"

	"github.com/stackrox/stackrox/central/cluster/datastore"
	"github.com/stackrox/stackrox/central/probesources"
	"github.com/stackrox/stackrox/central/risk/manager"
	v1 "github.com/stackrox/stackrox/generated/api/v1"
	"github.com/stackrox/stackrox/pkg/grpc"
)

// Service provides the interface to the microservice that serves alert data.
type Service interface {
	grpc.APIService

	AuthFuncOverride(ctx context.Context, fullMethodName string) (context.Context, error)

	v1.ClustersServiceServer
}

// New returns a new Service instance using the given DataStore.
func New(datastore datastore.DataStore, riskManager manager.Manager, probeSources probesources.ProbeSources) Service {
	return &serviceImpl{
		datastore:    datastore,
		riskManager:  riskManager,
		probeSources: probeSources,
	}
}
