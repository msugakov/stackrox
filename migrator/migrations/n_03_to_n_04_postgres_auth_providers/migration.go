// Code generated by pg-bindings generator. DO NOT EDIT.
package n3ton4

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/pkg/errors"
	"github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/migrator/migrations"
	"github.com/stackrox/rox/migrator/migrations/loghelper"
	"github.com/stackrox/rox/migrator/types"
	ops "github.com/stackrox/rox/pkg/metrics"
	pkgSchema "github.com/stackrox/rox/pkg/postgres/schema"
	"github.com/stackrox/rox/pkg/search"
	"github.com/stackrox/rox/pkg/search/postgres"
	bolt "go.etcd.io/bbolt"
	"gorm.io/gorm"
)

var (
	migration = types.Migration{
		StartingSeqNum: 100,
		VersionAfter:   storage.Version{SeqNum: 101},
		Run: func(databases *types.Databases) error {
			if err := moveAuthProviders(databases.BoltDB, databases.GormDB, databases.PostgresDB); err != nil {
				return errors.Wrap(err,
					"moving auth_providers from rocksdb to postgres")
			}
			return nil
		},
	}
	authProviderBucket = []byte("authProviders")
	batchSize          = 10000
	schema             = pkgSchema.AuthProvidersSchema
	log                = loghelper.LogWrapper{}
)

func moveAuthProviders(legacyDB *bolt.DB, gormDB *gorm.DB, postgresDB *pgxpool.Pool) error {
	ctx := context.Background()
	store := newStore(postgresDB, legacyDB)
	pkgSchema.ApplySchemaForTable(context.Background(), gormDB, schema.Table)

	var authProviders []*storage.AuthProvider
	var err error
	authProviders, err = store.GetAll(ctx)
	if err != nil {
		log.WriteToStderr("failed to fetch all authProviders")
		return err
	}
	if len(authProviders) > 0 {
		if err = store.copyFrom(ctx, authProviders...); err != nil {
			log.WriteToStderrf("failed to persist auth_providers to store %v", err)
			return err
		}
	}
	return nil
}

type storeImpl struct {
	db       *pgxpool.Pool // Postgres DB
	legacyDB *bolt.DB
}

// newStore returns a new Store instance using the provided sql instance.
func newStore(db *pgxpool.Pool, legacyDB *bolt.DB) *storeImpl {
	return &storeImpl{
		db: db, legacyDB: legacyDB,
	}
}

func (s *storeImpl) acquireConn(ctx context.Context, _ ops.Op, _ string) (*pgxpool.Conn, func(), error) {
	conn, err := s.db.Acquire(ctx)
	if err != nil {
		return nil, nil, err
	}
	return conn, conn.Release, nil
}
func (s *storeImpl) DeleteMany(ctx context.Context, ids []string) error {
	q := search.NewQueryBuilder().AddDocIDs(ids...).ProtoQuery()
	return postgres.RunDeleteRequestForSchema(schema, q, s.db)
}

func init() {
	migrations.MustRegisterMigration(migration)
}
