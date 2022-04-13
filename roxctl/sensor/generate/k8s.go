package generate

import (
	"github.com/spf13/cobra"
	"github.com/stackrox/stackrox/generated/storage"
	clusterValidation "github.com/stackrox/stackrox/pkg/cluster"
	"github.com/stackrox/stackrox/roxctl/common/util"
)

type sensorGenerateK8sCommand struct {
	*sensorGenerateCommand
}

func (s *sensorGenerateK8sCommand) ConstructK8s() {
	s.cluster.Type = storage.ClusterType_KUBERNETES_CLUSTER
	s.cluster.DynamicConfig.DisableAuditLogs = true
}

func k8s(generateCmd *sensorGenerateCommand) *cobra.Command {
	k8sCommand := sensorGenerateK8sCommand{sensorGenerateCommand: generateCmd}
	c := &cobra.Command{
		Use: "k8s",
		RunE: util.RunENoArgs(func(c *cobra.Command) error {
			k8sCommand.ConstructK8s()

			if err := clusterValidation.ValidatePartial(&k8sCommand.cluster); err.ToError() != nil {
				return err.ToError()
			}
			return k8sCommand.fullClusterCreation()
		}),
	}

	c.PersistentFlags().BoolVar(&k8sCommand.cluster.AdmissionControllerEvents, "admission-controller-listen-on-events", true, "enable admission controller webhook to listen on Kubernetes events")
	return c
}
