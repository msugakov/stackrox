package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
	"time"

	alertDatastore "github.com/stackrox/stackrox/central/alert/datastore"
	alertService "github.com/stackrox/stackrox/central/alert/service"
	apiTokenService "github.com/stackrox/stackrox/central/apitoken/service"
	"github.com/stackrox/stackrox/central/audit"
	authService "github.com/stackrox/stackrox/central/auth/service"
	"github.com/stackrox/stackrox/central/auth/userpass"
	authProviderDS "github.com/stackrox/stackrox/central/authprovider/datastore"
	authProviderService "github.com/stackrox/stackrox/central/authprovider/service"
	centralHealthService "github.com/stackrox/stackrox/central/centralhealth/service"
	"github.com/stackrox/stackrox/central/certgen"
	"github.com/stackrox/stackrox/central/cli"
	clusterDataStore "github.com/stackrox/stackrox/central/cluster/datastore"
	clusterService "github.com/stackrox/stackrox/central/cluster/service"
	"github.com/stackrox/stackrox/central/clusterinit/backend"
	clusterInitService "github.com/stackrox/stackrox/central/clusterinit/service"
	clustersHelmConfig "github.com/stackrox/stackrox/central/clusters/helmconfig"
	clustersZip "github.com/stackrox/stackrox/central/clusters/zip"
	complianceDatastore "github.com/stackrox/stackrox/central/compliance/datastore"
	complianceHandlers "github.com/stackrox/stackrox/central/compliance/handlers"
	complianceManager "github.com/stackrox/stackrox/central/compliance/manager"
	complianceManagerService "github.com/stackrox/stackrox/central/compliance/manager/service"
	complianceService "github.com/stackrox/stackrox/central/compliance/service"
	configService "github.com/stackrox/stackrox/central/config/service"
	credentialExpiryService "github.com/stackrox/stackrox/central/credentialexpiry/service"
	"github.com/stackrox/stackrox/central/cve/csv"
	"github.com/stackrox/stackrox/central/cve/fetcher"
	cveService "github.com/stackrox/stackrox/central/cve/service"
	"github.com/stackrox/stackrox/central/cve/suppress"
	debugService "github.com/stackrox/stackrox/central/debug/service"
	deploymentDatastore "github.com/stackrox/stackrox/central/deployment/datastore"
	deploymentService "github.com/stackrox/stackrox/central/deployment/service"
	detectionService "github.com/stackrox/stackrox/central/detection/service"
	developmentService "github.com/stackrox/stackrox/central/development/service"
	"github.com/stackrox/stackrox/central/docs"
	"github.com/stackrox/stackrox/central/endpoints"
	"github.com/stackrox/stackrox/central/enrichment"
	_ "github.com/stackrox/stackrox/central/externalbackups/plugins/all" // Import all of the external backup plugins
	backupService "github.com/stackrox/stackrox/central/externalbackups/service"
	featureFlagService "github.com/stackrox/stackrox/central/featureflags/service"
	"github.com/stackrox/stackrox/central/globaldb"
	dbAuthz "github.com/stackrox/stackrox/central/globaldb/authz"
	globaldbHandlers "github.com/stackrox/stackrox/central/globaldb/handlers"
	backupRestoreService "github.com/stackrox/stackrox/central/globaldb/v2backuprestore/service"
	graphqlHandler "github.com/stackrox/stackrox/central/graphql/handler"
	groupDataStore "github.com/stackrox/stackrox/central/group/datastore"
	groupService "github.com/stackrox/stackrox/central/group/service"
	"github.com/stackrox/stackrox/central/grpc/metrics"
	"github.com/stackrox/stackrox/central/helmcharts"
	imageDatastore "github.com/stackrox/stackrox/central/image/datastore"
	imageService "github.com/stackrox/stackrox/central/image/service"
	iiDatastore "github.com/stackrox/stackrox/central/imageintegration/datastore"
	iiService "github.com/stackrox/stackrox/central/imageintegration/service"
	iiStore "github.com/stackrox/stackrox/central/imageintegration/store"
	integrationHealthService "github.com/stackrox/stackrox/central/integrationhealth/service"
	"github.com/stackrox/stackrox/central/jwt"
	licenseService "github.com/stackrox/stackrox/central/license/service"
	licenseSingletons "github.com/stackrox/stackrox/central/license/singleton"
	logimbueHandler "github.com/stackrox/stackrox/central/logimbue/handler"
	logimbueStore "github.com/stackrox/stackrox/central/logimbue/store"
	metadataService "github.com/stackrox/stackrox/central/metadata/service"
	mitreService "github.com/stackrox/stackrox/central/mitre/service"
	namespaceService "github.com/stackrox/stackrox/central/namespace/service"
	networkBaselineDataStore "github.com/stackrox/stackrox/central/networkbaseline/datastore"
	networkBaselineService "github.com/stackrox/stackrox/central/networkbaseline/service"
	networkEntityDataStore "github.com/stackrox/stackrox/central/networkgraph/entity/datastore"
	"github.com/stackrox/stackrox/central/networkgraph/entity/gatherer"
	networkFlowService "github.com/stackrox/stackrox/central/networkgraph/service"
	networkPolicyService "github.com/stackrox/stackrox/central/networkpolicies/service"
	nodeService "github.com/stackrox/stackrox/central/node/service"
	"github.com/stackrox/stackrox/central/notifier/processor"
	notifierService "github.com/stackrox/stackrox/central/notifier/service"
	_ "github.com/stackrox/stackrox/central/notifiers/all" // These imports are required to register things from the respective packages.
	"github.com/stackrox/stackrox/central/option"
	pingService "github.com/stackrox/stackrox/central/ping/service"
	podService "github.com/stackrox/stackrox/central/pod/service"
	policyDataStore "github.com/stackrox/stackrox/central/policy/datastore"
	policyService "github.com/stackrox/stackrox/central/policy/service"
	probeUploadService "github.com/stackrox/stackrox/central/probeupload/service"
	processBaselineDataStore "github.com/stackrox/stackrox/central/processbaseline/datastore"
	processBaselineService "github.com/stackrox/stackrox/central/processbaseline/service"
	processIndicatorService "github.com/stackrox/stackrox/central/processindicator/service"
	"github.com/stackrox/stackrox/central/pruning"
	rbacService "github.com/stackrox/stackrox/central/rbac/service"
	reportConfigurationService "github.com/stackrox/stackrox/central/reportconfigurations/service"
	vulnReportScheduleManager "github.com/stackrox/stackrox/central/reports/manager"
	reportService "github.com/stackrox/stackrox/central/reports/service"
	"github.com/stackrox/stackrox/central/reprocessor"
	"github.com/stackrox/stackrox/central/risk/handlers/timeline"
	"github.com/stackrox/stackrox/central/role"
	roleDataStore "github.com/stackrox/stackrox/central/role/datastore"
	"github.com/stackrox/stackrox/central/role/mapper"
	"github.com/stackrox/stackrox/central/role/resources"
	roleService "github.com/stackrox/stackrox/central/role/service"
	centralSAC "github.com/stackrox/stackrox/central/sac"
	sacService "github.com/stackrox/stackrox/central/sac/service"
	"github.com/stackrox/stackrox/central/sac/transitional"
	"github.com/stackrox/stackrox/central/scanner"
	scannerDefinitionsHandler "github.com/stackrox/stackrox/central/scannerdefinitions/handler"
	searchService "github.com/stackrox/stackrox/central/search/service"
	secretService "github.com/stackrox/stackrox/central/secret/service"
	sensorService "github.com/stackrox/stackrox/central/sensor/service"
	"github.com/stackrox/stackrox/central/sensor/service/connection"
	"github.com/stackrox/stackrox/central/sensor/service/pipeline/all"
	sensorUpgradeControlService "github.com/stackrox/stackrox/central/sensorupgrade/controlservice"
	sensorUpgradeService "github.com/stackrox/stackrox/central/sensorupgrade/service"
	sensorUpgradeConfigStore "github.com/stackrox/stackrox/central/sensorupgradeconfig/datastore"
	serviceAccountService "github.com/stackrox/stackrox/central/serviceaccount/service"
	siStore "github.com/stackrox/stackrox/central/serviceidentities/datastore"
	siService "github.com/stackrox/stackrox/central/serviceidentities/service"
	signatureIntegrationService "github.com/stackrox/stackrox/central/signatureintegration/service"
	"github.com/stackrox/stackrox/central/splunk"
	summaryService "github.com/stackrox/stackrox/central/summary/service"
	"github.com/stackrox/stackrox/central/telemetry/gatherers"
	telemetryService "github.com/stackrox/stackrox/central/telemetry/service"
	"github.com/stackrox/stackrox/central/tlsconfig"
	"github.com/stackrox/stackrox/central/ui"
	userService "github.com/stackrox/stackrox/central/user/service"
	"github.com/stackrox/stackrox/central/version"
	vulnRequestManager "github.com/stackrox/stackrox/central/vulnerabilityrequest/manager/requestmgr"
	vulnRequestService "github.com/stackrox/stackrox/central/vulnerabilityrequest/service"
	"github.com/stackrox/stackrox/generated/storage"
	"github.com/stackrox/stackrox/pkg/auth/authproviders"
	"github.com/stackrox/stackrox/pkg/auth/authproviders/iap"
	"github.com/stackrox/stackrox/pkg/auth/authproviders/oidc"
	"github.com/stackrox/stackrox/pkg/auth/authproviders/openshift"
	"github.com/stackrox/stackrox/pkg/auth/authproviders/saml"
	authProviderUserpki "github.com/stackrox/stackrox/pkg/auth/authproviders/userpki"
	"github.com/stackrox/stackrox/pkg/auth/permissions"
	"github.com/stackrox/stackrox/pkg/concurrency"
	"github.com/stackrox/stackrox/pkg/config"
	"github.com/stackrox/stackrox/pkg/devbuild"
	"github.com/stackrox/stackrox/pkg/devmode"
	"github.com/stackrox/stackrox/pkg/env"
	"github.com/stackrox/stackrox/pkg/features"
	pkgGRPC "github.com/stackrox/stackrox/pkg/grpc"
	"github.com/stackrox/stackrox/pkg/grpc/authn"
	"github.com/stackrox/stackrox/pkg/grpc/authn/service"
	"github.com/stackrox/stackrox/pkg/grpc/authn/servicecerttoken"
	"github.com/stackrox/stackrox/pkg/grpc/authn/tokenbased"
	authnUserpki "github.com/stackrox/stackrox/pkg/grpc/authn/userpki"
	"github.com/stackrox/stackrox/pkg/grpc/authz"
	"github.com/stackrox/stackrox/pkg/grpc/authz/allow"
	"github.com/stackrox/stackrox/pkg/grpc/authz/or"
	"github.com/stackrox/stackrox/pkg/grpc/authz/perrpc"
	"github.com/stackrox/stackrox/pkg/grpc/authz/user"
	"github.com/stackrox/stackrox/pkg/grpc/errors"
	"github.com/stackrox/stackrox/pkg/grpc/routes"
	"github.com/stackrox/stackrox/pkg/httputil/proxy"
	"github.com/stackrox/stackrox/pkg/logging"
	pkgMetrics "github.com/stackrox/stackrox/pkg/metrics"
	"github.com/stackrox/stackrox/pkg/migrations"
	"github.com/stackrox/stackrox/pkg/osutils"
	"github.com/stackrox/stackrox/pkg/premain"
	"github.com/stackrox/stackrox/pkg/sac"
	"github.com/stackrox/stackrox/pkg/sac/observe"
	"github.com/stackrox/stackrox/pkg/sync"
	pkgVersion "github.com/stackrox/stackrox/pkg/version"
)

