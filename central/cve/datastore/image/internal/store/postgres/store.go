// Code generated by pg-bindings generator. DO NOT EDIT.

package postgres

import (
	"context"
	"reflect"
	"time"

	"github.com/gogo/protobuf/proto"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/stackrox/rox/central/globaldb"
	"github.com/stackrox/rox/central/metrics"
	"github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/pkg/logging"
	ops "github.com/stackrox/rox/pkg/metrics"
	"github.com/stackrox/rox/pkg/postgres/pgutils"
	"github.com/stackrox/rox/pkg/postgres/walker"
)

const (
	baseTable  = "image_cves"
	countStmt  = "SELECT COUNT(*) FROM image_cves"
	existsStmt = "SELECT EXISTS(SELECT 1 FROM image_cves WHERE Id = $1 AND OperatingSystem = $2)"

	getStmt    = "SELECT serialized FROM image_cves WHERE Id = $1 AND OperatingSystem = $2"
	deleteStmt = "DELETE FROM image_cves WHERE Id = $1 AND OperatingSystem = $2"
	walkStmt   = "SELECT serialized FROM image_cves"

	batchAfter = 100

	// using copyFrom, we may not even want to batch.  It would probably be simpler
	// to deal with failures if we just sent it all.  Something to think about as we
	// proceed and move into more e2e and larger performance testing
	batchSize = 10000
)

var (
	schema = walker.Walk(reflect.TypeOf((*storage.CVE)(nil)), baseTable)
	log    = logging.LoggerForModule()
)

func init() {
	globaldb.RegisterTable(schema)
}

type Store interface {
	Count(ctx context.Context) (int, error)
	Exists(ctx context.Context, id string, operatingSystem string) (bool, error)
	Get(ctx context.Context, id string, operatingSystem string) (*storage.CVE, bool, error)
	Upsert(ctx context.Context, obj *storage.CVE) error
	UpsertMany(ctx context.Context, objs []*storage.CVE) error
	Delete(ctx context.Context, id string, operatingSystem string) error

	Walk(ctx context.Context, fn func(obj *storage.CVE) error) error

	AckKeysIndexed(ctx context.Context, keys ...string) error
	GetKeysToIndex(ctx context.Context) ([]string, error)
}

type storeImpl struct {
	db *pgxpool.Pool
}

func createTableImageCves(ctx context.Context, db *pgxpool.Pool) {
	table := `
create table if not exists image_cves (
    Id varchar,
    OperatingSystem varchar,
    Cvss numeric,
    ImpactScore numeric,
    Summary varchar,
    Link varchar,
    PublishedOn timestamp,
    CreatedAt timestamp,
    LastModified timestamp,
    ScoreVersion integer,
    CvssV2_Vector varchar,
    CvssV2_AttackVector integer,
    CvssV2_AccessComplexity integer,
    CvssV2_Authentication integer,
    CvssV2_Confidentiality integer,
    CvssV2_Integrity integer,
    CvssV2_Availability integer,
    CvssV2_ExploitabilityScore numeric,
    CvssV2_ImpactScore numeric,
    CvssV2_Score numeric,
    CvssV2_Severity integer,
    CvssV3_Vector varchar,
    CvssV3_ExploitabilityScore numeric,
    CvssV3_ImpactScore numeric,
    CvssV3_AttackVector integer,
    CvssV3_AttackComplexity integer,
    CvssV3_PrivilegesRequired integer,
    CvssV3_UserInteraction integer,
    CvssV3_Scope integer,
    CvssV3_Confidentiality integer,
    CvssV3_Integrity integer,
    CvssV3_Availability integer,
    CvssV3_Score numeric,
    CvssV3_Severity integer,
    Suppressed bool,
    SuppressActivation timestamp,
    SuppressExpiry timestamp,
    Severity integer,
    serialized bytea,
    PRIMARY KEY(Id, OperatingSystem)
)
`

	_, err := db.Exec(ctx, table)
	if err != nil {
		log.Panicf("Error creating table %s: %v", table, err)
	}

	indexes := []string{}
	for _, index := range indexes {
		if _, err := db.Exec(ctx, index); err != nil {
			log.Panicf("Error creating index %s: %v", index, err)
		}
	}

	createTableImageCvesReferences(ctx, db)
}

