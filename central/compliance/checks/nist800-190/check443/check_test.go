package check443

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stackrox/stackrox/central/compliance/framework"
	complianceMocks "github.com/stackrox/stackrox/central/compliance/framework/mocks"
	"github.com/stackrox/stackrox/generated/storage"
	"github.com/stackrox/stackrox/pkg/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var (
	testCluster = &storage.Cluster{
		Id: uuid.NewV4().String(),
	}

	testDeployments = []*storage.Deployment{
		{
			Id: uuid.NewV4().String(),
		},
		{
			Id: uuid.NewV4().String(),
		},
	}

	testNodes = []*storage.Node{
		{
			Id: uuid.NewV4().String(),
		},
		{
			Id: uuid.NewV4().String(),
		},
	}

	domain = framework.NewComplianceDomain(testCluster, testNodes, testDeployments, nil, nil)
)

func TestNIST443_Success(t *testing.T) {
	t.Parallel()

	registry := framework.RegistrySingleton()
	check := registry.Lookup(standardID)
	require.NotNil(t, check)

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	data := complianceMocks.NewMockComplianceDataRepository(mockCtrl)
	data.EXPECT().Cluster().AnyTimes().Return(testCluster)
	data.EXPECT().CISDockerTriggered().AnyTimes().Return(false)
	data.EXPECT().CISKubernetesTriggered().AnyTimes().Return(true)

	run, err := framework.NewComplianceRun(check)
	require.NoError(t, err)
	err = run.Run(context.Background(), "standard", domain, data)
	require.NoError(t, err)

	results := run.GetAllResults()
	checkResults := results[standardID]
	require.NotNil(t, checkResults)

	require.Len(t, checkResults.Evidence(), 1)
	assert.Equal(t, framework.PassStatus, checkResults.Evidence()[0].Status)
}

func TestNIST443_Fail(t *testing.T) {
	t.Parallel()

	registry := framework.RegistrySingleton()
	check := registry.Lookup(standardID)
	require.NotNil(t, check)

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	data := complianceMocks.NewMockComplianceDataRepository(mockCtrl)
	data.EXPECT().Cluster().AnyTimes().Return(testCluster)
	data.EXPECT().CISDockerTriggered().AnyTimes().Return(false)
	data.EXPECT().CISKubernetesTriggered().AnyTimes().Return(false)

	run, err := framework.NewComplianceRun(check)
	require.NoError(t, err)
	err = run.Run(context.Background(), "standard", domain, data)
	require.NoError(t, err)

	results := run.GetAllResults()
	checkResults := results[standardID]
	require.NotNil(t, checkResults)

	require.Len(t, checkResults.Evidence(), 1)
	assert.Equal(t, framework.FailStatus, checkResults.Evidence()[0].Status)
}
