package sensor

import (
	"context"
	"os"

	"github.com/pkg/errors"
	"github.com/stackrox/stackrox/generated/internalapi/central"
	sensorInternal "github.com/stackrox/stackrox/generated/internalapi/sensor"
	"github.com/stackrox/stackrox/generated/storage"
	"github.com/stackrox/stackrox/pkg/centralsensor"
	"github.com/stackrox/stackrox/pkg/clusterid"
	"github.com/stackrox/stackrox/pkg/env"
	"github.com/stackrox/stackrox/pkg/expiringcache"
	"github.com/stackrox/stackrox/pkg/features"
	"github.com/stackrox/stackrox/pkg/grpc"
	"github.com/stackrox/stackrox/pkg/logging"
	"github.com/stackrox/stackrox/pkg/namespaces"
	"github.com/stackrox/stackrox/pkg/protoutils"
	"github.com/stackrox/stackrox/pkg/satoken"
	"github.com/stackrox/stackrox/sensor/common"
	"github.com/stackrox/stackrox/sensor/common/admissioncontroller"
	"github.com/stackrox/stackrox/sensor/common/centralclient"
	"github.com/stackrox/stackrox/sensor/common/certdistribution"
	"github.com/stackrox/stackrox/sensor/common/clusterentities"
	"github.com/stackrox/stackrox/sensor/common/compliance"
	"github.com/stackrox/stackrox/sensor/common/config"
	"github.com/stackrox/stackrox/sensor/common/deployment"
	"github.com/stackrox/stackrox/sensor/common/detector"
	"github.com/stackrox/stackrox/sensor/common/externalsrcs"
	"github.com/stackrox/stackrox/sensor/common/image"
	"github.com/stackrox/stackrox/sensor/common/networkflow/manager"
	"github.com/stackrox/stackrox/sensor/common/networkflow/service"
	"github.com/stackrox/stackrox/sensor/common/processfilter"
	"github.com/stackrox/stackrox/sensor/common/processsignal"
	"github.com/stackrox/stackrox/sensor/common/reprocessor"
	"github.com/stackrox/stackrox/sensor/common/sensor"
	"github.com/stackrox/stackrox/sensor/common/sensor/helmconfig"
	signalService "github.com/stackrox/stackrox/sensor/common/signal"
	k8sadmctrl "github.com/stackrox/stackrox/sensor/kubernetes/admissioncontroller"
	"github.com/stackrox/stackrox/sensor/kubernetes/client"
	"github.com/stackrox/stackrox/sensor/kubernetes/clusterhealth"
	"github.com/stackrox/stackrox/sensor/kubernetes/clusterstatus"
	"github.com/stackrox/stackrox/sensor/kubernetes/enforcer"
	"github.com/stackrox/stackrox/sensor/kubernetes/fake"
	"github.com/stackrox/stackrox/sensor/kubernetes/listener"
	"github.com/stackrox/stackrox/sensor/kubernetes/listener/resources"
	"github.com/stackrox/stackrox/sensor/kubernetes/localscanner"
	"github.com/stackrox/stackrox/sensor/kubernetes/networkpolicies"
	"github.com/stackrox/stackrox/sensor/kubernetes/orchestrator"
	"github.com/stackrox/stackrox/sensor/kubernetes/telemetry"
	"github.com/stackrox/stackrox/sensor/kubernetes/upgrade"
)

var (
	log = logging.LoggerForModule()
)

