// Code generated by pg-bindings generator. DO NOT EDIT.

//go:build sql_integration

package postgres

import (
	"context"
	"testing"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/stackrox/stackrox/generated/storage"
	"github.com/stackrox/stackrox/pkg/features"
	"github.com/stackrox/stackrox/pkg/postgres/pgtest"
	"github.com/stackrox/stackrox/pkg/sac"
	"github.com/stackrox/stackrox/pkg/testutils"
	"github.com/stackrox/stackrox/pkg/testutils/envisolator"
	"github.com/stretchr/testify/suite"
)

type IntegrationhealthStoreSuite struct {
	suite.Suite
	envIsolator *envisolator.EnvIsolator
}

func TestIntegrationhealthStore(t *testing.T) {
	suite.Run(t, new(IntegrationhealthStoreSuite))
}

func (s *IntegrationhealthStoreSuite) SetupTest() {
	s.envIsolator = envisolator.NewEnvIsolator(s.T())
	s.envIsolator.Setenv(features.PostgresDatastore.EnvVar(), "true")

	if !features.PostgresDatastore.Enabled() {
		s.T().Skip("Skip postgres store tests")
		s.T().SkipNow()
	}
}

func (s *IntegrationhealthStoreSuite) TearDownTest() {
	s.envIsolator.RestoreAll()
}

func (s *IntegrationhealthStoreSuite) TestStore() {
	ctx := sac.WithAllAccess(context.Background())

	source := pgtest.GetConnectionString(s.T())
	config, err := pgxpool.ParseConfig(source)
	s.Require().NoError(err)
	pool, err := pgxpool.ConnectConfig(ctx, config)
	s.NoError(err)
	defer pool.Close()

	Destroy(ctx, pool)
	store := New(ctx, pool)

	integrationHealth := &storage.IntegrationHealth{}
	s.NoError(testutils.FullInit(integrationHealth, testutils.SimpleInitializer(), testutils.JSONFieldsFilter))

	foundIntegrationHealth, exists, err := store.Get(ctx, integrationHealth.GetId())
	s.NoError(err)
	s.False(exists)
	s.Nil(foundIntegrationHealth)

	withNoAccessCtx := sac.WithNoAccess(ctx)

	s.NoError(store.Upsert(ctx, integrationHealth))
	foundIntegrationHealth, exists, err = store.Get(ctx, integrationHealth.GetId())
	s.NoError(err)
	s.True(exists)
	s.Equal(integrationHealth, foundIntegrationHealth)

	integrationHealthCount, err := store.Count(ctx)
	s.NoError(err)
	s.Equal(integrationHealthCount, 1)
	integrationHealthCount, err = store.Count(withNoAccessCtx)
	s.NoError(err)
	s.Zero(integrationHealthCount)

	integrationHealthExists, err := store.Exists(ctx, integrationHealth.GetId())
	s.NoError(err)
	s.True(integrationHealthExists)
	s.NoError(store.Upsert(ctx, integrationHealth))
	s.ErrorIs(store.Upsert(withNoAccessCtx, integrationHealth), sac.ErrResourceAccessDenied)

	foundIntegrationHealth, exists, err = store.Get(ctx, integrationHealth.GetId())
	s.NoError(err)
	s.True(exists)
	s.Equal(integrationHealth, foundIntegrationHealth)

	s.NoError(store.Delete(ctx, integrationHealth.GetId()))
	foundIntegrationHealth, exists, err = store.Get(ctx, integrationHealth.GetId())
	s.NoError(err)
	s.False(exists)
	s.Nil(foundIntegrationHealth)
	s.ErrorIs(store.Delete(withNoAccessCtx, integrationHealth.GetId()), sac.ErrResourceAccessDenied)

	var integrationHealths []*storage.IntegrationHealth
	for i := 0; i < 200; i++ {
		integrationHealth := &storage.IntegrationHealth{}
		s.NoError(testutils.FullInit(integrationHealth, testutils.UniqueInitializer(), testutils.JSONFieldsFilter))
		integrationHealths = append(integrationHealths, integrationHealth)
	}

	s.NoError(store.UpsertMany(ctx, integrationHealths))

	integrationHealthCount, err = store.Count(ctx)
	s.NoError(err)
	s.Equal(integrationHealthCount, 200)
}
