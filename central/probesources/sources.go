package probesources

import (
	"context"
	"net/http"
	"time"

	"github.com/hashicorp/go-multierror"
	"github.com/stackrox/stackrox/central/clusters"
	licenseManager "github.com/stackrox/stackrox/central/license/manager"
	licenseSingletons "github.com/stackrox/stackrox/central/license/singleton"
	probeUploadManager "github.com/stackrox/stackrox/central/probeupload/manager"
	"github.com/stackrox/stackrox/pkg/env"
	"github.com/stackrox/stackrox/pkg/httputil/proxy"
	"github.com/stackrox/stackrox/pkg/kocache"
	"github.com/stackrox/stackrox/pkg/logging"
	"github.com/stackrox/stackrox/pkg/probeupload"
	"github.com/stackrox/stackrox/pkg/sync"
)

const (
	httpTimeout = 30 * time.Second
)

var (
	log          = logging.LoggerForModule()
	instance     probeSources
	instanceInit sync.Once
)

//go:generate mockgen-wrapper

// ProbeSources interface provides the availability of the probes packages.
type ProbeSources interface {
	AnyAvailable(ctx context.Context) (bool, error)
	CopyAsSlice() []probeupload.ProbeSource
}

// probeSources contains the list of activated probe sources.
type probeSources struct {
	probeSources []probeupload.ProbeSource
}

// CopyAsSlice retrieves the activated probe sources as a slice backed by newly allocated memory.
func (s *probeSources) CopyAsSlice() []probeupload.ProbeSource {
	probeSources := make([]probeupload.ProbeSource, len(s.probeSources))
	copy(probeSources, s.probeSources)
	return probeSources
}

// AnyAvailable implements a simple heuristic for the availability of kernel probes.
// It returns true if any of the activated probe sources is available in the sense
// that it does support the transmitting of (some) kernel probes.
func (s *probeSources) AnyAvailable(ctx context.Context) (bool, error) {
	var finalErr error

	for _, source := range s.probeSources {
		isAvailable, err := source.IsAvailable(ctx)
		if err != nil {
			log.Warnf("Failed to check availability of kernel probe source %T: %v", source, err)
			finalErr = multierror.Append(finalErr, err)
		}
		if isAvailable {
			return true, nil
		}
	}

	return false, finalErr
}

func (s *probeSources) initializeStandardSources(probeUploadManager probeUploadManager.Manager, licenseMgr licenseManager.LicenseManager) {
	s.probeSources = make([]probeupload.ProbeSource, 0, 2)
	s.probeSources = append(s.probeSources, probeUploadManager)
	if env.OfflineModeEnv.BooleanSetting() {
		return
	}
	baseURL := clusters.CollectorModuleDownloadBaseURL.Setting()
	if baseURL == "" {
		return
	}

	opts := kocache.Options{}
	if licenseMgr != nil {
		opts.ModifyRequest = func(req *http.Request) {
			customerID := licenseMgr.GetActiveLicense().GetMetadata().GetLicensedForId()
			if customerID == "" {
				return
			}
			q := req.URL.Query()
			q.Set("cid", customerID)
			req.URL.RawQuery = q.Encode()
		}
	}

	httpClient := &http.Client{
		Transport: proxy.RoundTripper(),
		Timeout:   httpTimeout,
	}
	s.probeSources = append(s.probeSources, kocache.New(context.Background(), httpClient, baseURL, opts))
}

// Singleton returns the singleton instance for the ProbeSources entity.
func Singleton() ProbeSources {
	instanceInit.Do(func() {
		instance.initializeStandardSources(probeUploadManager.Singleton(), licenseSingletons.ManagerSingleton())
	})
	return &instance
}