// CreateSensor takes in a client interface and returns a sensor instantiation
func CreateSensor(client client.Interface, workloadHandler *fake.WorkloadManager) (*sensor.Sensor, error) {
	admCtrlSettingsMgr := admissioncontroller.NewSettingsManager(resources.DeploymentStoreSingleton(), resources.PodStoreSingleton())

	var helmManagedConfig *central.HelmManagedConfigInit
	if configFP := helmconfig.HelmConfigFingerprint.Setting(); configFP != "" {
		var err error
		helmManagedConfig, err = helmconfig.Load()
		if err != nil {
			return nil, errors.Wrap(err, "loading Helm cluster config")
		}
		if helmManagedConfig.GetClusterConfig().GetConfigFingerprint() != configFP {
			return nil, errors.Errorf("fingerprint %q of loaded config does not match expected fingerprint %q, config changes can only be applied via 'helm upgrade' or a similar chart-based mechanism", helmManagedConfig.GetClusterConfig().GetConfigFingerprint(), configFP)
		}
		log.Infof("Loaded Helm cluster configuration with fingerprint %q", configFP)

		if err := helmconfig.CheckEffectiveClusterName(helmManagedConfig); err != nil {
			return nil, errors.Wrap(err, "validating cluster name")
		}
	}

	if helmManagedConfig.GetClusterName() == "" {
		certClusterID, err := clusterid.ParseClusterIDFromServiceCert(storage.ServiceType_SENSOR_SERVICE)
		if err != nil {
			return nil, errors.Wrap(err, "parsing cluster ID from service certificate")
		}
		if centralsensor.IsInitCertClusterID(certClusterID) {
			return nil, errors.New("a sensor that uses certificates from an init bundle must have a cluster name specified")
		}
	}

	deploymentIdentification := fetchDeploymentIdentification(context.Background(), client.Kubernetes())
	log.Infof("Determined deployment identification: %s", protoutils.NewWrapper(deploymentIdentification))

	auditLogEventsInput := make(chan *sensorInternal.AuditEvents)
	auditLogCollectionManager := compliance.NewAuditLogCollectionManager()

	o := orchestrator.New(client.Kubernetes())
	complianceService := compliance.NewService(o, auditLogEventsInput, auditLogCollectionManager)

	configHandler := config.NewCommandHandler(admCtrlSettingsMgr, deploymentIdentification, helmManagedConfig, auditLogCollectionManager)
	enforcer, err := enforcer.New(client)
	if err != nil {
		return nil, errors.Wrap(err, "creating enforcer")
	}

	imageCache := expiringcache.NewExpiringCache(env.ReprocessInterval.DurationSetting())
	policyDetector := detector.New(enforcer, admCtrlSettingsMgr, resources.DeploymentStoreSingleton(), imageCache, auditLogEventsInput, auditLogCollectionManager, resources.NetworkPolicySingleton())
	admCtrlMsgForwarder := admissioncontroller.NewAdmCtrlMsgForwarder(admCtrlSettingsMgr, listener.New(client, configHandler, policyDetector, k8sNodeName.Setting()))

	upgradeCmdHandler, err := upgrade.NewCommandHandler(configHandler)
	if err != nil {
		return nil, errors.Wrap(err, "creating upgrade command handler")
	}

	imageService := image.NewService(imageCache)
	complianceCommandHandler := compliance.NewCommandHandler(complianceService)

	// Create Process Pipeline
	indicators := make(chan *central.MsgFromSensor)
	processPipeline := processsignal.NewProcessPipeline(indicators, clusterentities.StoreInstance(), processfilter.Singleton(), policyDetector)
	processSignals := signalService.New(processPipeline, indicators)
	networkFlowManager :=
		manager.NewManager(clusterentities.StoreInstance(), externalsrcs.StoreInstance(), policyDetector)
	components := []common.SensorComponent{
		admCtrlMsgForwarder,
		enforcer,
		networkFlowManager,
		networkpolicies.NewCommandHandler(client.Kubernetes()),
		clusterstatus.NewUpdater(client),
		clusterhealth.NewUpdater(client.Kubernetes(), 0),
		complianceCommandHandler,
		processSignals,
		telemetry.NewCommandHandler(client.Kubernetes()),
		upgradeCmdHandler,
		externalsrcs.Singleton(),
		admissioncontroller.AlertHandlerSingleton(),
		auditLogCollectionManager,
		reprocessor.NewHandler(admCtrlSettingsMgr, policyDetector, imageCache),
	}

	sensorNamespace, err := satoken.LoadNamespaceFromFile()
	if err != nil {
		log.Errorf("Failed to determine namespace from service account token file: %s", err)
	}
	if sensorNamespace == "" {
		sensorNamespace = os.Getenv("POD_NAMESPACE")
	}
	if sensorNamespace == "" {
		sensorNamespace = namespaces.StackRox
		log.Warnf("Unable to determine Sensor namespace, defaulting to %s", sensorNamespace)
	}

	if admCtrlSettingsMgr != nil {
		components = append(components, k8sadmctrl.NewConfigMapSettingsPersister(client.Kubernetes(), admCtrlSettingsMgr, sensorNamespace))
	}

	centralClient, err := centralclient.NewClient(env.CentralEndpoint.Setting())
	if err != nil {
		return nil, errors.Wrap(err, "creating central client")
	}

	// Local scanner can be started even if scanner-tls certs are available in the same namespace because
	// it ignores secrets not owned by Sensor.
	if features.LocalImageScanning.Enabled() && securedClusterIsNotManagedManually(helmManagedConfig) && env.LocalImageScanningEnabled.BooleanSetting() {
		podName := os.Getenv("POD_NAME")
		components = append(components,
			localscanner.NewLocalScannerTLSIssuer(client.Kubernetes(), sensorNamespace, podName))
	}

	s := sensor.NewSensor(
		configHandler,
		policyDetector,
		imageService,
		centralClient,
		components...,
	)

	if workloadHandler != nil {
		workloadHandler.SetSignalHandlers(processPipeline, networkFlowManager)
	}

	networkFlowService := service.NewService(networkFlowManager)
	apiServices := []grpc.APIService{
		networkFlowService,
		processSignals,
		complianceService,
		imageService,
		deployment.NewService(resources.DeploymentStoreSingleton(), resources.PodStoreSingleton()),
	}

	if admCtrlSettingsMgr != nil {
		apiServices = append(apiServices, admissioncontroller.NewManagementService(admCtrlSettingsMgr, admissioncontroller.AlertHandlerSingleton()))
	}

	apiServices = append(apiServices, certdistribution.NewService(client.Kubernetes(), sensorNamespace))

	s.AddAPIServices(apiServices...)
	return s, nil
}

func securedClusterIsNotManagedManually(helmManagedConfig *central.HelmManagedConfigInit) bool {
	return helmManagedConfig.GetManagedBy() != storage.ManagerType_MANAGER_TYPE_UNKNOWN &&
		helmManagedConfig.GetManagedBy() != storage.ManagerType_MANAGER_TYPE_MANUAL
}