var (
	log = logging.CreatePersistentLogger(logging.CurrentModule(), 0)

	authProviderBackendFactories = map[string]authproviders.BackendFactoryCreator{
		oidc.TypeName:                oidc.NewFactory,
		"auth0":                      oidc.NewFactory, // legacy
		saml.TypeName:                saml.NewFactory,
		authProviderUserpki.TypeName: authProviderUserpki.NewFactoryFactory(tlsconfig.ManagerInstance()),
		iap.TypeName:                 iap.NewFactory,
	}

	imageIntegrationContext = sac.WithGlobalAccessScopeChecker(context.Background(),
		sac.AllowFixedScopes(
			sac.AccessModeScopeKeys(storage.Access_READ_ACCESS, storage.Access_READ_WRITE_ACCESS),
			sac.ResourceScopeKeys(resources.ImageIntegration),
		))
)

const (
	ssoURLPathPrefix     = "/sso/"
	tokenRedirectURLPath = "/auth/response/generic"

	grpcServerWatchdogTimeout = 20 * time.Second

	maxServiceCertTokenLeeway = 1 * time.Minute

	proxyConfigPath = "/run/secrets/stackrox.io/proxy-config"
	proxyConfigFile = "config.yaml"
)

func init() {
	if !proxy.UseWithDefaultTransport() {
		log.Warn("Failed to use proxy transport with default HTTP transport. Some proxy features may not work.")
	}
}

