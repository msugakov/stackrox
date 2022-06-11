// Code generated by rocksdb-bindings generator. DO NOT EDIT.
package n30ton31
import (
	"context"

	"github.com/gogo/protobuf/proto"
	"github.com/stackrox/rox/generated/storage"
)

func alloc() proto.Message {
	return &storage.NetworkGraphConfig{}
}

func keyFunc(msg proto.Message) []byte {
	return []byte(msg.(*storage.NetworkGraphConfig).GetId())
}

// Walk iterates over all of the objects in the store and applies the closure
func (b *storeImpl) Walk(_ context.Context, fn func(obj *storage.NetworkGraphConfig) error) error {
	return b.crud.Walk(func(msg proto.Message) error {
		return fn(msg.(*storage.NetworkGraphConfig))
	})
}
