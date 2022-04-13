package node

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	nodeMultiplier "github.com/stackrox/stackrox/central/risk/multipliers/node"
	pkgScorer "github.com/stackrox/stackrox/central/risk/scorer"
	"github.com/stackrox/stackrox/generated/storage"
	"github.com/stretchr/testify/assert"
)

func TestScore(t *testing.T) {
	ctx := context.Background()

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	expectedRiskScore := 1.1027501
	expectedRiskResults := []*storage.Risk_Result{
		{
			Name: nodeMultiplier.VulnerabilitiesHeading,
			Factors: []*storage.Risk_Result_Factor{
				{Message: "Node \"node1\" contains 2 CVEs with severities ranging between Low and Moderate"},
			},
			Score: 1.1027501,
		},
	}

	scorer := NewNodeScorer()
	node := pkgScorer.GetMockNode()
	actualRisk := scorer.Score(ctx, node)
	assert.Equal(t, expectedRiskResults, actualRisk.GetResults())
	assert.InDelta(t, expectedRiskScore, actualRisk.GetScore(), 0.0001)
}
