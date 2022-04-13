package sensor

import (
	"github.com/stackrox/stackrox/pkg/concurrency"
	"github.com/stackrox/stackrox/sensor/common"
	"github.com/stackrox/stackrox/sensor/common/config"
	"github.com/stackrox/stackrox/sensor/common/detector"
	"google.golang.org/grpc"
)

// CentralCommunication interface allows you to start and stop the consumption/production loops.
type CentralCommunication interface {
	Start(centralConn grpc.ClientConnInterface, centralReachable *concurrency.Flag, handler config.Handler, detector detector.Detector)
	Stop(error)
	Stopped() concurrency.ReadOnlyErrorSignal
}

// NewCentralCommunication returns a new CentralCommunication.
func NewCentralCommunication(components ...common.SensorComponent) CentralCommunication {
	return &centralCommunicationImpl{
		receiver:   NewCentralReceiver(components...),
		sender:     NewCentralSender(components...),
		components: components,

		stopC:    concurrency.NewErrorSignal(),
		stoppedC: concurrency.NewErrorSignal(),
	}
}
