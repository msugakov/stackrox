package datastore

import (
	"context"

	"github.com/gogo/protobuf/types"
	"github.com/stackrox/stackrox/generated/storage"
	"github.com/stackrox/stackrox/pkg/timestamp"
)

// FlowDataStore stores all of the flows for a single cluster.
//go:generate mockgen-wrapper
type FlowDataStore interface {
	GetAllFlows(ctx context.Context, since *types.Timestamp) ([]*storage.NetworkFlow, types.Timestamp, error)
	GetMatchingFlows(ctx context.Context, pred func(*storage.NetworkFlowProperties) bool, since *types.Timestamp) ([]*storage.NetworkFlow, types.Timestamp, error)
	// UpsertFlows upserts the given flows to the store. The flows slice might be modified by this function, so if you
	// need to use it afterwards, create a copy.
	UpsertFlows(ctx context.Context, flows []*storage.NetworkFlow, lastUpdateTS timestamp.MicroTS) error
	RemoveFlowsForDeployment(ctx context.Context, id string) error
	RemoveMatchingFlows(ctx context.Context, keyMatchFn func(props *storage.NetworkFlowProperties) bool, valueMatchFn func(flow *storage.NetworkFlow) bool) error
}
