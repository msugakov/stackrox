package deploytime

import (
	"github.com/stackrox/stackrox/generated/storage"
	"github.com/stackrox/stackrox/pkg/booleanpolicy"
	"github.com/stackrox/stackrox/pkg/detection"
	"github.com/stackrox/stackrox/pkg/detection/deploytime"
)

// Detector provides an interface for getting and managing alerts and enforcements on deployments.
type Detector interface {
	PolicySet() detection.PolicySet

	Detect(ctx deploytime.DetectionContext, enhancedDeployment booleanpolicy.EnhancedDeployment, policyFilters ...detection.FilterOption) ([]*storage.Alert, error)
}

// NewDetector returns a new instance of a Detector.
func NewDetector(policySet detection.PolicySet) Detector {
	return &detectorImpl{
		policySet:      policySet,
		singleDetector: deploytime.NewDetector(policySet),
	}
}
