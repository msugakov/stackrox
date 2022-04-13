package csv

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stackrox/stackrox/central/audit"
	clusterMocks "github.com/stackrox/stackrox/central/cluster/datastore/mocks"
	cveMocks "github.com/stackrox/stackrox/central/cve/datastore/mocks"
	deploymentMocks "github.com/stackrox/stackrox/central/deployment/datastore/mocks"
	"github.com/stackrox/stackrox/central/graphql/resolvers"
	imageMocks "github.com/stackrox/stackrox/central/image/datastore/mocks"
	componentMocks "github.com/stackrox/stackrox/central/imagecomponent/datastore/mocks"
	nsMocks "github.com/stackrox/stackrox/central/namespace/datastore/mocks"
	nodeMocks "github.com/stackrox/stackrox/central/node/globaldatastore/mocks"
	notifierMocks "github.com/stackrox/stackrox/central/notifier/processor/mocks"
	v1 "github.com/stackrox/stackrox/generated/api/v1"
	"github.com/stackrox/stackrox/pkg/sac"
	"github.com/stackrox/stackrox/pkg/search"
	"github.com/stackrox/stackrox/pkg/search/scoped"
	"github.com/stretchr/testify/suite"
)

func TestCVEScoping(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(CVEScopingTestSuite))
}

type CVEScopingTestSuite struct {
	suite.Suite
	ctx                 context.Context
	mockCtrl            *gomock.Controller
	clusterDataStore    *clusterMocks.MockDataStore
	nsDataStore         *nsMocks.MockDataStore
	deploymentDataStore *deploymentMocks.MockDataStore
	imageDataStore      *imageMocks.MockDataStore
	nodeDataStore       *nodeMocks.MockGlobalDataStore
	componentDataStore  *componentMocks.MockDataStore
	cveDataStore        *cveMocks.MockDataStore
	resolver            *resolvers.Resolver
	handler             *handlerImpl
}

func (suite *CVEScopingTestSuite) SetupTest() {
	suite.mockCtrl = gomock.NewController(suite.T())
	suite.clusterDataStore = clusterMocks.NewMockDataStore(suite.mockCtrl)
	suite.nsDataStore = nsMocks.NewMockDataStore(suite.mockCtrl)
	suite.deploymentDataStore = deploymentMocks.NewMockDataStore(suite.mockCtrl)
	suite.imageDataStore = imageMocks.NewMockDataStore(suite.mockCtrl)
	suite.nodeDataStore = nodeMocks.NewMockGlobalDataStore(suite.mockCtrl)
	suite.componentDataStore = componentMocks.NewMockDataStore(suite.mockCtrl)
	suite.cveDataStore = cveMocks.NewMockDataStore(suite.mockCtrl)
	notifierMock := notifierMocks.NewMockProcessor(suite.mockCtrl)

	notifierMock.EXPECT().HasEnabledAuditNotifiers().Return(false).AnyTimes()

	suite.resolver = &resolvers.Resolver{
		ClusterDataStore:        suite.clusterDataStore,
		NamespaceDataStore:      suite.nsDataStore,
		DeploymentDataStore:     suite.deploymentDataStore,
		ImageDataStore:          suite.imageDataStore,
		NodeGlobalDataStore:     suite.nodeDataStore,
		ImageComponentDataStore: suite.componentDataStore,
		CVEDataStore:            suite.cveDataStore,
		AuditLogger:             audit.New(notifierMock),
	}

	suite.handler = newHandler(suite.resolver)

	suite.ctx = sac.WithGlobalAccessScopeChecker(context.Background(), sac.AllowAllAccessScopeChecker())
}

func (suite *CVEScopingTestSuite) TearDownTest() {
	suite.mockCtrl.Finish()
}

func (suite *CVEScopingTestSuite) TestSingleResourceQuery() {
	imgSha := "img1"

	query := search.ConjunctionQuery(
		search.NewQueryBuilder().AddStrings(search.ImageSHA, imgSha).ProtoQuery(),
		search.NewQueryBuilder().AddBools(search.Fixable, true).ProtoQuery())

	suite.imageDataStore.EXPECT().
		Search(suite.ctx, search.NewQueryBuilder().AddStrings(search.ImageSHA, imgSha).ProtoQuery()).
		Return([]search.Result{{ID: imgSha}}, nil)

	expected := scoped.Context(suite.ctx, scoped.Scope{
		Level: v1.SearchCategory_IMAGES,
		ID:    imgSha,
	})
	actual, err := suite.handler.getScopeContext(suite.ctx, query)
	suite.NoError(err)
	suite.Equal(expected, actual)
}

func (suite *CVEScopingTestSuite) TestMultipleResourceQuery() {
	imgSha := "img1"

	query := search.ConjunctionQuery(
		search.NewQueryBuilder().AddStrings(search.DeploymentName, "dep").ProtoQuery(),
		search.NewQueryBuilder().AddStrings(search.ImageSHA, imgSha).ProtoQuery(),
		search.NewQueryBuilder().AddBools(search.Fixable, true).ProtoQuery())

	suite.imageDataStore.EXPECT().
		Search(suite.ctx, search.NewQueryBuilder().AddStrings(search.ImageSHA, imgSha).ProtoQuery()).
		Return([]search.Result{{ID: imgSha}}, nil)

	expected := scoped.Context(suite.ctx, scoped.Scope{
		Level: v1.SearchCategory_IMAGES,
		ID:    imgSha,
	})
	// Lowest resource scope should be applied.
	actual, err := suite.handler.getScopeContext(suite.ctx, query)
	suite.NoError(err)
	suite.Equal(expected, actual)
}

func (suite *CVEScopingTestSuite) TestMultipleMatchesQuery() {
	img := "img"

	query := search.ConjunctionQuery(
		search.NewQueryBuilder().AddStrings(search.ImageName, img).ProtoQuery(),
		search.NewQueryBuilder().AddBools(search.Fixable, true).ProtoQuery())

	suite.imageDataStore.EXPECT().
		Search(suite.ctx, search.NewQueryBuilder().AddStrings(search.ImageName, img).ProtoQuery()).
		Return([]search.Result{{ID: "img1"}, {ID: "img2"}}, nil)

	// No scope should be applied.
	actual, err := suite.handler.getScopeContext(suite.ctx, query)
	suite.NoError(err)
	suite.Equal(suite.ctx, actual)
}

func (suite *CVEScopingTestSuite) TestNoReScope() {
	img := "img"

	query := search.ConjunctionQuery(
		search.NewQueryBuilder().AddStrings(search.ImageName, img).ProtoQuery(),
		search.NewQueryBuilder().AddBools(search.Fixable, true).ProtoQuery())

	expected := scoped.Context(suite.ctx, scoped.Scope{
		Level: v1.SearchCategory_DEPLOYMENTS,
		ID:    "dep",
	})
	actual, err := suite.handler.getScopeContext(expected, query)
	suite.NoError(err)
	suite.Equal(expected, actual)
}
