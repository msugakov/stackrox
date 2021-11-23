// Code generated by pg-bindings generator. DO NOT EDIT.

package postgres

import (
	"bytes"
	"context"
	"reflect"
	"time"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/gogo/protobuf/proto"
	"github.com/gogo/protobuf/types"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/stackrox/rox/central/globaldb"
	"github.com/stackrox/rox/central/metrics"
	"github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/pkg/logging"
	ops "github.com/stackrox/rox/pkg/metrics"
	"github.com/stackrox/rox/pkg/set"
)

const (
		countStmt = "select count(*) from ProcessIndicator"
		existsStmt = "select exists(select 1 from ProcessIndicator where id = $1)"
		getIDsStmt = "select id from ProcessIndicator"
		getStmt = "select serialized from ProcessIndicator where id = $1"
		getManyStmt = "select serialized from ProcessIndicator where id = ANY($1::text[])"
		deleteStmt = "delete from ProcessIndicator where id = $1"
		deleteManyStmt = "delete from ProcessIndicator where id = ANY($1::text[])"
		walkStmt = "select serialized from ProcessIndicator"
		walkWithIDStmt = "select id, serialized from ProcessIndicator"
)

var (
	log = logging.LoggerForModule()

	table = "ProcessIndicator"

	marshaler = &jsonpb.Marshaler{EnumsAsInts: true, EmitDefaults: true}
)

type Store interface {
	Count() (int, error)
	Exists(id string) (bool, error)
	GetIDs() ([]string, error)
	Get(id string) (*storage.ProcessIndicator, bool, error)
	GetMany(ids []string) ([]*storage.ProcessIndicator, []int, error)
	Upsert(obj *storage.ProcessIndicator) error
	UpsertMany(objs []*storage.ProcessIndicator) error
	Delete(id string) error
	DeleteMany(ids []string) error
	Walk(fn func(obj *storage.ProcessIndicator) error) error
	AckKeysIndexed(keys ...string) error
	GetKeysToIndex() ([]string, error)
}

type storeImpl struct {
	db *pgxpool.Pool
}

func alloc() proto.Message {
	return &storage.ProcessIndicator{}
}

func keyFunc(msg proto.Message) string {
	return msg.(*storage.ProcessIndicator).GetId()
}

const (
	batchInsertTemplate = "<no value>"
)

// New returns a new Store instance using the provided sql instance.
func New(db *pgxpool.Pool) Store {
	globaldb.RegisterTable(table, "ProcessIndicator")

	for _, table := range []string {
		"create table if not exists ProcessIndicator(serialized jsonb not null, Id varchar, DeploymentId varchar, ContainerName varchar, PodId varchar, PodUid varchar, ClusterId varchar, Namespace varchar, Signal_ContainerId varchar, Signal_Name varchar, Signal_Args varchar, Signal_ExecFilePath varchar, Signal_Uid numeric, PRIMARY KEY (Id));",
		"create index if not exists ProcessIndicator_DeploymentId on ProcessIndicator using hash(DeploymentId)",
		"create index if not exists ProcessIndicator_ContainerName on ProcessIndicator using hash(ContainerName)",
		"create index if not exists ProcessIndicator_PodUid on ProcessIndicator using hash(PodUid)",
		
	} {
		_, err := db.Exec(context.Background(), table)
		if err != nil {
			panic("error creating table: " + table)
		}
	}

//
	return &storeImpl{
		db: db,
	}
//
}

// Count returns the number of objects in the store
func (s *storeImpl) Count() (int, error) {
	defer metrics.SetPostgresOperationDurationTime(time.Now(), ops.Count, "ProcessIndicator")

	row := s.db.QueryRow(context.Background(), countStmt)
	var count int
	if err := row.Scan(&count); err != nil {
		return 0, err
	}
	return count, nil
}

// Exists returns if the id exists in the store
func (s *storeImpl) Exists(id string) (bool, error) {
	defer metrics.SetPostgresOperationDurationTime(time.Now(), ops.Exists, "ProcessIndicator")

	row := s.db.QueryRow(context.Background(), existsStmt, id)
	var exists bool
	if err := row.Scan(&exists); err != nil {
		return false, nilNoRows(err)
	}
	return exists, nil
}

// GetIDs returns all the IDs for the store
func (s *storeImpl) GetIDs() ([]string, error) {
	defer metrics.SetPostgresOperationDurationTime(time.Now(), ops.GetAll, "ProcessIndicatorIDs")

	rows, err := s.db.Query(context.Background(), getIDsStmt)
	if err != nil {
		return nil, nilNoRows(err)
	}
	defer rows.Close()
	var ids []string
	for rows.Next() {
		var id string
		if err := rows.Scan(&id); err != nil {
			return nil, err
		}
		ids = append(ids, id)
	}
	return ids, nil
}

func nilNoRows(err error) error {
	if err == pgx.ErrNoRows {
		return nil
	}
	return err
}

