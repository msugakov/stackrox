package orchestrator

import (
	"context"

	"github.com/stackrox/stackrox/pkg/logging"
	"github.com/stackrox/stackrox/sensor/common/orchestrator"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	coreV1Listers "k8s.io/client-go/listers/core/v1"
)

var (
	log = logging.LoggerForModule()
)

type kubernetesOrchestrator struct {
	client     kubernetes.Interface
	nodeLister coreV1Listers.NodeLister
}

// New returns a new kubernetes orchestrator client.
func New(kubernetes kubernetes.Interface) orchestrator.Orchestrator {
	sif := informers.NewSharedInformerFactory(kubernetes, 0)
	nodeLister := sif.Core().V1().Nodes().Lister()
	sif.Start(context.Background().Done())

	return &kubernetesOrchestrator{
		client:     kubernetes,
		nodeLister: nodeLister,
	}
}

func (k *kubernetesOrchestrator) GetNodeScrapeConfig(nodeName string) (*orchestrator.NodeScrapeConfig, error) {
	node, err := k.nodeLister.Get(nodeName)
	if err != nil {
		return nil, err
	}

	_, isMasterNode := node.GetLabels()["node-role.kubernetes.io/master"]

	return &orchestrator.NodeScrapeConfig{
		ContainerRuntimeVersion: node.Status.NodeInfo.ContainerRuntimeVersion,
		IsMasterNode:            isMasterNode,
	}, nil
}