func runSafeMode() {
	log.Info("Started Central up in safe mode. Sleeping forever...")

	signalsC := make(chan os.Signal, 1)
	signal.Notify(signalsC, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)

	sig := <-signalsC
	log.Infof("Caught %s signal", sig)
	log.Info("Central terminated")
}

func main() {
	premain.StartMain()

	conf := config.GetConfig()
	if conf == nil || conf.Maintenance.SafeMode {
		if conf == nil {
			log.Error("cannot get central configuration. Starting up in safe mode")
		}
		runSafeMode()
		return
	}

	proxy.WatchProxyConfig(context.Background(), proxyConfigPath, proxyConfigFile, true)

	devmode.StartOnDevBuilds("central")

	log.Infof("Running StackRox Version: %s", pkgVersion.GetMainVersion())
	ensureDB()

	// Now that we verified that the DB can be loaded, remove the .backup directory
	if err := migrations.SafeRemoveDBWithSymbolicLink(filepath.Join(migrations.DBMountPath(), ".backup")); err != nil {
		log.Errorf("Failed to remove backup DB: %v", err)
	}

	// Update last associated software version on DBs.
	migrations.SetCurrent(option.CentralOptions.DBPathBase)

	// Start the prometheus metrics server
	pkgMetrics.NewDefaultHTTPServer().RunForever()
	pkgMetrics.GatherThrottleMetricsForever(pkgMetrics.CentralSubsystem.String())

	licenseMgr := licenseSingletons.ManagerSingleton()
	_, err := licenseMgr.Initialize()
	if err != nil {
		log.Warnf("Could not initialize license manager: %v", err)
	}

	go startGRPCServer()

	waitForTerminationSignal()
}