// Get returns the object, if it exists from the store
func (s *storeImpl) Get(id string) (*storage.ProcessIndicator, bool, error) {
	defer metrics.SetPostgresOperationDurationTime(time.Now(), ops.Get, "ProcessIndicator")

	conn, release := s.acquireConn(ops.Get, "ProcessIndicator")
	defer release()

	row := conn.QueryRow(context.Background(), getStmt, id)
	var data []byte
	if err := row.Scan(&data); err != nil {
		return nil, false, nilNoRows(err)
	}

	msg := alloc()
	buf := bytes.NewBuffer(data)
	defer metrics.SetJSONPBOperationDurationTime(time.Now(), "Unmarshal", "ProcessIndicator")
	if err := jsonpb.Unmarshal(buf, msg); err != nil {
		return nil, false, err
	}
	return msg.(*storage.ProcessIndicator), true, nil
}

// GetMany returns the objects specified by the IDs or the index in the missing indices slice 
func (s *storeImpl) GetMany(ids []string) ([]*storage.ProcessIndicator, []int, error) {
	defer metrics.SetPostgresOperationDurationTime(time.Now(), ops.GetMany, "ProcessIndicator")

	conn, release := s.acquireConn(ops.GetMany, "ProcessIndicator")
	defer release()

	rows, err := conn.Query(context.Background(), getManyStmt, ids)
	if err != nil {
		if err == pgx.ErrNoRows {
			missingIndices := make([]int, 0, len(ids))
			for i := range ids {
				missingIndices = append(missingIndices, i)
			}
			return nil, missingIndices, nil
		}
		return nil, nil, err
	}
	defer rows.Close()
	elems := make([]*storage.ProcessIndicator, 0, len(ids))
	foundSet := set.NewStringSet()
	for rows.Next() {
		var data []byte
		if err := rows.Scan(&data); err != nil {
			return nil, nil, err
		}
		msg := alloc()
		buf := bytes.NewBuffer(data)
		t := time.Now()
		if err := jsonpb.Unmarshal(buf, msg); err != nil {
			return nil, nil, err
		}
		metrics.SetJSONPBOperationDurationTime(t, "Unmarshal", "ProcessIndicator")
		elem := msg.(*storage.ProcessIndicator)
		foundSet.Add(elem.GetId())
		elems = append(elems, elem)
	}
	missingIndices := make([]int, 0, len(ids)-len(foundSet))
	for i, id := range ids {
		if !foundSet.Contains(id) {
			missingIndices = append(missingIndices, i)
		}
	}
	return elems, missingIndices, nil
}

func convertEnumSliceToIntArray(i interface{}) []int32 {
	enumSlice := reflect.ValueOf(i)
	enumSliceLen := enumSlice.Len()
	resultSlice := make([]int32, 0, enumSliceLen)
	for i := 0; i < enumSlice.Len(); i++ {
		resultSlice = append(resultSlice, int32(enumSlice.Index(i).Int()))
	}
	return resultSlice
}

func nilOrStringTimestamp(t *types.Timestamp) *string {
  if t == nil {
    return nil
  }
  s := t.String()
  return &s
}

func (s *storeImpl) upsert(id string, obj0 *storage.ProcessIndicator) error {
	t := time.Now()
	serialized, err := marshaler.MarshalToString(obj0)
	if err != nil {
		return err
	}
	metrics.SetJSONPBOperationDurationTime(t, "Marshal", "ProcessIndicator")
	conn, release := s.acquireConn(ops.Add, "ProcessIndicator")
	defer release()

	tx, err := conn.BeginTx(context.Background(), pgx.TxOptions{})
	if err != nil {
		return err
	}
    doRollback := true
	defer func() {
		if doRollback {
			if rollbackErr := tx.Rollback(context.Background()); rollbackErr != nil {
				log.Errorf("error rolling backing: %v", err)
			}
		}
	}()

	localQuery := "insert into ProcessIndicator(serialized, Id, DeploymentId, ContainerName, PodId, PodUid, ClusterId, Namespace, Signal_ContainerId, Signal_Name, Signal_Args, Signal_ExecFilePath, Signal_Uid) values($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13) on conflict(Id) do update set serialized = EXCLUDED.serialized, Id = EXCLUDED.Id, DeploymentId = EXCLUDED.DeploymentId, ContainerName = EXCLUDED.ContainerName, PodId = EXCLUDED.PodId, PodUid = EXCLUDED.PodUid, ClusterId = EXCLUDED.ClusterId, Namespace = EXCLUDED.Namespace, Signal_ContainerId = EXCLUDED.Signal_ContainerId, Signal_Name = EXCLUDED.Signal_Name, Signal_Args = EXCLUDED.Signal_Args, Signal_ExecFilePath = EXCLUDED.Signal_ExecFilePath, Signal_Uid = EXCLUDED.Signal_Uid"
	_, err = tx.Exec(context.Background(),localQuery, serialized, obj0.GetId(), obj0.GetDeploymentId(), obj0.GetContainerName(), obj0.GetPodId(), obj0.GetPodUid(), obj0.GetClusterId(), obj0.GetNamespace(), obj0.GetSignal().GetContainerId(), obj0.GetSignal().GetName(), obj0.GetSignal().GetArgs(), obj0.GetSignal().GetExecFilePath(), obj0.GetSignal().GetUid())
	if err != nil {
    	return err
  	}

    doRollback = false
	return tx.Commit(context.Background())
}

