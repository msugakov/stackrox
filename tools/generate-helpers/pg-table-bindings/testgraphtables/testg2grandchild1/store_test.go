// Code generated by pg-bindings generator. DO NOT EDIT.

//go:build sql_integration

package postgres

import (
	"context"
	"testing"

	"github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/pkg/features"
	"github.com/stackrox/rox/pkg/postgres/pgtest"
	"github.com/stackrox/rox/pkg/sac"
	"github.com/stackrox/rox/pkg/testutils"
	"github.com/stackrox/rox/pkg/testutils/envisolator"
	"github.com/stretchr/testify/suite"
)

type TestG2GrandChild1StoreSuite struct {
	suite.Suite
	envIsolator *envisolator.EnvIsolator
	store       Store
	testDB      *pgtest.TestPostgres
}

func TestTestG2GrandChild1Store(t *testing.T) {
	suite.Run(t, new(TestG2GrandChild1StoreSuite))
}

func (s *TestG2GrandChild1StoreSuite) SetupSuite() {
	s.envIsolator = envisolator.NewEnvIsolator(s.T())
	s.envIsolator.Setenv(features.PostgresDatastore.EnvVar(), "true")

	if !features.PostgresDatastore.Enabled() {
		s.T().Skip("Skip postgres store tests")
		s.T().SkipNow()
	}

	s.testDB = pgtest.ForT(s.T())
	s.store = New(s.testDB.Pool)
}

func (s *TestG2GrandChild1StoreSuite) SetupTest() {
	ctx := sac.WithAllAccess(context.Background())
	tag, err := s.testDB.Exec(ctx, "TRUNCATE test_g2_grand_child1 CASCADE")
	s.T().Log("test_g2_grand_child1", tag)
	s.NoError(err)
}

func (s *TestG2GrandChild1StoreSuite) TearDownSuite() {
	s.testDB.Teardown(s.T())
	s.envIsolator.RestoreAll()
}

func (s *TestG2GrandChild1StoreSuite) TestStore() {
	ctx := sac.WithAllAccess(context.Background())

	store := s.store

	testG2GrandChild1 := &storage.TestG2GrandChild1{}
	s.NoError(testutils.FullInit(testG2GrandChild1, testutils.SimpleInitializer(), testutils.JSONFieldsFilter))

	foundTestG2GrandChild1, exists, err := store.Get(ctx, testG2GrandChild1.GetId())
	s.NoError(err)
	s.False(exists)
	s.Nil(foundTestG2GrandChild1)

}
