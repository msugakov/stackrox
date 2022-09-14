// Code generated by pg-bindings generator. DO NOT EDIT.
package n7ton8

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/pkg/errors"
	"github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/migrator/migrations"
	"github.com/stackrox/rox/migrator/migrations/loghelper"
	legacy "github.com/stackrox/rox/migrator/migrations/n_07_to_n_08_postgres_api_tokens/legacy"
	pgStore "github.com/stackrox/rox/migrator/migrations/n_07_to_n_08_postgres_api_tokens/postgres"
	"github.com/stackrox/rox/migrator/types"
	pkgMigrations "github.com/stackrox/rox/pkg/migrations"
	pkgSchema "github.com/stackrox/rox/pkg/postgres/schema"
	"github.com/stackrox/rox/pkg/sac"
	"gorm.io/gorm"
)

var (
	migration = types.Migration{
		StartingSeqNum: pkgMigrations.CurrentDBVersionSeqNumWithoutPostgres() + 7,
		VersionAfter:   storage.Version{SeqNum: int32(pkgMigrations.CurrentDBVersionSeqNumWithoutPostgres()) + 8},
		Run: func(databases *types.Databases) error {
			legacyStore, err := legacy.New(databases.PkgRocksDB)
			if err != nil {
				return err
			}
			if err := move(databases.GormDB, databases.PostgresDB, legacyStore); err != nil {
				return errors.Wrap(err,
					"moving api_tokens from rocksdb to postgres")
			}
			return nil
		},
	}
	batchSize = 10000
	schema    = pkgSchema.ApiTokensSchema
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
	var apiTokens []*storage.TokenMetadata
	err = walk(ctx, legacyStore, func(obj *storage.TokenMetadata) error {
		apiTokens = append(apiTokens, obj)
		if len(apiTokens) == batchSize {
			if err := store.UpsertMany(ctx, apiTokens); err != nil {
				log.WriteToStderrf("failed to persist api_tokens to store %v", err)
				return err
			}
			apiTokens = apiTokens[:0]
		}
		return nil
	})
	if err != nil {
		return err
	}
	if len(apiTokens) > 0 {
		if err = store.UpsertMany(ctx, apiTokens); err != nil {
			log.WriteToStderrf("failed to persist api_tokens to store %v", err)
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

func walk(ctx context.Context, s legacy.Store, fn func(obj *storage.TokenMetadata) error) error {
	return s.Walk(ctx, fn)
}

func init() {
	migrations.MustRegisterMigration(migration)
}