// Upsert inserts the object into the DB
func (s *storeImpl) Upsert(obj *storage.ProcessIndicator) error {
	defer metrics.SetPostgresOperationDurationTime(time.Now(), ops.Add, "ProcessIndicator")
	return s.upsert(keyFunc(obj), obj)
}

func (s *storeImpl) acquireConn(op ops.Op, typ string) (*pgxpool.Conn, func()) {
	defer metrics.SetAcquireDuration(time.Now(), op, typ)
	conn, err := s.db.Acquire(context.Background())
	if err != nil {
		panic(err)
	}
	return conn, conn.Release
}

// UpsertMany batches objects into the DB
func (s *storeImpl) UpsertMany(objs []*storage.ProcessIndicator) error {
	log.Infof("Upserting: %d indicators in batch", len(objs))
	if len(objs) == 0 {
		return nil
	}

	batch := &pgx.Batch{}
	defer func(now time.Time) {
		log.Infof("Upserting: %d indicators in batch - %d ms", len(objs), time.Since(now).Milliseconds())
	}(time.Now())
	defer metrics.SetPostgresOperationDurationTime(time.Now(), ops.AddMany, "ProcessIndicator")
	for _, obj0 := range objs {
		t := time.Now()
		serialized, err := marshaler.MarshalToString(obj0)
		if err != nil {
			return err
		}
		metrics.SetJSONPBOperationDurationTime(t, "Marshal", "ProcessIndicator")
		localQuery := "insert into ProcessIndicator(serialized, Id, DeploymentId, ContainerName, PodId, PodUid, ClusterId, Namespace, Signal_ContainerId, Signal_Name, Signal_Args, Signal_ExecFilePath, Signal_Uid) values($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13) on conflict(Id) do update set serialized = EXCLUDED.serialized, Id = EXCLUDED.Id, DeploymentId = EXCLUDED.DeploymentId, ContainerName = EXCLUDED.ContainerName, PodId = EXCLUDED.PodId, PodUid = EXCLUDED.PodUid, ClusterId = EXCLUDED.ClusterId, Namespace = EXCLUDED.Namespace, Signal_ContainerId = EXCLUDED.Signal_ContainerId, Signal_Name = EXCLUDED.Signal_Name, Signal_Args = EXCLUDED.Signal_Args, Signal_ExecFilePath = EXCLUDED.Signal_ExecFilePath, Signal_Uid = EXCLUDED.Signal_Uid"
		batch.Queue(localQuery, serialized, obj0.GetId(), obj0.GetDeploymentId(), obj0.GetContainerName(), obj0.GetPodId(), obj0.GetPodUid(), obj0.GetClusterId(), obj0.GetNamespace(), obj0.GetSignal().GetContainerId(), obj0.GetSignal().GetName(), obj0.GetSignal().GetArgs(), obj0.GetSignal().GetExecFilePath(), obj0.GetSignal().GetUid())
	}

	conn, release := s.acquireConn(ops.AddMany, "ProcessIndicator")
	defer release()

	results := conn.SendBatch(context.Background(), batch)
	if err := results.Close(); err != nil {
		return err
	}
	return nil
}

// Delete removes the specified ID from the store
func (s *storeImpl) Delete(id string) error {
	defer metrics.SetPostgresOperationDurationTime(time.Now(), ops.Remove, "ProcessIndicator")

	conn, release := s.acquireConn(ops.Remove, "ProcessIndicator")
	defer release()

	if _, err := conn.Exec(context.Background(), deleteStmt, id); err != nil {
		return err
	}
	return nil
}

// Delete removes the specified IDs from the store
func (s *storeImpl) DeleteMany(ids []string) error {
	defer metrics.SetPostgresOperationDurationTime(time.Now(), ops.RemoveMany, "ProcessIndicator")

	conn, release := s.acquireConn(ops.RemoveMany, "ProcessIndicator")
	defer release()
	if _, err := conn.Exec(context.Background(), deleteManyStmt, ids); err != nil {
		return err
	}
	return nil
}

// Walk iterates over all of the objects in the store and applies the closure
func (s *storeImpl) Walk(fn func(obj *storage.ProcessIndicator) error) error {
	rows, err := s.db.Query(context.Background(), walkStmt)
	if err != nil {
		return nilNoRows(err)
	}
	defer rows.Close()
	for rows.Next() {
		var data []byte
		if err := rows.Scan(&data); err != nil {
			return err
		}
		msg := alloc()
		buf := bytes.NewReader(data)
		if err := jsonpb.Unmarshal(buf, msg); err != nil {
			return err
		}
		return fn(msg.(*storage.ProcessIndicator))
	}
	return nil
}

// AckKeysIndexed acknowledges the passed keys were indexed
func (s *storeImpl) AckKeysIndexed(keys ...string) error {
	return nil
}

// GetKeysToIndex returns the keys that need to be indexed
func (s *storeImpl) GetKeysToIndex() ([]string, error) {
	return nil, nil
}