func createTableImageCvesReferences(ctx context.Context, db *pgxpool.Pool) {
	table := `
create table if not exists image_cves_References (
    image_cves_Id varchar,
    image_cves_OperatingSystem varchar,
    idx integer,
    URI varchar,
    Tags text[],
    PRIMARY KEY(image_cves_Id, image_cves_OperatingSystem, idx),
    CONSTRAINT fk_parent_table_0 FOREIGN KEY (image_cves_Id, image_cves_OperatingSystem) REFERENCES image_cves(Id, OperatingSystem) ON DELETE CASCADE
)
`

	_, err := db.Exec(ctx, table)
	if err != nil {
		log.Panicf("Error creating table %s: %v", table, err)
	}

	indexes := []string{

		"create index if not exists imageCvesReferences_idx on image_cves_References using btree(idx)",
	}
	for _, index := range indexes {
		if _, err := db.Exec(ctx, index); err != nil {
			log.Panicf("Error creating index %s: %v", index, err)
		}
	}

}

func insertIntoImageCves(ctx context.Context, tx pgx.Tx, obj *storage.CVE) error {

	serialized, marshalErr := obj.Marshal()
	if marshalErr != nil {
		return marshalErr
	}

	values := []interface{}{
		// parent primary keys start
		obj.GetId(),
		obj.GetOperatingSystem(),
		obj.GetCvss(),
		obj.GetImpactScore(),
		obj.GetSummary(),
		obj.GetLink(),
		pgutils.NilOrTime(obj.GetPublishedOn()),
		pgutils.NilOrTime(obj.GetCreatedAt()),
		pgutils.NilOrTime(obj.GetLastModified()),
		obj.GetScoreVersion(),
		obj.GetCvssV2().GetVector(),
		obj.GetCvssV2().GetAttackVector(),
		obj.GetCvssV2().GetAccessComplexity(),
		obj.GetCvssV2().GetAuthentication(),
		obj.GetCvssV2().GetConfidentiality(),
		obj.GetCvssV2().GetIntegrity(),
		obj.GetCvssV2().GetAvailability(),
		obj.GetCvssV2().GetExploitabilityScore(),
		obj.GetCvssV2().GetImpactScore(),
		obj.GetCvssV2().GetScore(),
		obj.GetCvssV2().GetSeverity(),
		obj.GetCvssV3().GetVector(),
		obj.GetCvssV3().GetExploitabilityScore(),
		obj.GetCvssV3().GetImpactScore(),
		obj.GetCvssV3().GetAttackVector(),
		obj.GetCvssV3().GetAttackComplexity(),
		obj.GetCvssV3().GetPrivilegesRequired(),
		obj.GetCvssV3().GetUserInteraction(),
		obj.GetCvssV3().GetScope(),
		obj.GetCvssV3().GetConfidentiality(),
		obj.GetCvssV3().GetIntegrity(),
		obj.GetCvssV3().GetAvailability(),
		obj.GetCvssV3().GetScore(),
		obj.GetCvssV3().GetSeverity(),
		obj.GetSuppressed(),
		pgutils.NilOrTime(obj.GetSuppressActivation()),
		pgutils.NilOrTime(obj.GetSuppressExpiry()),
		obj.GetSeverity(),
		serialized,
	}

	finalStr := "INSERT INTO image_cves (Id, OperatingSystem, Cvss, ImpactScore, Summary, Link, PublishedOn, CreatedAt, LastModified, ScoreVersion, CvssV2_Vector, CvssV2_AttackVector, CvssV2_AccessComplexity, CvssV2_Authentication, CvssV2_Confidentiality, CvssV2_Integrity, CvssV2_Availability, CvssV2_ExploitabilityScore, CvssV2_ImpactScore, CvssV2_Score, CvssV2_Severity, CvssV3_Vector, CvssV3_ExploitabilityScore, CvssV3_ImpactScore, CvssV3_AttackVector, CvssV3_AttackComplexity, CvssV3_PrivilegesRequired, CvssV3_UserInteraction, CvssV3_Scope, CvssV3_Confidentiality, CvssV3_Integrity, CvssV3_Availability, CvssV3_Score, CvssV3_Severity, Suppressed, SuppressActivation, SuppressExpiry, Severity, serialized) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25, $26, $27, $28, $29, $30, $31, $32, $33, $34, $35, $36, $37, $38, $39) ON CONFLICT(Id, OperatingSystem) DO UPDATE SET Id = EXCLUDED.Id, OperatingSystem = EXCLUDED.OperatingSystem, Cvss = EXCLUDED.Cvss, ImpactScore = EXCLUDED.ImpactScore, Summary = EXCLUDED.Summary, Link = EXCLUDED.Link, PublishedOn = EXCLUDED.PublishedOn, CreatedAt = EXCLUDED.CreatedAt, LastModified = EXCLUDED.LastModified, ScoreVersion = EXCLUDED.ScoreVersion, CvssV2_Vector = EXCLUDED.CvssV2_Vector, CvssV2_AttackVector = EXCLUDED.CvssV2_AttackVector, CvssV2_AccessComplexity = EXCLUDED.CvssV2_AccessComplexity, CvssV2_Authentication = EXCLUDED.CvssV2_Authentication, CvssV2_Confidentiality = EXCLUDED.CvssV2_Confidentiality, CvssV2_Integrity = EXCLUDED.CvssV2_Integrity, CvssV2_Availability = EXCLUDED.CvssV2_Availability, CvssV2_ExploitabilityScore = EXCLUDED.CvssV2_ExploitabilityScore, CvssV2_ImpactScore = EXCLUDED.CvssV2_ImpactScore, CvssV2_Score = EXCLUDED.CvssV2_Score, CvssV2_Severity = EXCLUDED.CvssV2_Severity, CvssV3_Vector = EXCLUDED.CvssV3_Vector, CvssV3_ExploitabilityScore = EXCLUDED.CvssV3_ExploitabilityScore, CvssV3_ImpactScore = EXCLUDED.CvssV3_ImpactScore, CvssV3_AttackVector = EXCLUDED.CvssV3_AttackVector, CvssV3_AttackComplexity = EXCLUDED.CvssV3_AttackComplexity, CvssV3_PrivilegesRequired = EXCLUDED.CvssV3_PrivilegesRequired, CvssV3_UserInteraction = EXCLUDED.CvssV3_UserInteraction, CvssV3_Scope = EXCLUDED.CvssV3_Scope, CvssV3_Confidentiality = EXCLUDED.CvssV3_Confidentiality, CvssV3_Integrity = EXCLUDED.CvssV3_Integrity, CvssV3_Availability = EXCLUDED.CvssV3_Availability, CvssV3_Score = EXCLUDED.CvssV3_Score, CvssV3_Severity = EXCLUDED.CvssV3_Severity, Suppressed = EXCLUDED.Suppressed, SuppressActivation = EXCLUDED.SuppressActivation, SuppressExpiry = EXCLUDED.SuppressExpiry, Severity = EXCLUDED.Severity, serialized = EXCLUDED.serialized"
	_, err := tx.Exec(ctx, finalStr, values...)
	if err != nil {
		return err
	}

	var query string

	for childIdx, child := range obj.GetReferences() {
		if err := insertIntoImageCvesReferences(ctx, tx, child, obj.GetId(), obj.GetOperatingSystem(), childIdx); err != nil {
			return err
		}
	}

	query = "delete from image_cves_References where image_cves_Id = $1 AND image_cves_OperatingSystem = $2 AND idx >= $3"
	_, err = tx.Exec(ctx, query, obj.GetId(), obj.GetOperatingSystem(), len(obj.GetReferences()))
	if err != nil {
		return err
	}
	return nil
}

