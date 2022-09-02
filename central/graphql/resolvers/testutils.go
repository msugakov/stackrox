package resolvers

import (
	"context"
	"fmt"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/graph-gophers/graphql-go"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/pkg/errors"
	imageComponentCVEEdgeDS "github.com/stackrox/rox/central/componentcveedge/datastore"
	imageComponentCVEEdgePostgres "github.com/stackrox/rox/central/componentcveedge/datastore/store/postgres"
	imageComponentCVEEdgeSearch "github.com/stackrox/rox/central/componentcveedge/search"
	imageCVEDS "github.com/stackrox/rox/central/cve/image/datastore"
	imageCVESearch "github.com/stackrox/rox/central/cve/image/datastore/search"
	imageCVEPostgres "github.com/stackrox/rox/central/cve/image/datastore/store/postgres"
	nodeCVEDS "github.com/stackrox/rox/central/cve/node/datastore"
	nodeCVESearch "github.com/stackrox/rox/central/cve/node/datastore/search"
	nodeCVEPostgres "github.com/stackrox/rox/central/cve/node/datastore/store/postgres"
	"github.com/stackrox/rox/central/graphql/resolvers/loaders"
	imageDS "github.com/stackrox/rox/central/image/datastore"
	imagePostgres "github.com/stackrox/rox/central/image/datastore/store/postgres"
	imageComponentDS "github.com/stackrox/rox/central/imagecomponent/datastore"
	imageComponentPostgres "github.com/stackrox/rox/central/imagecomponent/datastore/store/postgres"
	imageComponentSearch "github.com/stackrox/rox/central/imagecomponent/search"
	nodeDS "github.com/stackrox/rox/central/node/datastore/dackbox/datastore"
	nodeGlobalDS "github.com/stackrox/rox/central/node/datastore/dackbox/globaldatastore"
	nodeSearch "github.com/stackrox/rox/central/node/datastore/search"
	nodePostgres "github.com/stackrox/rox/central/node/datastore/store/postgres"
	nodeComponentDS "github.com/stackrox/rox/central/nodecomponent/datastore"
	nodeComponentSearch "github.com/stackrox/rox/central/nodecomponent/datastore/search"
	nodeComponentPostgres "github.com/stackrox/rox/central/nodecomponent/datastore/store/postgres"
	nodeComponentCVEEdgeDS "github.com/stackrox/rox/central/nodecomponentcveedge/datastore"
	nodeComponentCVEEdgeSearch "github.com/stackrox/rox/central/nodecomponentcveedge/datastore/search"
	nodeComponentCVEEdgePostgres "github.com/stackrox/rox/central/nodecomponentcveedge/datastore/store/postgres"
	"github.com/stackrox/rox/central/ranking"
	mockRisks "github.com/stackrox/rox/central/risk/datastore/mocks"
	"github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/pkg/features"
	"github.com/stackrox/rox/pkg/fixtures"
	"github.com/stackrox/rox/pkg/grpc/authn"
	mockIdentity "github.com/stackrox/rox/pkg/grpc/authn/mocks"
	nodeConverter "github.com/stackrox/rox/pkg/nodes/converter"
	"github.com/stackrox/rox/pkg/postgres/pgtest"
	"github.com/stackrox/rox/pkg/sac"
	"github.com/stackrox/rox/pkg/utils"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func setupPostgresConn(t testing.TB) (*pgxpool.Pool, *gorm.DB) {
	source := pgtest.GetConnectionString(t)
	config, err := pgxpool.ParseConfig(source)
	assert.NoError(t, err)

	pool, err := pgxpool.ConnectConfig(context.Background(), config)
	assert.NoError(t, err)

	gormDB := pgtest.OpenGormDB(t, source)

	return pool, gormDB
}

func setupResolverForImageGraphQLTestsWithPostgres(
	t testing.TB,
	imageDataStore imageDS.DataStore,
	imageComponentDataStore imageComponentDS.DataStore,
	cveDataStore imageCVEDS.DataStore,
	imageComponentCVEEdgeDatastore imageComponentCVEEdgeDS.DataStore,
) (*Resolver, *graphql.Schema) {
	// loaders used by graphql layer
	registerImageLoader(t, imageDataStore)
	registerImageComponentLoader(t, imageComponentDataStore)
	registerImageCveLoader(t, cveDataStore)

	resolver := &Resolver{
		ImageDataStore:            imageDataStore,
		ImageComponentDataStore:   imageComponentDataStore,
		ImageCVEDataStore:         cveDataStore,
		ComponentCVEEdgeDataStore: imageComponentCVEEdgeDatastore,
	}

	schema, err := graphql.ParseSchema(Schema(), resolver)
	assert.NoError(t, err)

	return resolver, schema
}

func createImageDatastoreForPostgres(_ testing.TB, ctrl *gomock.Controller, db *pgxpool.Pool, gormDB *gorm.DB) imageDS.DataStore {
	ctx := context.Background()
	imagePostgres.Destroy(ctx, db)

	return imageDS.NewWithPostgres(
		imagePostgres.CreateTableAndNewStore(ctx, db, gormDB, false),
		imagePostgres.NewIndexer(db),
		mockRisks.NewMockDataStore(ctrl),
		ranking.NewRanker(),
		ranking.NewRanker(),
	)
}

func createImageComponentDatastoreForPostgres(_ testing.TB, ctrl *gomock.Controller, db *pgxpool.Pool, gormDB *gorm.DB) imageComponentDS.DataStore {
	ctx := context.Background()
	imageComponentPostgres.Destroy(ctx, db)

	mockRisk := mockRisks.NewMockDataStore(ctrl)
	store := imageComponentPostgres.CreateTableAndNewStore(ctx, db, gormDB)
	indexer := imageComponentPostgres.NewIndexer(db)
	searcher := imageComponentSearch.NewV2(store, indexer)

	return imageComponentDS.New(
		nil, store, indexer, searcher, mockRisk, ranking.NewRanker(),
	)
}

func createImageCVEDatastoreForPostgres(t testing.TB, db *pgxpool.Pool, gormDB *gorm.DB) imageCVEDS.DataStore {
	ctx := context.Background()
	imageCVEPostgres.Destroy(ctx, db)

	store := imageCVEPostgres.CreateTableAndNewStore(ctx, db, gormDB)
	indexer := imageCVEPostgres.NewIndexer(db)
	searcher := imageCVESearch.New(store, indexer)
	datastore, err := imageCVEDS.New(store, indexer, searcher, nil)
	assert.NoError(t, err)

	return datastore
}

func createImageComponentCVEEdgeDatastoreForPostgres(_ testing.TB, db *pgxpool.Pool, gormDB *gorm.DB) imageComponentCVEEdgeDS.DataStore {
	ctx := context.Background()
	imageComponentCVEEdgePostgres.Destroy(ctx, db)

	store := imageComponentCVEEdgePostgres.CreateTableAndNewStore(ctx, db, gormDB)
	indexer := imageComponentCVEEdgePostgres.NewIndexer(db)
	searcher := imageComponentCVEEdgeSearch.NewV2(store, indexer)

	return imageComponentCVEEdgeDS.New(nil, store, indexer, searcher)
}

func registerImageLoader(_ testing.TB, ds imageDS.DataStore) {
	loaders.RegisterTypeFactory(reflect.TypeOf(storage.Image{}), func() interface{} {
		return loaders.NewImageLoader(ds)
	})
}

func registerImageComponentLoader(_ testing.TB, ds imageComponentDS.DataStore) {
	loaders.RegisterTypeFactory(reflect.TypeOf(storage.ImageComponent{}), func() interface{} {
		return loaders.NewComponentLoader(ds)
	})
}

func registerImageCveLoader(_ testing.TB, ds imageCVEDS.DataStore) {
	loaders.RegisterTypeFactory(reflect.TypeOf(storage.ImageCVE{}), func() interface{} {
		return loaders.NewImageCVELoader(ds)
	})
}

func getTestImages(imageCount int) []*storage.Image {
	images := make([]*storage.Image, 0, imageCount)
	for i := 0; i < imageCount; i++ {
		img := fixtures.GetImageWithUniqueComponents(100)
		id := fmt.Sprintf("%d", i)
		img.Id = id
		images = append(images, img)
	}
	return images
}

func contextWithImagePerm(t testing.TB, ctrl *gomock.Controller) context.Context {
	id := mockIdentity.NewMockIdentity(ctrl)
	id.EXPECT().Permissions().Return(map[string]storage.Access{"Image": storage.Access_READ_ACCESS}).AnyTimes()
	return authn.ContextWithIdentity(sac.WithAllAccess(loaders.WithLoaderContext(context.Background())), id, t)
}

func contextWithNodePerm(t testing.TB, ctrl *gomock.Controller) context.Context {
	id := mockIdentity.NewMockIdentity(ctrl)
	id.EXPECT().Permissions().Return(map[string]storage.Access{"Node": storage.Access_READ_ACCESS}).AnyTimes()
	return authn.ContextWithIdentity(sac.WithAllAccess(loaders.WithLoaderContext(context.Background())), id, t)
}

func setupResolverForNodeGraphQLTestsWithPostgres(
	t testing.TB,
	nodeDataStore nodeDS.DataStore,
	nodeComponentDataStore nodeComponentDS.DataStore,
	cveDataStore nodeCVEDS.DataStore,
	nodeComponentCVEEdgeDatastore nodeComponentCVEEdgeDS.DataStore,
) *graphql.Schema {
	// loaders used by graphql layer
	registerNodeLoader(t, nodeDataStore)
	registerNodeComponentLoader(t, nodeComponentDataStore)
	registerNodeCveLoader(t, cveDataStore)

	nodeGlobalDatastore, err := nodeGlobalDS.New(nodeDataStore)
	assert.NoError(t, err)
	resolver := &Resolver{
		NodeGlobalDataStore:           nodeGlobalDatastore,
		NodeComponentDataStore:        nodeComponentDataStore,
		NodeCVEDataStore:              cveDataStore,
		NodeComponentCVEEdgeDataStore: nodeComponentCVEEdgeDatastore,
	}

	schema, err := graphql.ParseSchema(Schema(), resolver)
	assert.NoError(t, err)

	return schema
}

func createNodeDatastoreForPostgres(_ testing.TB, ctrl *gomock.Controller, db *pgxpool.Pool, gormDB *gorm.DB) nodeDS.DataStore {
	ctx := context.Background()
	nodePostgres.Destroy(ctx, db)

	mockRisk := mockRisks.NewMockDataStore(ctrl)
	store := nodePostgres.CreateTableAndNewStore(ctx, db, gormDB, false)
	indexer := nodePostgres.NewIndexer(db)
	searcher := nodeSearch.NewV2(store, indexer)
	return nodeDS.NewWithPostgres(store, indexer, searcher, mockRisk, ranking.NewRanker(), ranking.NewRanker())
}

func createNodeComponentDatastoreForPostgres(_ testing.TB, ctrl *gomock.Controller, db *pgxpool.Pool, gormDB *gorm.DB) nodeComponentDS.DataStore {
	if !features.PostgresDatastore.Enabled() {
		utils.CrashOnError(errors.New("Cannot create datastore with postgres disabled"))
	}
	ctx := context.Background()
	nodeComponentPostgres.Destroy(ctx, db)

	mockRisk := mockRisks.NewMockDataStore(ctrl)
	store := nodeComponentPostgres.CreateTableAndNewStore(ctx, db, gormDB)
	indexer := nodeComponentPostgres.NewIndexer(db)
	searcher := nodeComponentSearch.New(store, indexer)

	return nodeComponentDS.New(store, indexer, searcher, mockRisk, ranking.NewRanker())
}

func createNodeCVEDatastoreForPostgres(t testing.TB, db *pgxpool.Pool, gormDB *gorm.DB) nodeCVEDS.DataStore {
	if !features.PostgresDatastore.Enabled() {
		utils.CrashOnError(errors.New("Cannot create datastore with postgres disabled"))
	}
	ctx := context.Background()
	nodeCVEPostgres.Destroy(ctx, db)

	store := nodeCVEPostgres.CreateTableAndNewStore(ctx, db, gormDB)
	indexer := nodeCVEPostgres.NewIndexer(db)
	searcher := nodeCVESearch.New(store, indexer)
	datastore, err := nodeCVEDS.New(store, indexer, searcher, nil)
	assert.NoError(t, err)

	return datastore
}

func NodeComponentCVEEdgeDatastoreForPostgres(_ testing.TB, db *pgxpool.Pool, gormDB *gorm.DB) nodeComponentCVEEdgeDS.DataStore {
	if !features.PostgresDatastore.Enabled() {
		utils.CrashOnError(errors.New("Cannot create datastore with postgres disabled"))
	}
	ctx := context.Background()
	nodeComponentCVEEdgePostgres.Destroy(ctx, db)

	store := nodeComponentCVEEdgePostgres.CreateTableAndNewStore(ctx, db, gormDB)
	indexer := nodeComponentCVEEdgePostgres.NewIndexer(db)
	searcher := nodeComponentCVEEdgeSearch.New(store, indexer)

	return nodeComponentCVEEdgeDS.New(store, indexer, searcher)
}

func registerNodeLoader(_ testing.TB, ds nodeDS.DataStore) {
	loaders.RegisterTypeFactory(reflect.TypeOf(storage.Node{}), func() interface{} {
		return loaders.NewNodeLoader(ds)
	})
}

func registerNodeComponentLoader(_ testing.TB, ds nodeComponentDS.DataStore) {
	loaders.RegisterTypeFactory(reflect.TypeOf(storage.NodeComponent{}), func() interface{} {
		return loaders.NewNodeComponentLoader(ds)
	})
}

func registerNodeCveLoader(_ testing.TB, ds nodeCVEDS.DataStore) {
	loaders.RegisterTypeFactory(reflect.TypeOf(storage.NodeCVE{}), func() interface{} {
		return loaders.NewNodeCVELoader(ds)
	})
}

func getTestNodesForPostgres(nodeCount int) []*storage.Node {
	nodes := make([]*storage.Node, 0, nodeCount)
	for i := 0; i < nodeCount; i++ {
		node := fixtures.GetNodeWithUniqueComponents(100)
		nodeConverter.MoveNodeVulnsToNewField(node)
		id := fmt.Sprintf("%d", i)
		node.Id = id
		nodes = append(nodes, node)
	}
	return nodes
}
