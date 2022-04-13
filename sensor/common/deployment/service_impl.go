package deployment

import (
	"context"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/pkg/errors"
	"github.com/stackrox/stackrox/generated/internalapi/sensor"
	"github.com/stackrox/stackrox/generated/storage"
	"github.com/stackrox/stackrox/pkg/errorhelpers"
	grpcPkg "github.com/stackrox/stackrox/pkg/grpc"
	"github.com/stackrox/stackrox/pkg/grpc/authz/idcheck"
	"github.com/stackrox/stackrox/sensor/common/store"
	"google.golang.org/grpc"
)

// Service is an interface provides functionality to get deployments from Sensor.
type Service interface {
	grpcPkg.APIService
	sensor.DeploymentServiceServer
	AuthFuncOverride(ctx context.Context, fullMethodName string) (context.Context, error)
}

// NewService returns the DeploymentServiceServer API for Sensor to use.
func NewService(deployments store.DeploymentStore, pods store.PodStore) Service {
	return &serviceImpl{
		deployments: deployments,
		pods:        pods,
	}
}

type serviceImpl struct {
	deployments store.DeploymentStore
	pods        store.PodStore
}

// RegisterServiceServer registers this service with the given gRPC Server.
func (s *serviceImpl) RegisterServiceServer(grpcServer *grpc.Server) {
	sensor.RegisterDeploymentServiceServer(grpcServer, s)
}

// RegisterServiceHandler implements the APIService interface, but the agent does not accept calls over the gRPC gateway
func (s *serviceImpl) RegisterServiceHandler(context.Context, *runtime.ServeMux, *grpc.ClientConn) error {
	return nil
}

// AuthFuncOverride specifies the auth criteria for this API.
func (s *serviceImpl) AuthFuncOverride(ctx context.Context, fullMethodName string) (context.Context, error) {
	return ctx, idcheck.AdmissionControlOnly().Authorized(ctx, fullMethodName)
}

func (s *serviceImpl) GetDeploymentForPod(ctx context.Context, req *sensor.GetDeploymentForPodRequest) (*storage.Deployment, error) {
	if req.GetPodName() == "" || req.GetNamespace() == "" {
		return nil, errors.Wrap(errorhelpers.ErrInvalidArgs, "pod namespace and pod name must be provided")
	}

	pod := s.pods.GetByName(req.GetPodName(), req.GetNamespace())
	if pod == nil {
		return nil, errors.Wrapf(errorhelpers.ErrNotFound,
			"namespace/%s/pods/%s not found",
			req.GetNamespace(), req.GetPodName())
	}

	dep := s.deployments.Get(pod.GetDeploymentId())
	if dep == nil {
		return nil, errors.Wrapf(errorhelpers.ErrNotFound,
			"no containing deployment found for namespace/%s/pods/%s",
			req.GetNamespace(), req.GetPodName())
	}
	return dep, nil
}
