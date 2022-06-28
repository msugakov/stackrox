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

type ImageCveEdgesStoreSuite struct {
	suite.Suite
	envIsolator *envisolator.EnvIsolator
	store       Store
	testDB      *pgtest.TestPostgres
}

func TestImageCveEdgesStore(t *testing.T) {
	suite.Run(t, new(ImageCveEdgesStoreSuite))
}

func (s *ImageCveEdgesStoreSuite) SetupSuite() {
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

func (s *ImageCveEdgesStoreSuite) SetupTest() {
	ctx := sac.WithAllAccess(context.Background())
	tag, err := s.testDB.Exec(ctx, "TRUNCATE image_cve_edges CASCADE")
	s.T().Log("image_cve_edges", tag)
	s.NoError(err)
}

func (s *ImageCveEdgesStoreSuite) TearDownSuite() {
	s.testDB.Teardown(s.T())
	s.envIsolator.RestoreAll()
}

func (s *ImageCveEdgesStoreSuite) TestStore() {
	ctx := sac.WithAllAccess(context.Background())

	store := s.store

	imageCVEEdge := &storage.ImageCVEEdge{}
	s.NoError(testutils.FullInit(imageCVEEdge, testutils.SimpleInitializer(), testutils.JSONFieldsFilter))

	foundImageCVEEdge, exists, err := store.Get(ctx, imageCVEEdge.GetId())
	s.NoError(err)
	s.False(exists)
	s.Nil(foundImageCVEEdge)

}
