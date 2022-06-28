// Code generated by pg-bindings generator. DO NOT EDIT.

//go:build sql_integration

package postgres

import (
	"context"
	"testing"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/pkg/features"
	"github.com/stackrox/rox/pkg/postgres/pgtest"
	"github.com/stackrox/rox/pkg/sac"
	"github.com/stackrox/rox/pkg/testutils"
	"github.com/stackrox/rox/pkg/testutils/envisolator"
	"github.com/stretchr/testify/suite"
)

type NetworkGraphConfigsStoreSuite struct {
	suite.Suite
	envIsolator *envisolator.EnvIsolator
	store       Store
	testDB      *pgtest.TestPostgres
}

func TestNetworkGraphConfigsStore(t *testing.T) {
	suite.Run(t, new(NetworkGraphConfigsStoreSuite))
}

func (s *NetworkGraphConfigsStoreSuite) SetupSuite() {
	s.envIsolator = envisolator.NewEnvIsolator(s.T())
	s.envIsolator.Setenv(features.PostgresDatastore.EnvVar(), "true")

	if !features.PostgresDatastore.Enabled() {
		s.T().Skip("Skip postgres store tests")
		s.T().SkipNow()
	}

	ctx := sac.WithAllAccess(context.Background())

	source := pgtest.GetConnectionString(s.T())
	config, err := pgxpool.ParseConfig(source)
	s.Require().NoError(err)
	pool, err := pgxpool.ConnectConfig(ctx, config)
	s.Require().NoError(err)

	Destroy(ctx, pool)

	s.pool = pool
	gormDB := pgtest.OpenGormDB(s.T(), source, false)
	defer pgtest.CloseGormDB(s.T(), gormDB)
	s.testDB = pgtest.ForT(s.T())
	s.store = New(s.testDB.Pool)
}

func (s *NetworkGraphConfigsStoreSuite) SetupTest() {
	ctx := sac.WithAllAccess(context.Background())
	tag, err := s.testDB.Exec(ctx, "TRUNCATE network_graph_configs CASCADE")
	s.T().Log("network_graph_configs", tag)
	s.NoError(err)
}

func (s *NetworkGraphConfigsStoreSuite) TearDownSuite() {
	s.testDB.Teardown(s.T())
	s.envIsolator.RestoreAll()
}

func (s *NetworkGraphConfigsStoreSuite) TestStore() {
	ctx := sac.WithAllAccess(context.Background())

	store := s.store

	networkGraphConfig := &storage.NetworkGraphConfig{}
	s.NoError(testutils.FullInit(networkGraphConfig, testutils.SimpleInitializer(), testutils.JSONFieldsFilter))

	foundNetworkGraphConfig, exists, err := store.Get(ctx, networkGraphConfig.GetId())
	s.NoError(err)
	s.False(exists)
	s.Nil(foundNetworkGraphConfig)

	withNoAccessCtx := sac.WithNoAccess(ctx)

	s.NoError(store.Upsert(ctx, networkGraphConfig))
	foundNetworkGraphConfig, exists, err = store.Get(ctx, networkGraphConfig.GetId())
	s.NoError(err)
	s.True(exists)
	s.Equal(networkGraphConfig, foundNetworkGraphConfig)

	networkGraphConfigCount, err := store.Count(ctx)
	s.NoError(err)
	s.Equal(1, networkGraphConfigCount)
	networkGraphConfigCount, err = store.Count(withNoAccessCtx)
	s.NoError(err)
	s.Zero(networkGraphConfigCount)

	networkGraphConfigExists, err := store.Exists(ctx, networkGraphConfig.GetId())
	s.NoError(err)
	s.True(networkGraphConfigExists)
	s.NoError(store.Upsert(ctx, networkGraphConfig))
	s.ErrorIs(store.Upsert(withNoAccessCtx, networkGraphConfig), sac.ErrResourceAccessDenied)

	foundNetworkGraphConfig, exists, err = store.Get(ctx, networkGraphConfig.GetId())
	s.NoError(err)
	s.True(exists)
	s.Equal(networkGraphConfig, foundNetworkGraphConfig)

	s.NoError(store.Delete(ctx, networkGraphConfig.GetId()))
	foundNetworkGraphConfig, exists, err = store.Get(ctx, networkGraphConfig.GetId())
	s.NoError(err)
	s.False(exists)
	s.Nil(foundNetworkGraphConfig)
	s.ErrorIs(store.Delete(withNoAccessCtx, networkGraphConfig.GetId()), sac.ErrResourceAccessDenied)

	var networkGraphConfigs []*storage.NetworkGraphConfig
	for i := 0; i < 200; i++ {
		networkGraphConfig := &storage.NetworkGraphConfig{}
		s.NoError(testutils.FullInit(networkGraphConfig, testutils.UniqueInitializer(), testutils.JSONFieldsFilter))
		networkGraphConfigs = append(networkGraphConfigs, networkGraphConfig)
	}

	s.NoError(store.UpsertMany(ctx, networkGraphConfigs))

	networkGraphConfigCount, err = store.Count(ctx)
	s.NoError(err)
	s.Equal(200, networkGraphConfigCount)
}