func ensureDB() {
	err := version.Ensure(globaldb.GetGlobalDB(), globaldb.GetRocksDB())
	if err != nil {
		log.Panicf("DB version check failed. You may need to run migrations: %v", err)
	}
}

func startServices() {
	if err := complianceManager.Singleton().Start(); err != nil {
		log.Panicf("could not start compliance manager: %v", err)
	}
	reprocessor.Singleton().Start()
	suppress.Singleton().Start()
	pruning.Singleton().Start()
	gatherer.Singleton().Start()
	vulnRequestManager.Singleton().Start()

	go registerDelayedIntegrations(iiStore.DelayedIntegrations)
}

func servicesToRegister(registry authproviders.Registry, authzTraceSink observe.AuthzTraceSink) []pkgGRPC.APIService {
	// PLEASE KEEP THE FOLLOWING LIST SORTED.
	servicesToRegister := []pkgGRPC.APIService{
		alertService.Singleton(),
		apiTokenService.Singleton(),
		authService.New(),
		authProviderService.New(registry, groupDataStore.Singleton()),
		backupRestoreService.Singleton(),
		backupService.Singleton(),
		centralHealthService.Singleton(),
		certgen.ServiceSingleton(),
		clusterInitService.Singleton(),
		clusterService.Singleton(),
		complianceManagerService.Singleton(),
		complianceService.Singleton(),
		configService.Singleton(),
		credentialExpiryService.Singleton(),
		cveService.Singleton(),
		debugService.New(
			clusterDataStore.Singleton(),
			connection.ManagerSingleton(),
			gatherers.Singleton(),
			logimbueStore.Singleton(),
			authzTraceSink,
		),
		deploymentService.Singleton(),
		detectionService.Singleton(),
		featureFlagService.Singleton(),
		groupService.Singleton(),
		helmcharts.NewService(),
		imageService.Singleton(),
		iiService.Singleton(),
		licenseService.New(false, licenseSingletons.ManagerSingleton()),
		integrationHealthService.Singleton(),
		metadataService.New(),
		mitreService.Singleton(),
		namespaceService.Singleton(),
		networkBaselineService.Singleton(),
		networkFlowService.Singleton(),
		networkPolicyService.Singleton(),
		nodeService.Singleton(),
		notifierService.Singleton(),
		pingService.Singleton(),
		podService.Singleton(),
		policyService.Singleton(),
		probeUploadService.Singleton(),
		processIndicatorService.Singleton(),
		processBaselineService.Singleton(),
		rbacService.Singleton(),
		roleService.Singleton(),
		sacService.Singleton(),
		searchService.Singleton(),
		secretService.Singleton(),
		sensorService.New(connection.ManagerSingleton(), all.Singleton(), clusterDataStore.Singleton()),
		sensorUpgradeControlService.Singleton(),
		sensorUpgradeService.Singleton(),
		serviceAccountService.Singleton(),
		siService.Singleton(),
		summaryService.Singleton(),
		telemetryService.Singleton(),
		userService.Singleton(),
		vulnRequestService.Singleton(),
	}

	if features.VulnReporting.Enabled() {
		servicesToRegister = append(servicesToRegister,
			reportConfigurationService.Singleton(),
			reportService.Singleton())
	}

	if features.ImageSignatureVerification.Enabled() {
		servicesToRegister = append(servicesToRegister, signatureIntegrationService.Singleton())
	}

	autoTriggerUpgrades := sensorUpgradeConfigStore.Singleton().AutoTriggerSetting()
	if err := connection.ManagerSingleton().Start(
		clusterDataStore.Singleton(),
		networkEntityDataStore.Singleton(),
		policyDataStore.Singleton(),
		processBaselineDataStore.Singleton(),
		networkBaselineDataStore.Singleton(),
		autoTriggerUpgrades,
	); err != nil {
		log.Panicf("Couldn't start sensor connection manager: %v", err)
	}

	if !env.OfflineModeEnv.BooleanSetting() {
		go fetcher.SingletonManager().Start()
	}

	if devbuild.IsEnabled() {
		servicesToRegister = append(servicesToRegister, developmentService.Singleton())
	}

	return servicesToRegister
}

