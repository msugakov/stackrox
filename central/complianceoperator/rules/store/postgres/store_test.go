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

type ComplianceOperatorRulesStoreSuite struct {
	suite.Suite
	envIsolator *envisolator.EnvIsolator
	store       Store
	testDB      *pgtest.TestPostgres
}

func TestComplianceOperatorRulesStore(t *testing.T) {
	suite.Run(t, new(ComplianceOperatorRulesStoreSuite))
}

func (s *ComplianceOperatorRulesStoreSuite) SetupSuite() {
	s.envIsolator = envisolator.NewEnvIsolator(s.T())
	s.envIsolator.Setenv(features.PostgresDatastore.EnvVar(), "true")

	if !features.PostgresDatastore.Enabled() {
		s.T().Skip("Skip postgres store tests")
		s.T().SkipNow()
	}

	s.testDB = pgtest.ForT(s.T())
	s.store = New(s.testDB.Pool)
}

func (s *ComplianceOperatorRulesStoreSuite) SetupTest() {
	ctx := sac.WithAllAccess(context.Background())
	tag, err := s.testDB.Exec(ctx, "TRUNCATE compliance_operator_rules CASCADE")
	s.T().Log("compliance_operator_rules", tag)
	s.NoError(err)
}

func (s *ComplianceOperatorRulesStoreSuite) TearDownSuite() {
	s.testDB.Teardown(s.T())
	s.envIsolator.RestoreAll()
}

func (s *ComplianceOperatorRulesStoreSuite) TestStore() {
	ctx := sac.WithAllAccess(context.Background())

	store := s.store

	complianceOperatorRule := &storage.ComplianceOperatorRule{}
	s.NoError(testutils.FullInit(complianceOperatorRule, testutils.SimpleInitializer(), testutils.JSONFieldsFilter))

	foundComplianceOperatorRule, exists, err := store.Get(ctx, complianceOperatorRule.GetId())
	s.NoError(err)
	s.False(exists)
	s.Nil(foundComplianceOperatorRule)

	withNoAccessCtx := sac.WithNoAccess(ctx)

	s.NoError(store.Upsert(ctx, complianceOperatorRule))
	foundComplianceOperatorRule, exists, err = store.Get(ctx, complianceOperatorRule.GetId())
	s.NoError(err)
	s.True(exists)
	s.Equal(complianceOperatorRule, foundComplianceOperatorRule)

	complianceOperatorRuleCount, err := store.Count(ctx)
	s.NoError(err)
	s.Equal(1, complianceOperatorRuleCount)
	complianceOperatorRuleCount, err = store.Count(withNoAccessCtx)
	s.NoError(err)
	s.Zero(complianceOperatorRuleCount)

	complianceOperatorRuleExists, err := store.Exists(ctx, complianceOperatorRule.GetId())
	s.NoError(err)
	s.True(complianceOperatorRuleExists)
	s.NoError(store.Upsert(ctx, complianceOperatorRule))
	s.ErrorIs(store.Upsert(withNoAccessCtx, complianceOperatorRule), sac.ErrResourceAccessDenied)

	foundComplianceOperatorRule, exists, err = store.Get(ctx, complianceOperatorRule.GetId())
	s.NoError(err)
	s.True(exists)
	s.Equal(complianceOperatorRule, foundComplianceOperatorRule)

	s.NoError(store.Delete(ctx, complianceOperatorRule.GetId()))
	foundComplianceOperatorRule, exists, err = store.Get(ctx, complianceOperatorRule.GetId())
	s.NoError(err)
	s.False(exists)
	s.Nil(foundComplianceOperatorRule)
	s.ErrorIs(store.Delete(withNoAccessCtx, complianceOperatorRule.GetId()), sac.ErrResourceAccessDenied)

	var complianceOperatorRules []*storage.ComplianceOperatorRule
	for i := 0; i < 200; i++ {
		complianceOperatorRule := &storage.ComplianceOperatorRule{}
		s.NoError(testutils.FullInit(complianceOperatorRule, testutils.UniqueInitializer(), testutils.JSONFieldsFilter))
		complianceOperatorRules = append(complianceOperatorRules, complianceOperatorRule)
	}

	s.NoError(store.UpsertMany(ctx, complianceOperatorRules))

	complianceOperatorRuleCount, err = store.Count(ctx)
	s.NoError(err)
	s.Equal(200, complianceOperatorRuleCount)
}
