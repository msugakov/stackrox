package dackbox

import (
	"time"

	"github.com/gogo/protobuf/proto"
	vulnDackBox "github.com/stackrox/stackrox/central/cve/dackbox"
	"github.com/stackrox/stackrox/central/cve/store"
	"github.com/stackrox/stackrox/central/metrics"
	"github.com/stackrox/stackrox/generated/storage"
	"github.com/stackrox/stackrox/pkg/batcher"
	"github.com/stackrox/stackrox/pkg/concurrency"
	"github.com/stackrox/stackrox/pkg/dackbox"
	ops "github.com/stackrox/stackrox/pkg/metrics"
)

const batchSize = 100

type storeImpl struct {
	keyFence concurrency.KeyFence
	dacky    *dackbox.DackBox
}

// New returns a new Store instance.
func New(dacky *dackbox.DackBox, keyFence concurrency.KeyFence) store.Store {
	return &storeImpl{
		keyFence: keyFence,
		dacky:    dacky,
	}
}

func (b *storeImpl) Exists(id string) (bool, error) {
	dackTxn, err := b.dacky.NewReadOnlyTransaction()
	if err != nil {
		return false, err
	}
	defer dackTxn.Discard()

	exists, err := vulnDackBox.Reader.ExistsIn(vulnDackBox.BucketHandler.GetKey(id), dackTxn)
	if err != nil {
		return false, err
	}

	return exists, nil
}

func (b *storeImpl) Count() (int, error) {
	defer metrics.SetDackboxOperationDurationTime(time.Now(), ops.Count, "CVE")

	dackTxn, err := b.dacky.NewReadOnlyTransaction()
	if err != nil {
		return 0, err
	}
	defer dackTxn.Discard()

	count, err := vulnDackBox.Reader.CountIn(vulnDackBox.Bucket, dackTxn)
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (b *storeImpl) GetAll() ([]*storage.CVE, error) {
	defer metrics.SetDackboxOperationDurationTime(time.Now(), ops.GetAll, "CVE")

	dackTxn, err := b.dacky.NewReadOnlyTransaction()
	if err != nil {
		return nil, err
	}
	defer dackTxn.Discard()

	msgs, err := vulnDackBox.Reader.ReadAllIn(vulnDackBox.Bucket, dackTxn)
	if err != nil {
		return nil, err
	}
	ret := make([]*storage.CVE, 0, len(msgs))
	for _, msg := range msgs {
		ret = append(ret, msg.(*storage.CVE))
	}

	return ret, nil
}

func (b *storeImpl) Get(id string) (cve *storage.CVE, exists bool, err error) {
	defer metrics.SetDackboxOperationDurationTime(time.Now(), ops.Get, "CVE")

	dackTxn, err := b.dacky.NewReadOnlyTransaction()
	if err != nil {
		return nil, false, err
	}
	defer dackTxn.Discard()

	msg, err := vulnDackBox.Reader.ReadIn(vulnDackBox.BucketHandler.GetKey(id), dackTxn)
	if err != nil || msg == nil {
		return nil, false, err
	}

	return msg.(*storage.CVE), true, err
}

func (b *storeImpl) GetBatch(ids []string) ([]*storage.CVE, []int, error) {
	defer metrics.SetDackboxOperationDurationTime(time.Now(), ops.GetMany, "CVE")

	dackTxn, err := b.dacky.NewReadOnlyTransaction()
	if err != nil {
		return nil, nil, err
	}
	defer dackTxn.Discard()

	msgs := make([]proto.Message, 0, len(ids))
	missing := make([]int, 0, len(ids)/2)
	for idx, id := range ids {
		msg, err := vulnDackBox.Reader.ReadIn(vulnDackBox.BucketHandler.GetKey(id), dackTxn)
		if err != nil {
			return nil, nil, err
		}
		if msg != nil {
			msgs = append(msgs, msg)
		} else {
			missing = append(missing, idx)
		}
	}

	ret := make([]*storage.CVE, 0, len(msgs))
	for _, msg := range msgs {
		ret = append(ret, msg.(*storage.CVE))
	}

	return ret, missing, nil
}

func (b *storeImpl) Upsert(cves ...*storage.CVE) error {
	defer metrics.SetDackboxOperationDurationTime(time.Now(), ops.Upsert, "CVE")

	keysToUpsert := make([][]byte, 0, len(cves))
	for _, vuln := range cves {
		keysToUpsert = append(keysToUpsert, vulnDackBox.KeyFunc(vuln))
	}
	lockedKeySet := concurrency.DiscreteKeySet(keysToUpsert...)

	return b.keyFence.DoStatusWithLock(lockedKeySet, func() error {
		batch := batcher.New(len(cves), batchSize)
		for {
			start, end, ok := batch.Next()
			if !ok {
				break
			}

			if err := b.upsertNoBatch(cves[start:end]...); err != nil {
				return err
			}
		}
		return nil
	})
}

func (b *storeImpl) upsertNoBatch(cves ...*storage.CVE) error {
	dackTxn, err := b.dacky.NewTransaction()
	if err != nil {
		return err
	}
	defer dackTxn.Discard()

	for _, cve := range cves {
		err := vulnDackBox.Upserter.UpsertIn(nil, cve, dackTxn)
		if err != nil {
			return err
		}
	}

	if err := dackTxn.Commit(); err != nil {
		return err
	}
	return nil
}

func (b *storeImpl) Delete(ids ...string) error {
	defer metrics.SetDackboxOperationDurationTime(time.Now(), ops.RemoveMany, "CVE")

	keysToUpsert := make([][]byte, 0, len(ids))
	for _, id := range ids {
		keysToUpsert = append(keysToUpsert, vulnDackBox.BucketHandler.GetKey(id))
	}
	lockedKeySet := concurrency.DiscreteKeySet(keysToUpsert...)

	return b.keyFence.DoStatusWithLock(lockedKeySet, func() error {
		batch := batcher.New(len(ids), batchSize)
		for {
			start, end, ok := batch.Next()
			if !ok {
				break
			}

			if err := b.deleteNoBatch(ids[start:end]...); err != nil {
				return err
			}
		}
		return nil
	})
}

func (b *storeImpl) deleteNoBatch(ids ...string) error {
	dackTxn, err := b.dacky.NewTransaction()
	if err != nil {
		return err
	}
	defer dackTxn.Discard()

	for _, id := range ids {
		if err := vulnDackBox.Deleter.DeleteIn(vulnDackBox.BucketHandler.GetKey(id), dackTxn); err != nil {
			return err
		}
	}

	if err := dackTxn.Commit(); err != nil {
		return err
	}
	return nil
}
