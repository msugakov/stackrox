package reprocessing

import (
	"context"

	"github.com/stackrox/stackrox/central/deployment/datastore"
	"github.com/stackrox/stackrox/central/detection/lifecycle"
	countMetrics "github.com/stackrox/stackrox/central/metrics"
	"github.com/stackrox/stackrox/central/reprocessor"
	riskManager "github.com/stackrox/stackrox/central/risk/manager"
	"github.com/stackrox/stackrox/central/sensor/service/common"
	"github.com/stackrox/stackrox/central/sensor/service/pipeline"
	"github.com/stackrox/stackrox/central/sensor/service/pipeline/reconciliation"
	"github.com/stackrox/stackrox/generated/internalapi/central"
	"github.com/stackrox/stackrox/pkg/logging"
	"github.com/stackrox/stackrox/pkg/metrics"
	"github.com/stackrox/stackrox/pkg/search"
)

var (
	log = logging.LoggerForModule()
)

// GetPipeline returns an instantiation of this particular pipeline
func GetPipeline() pipeline.Fragment {
	return NewPipeline(datastore.Singleton(), lifecycle.SingletonManager(), riskManager.Singleton(), reprocessor.Singleton())
}

// NewPipeline returns a new instance of Pipeline.
func NewPipeline(deployments datastore.DataStore, manager lifecycle.Manager, riskManager riskManager.Manager, riskReprocessor reprocessor.Loop) pipeline.Fragment {
	return &pipelineImpl{
		riskManager:     riskManager,
		riskReprocessor: riskReprocessor,
		manager:         manager,
		deployments:     deployments,
	}
}

type pipelineImpl struct {
	deployments     datastore.DataStore
	riskManager     riskManager.Manager
	riskReprocessor reprocessor.Loop
	manager         lifecycle.Manager
}

func (s *pipelineImpl) Reconcile(ctx context.Context, clusterID string, _ *reconciliation.StoreMap) error {
	// Run reprocessing once sync has completed
	query := search.NewQueryBuilder().AddExactMatches(search.ClusterID, clusterID).ProtoQuery()
	results, err := s.deployments.Search(ctx, query)
	if err != nil {
		return err
	}
	s.riskReprocessor.ReprocessRiskForDeployments(search.ResultsToIDs(results)...)
	return nil
}

func (s *pipelineImpl) Match(msg *central.MsgFromSensor) bool {
	return msg.GetEvent().GetReprocessDeployment() != nil
}

// Run runs the pipeline template on the input and returns the output.
func (s *pipelineImpl) Run(ctx context.Context, clusterID string, msg *central.MsgFromSensor, injector common.MessageInjector) error {
	defer countMetrics.IncrementResourceProcessedCounter(pipeline.ActionToOperation(msg.GetEvent().GetAction()), metrics.DeploymentReprocess)

	reprocessMsg := msg.GetEvent().GetReprocessDeployment()

	deployment, exists, err := s.deployments.GetDeployment(ctx, reprocessMsg.GetDeploymentId())
	if err != nil || !exists {
		return err
	}

	s.riskManager.ReprocessDeploymentRisk(deployment)
	return nil
}

func (s *pipelineImpl) OnFinish(_ string) {}