func insertIntoImageCvesReferences(ctx context.Context, tx pgx.Tx, obj *storage.CVE_Reference, image_cves_Id string, image_cves_OperatingSystem string, idx int) error {

	values := []interface{}{
		// parent primary keys start
		image_cves_Id,
		image_cves_OperatingSystem,
		idx,
		obj.GetURI(),
		obj.GetTags(),
	}

	finalStr := "INSERT INTO image_cves_References (image_cves_Id, image_cves_OperatingSystem, idx, URI, Tags) VALUES($1, $2, $3, $4, $5) ON CONFLICT(image_cves_Id, image_cves_OperatingSystem, idx) DO UPDATE SET image_cves_Id = EXCLUDED.image_cves_Id, image_cves_OperatingSystem = EXCLUDED.image_cves_OperatingSystem, idx = EXCLUDED.idx, URI = EXCLUDED.URI, Tags = EXCLUDED.Tags"
	_, err := tx.Exec(ctx, finalStr, values...)
	if err != nil {
		return err
	}

	return nil
}

func (s *storeImpl) copyFromImageCves(ctx context.Context, tx pgx.Tx, objs ...*storage.CVE) error {

	inputRows := [][]interface{}{}

	var err error

	copyCols := []string{

		"id",

		"operatingsystem",

		"cvss",

		"impactscore",

		"summary",

		"link",

		"publishedon",

		"createdat",

		"lastmodified",

		"scoreversion",

		"cvssv2_vector",

		"cvssv2_attackvector",

		"cvssv2_accesscomplexity",

		"cvssv2_authentication",

		"cvssv2_confidentiality",

		"cvssv2_integrity",

		"cvssv2_availability",

		"cvssv2_exploitabilityscore",

		"cvssv2_impactscore",

		"cvssv2_score",

		"cvssv2_severity",

		"cvssv3_vector",

		"cvssv3_exploitabilityscore",

		"cvssv3_impactscore",

		"cvssv3_attackvector",

		"cvssv3_attackcomplexity",

		"cvssv3_privilegesrequired",

		"cvssv3_userinteraction",

		"cvssv3_scope",

		"cvssv3_confidentiality",

		"cvssv3_integrity",

		"cvssv3_availability",

		"cvssv3_score",

		"cvssv3_severity",

		"suppressed",

		"suppressactivation",

		"suppressexpiry",

		"severity",

		"serialized",
	}

	for idx, obj := range objs {
		// Todo: ROX-9499 Figure out how to more cleanly template around this issue.
		log.Debugf("This is here for now because there is an issue with pods_TerminatedInstances where the obj in the loop is not used as it only consists of the parent id and the idx.  Putting this here as a stop gap to simply use the object.  %s", obj)

		serialized, marshalErr := obj.Marshal()
		if marshalErr != nil {
			return marshalErr
		}

		inputRows = append(inputRows, []interface{}{

			obj.GetId(),

			obj.GetOperatingSystem(),

			obj.GetCvss(),

			obj.GetImpactScore(),

			obj.GetSummary(),

			obj.GetLink(),

			pgutils.NilOrTime(obj.GetPublishedOn()),

			pgutils.NilOrTime(obj.GetCreatedAt()),

			pgutils.NilOrTime(obj.GetLastModified()),

			obj.GetScoreVersion(),

			obj.GetCvssV2().GetVector(),

			obj.GetCvssV2().GetAttackVector(),

			obj.GetCvssV2().GetAccessComplexity(),

			obj.GetCvssV2().GetAuthentication(),

			obj.GetCvssV2().GetConfidentiality(),

			obj.GetCvssV2().GetIntegrity(),

			obj.GetCvssV2().GetAvailability(),

			obj.GetCvssV2().GetExploitabilityScore(),

			obj.GetCvssV2().GetImpactScore(),

			obj.GetCvssV2().GetScore(),

			obj.GetCvssV2().GetSeverity(),

			obj.GetCvssV3().GetVector(),

			obj.GetCvssV3().GetExploitabilityScore(),

			obj.GetCvssV3().GetImpactScore(),

			obj.GetCvssV3().GetAttackVector(),

			obj.GetCvssV3().GetAttackComplexity(),

			obj.GetCvssV3().GetPrivilegesRequired(),

			obj.GetCvssV3().GetUserInteraction(),

			obj.GetCvssV3().GetScope(),

			obj.GetCvssV3().GetConfidentiality(),

			obj.GetCvssV3().GetIntegrity(),

			obj.GetCvssV3().GetAvailability(),

			obj.GetCvssV3().GetScore(),

			obj.GetCvssV3().GetSeverity(),

			obj.GetSuppressed(),

			pgutils.NilOrTime(obj.GetSuppressActivation()),

			pgutils.NilOrTime(obj.GetSuppressExpiry()),

			obj.GetSeverity(),

			serialized,
		})

		if _, err := tx.Exec(ctx, deleteStmt, obj.GetId(), obj.GetOperatingSystem()); err != nil {
			return err
		}

		// if we hit our batch size we need to push the data
		if (idx+1)%batchSize == 0 || idx == len(objs)-1 {
			// copy does not upsert so have to delete first.  parent deletion cascades so only need to
			// delete for the top level parent

			_, err = tx.CopyFrom(ctx, pgx.Identifier{"image_cves"}, copyCols, pgx.CopyFromRows(inputRows))

			if err != nil {
				return err
			}

			// clear the input rows for the next batch
			inputRows = inputRows[:0]
		}
	}

	for _, obj := range objs {

		if err = s.copyFromImageCvesReferences(ctx, tx, obj.GetId(), obj.GetOperatingSystem(), obj.GetReferences()...); err != nil {
			return err
		}
	}

	return err
}

