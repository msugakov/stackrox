package securedclusterservices

import (
	"testing"

	helmTest "github.com/stackrox/helmtest/pkg/framework"
	"github.com/stackrox/stackrox/image"
	"github.com/stackrox/stackrox/pkg/buildinfo"
	"github.com/stackrox/stackrox/pkg/features"
	"github.com/stackrox/stackrox/pkg/helm/charts"
	helmChartTestUtils "github.com/stackrox/stackrox/pkg/helm/charts/testutils"
)

func TestWithHelmtest(t *testing.T) {
	additionalTestDirs := []string{}
	if !buildinfo.ReleaseBuild {
		additionalTestDirs = append(additionalTestDirs, "testdata/scanner-slim")
	}

	helmChartTestUtils.RunHelmTestSuite(t, "testdata/helmtest", image.SecuredClusterServicesChartPrefix, helmChartTestUtils.RunHelmTestSuiteOpts{
		HelmTestOpts: []helmTest.LoaderOpt{helmTest.WithAdditionalTestDirs(additionalTestDirs...)},
		MetaValuesOverridesFunc: func(values *charts.MetaValues) {
			// TODO(ROX-8793): The tests will be enabled in a follow-up ticket because the current implementation breaks helm chart rendering.
			if !buildinfo.ReleaseBuild {
				values.FeatureFlags[features.LocalImageScanning.EnvVar()] = true
			}
		}},
	)
}
