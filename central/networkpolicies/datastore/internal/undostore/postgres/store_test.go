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

type NetworkpolicyapplicationundorecordsStoreSuite struct {
	suite.Suite
	envIsolator *envisolator.EnvIsolator
	store       Store
	pool        *pgxpool.Pool
}

func TestNetworkpolicyapplicationundorecordsStore(t *testing.T) {
	suite.Run(t, new(NetworkpolicyapplicationundorecordsStoreSuite))
}

func (s *NetworkpolicyapplicationundorecordsStoreSuite) SetupSuite() {
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
	s.store = CreateTableAndNewStore(ctx, pool, gormDB)
}

func (s *NetworkpolicyapplicationundorecordsStoreSuite) SetupTest() {
	ctx := sac.WithAllAccess(context.Background())
	tag, err := s.pool.Exec(ctx, "TRUNCATE networkpolicyapplicationundorecords CASCADE")
	s.T().Log("networkpolicyapplicationundorecords", tag)
	s.NoError(err)
}

func (s *NetworkpolicyapplicationundorecordsStoreSuite) TearDownSuite() {
	if s.pool != nil {
		s.pool.Close()
	}
	s.envIsolator.RestoreAll()
}

func (s *NetworkpolicyapplicationundorecordsStoreSuite) TestStore() {
	ctx := sac.WithAllAccess(context.Background())

	store := s.store

	networkPolicyApplicationUndoRecord := &storage.NetworkPolicyApplicationUndoRecord{}
	s.NoError(testutils.FullInit(networkPolicyApplicationUndoRecord, testutils.SimpleInitializer(), testutils.JSONFieldsFilter))

	foundNetworkPolicyApplicationUndoRecord, exists, err := store.Get(ctx, networkPolicyApplicationUndoRecord.GetClusterId())
	s.NoError(err)
	s.False(exists)
	s.Nil(foundNetworkPolicyApplicationUndoRecord)

	withNoAccessCtx := sac.WithNoAccess(ctx)

	s.NoError(store.Upsert(ctx, networkPolicyApplicationUndoRecord))
	foundNetworkPolicyApplicationUndoRecord, exists, err = store.Get(ctx, networkPolicyApplicationUndoRecord.GetClusterId())
	s.NoError(err)
	s.True(exists)
	s.Equal(networkPolicyApplicationUndoRecord, foundNetworkPolicyApplicationUndoRecord)

	networkPolicyApplicationUndoRecordCount, err := store.Count(ctx)
	s.NoError(err)
	s.Equal(1, networkPolicyApplicationUndoRecordCount)
	networkPolicyApplicationUndoRecordCount, err = store.Count(withNoAccessCtx)
	s.NoError(err)
	s.Zero(networkPolicyApplicationUndoRecordCount)

	networkPolicyApplicationUndoRecordExists, err := store.Exists(ctx, networkPolicyApplicationUndoRecord.GetClusterId())
	s.NoError(err)
	s.True(networkPolicyApplicationUndoRecordExists)
	s.NoError(store.Upsert(ctx, networkPolicyApplicationUndoRecord))
	s.ErrorIs(store.Upsert(withNoAccessCtx, networkPolicyApplicationUndoRecord), sac.ErrResourceAccessDenied)

	foundNetworkPolicyApplicationUndoRecord, exists, err = store.Get(ctx, networkPolicyApplicationUndoRecord.GetClusterId())
	s.NoError(err)
	s.True(exists)
	s.Equal(networkPolicyApplicationUndoRecord, foundNetworkPolicyApplicationUndoRecord)

	s.NoError(store.Delete(ctx, networkPolicyApplicationUndoRecord.GetClusterId()))
	foundNetworkPolicyApplicationUndoRecord, exists, err = store.Get(ctx, networkPolicyApplicationUndoRecord.GetClusterId())
	s.NoError(err)
	s.False(exists)
	s.Nil(foundNetworkPolicyApplicationUndoRecord)
	s.NoError(store.Delete(withNoAccessCtx, networkPolicyApplicationUndoRecord.GetClusterId()))

	var networkPolicyApplicationUndoRecords []*storage.NetworkPolicyApplicationUndoRecord
	for i := 0; i < 200; i++ {
		networkPolicyApplicationUndoRecord := &storage.NetworkPolicyApplicationUndoRecord{}
		s.NoError(testutils.FullInit(networkPolicyApplicationUndoRecord, testutils.UniqueInitializer(), testutils.JSONFieldsFilter))
		networkPolicyApplicationUndoRecords = append(networkPolicyApplicationUndoRecords, networkPolicyApplicationUndoRecord)
	}

	s.NoError(store.UpsertMany(ctx, networkPolicyApplicationUndoRecords))

	networkPolicyApplicationUndoRecordCount, err = store.Count(ctx)
	s.NoError(err)
	s.Equal(200, networkPolicyApplicationUndoRecordCount)
}
