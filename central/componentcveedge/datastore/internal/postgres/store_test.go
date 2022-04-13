// Code generated by pg-bindings generator. DO NOT EDIT.

//go:build sql_integration

package postgres

import (
	"context"
	"testing"

	"github.com/jackc/pgx/v4/pgxpool"
	storage "github.com/stackrox/stackrox/generated/storage"
	"github.com/stackrox/stackrox/pkg/features"
	"github.com/stackrox/stackrox/pkg/postgres/pgtest"
	"github.com/stackrox/stackrox/pkg/testutils"
	"github.com/stackrox/stackrox/pkg/testutils/envisolator"
	"github.com/stretchr/testify/suite"
)

type ImageComponentCveRelationsStoreSuite struct {
	suite.Suite
	envIsolator *envisolator.EnvIsolator
}

func TestImageComponentCveRelationsStore(t *testing.T) {
	suite.Run(t, new(ImageComponentCveRelationsStoreSuite))
}

func (s *ImageComponentCveRelationsStoreSuite) SetupTest() {
	s.envIsolator = envisolator.NewEnvIsolator(s.T())
	s.envIsolator.Setenv(features.PostgresDatastore.EnvVar(), "true")

	if !features.PostgresDatastore.Enabled() {
		s.T().Skip("Skip postgres store tests")
		s.T().SkipNow()
	}
}

func (s *ImageComponentCveRelationsStoreSuite) TearDownTest() {
	s.envIsolator.RestoreAll()
}

func (s *ImageComponentCveRelationsStoreSuite) TestStore() {
	ctx := context.Background()

	source := pgtest.GetConnectionString(s.T())
	config, err := pgxpool.ParseConfig(source)
	s.Require().NoError(err)
	pool, err := pgxpool.ConnectConfig(ctx, config)
	s.NoError(err)
	defer pool.Close()

	Destroy(ctx, pool)
	store := New(ctx, pool)

	componentCVEEdge := &storage.ComponentCVEEdge{}
	s.NoError(testutils.FullInit(componentCVEEdge, testutils.SimpleInitializer(), testutils.JSONFieldsFilter))

	foundComponentCVEEdge, exists, err := store.Get(ctx, componentCVEEdge.GetId(), componentCVEEdge.GetImageComponentId(), componentCVEEdge.GetCveId())
	s.NoError(err)
	s.False(exists)
	s.Nil(foundComponentCVEEdge)

}