func watchdog(signal *concurrency.Signal, timeout time.Duration) {
	if !concurrency.WaitWithTimeout(signal, timeout) {
		log.Errorf("API server failed to start within %v!", timeout)
		log.Error("This usually means something is *very* wrong. Terminating ...")
		if err := syscall.Kill(syscall.Getpid(), syscall.SIGABRT); err != nil {
			panic(err)
		}
	}
}

func startGRPCServer() {
	// Temporarily elevate permissions to modify auth providers.
	authProviderRegisteringCtx := sac.WithGlobalAccessScopeChecker(context.Background(),
		sac.AllowFixedScopes(
			sac.AccessModeScopeKeys(storage.Access_READ_ACCESS, storage.Access_READ_WRITE_ACCESS),
			sac.ResourceScopeKeys(resources.AuthProvider)))

	// Create the registry of applied auth providers.
	registry, err := authproviders.NewStoreBackedRegistry(
		ssoURLPathPrefix, tokenRedirectURLPath,
		authProviderDS.Singleton(), jwt.IssuerFactorySingleton(),
		mapper.FactorySingleton())
	if err != nil {
		log.Panicf("Could not create auth provider registry: %v", err)
	}

	// env.EnableOpenShiftAuth signals the desire but does not guarantee Central
	// is configured correctly to talk to the OpenShift's OAuth server. If this
	// is the case, we can be setting up an auth providers which won't work.
	if env.EnableOpenShiftAuth.BooleanSetting() {
		authProviderBackendFactories[openshift.TypeName] = openshift.NewFactory
	}

	for typeName, factoryCreator := range authProviderBackendFactories {
		if err := registry.RegisterBackendFactory(authProviderRegisteringCtx, typeName, factoryCreator); err != nil {
			log.Panicf("Could not register %s auth provider factory: %v", typeName, err)
		}
	}
	if err := registry.Init(); err != nil {
		log.Panicf("Could not initialize auth provider registry: %v", err)
	}

	basicAuthMgr, err := userpass.CreateManager(roleDataStore.Singleton())
	if err != nil {
		log.Panicf("Could not create basic auth manager: %v", err)
	}

	basicAuthProvider := userpass.RegisterAuthProviderOrPanic(authProviderRegisteringCtx, basicAuthMgr, registry)

	clusterInitBackend := backend.Singleton()
	serviceMTLSExtractor, err := service.NewExtractorWithCertValidation(clusterInitBackend)
	if err != nil {
		log.Panicf("Could not create mTLS-based service identity extractor: %v", err)
	}

	serviceTokenExtractor, err := servicecerttoken.NewExtractorWithCertValidation(maxServiceCertTokenLeeway, clusterInitBackend)
	if err != nil {
		log.Panicf("Could not create ServiceCert token-based identity extractor: %v", err)
	}

	idExtractors := []authn.IdentityExtractor{
		serviceMTLSExtractor, // internal services
		tokenbased.NewExtractor(roleDataStore.Singleton(), jwt.ValidatorSingleton()), // JWT tokens
		userpass.IdentityExtractorOrPanic(roleDataStore.Singleton(), basicAuthMgr, basicAuthProvider),
		serviceTokenExtractor,
		authnUserpki.NewExtractor(tlsconfig.ManagerInstance()),
	}

	endpointCfgs, err := endpoints.InstantiateAll(tlsconfig.ManagerInstance())
	if err != nil {
		log.Panicf("Could not instantiate endpoint configs: %v", err)
	}

	config := pkgGRPC.Config{
		CustomRoutes:       customRoutes(),
		IdentityExtractors: idExtractors,
		AuthProviders:      registry,
		Auditor:            audit.New(processor.Singleton()),
		GRPCMetrics:        metrics.GRPCSingleton(),
		HTTPMetrics:        metrics.HTTPSingleton(),
		Endpoints:          endpointCfgs,
	}

	if devbuild.IsEnabled() {
		config.UnaryInterceptors = append(config.UnaryInterceptors,
			errors.LogInternalErrorInterceptor,
			errors.PanicOnInvariantViolationUnaryInterceptor,
		)
		config.StreamInterceptors = append(config.StreamInterceptors,
			errors.LogInternalErrorStreamInterceptor,
			errors.PanicOnInvariantViolationStreamInterceptor,
		)
		// This helps validate that SAC is being used correctly.
		config.UnaryInterceptors = append(config.UnaryInterceptors, transitional.VerifySACScopeChecksInterceptor)
	}

	// This adds an on-demand global tracing for the built-in authorization.
	authzTraceSink := observe.NewAuthzTraceSink()
	config.UnaryInterceptors = append(config.UnaryInterceptors,
		observe.AuthzTraceInterceptor(authzTraceSink),
	)
	config.HTTPInterceptors = append(config.HTTPInterceptors, observe.AuthzTraceHTTPInterceptor(authzTraceSink))

	// Before authorization is checked, we want to inject the sac client into the context.
	config.PreAuthContextEnrichers = append(config.PreAuthContextEnrichers,
		centralSAC.GetEnricher().GetPreAuthContextEnricher(authzTraceSink),
	)

	server := pkgGRPC.NewAPI(config)
	server.Register(servicesToRegister(registry, authzTraceSink)...)

	startServices()
	startedSig := server.Start()

	go watchdog(startedSig, grpcServerWatchdogTimeout)
}

