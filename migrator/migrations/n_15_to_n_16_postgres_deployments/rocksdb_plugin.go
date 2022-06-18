package n15ton16

import (
	"context"

	"github.com/gogo/protobuf/proto"
	"github.com/stackrox/rox/generated/storage"
)

func alloc() proto.Message {
	return &storage.Deployment{}
}

func keyFunc(msg proto.Message) []byte {
	return []byte(msg.(*storage.Deployment).GetId())
}

// Walk iterates over all of the objects in the store and applies the closure
func (b *storeImpl) Walk(_ context.Context, fn func(obj *storage.Deployment) error) error {
	return b.crud.Walk(func(msg proto.Message) error {
		return fn(msg.(*storage.Deployment))
	})
}
