package service

import (
	v1 "github.com/stackrox/stackrox/generated/api/v1"
	"github.com/stackrox/stackrox/pkg/grpc"
	"github.com/stackrox/stackrox/pkg/logging"
)

var (
	log = logging.LoggerForModule()
)

// Service provides the service interface.
type Service interface {
	grpc.APIService

	v1.FeatureFlagServiceServer
}

// New returns a new Service instance.
func New() Service {
	return &serviceImpl{}
}
