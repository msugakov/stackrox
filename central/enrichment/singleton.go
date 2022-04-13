package enrichment

import (
	"time"

	cveDataStore "github.com/stackrox/stackrox/central/cve/datastore"
	"github.com/stackrox/stackrox/central/cve/fetcher"
	"github.com/stackrox/stackrox/central/image/datastore"
	"github.com/stackrox/stackrox/central/imageintegration"
	"github.com/stackrox/stackrox/central/integrationhealth/reporter"
	signatureIntegrationDataStore "github.com/stackrox/stackrox/central/signatureintegration/datastore"
	"github.com/stackrox/stackrox/central/vulnerabilityrequest/suppressor"
	"github.com/stackrox/stackrox/pkg/expiringcache"
	imageEnricher "github.com/stackrox/stackrox/pkg/images/enricher"
	"github.com/stackrox/stackrox/pkg/metrics"
	nodeEnricher "github.com/stackrox/stackrox/pkg/nodes/enricher"
	"github.com/stackrox/stackrox/pkg/sync"
)

var (
	once sync.Once

	ie      imageEnricher.ImageEnricher
	ne      nodeEnricher.NodeEnricher
	en      Enricher
	cf      fetcher.OrchestratorIstioCVEManager
	manager Manager

	metadataCacheOnce sync.Once
	metadataCache     expiringcache.Cache

	imageCacheExpiryDuration = 4 * time.Hour
)

func initialize() {
	ie = imageEnricher.New(cveDataStore.Singleton(), suppressor.Singleton(), imageintegration.Set(),
		metrics.CentralSubsystem, ImageMetadataCacheSingleton(), datastore.Singleton().GetImage, reporter.Singleton(),
		signatureIntegrationDataStore.Singleton().GetAllSignatureIntegrations)
	ne = nodeEnricher.New(cveDataStore.Singleton(), metrics.CentralSubsystem)
	en = New(datastore.Singleton(), ie)
	cf = fetcher.SingletonManager()
	manager = newManager(imageintegration.Set(), ne, cf)
}

// Singleton provides the singleton Enricher to use.
func Singleton() Enricher {
	once.Do(initialize)
	return en
}

// ImageEnricherSingleton provides the singleton ImageEnricher to use.
func ImageEnricherSingleton() imageEnricher.ImageEnricher {
	once.Do(initialize)
	return ie
}

// ImageMetadataCacheSingleton returns the cache for image metadata
func ImageMetadataCacheSingleton() expiringcache.Cache {
	metadataCacheOnce.Do(func() {
		metadataCache = expiringcache.NewExpiringCache(imageCacheExpiryDuration, expiringcache.UpdateExpirationOnGets)
	})
	return metadataCache
}

// NodeEnricherSingleton provides the singleton NodeEnricher to use.
func NodeEnricherSingleton() nodeEnricher.NodeEnricher {
	once.Do(initialize)
	return ne
}

// ManagerSingleton returns the multiplexing manager
func ManagerSingleton() Manager {
	once.Do(initialize)
	return manager
}
