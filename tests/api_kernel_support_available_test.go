package tests

import (
	"context"
	"testing"
	"time"

	v1 "github.com/stackrox/stackrox/generated/api/v1"
	"github.com/stackrox/stackrox/pkg/testutils"
	"github.com/stretchr/testify/assert"
)

func TestKernelSupportAvailableApi(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	conn := testutils.GRPCConnectionToCentral(t)

	service := v1.NewClustersServiceClient(conn)
	resp, err := service.GetKernelSupportAvailable(ctx, &v1.Empty{})

	// Central in CI is deployed in online mode, hence the expectation is
	// that kernel support is available via the HTTP download site.
	assert.NoError(t, err)
	assert.True(t, resp.KernelSupportAvailable)
}
