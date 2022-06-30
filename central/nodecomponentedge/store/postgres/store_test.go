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

type NodeComponentEdgesStoreSuite struct {
	suite.Suite
	envIsolator *envisolator.EnvIsolator
	store       Store
	testDB      *pgtest.TestPostgres
}

func TestNodeComponentEdgesStore(t *testing.T) {
	suite.Run(t, new(NodeComponentEdgesStoreSuite))
}

func (s *NodeComponentEdgesStoreSuite) SetupSuite() {
	s.envIsolator = envisolator.NewEnvIsolator(s.T())
	s.envIsolator.Setenv(features.PostgresDatastore.EnvVar(), "true")

	if !features.PostgresDatastore.Enabled() {
		s.T().Skip("Skip postgres store tests")
		s.T().SkipNow()
	}

	s.testDB = pgtest.ForT(s.T())
	s.store = New(s.testDB.Pool)
}

func (s *NodeComponentEdgesStoreSuite) SetupTest() {
	ctx := sac.WithAllAccess(context.Background())
	tag, err := s.testDB.Exec(ctx, "TRUNCATE node_component_edges CASCADE")
	s.T().Log("node_component_edges", tag)
	s.NoError(err)
}

func (s *NodeComponentEdgesStoreSuite) TearDownSuite() {
	s.testDB.Teardown(s.T())
	s.envIsolator.RestoreAll()
}

func (s *NodeComponentEdgesStoreSuite) TestStore() {
	ctx := sac.WithAllAccess(context.Background())

	store := s.store

	nodeComponentEdge := &storage.NodeComponentEdge{}
	s.NoError(testutils.FullInit(nodeComponentEdge, testutils.SimpleInitializer(), testutils.JSONFieldsFilter))

	foundNodeComponentEdge, exists, err := store.Get(ctx, nodeComponentEdge.GetId())
	s.NoError(err)
	s.False(exists)
	s.Nil(foundNodeComponentEdge)

}
