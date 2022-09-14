// Code generated by pg-bindings generator. DO NOT EDIT.
package n39ton40

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/pkg/errors"
	"github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/migrator/migrations"
	"github.com/stackrox/rox/migrator/migrations/loghelper"
	legacy "github.com/stackrox/rox/migrator/migrations/n_39_to_n_40_postgres_policies/legacy"
	pgStore "github.com/stackrox/rox/migrator/migrations/n_39_to_n_40_postgres_policies/postgres"
	"github.com/stackrox/rox/migrator/types"
	pkgMigrations "github.com/stackrox/rox/pkg/migrations"
	pkgSchema "github.com/stackrox/rox/pkg/postgres/schema"
	"github.com/stackrox/rox/pkg/sac"
	"gorm.io/gorm"
)

var (
	migration = types.Migration{
		StartingSeqNum: pkgMigrations.CurrentDBVersionSeqNumWithoutPostgres() + 39,
		VersionAfter:   storage.Version{SeqNum: int32(pkgMigrations.CurrentDBVersionSeqNumWithoutPostgres()) + 40},
		Run: func(databases *types.Databases) error {
			legacyStore := legacy.New(databases.BoltDB)
			if err := move(databases.GormDB, databases.PostgresDB, legacyStore); err != nil {
				return errors.Wrap(err,
					"moving policies from rocksdb to postgres")
			}
			return nil
		},
	}
	batchSize = 10000
	schema    = pkgSchema.PoliciesSchema
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
	policies, err := legacyStore.GetAll(ctx)
	if err != nil {
		return err
	}
	if len(policies) > 0 {
		if err = store.UpsertMany(ctx, policies); err != nil {
			log.WriteToStderrf("failed to persist policies to store %v", err)
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

func init() {
	migrations.MustRegisterMigration(migration)
}
