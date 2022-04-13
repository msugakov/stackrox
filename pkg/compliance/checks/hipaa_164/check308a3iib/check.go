package check308a3iib

import (
	"github.com/stackrox/stackrox/pkg/compliance/checks/common"
	"github.com/stackrox/stackrox/pkg/compliance/checks/standards"
)

func init() {
	standards.RegisterChecksForStandard(standards.Hipaa164, map[string]*standards.CheckAndMetadata{
		standards.HIPAA164CheckName("308_a_3_ii_b"): clusterIsCompliant(),
	})
}

func clusterIsCompliant() *standards.CheckAndMetadata {
	// This is a partial check.  The evidence from this check will be folded together with evidence generated in central
	checkAndMetadata := common.MasterAPIServerRBACConfigurationCommandLine()
	checkAndMetadata.Metadata.InterpretationText = interpretationText
	return checkAndMetadata
}