func (s *storeImpl) copyFromImageCvesReferences(ctx context.Context, tx pgx.Tx, image_cves_Id string, image_cves_OperatingSystem string, objs ...*storage.CVE_Reference) error {

	inputRows := [][]interface{}{}

	var err error

	copyCols := []string{

		"image_cves_id",

		"image_cves_operatingsystem",

		"idx",

		"uri",

		"tags",
	}

	for idx, obj := range objs {
		// Todo: ROX-9499 Figure out how to more cleanly template around this issue.
		log.Debugf("This is here for now because there is an issue with pods_TerminatedInstances where the obj in the loop is not used as it only consists of the parent id and the idx.  Putting this here as a stop gap to simply use the object.  %s", obj)

		inputRows = append(inputRows, []interface{}{

			image_cves_Id,

			image_cves_OperatingSystem,

			idx,

			obj.GetURI(),

			obj.GetTags(),
		})

		// if we hit our batch size we need to push the data
		if (idx+1)%batchSize == 0 || idx == len(objs)-1 {
			// copy does not upsert so have to delete first.  parent deletion cascades so only need to
			// delete for the top level parent

			_, err = tx.CopyFrom(ctx, pgx.Identifier{"image_cves_references"}, copyCols, pgx.CopyFromRows(inputRows))

			if err != nil {
				return err
			}

			// clear the input rows for the next batch
			inputRows = inputRows[:0]
		}
	}

	return err
}

