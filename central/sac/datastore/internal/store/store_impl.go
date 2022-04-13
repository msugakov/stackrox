// Code generated by boltbindings generator. DO NOT EDIT.

package store

import (
	proto "github.com/gogo/protobuf/proto"
	metrics "github.com/stackrox/stackrox/central/metrics"
	storage "github.com/stackrox/stackrox/generated/storage"
	protoCrud "github.com/stackrox/stackrox/pkg/bolthelper/crud/proto"
	ops "github.com/stackrox/stackrox/pkg/metrics"
	bbolt "go.etcd.io/bbolt"
	"time"
)

var (
	bucketName = []byte("authzPlugins")
)

type store struct {
	crud protoCrud.MessageCrud
}

func key(msg proto.Message) []byte {
	return []byte(msg.(*storage.AuthzPluginConfig).GetId())
}

func alloc() proto.Message {
	return new(storage.AuthzPluginConfig)
}

func newStore(db *bbolt.DB) (*store, error) {
	newCrud, err := protoCrud.NewMessageCrud(db, bucketName, key, alloc)
	if err != nil {
		return nil, err
	}
	return &store{crud: newCrud}, nil
}

func (s *store) DeleteAuthzPluginConfig(id string) error {
	defer metrics.SetBoltOperationDurationTime(time.Now(), ops.Remove, "AuthzPluginConfig")
	_, _, err := s.crud.Delete(id)
	return err
}

func (s *store) GetAuthzPluginConfig(id string) (*storage.AuthzPluginConfig, error) {
	defer metrics.SetBoltOperationDurationTime(time.Now(), ops.Get, "AuthzPluginConfig")
	msg, err := s.crud.Read(id)
	if err != nil {
		return nil, err
	}
	if msg == nil {
		return nil, nil
	}
	authzpluginconfig := msg.(*storage.AuthzPluginConfig)
	return authzpluginconfig, nil
}

func (s *store) ListAuthzPluginConfigs() ([]*storage.AuthzPluginConfig, error) {
	defer metrics.SetBoltOperationDurationTime(time.Now(), ops.GetAll, "AuthzPluginConfig")
	msgs, err := s.crud.ReadAll()
	if err != nil {
		return nil, err
	}
	storedKeys := make([]*storage.AuthzPluginConfig, len(msgs))
	for i, msg := range msgs {
		storedKeys[i] = msg.(*storage.AuthzPluginConfig)
	}
	return storedKeys, nil
}

func (s *store) UpsertAuthzPluginConfig(authzpluginconfig *storage.AuthzPluginConfig) error {
	defer metrics.SetBoltOperationDurationTime(time.Now(), ops.Upsert, "AuthzPluginConfig")
	_, _, err := s.crud.Upsert(authzpluginconfig)
	return err
}

func (s *store) UpsertAuthzPluginConfigs(authzpluginconfigs []*storage.AuthzPluginConfig) error {
	msgs := make([]proto.Message, len(authzpluginconfigs))
	for i, key := range authzpluginconfigs {
		msgs[i] = key
	}
	_, _, err := s.crud.UpsertBatch(msgs)
	return err
}
