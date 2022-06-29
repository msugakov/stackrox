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

type TestGrandChild1StoreSuite struct {
	suite.Suite
	envIsolator *envisolator.EnvIsolator
	store       Store
	testDB      *pgtest.TestPostgres
}

func TestTestGrandChild1Store(t *testing.T) {
	suite.Run(t, new(TestGrandChild1StoreSuite))
}

func (s *TestGrandChild1StoreSuite) SetupSuite() {
	s.envIsolator = envisolator.NewEnvIsolator(s.T())
	s.envIsolator.Setenv(features.PostgresDatastore.EnvVar(), "true")

	if !features.PostgresDatastore.Enabled() {
		s.T().Skip("Skip postgres store tests")
		s.T().SkipNow()
	}

	s.testDB = pgtest.ForT(s.T())
	s.store = New(s.testDB.Pool)
}

func (s *TestGrandChild1StoreSuite) SetupTest() {
	ctx := sac.WithAllAccess(context.Background())
	tag, err := s.testDB.Exec(ctx, "TRUNCATE test_grand_child1 CASCADE")
	s.T().Log("test_grand_child1", tag)
	s.NoError(err)
}

func (s *TestGrandChild1StoreSuite) TearDownSuite() {
	s.testDB.Teardown(s.T())
	s.envIsolator.RestoreAll()
}

func (s *TestGrandChild1StoreSuite) TestStore() {
	ctx := sac.WithAllAccess(context.Background())

	store := s.store

	testGrandChild1 := &storage.TestGrandChild1{}
	s.NoError(testutils.FullInit(testGrandChild1, testutils.SimpleInitializer(), testutils.JSONFieldsFilter))

	foundTestGrandChild1, exists, err := store.Get(ctx, testGrandChild1.GetId())
	s.NoError(err)
	s.False(exists)
	s.Nil(foundTestGrandChild1)

}
