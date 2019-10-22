package imagecomponent

import (
	"context"
	"fmt"

	"github.com/stackrox/rox/central/risk/multipliers"
	"github.com/stackrox/rox/central/risk/scorer/vulns"
	"github.com/stackrox/rox/generated/storage"
)

const (
	// ImageComponentVulnerabilitiesHeading is the risk result name for scores calculated by this multiplier.
	ImageComponentVulnerabilitiesHeading = "Image Component Vulnerabilities"

	vulnSaturation = 100
	vulnMaxScore   = 4
)

// vulnerabilitiesMultiplier is a scorer for the vulnerabilities in a image component
type vulnerabilitiesMultiplier struct{}

// NewVulnerabilities provides a multiplier that scores the data based on the CVSS scores and number of CVEs
func NewVulnerabilities() Multiplier {
	return &vulnerabilitiesMultiplier{}
}

// Score takes an image component and evaluates its risk based on vulnerabilties
func (c *vulnerabilitiesMultiplier) Score(ctx context.Context, imageComponent *storage.EmbeddedImageScanComponent) *storage.Risk_Result {
	min, max, sum, numCVEs := vulns.ProcessComponent(imageComponent)
	if numCVEs == 0 {
		return nil
	}

	return &storage.Risk_Result{
		Name: ImageComponentVulnerabilitiesHeading,
		Factors: []*storage.Risk_Result_Factor{
			{
				Message: fmt.Sprintf("Image Component %s version %s contains %d CVEs with CVSS scores ranging between %0.1f and %0.1f",
					imageComponent.GetName(), imageComponent.GetVersion(), numCVEs, min, max),
			},
		},
		Score: multipliers.NormalizeScore(sum, vulnSaturation, vulnMaxScore),
	}
}