// New returns a new Store instance using the provided sql instance.
func New(ctx context.Context, db *pgxpool.Pool) Store {
	createTableImageCves(ctx, db)

	return &storeImpl{
		db: db,
	}
}

func (s *storeImpl) copyFrom(ctx context.Context, objs ...*storage.CVE) error {
	conn, release := s.acquireConn(ctx, ops.Get, "CVE")
	defer release()

	tx, err := conn.Begin(ctx)
	if err != nil {
		return err
	}

	if err := s.copyFromImageCves(ctx, tx, objs...); err != nil {
		if err := tx.Rollback(ctx); err != nil {
			return err
		}
		return err
	}
	if err := tx.Commit(ctx); err != nil {
		return err
	}
	return nil
}

func (s *storeImpl) upsert(ctx context.Context, objs ...*storage.CVE) error {
	conn, release := s.acquireConn(ctx, ops.Get, "CVE")
	defer release()

	for _, obj := range objs {
		tx, err := conn.Begin(ctx)
		if err != nil {
			return err
		}

		if err := insertIntoImageCves(ctx, tx, obj); err != nil {
			if err := tx.Rollback(ctx); err != nil {
				return err
			}
			return err
		}
		if err := tx.Commit(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (s *storeImpl) Upsert(ctx context.Context, obj *storage.CVE) error {
	defer metrics.SetPostgresOperationDurationTime(time.Now(), ops.Upsert, "CVE")

	return s.upsert(ctx, obj)
}

func (s *storeImpl) UpsertMany(ctx context.Context, objs []*storage.CVE) error {
	defer metrics.SetPostgresOperationDurationTime(time.Now(), ops.UpdateMany, "CVE")

	if len(objs) < batchAfter {
		return s.upsert(ctx, objs...)
	} else {
		return s.copyFrom(ctx, objs...)
	}
}

// Count returns the number of objects in the store
func (s *storeImpl) Count(ctx context.Context) (int, error) {
	defer metrics.SetPostgresOperationDurationTime(time.Now(), ops.Count, "CVE")

	row := s.db.QueryRow(ctx, countStmt)
	var count int
	if err := row.Scan(&count); err != nil {
		return 0, err
	}
	return count, nil
}

// Exists returns if the id exists in the store
func (s *storeImpl) Exists(ctx context.Context, id string, operatingSystem string) (bool, error) {
	defer metrics.SetPostgresOperationDurationTime(time.Now(), ops.Exists, "CVE")

	row := s.db.QueryRow(ctx, existsStmt, id, operatingSystem)
	var exists bool
	if err := row.Scan(&exists); err != nil {
		return false, pgutils.ErrNilIfNoRows(err)
	}
	return exists, nil
}

// Get returns the object, if it exists from the store
func (s *storeImpl) Get(ctx context.Context, id string, operatingSystem string) (*storage.CVE, bool, error) {
	defer metrics.SetPostgresOperationDurationTime(time.Now(), ops.Get, "CVE")

	conn, release := s.acquireConn(ctx, ops.Get, "CVE")
	defer release()

	row := conn.QueryRow(ctx, getStmt, id, operatingSystem)
	var data []byte
	if err := row.Scan(&data); err != nil {
		return nil, false, pgutils.ErrNilIfNoRows(err)
	}

	var msg storage.CVE
	if err := proto.Unmarshal(data, &msg); err != nil {
		return nil, false, err
	}
	return &msg, true, nil
}

func (s *storeImpl) acquireConn(ctx context.Context, op ops.Op, typ string) (*pgxpool.Conn, func()) {
	defer metrics.SetAcquireDBConnDuration(time.Now(), op, typ)
	conn, err := s.db.Acquire(ctx)
	if err != nil {
		panic(err)
	}
	return conn, conn.Release
}

// Delete removes the specified ID from the store
func (s *storeImpl) Delete(ctx context.Context, id string, operatingSystem string) error {
	defer metrics.SetPostgresOperationDurationTime(time.Now(), ops.Remove, "CVE")

	conn, release := s.acquireConn(ctx, ops.Remove, "CVE")
	defer release()

	if _, err := conn.Exec(ctx, deleteStmt, id, operatingSystem); err != nil {
		return err
	}
	return nil
}

// Walk iterates over all of the objects in the store and applies the closure
func (s *storeImpl) Walk(ctx context.Context, fn func(obj *storage.CVE) error) error {
	rows, err := s.db.Query(ctx, walkStmt)
	if err != nil {
		return pgutils.ErrNilIfNoRows(err)
	}
	defer rows.Close()
	for rows.Next() {
		var data []byte
		if err := rows.Scan(&data); err != nil {
			return err
		}
		var msg storage.CVE
		if err := proto.Unmarshal(data, &msg); err != nil {
			return err
		}
		if err := fn(&msg); err != nil {
			return err
		}
	}
	return nil
}

//// Used for testing

func dropTableImageCves(ctx context.Context, db *pgxpool.Pool) {
	_, _ = db.Exec(ctx, "DROP TABLE IF EXISTS image_cves CASCADE")
	dropTableImageCvesReferences(ctx, db)

}

func dropTableImageCvesReferences(ctx context.Context, db *pgxpool.Pool) {
	_, _ = db.Exec(ctx, "DROP TABLE IF EXISTS image_cves_References CASCADE")

}

func Destroy(ctx context.Context, db *pgxpool.Pool) {
	dropTableImageCves(ctx, db)
}

//// Stubs for satisfying legacy interfaces

// AckKeysIndexed acknowledges the passed keys were indexed
func (s *storeImpl) AckKeysIndexed(ctx context.Context, keys ...string) error {
	return nil
}

// GetKeysToIndex returns the keys that need to be indexed
func (s *storeImpl) GetKeysToIndex(ctx context.Context) ([]string, error) {
	return nil, nil
}
