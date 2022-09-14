// Code generated by pg-bindings generator. DO NOT EDIT.
package n10ton11

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/pkg/errors"
	"github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/migrator/migrations"
	"github.com/stackrox/rox/migrator/migrations/loghelper"
	legacy "github.com/stackrox/rox/migrator/migrations/n_10_to_n_11_postgres_cluster_health_statuses/legacy"
	pgStore "github.com/stackrox/rox/migrator/migrations/n_10_to_n_11_postgres_cluster_health_statuses/postgres"
	"github.com/stackrox/rox/migrator/types"
	pkgMigrations "github.com/stackrox/rox/pkg/migrations"
	pkgSchema "github.com/stackrox/rox/pkg/postgres/schema"
	"github.com/stackrox/rox/pkg/sac"
	"gorm.io/gorm"
)

var (
	migration = types.Migration{
		StartingSeqNum: pkgMigrations.CurrentDBVersionSeqNumWithoutPostgres() + 10,
		VersionAfter:   storage.Version{SeqNum: int32(pkgMigrations.CurrentDBVersionSeqNumWithoutPostgres()) + 11},
		Run: func(databases *types.Databases) error {
			legacyStore, err := legacy.New(databases.PkgRocksDB)
			if err != nil {
				return err
			}
			if err := move(databases.GormDB, databases.PostgresDB, legacyStore); err != nil {
				return errors.Wrap(err,
					"moving cluster_health_statuses from rocksdb to postgres")
			}
			return nil
		},
	}
	batchSize = 10000
	schema    = pkgSchema.ClusterHealthStatusesSchema
	log       = loghelper.LogWrapper{}
)

func move(gormDB *gorm.DB, postgresDB *pgxpool.Pool, legacyStore legacy.Store) error {
	ctx := sac.WithAllAccess(context.Background())
	store := pgStore.New(postgresDB)
	pkgSchema.ApplySchemaForTable(context.Background(), gormDB, schema.Table)
	_, err := postgresDB.Exec(ctx, fmt.Sprintf("ALTER TABLE %s DISABLE TRIGGER ALL", schema.Table))
	if err != nil {
		log.WriteToStderrf("failed to disable triggers for %s", schema.Table)
		return err
	}
	var clusterHealthStatuses []*storage.ClusterHealthStatus
	err = walk(ctx, legacyStore, func(obj *storage.ClusterHealthStatus) error {
		clusterHealthStatuses = append(clusterHealthStatuses, obj)
		if len(clusterHealthStatuses) == batchSize {
			if err := store.UpsertMany(ctx, clusterHealthStatuses); err != nil {
				log.WriteToStderrf("failed to persist cluster_health_statuses to store %v", err)
				return err
			}
			clusterHealthStatuses = clusterHealthStatuses[:0]
		}
		return nil
	})
	if err != nil {
		return err
	}
	if len(clusterHealthStatuses) > 0 {
		if err = store.UpsertMany(ctx, clusterHealthStatuses); err != nil {
			log.WriteToStderrf("failed to persist cluster_health_statuses to store %v", err)
			return err
		}
	}
	_, err = postgresDB.Exec(ctx, fmt.Sprintf("ALTER TABLE %s ENABLE TRIGGER ALL", schema.Table))
	if err != nil {
		log.WriteToStderrf("failed to enable triggers for %s", schema.Table)
		return err
	}
	return nil
}

func walk(ctx context.Context, s legacy.Store, fn func(obj *storage.ClusterHealthStatus) error) error {
	return s.Walk(ctx, fn)
}

func init() {
	migrations.MustRegisterMigration(migration)
}