func registerDelayedIntegrations(integrationsInput []iiStore.DelayedIntegration) {
	integrationManager := enrichment.ManagerSingleton()

	integrations := make(map[int]iiStore.DelayedIntegration, len(integrationsInput))
	for k, v := range integrationsInput {
		integrations[k] = v
	}
	ds := iiDatastore.Singleton()
	for len(integrations) > 0 {
		for idx, integration := range integrations {
			_, exists, _ := ds.GetImageIntegration(imageIntegrationContext, integration.Integration.GetId())
			if exists {
				delete(integrations, idx)
				continue
			}
			ready := integration.Trigger()
			if !ready {
				continue
			}
			// add the integration first, which is more likely to fail. If it does, no big deal -- you can still try to
			// manually add it and get the error message.
			err := integrationManager.Upsert(integration.Integration)
			if err == nil {
				err = ds.UpdateImageIntegration(imageIntegrationContext, integration.Integration)
				if err != nil {
					// so, we added the integration to the set but we weren't able to save it.
					// This is ok -- the image scanner will "work" and after a restart we'll try to save it again.
					log.Errorf("We added the %q integration, but saving it failed with: %v. We'll try again next restart", integration.Integration.GetName(), err)
				} else {
					log.Infof("Registered integration %q", integration.Integration.GetName())
				}
				reprocessor.Singleton().ShortCircuit()
			} else {
				log.Errorf("Unable to register integration %q: %v", integration.Integration.GetName(), err)
			}
			// either way, time to stop watching this entry
			delete(integrations, idx)
		}
		time.Sleep(5 * time.Second)
	}
	log.Debug("All dynamic integrations registered, exiting")
}

