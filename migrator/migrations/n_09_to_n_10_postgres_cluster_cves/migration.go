// Code originally generated by pg-bindings generator.

package n9ton10

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/pkg/errors"
	"github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/migrator/migrations"
	cveUtil "github.com/stackrox/rox/migrator/migrations/cvehelper"
	"github.com/stackrox/rox/migrator/migrations/loghelper"
	legacy "github.com/stackrox/rox/migrator/migrations/n_09_to_n_10_postgres_cluster_cves/legacy"
	pgStore "github.com/stackrox/rox/migrator/migrations/n_09_to_n_10_postgres_cluster_cves/postgres"
	"github.com/stackrox/rox/migrator/types"
	rawDackbox "github.com/stackrox/rox/pkg/dackbox/raw"
	pkgMigrations "github.com/stackrox/rox/pkg/migrations"
	pkgSchema "github.com/stackrox/rox/pkg/postgres/schema"
	"github.com/stackrox/rox/pkg/sac"
	"gorm.io/gorm"
)

var (
	migration = types.Migration{
		StartingSeqNum: pkgMigrations.CurrentDBVersionSeqNumWithoutPostgres() + 9,
		VersionAfter:   storage.Version{SeqNum: int32(pkgMigrations.CurrentDBVersionSeqNumWithoutPostgres()) + 10},
		Run: func(databases *types.Databases) error {
			legacyStore := legacy.New(rawDackbox.GetGlobalDackBox(), rawDackbox.GetKeyFence())
			if err := move(databases.GormDB, databases.PostgresDB, legacyStore); err != nil {
				return errors.Wrap(err,
					"moving cluster_cves from rocksdb to postgres")
			}
			return nil
		},
	}
	batchSize = 10000
	schema    = pkgSchema.ClusterCvesSchema
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
	var clusterCves []*storage.ClusterCVE
	err = walk(ctx, legacyStore, func(obj *storage.CVE) error {
		clusterCves = append(clusterCves, convertCVEToClusterCVEs(obj)...)
		if len(clusterCves) == batchSize {
			if err := store.UpsertMany(ctx, clusterCves); err != nil {
				log.WriteToStderrf("failed to persist cluster_cves to store %v", err)
				return err
			}
			clusterCves = clusterCves[:0]
		}
		return nil
	})
	if err != nil {
		return err
	}
	if len(clusterCves) > 0 {
		if err = store.UpsertMany(ctx, clusterCves); err != nil {
			log.WriteToStderrf("failed to persist cluster_cves to store %v", err)
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

func walk(ctx context.Context, s legacy.Store, fn func(obj *storage.CVE) error) error {
	ids, err := s.GetIDs(ctx)
	if err != nil {
		return err
	}

	for i := 0; i < len(ids); i += batchSize {
		end := i + batchSize

		if end > len(ids) {
			end = len(ids)
		}
		objs, _, err := s.GetMany(ctx, ids[i:end])
		if err != nil {
			return err
		}
		for _, obj := range objs {
			if err = fn(obj); err != nil {
				return err
			}
		}
	}
	return nil
}

func init() {
	migrations.MustRegisterMigration(migration)
}

// convertCVEToClusterCVEs split a storage.CVE to potential multiple storage.ClusterCVE(s)
// for each type
func convertCVEToClusterCVEs(cve *storage.CVE) []*storage.ClusterCVE {
	clusterCVE := &storage.ClusterCVE{
		CveBaseInfo: &storage.CVEInfo{
			Cve:          cve.GetId(),
			Summary:      cve.GetSummary(),
			Link:         cve.GetLink(),
			PublishedOn:  cve.GetPublishedOn(),
			CreatedAt:    cve.GetCreatedAt(),
			LastModified: cve.GetLastModified(),
			CvssV2:       cve.GetCvssV2(),
			CvssV3:       cve.GetCvssV3(),
		},
		Cvss:         cve.GetCvss(),
		Severity:     cve.GetSeverity(),
		ImpactScore:  cve.GetImpactScore(),
		Snoozed:      cve.GetSuppressed(),
		SnoozeStart:  cve.GetSuppressActivation(),
		SnoozeExpiry: cve.GetSuppressExpiry(),
	}
	if clusterCVE.GetCveBaseInfo().GetCvssV3() != nil {
		clusterCVE.CveBaseInfo.ScoreVersion = storage.CVEInfo_V3
		clusterCVE.ImpactScore = cve.GetCvssV3().GetImpactScore()
	} else if clusterCVE.GetCveBaseInfo().GetCvssV2() != nil {
		clusterCVE.CveBaseInfo.ScoreVersion = storage.CVEInfo_V2
		clusterCVE.ImpactScore = cve.GetCvssV2().GetImpactScore()
	}

	ret := make([]*storage.ClusterCVE, 0, len(cve.GetTypes()))
	for _, typ := range cve.GetTypes() {
		cloned := clusterCVE.Clone()
		cloned.Id = cveUtil.ID(cve.GetId(), typ.String())
		cloned.Type = typ
		ret = append(ret, cloned)
	}
	return ret
}
