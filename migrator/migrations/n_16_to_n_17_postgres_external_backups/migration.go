// Code generated by pg-bindings generator. DO NOT EDIT.
package n16ton17

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/pkg/errors"
	"github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/migrator/migrations"
	"github.com/stackrox/rox/migrator/migrations/loghelper"
	legacy "github.com/stackrox/rox/migrator/migrations/n_16_to_n_17_postgres_external_backups/legacy"
	pgStore "github.com/stackrox/rox/migrator/migrations/n_16_to_n_17_postgres_external_backups/postgres"
	"github.com/stackrox/rox/migrator/types"
	pkgSchema "github.com/stackrox/rox/pkg/postgres/schema"
	"github.com/stackrox/rox/pkg/sac"
	bolt "go.etcd.io/bbolt"
	"gorm.io/gorm"
)

var (
	migration = types.Migration{
		StartingSeqNum: 100,
		VersionAfter:   storage.Version{SeqNum: 101},
		Run: func(databases *types.Databases) error {
			legacyStore := legacy.New(databases.BoltDB)
			if err := move(databases.BoltDB, databases.GormDB, databases.PostgresDB, legacyStore); err != nil {
				return errors.Wrap(err,
					"moving external_backups from rocksdb to postgres")
			}
			return nil
		},
	}
	batchSize = 10000
	schema    = pkgSchema.ExternalBackupsSchema
	log       = loghelper.LogWrapper{}
)

func move(legacyDB *bolt.DB, gormDB *gorm.DB, postgresDB *pgxpool.Pool, legacyStore legacy.Store) error {
	ctx := sac.WithGlobalAccessScopeChecker(context.Background(), sac.AllowAllAccessScopeChecker())
	store := pgStore.New(postgresDB)
	pkgSchema.ApplySchemaForTable(context.Background(), gormDB, schema.Table)
	var externalBackups []*storage.ExternalBackup
	var err error
	externalBackups, err = legacyStore.GetAll(ctx)
	if err != nil {
		log.WriteToStderr("failed to fetch all externalBackups")
		return err
	}
	if len(externalBackups) > 0 {
		if err = store.UpsertMany(ctx, externalBackups); err != nil {
			log.WriteToStderrf("failed to persist external_backups to store %v", err)
			return err
		}
	}
	return nil
}

func init() {
	migrations.MustRegisterMigration(migration)
}