func uiRoute() routes.CustomRoute {
	return routes.CustomRoute{
		Route:         "/",
		Authorizer:    allow.Anonymous(),
		ServerHandler: ui.Mux(),
		Compression:   true,
	}
}

func customRoutes() (customRoutes []routes.CustomRoute) {
	customRoutes = []routes.CustomRoute{
		uiRoute(),
		{
			Route:         "/api/extensions/clusters/zip",
			Authorizer:    or.SensorOrAuthorizer(user.With(permissions.View(resources.Cluster), permissions.View(resources.ServiceIdentity))),
			ServerHandler: clustersZip.Handler(clusterDataStore.Singleton(), siStore.Singleton()),
			Compression:   false,
		},
		{
			Route:         "/api/extensions/scanner/zip",
			Authorizer:    user.With(permissions.View(resources.ScannerBundle)),
			ServerHandler: scanner.Handler(),
			Compression:   false,
		},
		{
			Route:         "/api/cli/download/",
			Authorizer:    user.With(),
			ServerHandler: cli.Handler(),
			Compression:   true,
		},
		{
			Route:         "/db/backup",
			Authorizer:    dbAuthz.DBReadAccessAuthorizer(),
			ServerHandler: globaldbHandlers.BackupDB(globaldb.GetGlobalDB(), globaldb.GetRocksDB(), false),
			Compression:   true,
		},
		{
			Route:         "/api/extensions/backup",
			Authorizer:    user.WithRole(role.Admin),
			ServerHandler: globaldbHandlers.BackupDB(globaldb.GetGlobalDB(), globaldb.GetRocksDB(), true),
			Compression:   true,
		},
		{
			Route:         "/db/restore",
			Authorizer:    dbAuthz.DBWriteAccessAuthorizer(),
			ServerHandler: globaldbHandlers.RestoreDB(globaldb.GetGlobalDB(), globaldb.GetRocksDB()),
		},
		{
			Route:         "/api/docs/swagger",
			Authorizer:    user.With(permissions.View(resources.APIToken)),
			ServerHandler: docs.Swagger(),
			Compression:   true,
		},
		{
			Route:         "/api/graphql",
			Authorizer:    user.With(), // graphql enforces permissions internally
			ServerHandler: graphqlHandler.Handler(),
			Compression:   true,
		},
		{
			Route:         "/api/compliance/export/csv",
			Authorizer:    user.With(permissions.View(resources.Compliance)),
			ServerHandler: complianceHandlers.CSVHandler(),
			Compression:   true,
		},
		{
			Route:         "/api/risk/timeline/export/csv",
			Authorizer:    user.With(permissions.View(resources.Deployment), permissions.View(resources.Indicator), permissions.View(resources.ProcessWhitelist)),
			ServerHandler: timeline.CSVHandler(),
			Compression:   true,
		},
		{
			Route:         "/api/vm/export/csv",
			Authorizer:    user.With(permissions.View(resources.Image), permissions.View(resources.Deployment), permissions.View(resources.Node)),
			ServerHandler: csv.CVECSVHandler(),
			Compression:   true,
		},
		{
			Route:         "/api/splunk/ta/vulnmgmt",
			Authorizer:    user.With(permissions.View(resources.Image), permissions.View(resources.Deployment)),
			ServerHandler: splunk.NewVulnMgmtHandler(deploymentDatastore.Singleton(), imageDatastore.Singleton()),
			Compression:   true,
		},
		{
			Route:         "/api/splunk/ta/compliance",
			Authorizer:    user.With(permissions.View(resources.Compliance)),
			ServerHandler: splunk.NewComplianceHandler(complianceDatastore.Singleton()),
			Compression:   true,
		},
		{
			Route:         "/api/splunk/ta/violations",
			Authorizer:    user.With(permissions.View(resources.Alert)),
			ServerHandler: splunk.NewViolationsHandler(alertDatastore.Singleton()),
			Compression:   true,
		},
		{
			Route:         "/db/v2/restore",
			Authorizer:    dbAuthz.DBWriteAccessAuthorizer(),
			ServerHandler: backupRestoreService.Singleton().RestoreHandler(),
		},
		{
			Route:         "/db/v2/resumerestore",
			Authorizer:    dbAuthz.DBWriteAccessAuthorizer(),
			ServerHandler: backupRestoreService.Singleton().ResumeRestoreHandler(),
		},
		{
			Route:         "/api/logimbue",
			Authorizer:    user.With(),
			ServerHandler: logimbueHandler.Singleton(),
			Compression:   false,
		},
	}

	customRoutes = append(customRoutes, routes.CustomRoute{
		Route:         "/api/extensions/clusters/helm-config.yaml",
		Authorizer:    or.SensorOrAuthorizer(user.With(permissions.View(resources.Cluster))),
		ServerHandler: clustersHelmConfig.Handler(clusterDataStore.Singleton()),
		Compression:   true,
	})

	scannerDefinitionsRoute := "/api/extensions/scannerdefinitions"
	customRoutes = append(customRoutes,
		routes.CustomRoute{
			Route: scannerDefinitionsRoute,
			Authorizer: perrpc.FromMap(map[authz.Authorizer][]string{
				or.ScannerOr(user.With(permissions.View(resources.ScannerDefinitions))): {
					routes.RPCNameForHTTP(scannerDefinitionsRoute, http.MethodGet),
				},
				user.With(permissions.Modify(resources.ScannerDefinitions)): {
					routes.RPCNameForHTTP(scannerDefinitionsRoute, http.MethodPost),
				},
			}),
			ServerHandler: scannerDefinitionsHandler.Singleton(),
			Compression:   false,
		},
	)

	customRoutes = append(customRoutes, debugRoutes()...)
	return
}

