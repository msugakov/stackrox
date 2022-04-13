package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/pkg/errors"
	"github.com/stackrox/stackrox/central/notifiers"
	"github.com/stackrox/stackrox/central/notifiers/cscc/findings"
	"github.com/stackrox/stackrox/pkg/httputil/proxy"
	"github.com/stackrox/stackrox/pkg/utils"
	"golang.org/x/oauth2/google"
)

const (
	cloudPlatformScope = "https://www.googleapis.com/auth/cloud-platform"
	timeout            = 5 * time.Second
)

var (
	client = &http.Client{
		Transport: proxy.RoundTripper(),
	}
)

// Logger is the minimal interface we need to use to log data.
type Logger interface {
	Warnf(format string, args ...interface{})
	Infof(format string, args ...interface{})
	Debugf(format string, args ...interface{})
}

// A Config contains the necessary information to make a Cloud SCC request.
type Config struct {
	SourceID       string // The Source ID is assigned by SCC when a Source is created for StackRox.
	ServiceAccount []byte
	Logger         Logger
}

func (c *Config) postURL(findingID string) string {
	return fmt.Sprintf(
		"https://securitycenter.googleapis.com/v1beta1/%s/findings?findingId=%s",
		c.SourceID,
		findingID,
	)
}

// CreateFinding creates the provided Finding.
func (c *Config) CreateFinding(ctx context.Context, finding *findings.Finding, id string) error {
	req, err := c.request(finding, id)
	if err != nil {
		return errors.Wrap(err, "request creation")
	}

	ctx, cancel := timeoutContext(ctx)
	defer cancel()
	tokenSource, err := c.getTokenSource(ctx)
	if err != nil {
		return errors.Wrap(err, "token source retrieval")
	}

	token, err := tokenSource.TokenSource.Token()
	if err != nil {
		return errors.Wrap(err, "token retrieval")
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token.AccessToken))

	resp, err := client.Do(req.WithContext(ctx))
	if err != nil {
		return errors.Wrap(err, "request")
	}
	defer utils.IgnoreError(resp.Body.Close)

	return notifiers.CreateError("Cloud SCC", resp)
}

func (c *Config) request(finding *findings.Finding, id string) (*http.Request, error) {
	b, err := json.Marshal(finding)
	if err != nil {
		return nil, errors.Wrap(err, "marshal")
	}
	c.Logger.Debugf("Request: %s", string(b))

	req, err := http.NewRequest("POST", c.postURL(id), bytes.NewReader(b))
	if err != nil {
		return nil, errors.Wrap(err, "build")
	}
	return req, nil
}

func (c *Config) getTokenSource(ctx context.Context) (*google.DefaultCredentials, error) {
	cfg, err := google.JWTConfigFromJSON(c.ServiceAccount, cloudPlatformScope)
	if err != nil {
		return nil, errors.Wrap(err, "google.JWTConfigFromJSON")
	}
	pid, err := c.embeddedProjectID()
	if err != nil {
		return nil, errors.Wrap(err, "project ID retrieval")
	}
	return &google.DefaultCredentials{
		ProjectID:   pid,
		TokenSource: cfg.TokenSource(ctx),
		JSON:        c.ServiceAccount,
	}, nil
}

func timeoutContext(ctx context.Context) (context.Context, context.CancelFunc) {
	return context.WithTimeout(ctx, timeout)
}

func (c *Config) embeddedProjectID() (string, error) {
	// jwt.Config does not expose the project ID, so re-unmarshal to get it.
	var pid struct {
		ProjectID string `json:"project_id"`
	}
	if err := json.Unmarshal(c.ServiceAccount, &pid); err != nil {
		return "", err
	}
	return pid.ProjectID, nil
}
