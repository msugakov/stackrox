// Code generated by pg-bindings generator. DO NOT EDIT.
package n57ton58

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/pkg/errors"
	"github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/migrator/migrations"
	"github.com/stackrox/rox/migrator/migrations/loghelper"
	"github.com/stackrox/rox/migrator/types"
	"github.com/stackrox/rox/pkg/db"
	ops "github.com/stackrox/rox/pkg/metrics"
	pkgSchema "github.com/stackrox/rox/pkg/postgres/schema"
	"github.com/stackrox/rox/pkg/rocksdb"
	generic "github.com/stackrox/rox/pkg/rocksdb/crud"
	"github.com/stackrox/rox/pkg/search"
	"github.com/stackrox/rox/pkg/search/postgres"
	"gorm.io/gorm"
)

var (
	migration = types.Migration{
		StartingSeqNum: 100,
		VersionAfter:   storage.Version{SeqNum: 101},
		Run: func(databases *types.Databases) error {
			if err := moveWatchedImages(databases.PkgRocksDB, databases.GormDB, databases.PostgresDB); err != nil {
				return errors.Wrap(err,
					"moving watched_images from rocksdb to postgres")
			}
			return nil
		},
	}
	rocksdbBucket = []byte("watchedimages")
	batchSize     = 10000
	schema        = pkgSchema.WatchedImagesSchema
	log           = loghelper.LogWrapper{}
)

func moveWatchedImages(rocksDB *rocksdb.RocksDB, gormDB *gorm.DB, postgresDB *pgxpool.Pool) error {
	ctx := context.Background()
	store := newStore(postgresDB, generic.NewCRUD(rocksDB, rocksdbBucket, keyFunc, alloc, false))
	pkgSchema.ApplySchemaForTable(context.Background(), gormDB, schema.Table)

	var watchedImages []*storage.WatchedImage
	store.Walk(ctx, func(obj *storage.WatchedImage) error {
		watchedImages = append(watchedImages, obj)
		if len(watchedImages) == 10*batchSize {
			if err := store.copyFrom(ctx, watchedImages...); err != nil {
				log.WriteToStderrf("failed to persist watched_images to store %v", err)
				return err
			}
			watchedImages = watchedImages[:0]
		}
		return nil
	})
	if len(watchedImages) > 0 {
		if err := store.copyFrom(ctx, watchedImages...); err != nil {
			log.WriteToStderrf("failed to persist watched_images to store %v", err)
			return err
		}
	}
	return nil
}

type storeImpl struct {
	db   *pgxpool.Pool // Postgres DB
	crud db.Crud       // Rocksdb DB crud
}

// newStore returns a new Store instance using the provided sql instance.
func newStore(db *pgxpool.Pool, crud db.Crud) *storeImpl {
	return &storeImpl{
		db:   db,
		crud: crud,
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