func debugRoutes() []routes.CustomRoute {
	customRoutes := make([]routes.CustomRoute, 0, len(routes.DebugRoutes))

	for r, h := range routes.DebugRoutes {
		customRoutes = append(customRoutes, routes.CustomRoute{
			Route:         r,
			Authorizer:    user.WithRole(role.Admin),
			ServerHandler: h,
			Compression:   true,
		})
	}
	return customRoutes
}

type stoppable interface {
	Stop()
}

type stoppableWithName struct {
	obj  stoppable
	name string
}

func waitForTerminationSignal() {
	signalsC := make(chan os.Signal, 1)
	signal.Notify(signalsC, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	sig := <-signalsC
	log.Infof("Caught %s signal", sig)

	stoppables := []stoppableWithName{
		{reprocessor.Singleton(), "reprocessor loop"},
		{suppress.Singleton(), "cve unsuppress loop"},
		{pruning.Singleton(), "gargage collector"},
		{gatherer.Singleton(), "network graph default external sources gatherer"},
		{vulnRequestManager.Singleton(), "vuln deferral requests expiry loop"},
	}

	if features.VulnReporting.Enabled() {
		stoppables = append(stoppables, stoppableWithName{vulnReportScheduleManager.Singleton(), "vuln reports schedule manager"})
	}

	var wg sync.WaitGroup
	for _, stoppable := range stoppables {
		wg.Add(1)
		go func(s stoppableWithName) {
			defer wg.Done()
			s.obj.Stop()
			log.Infof("Stopped %s", s.name)
		}(stoppable)
	}
	wg.Wait()

	globaldb.Close()

	if sig == syscall.SIGHUP {
		log.Info("Restarting central")
		osutils.Restart()
	}
	log.Info("Central terminated")
}
