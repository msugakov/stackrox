// Code generated by notifier generator. DO NOT EDIT.

package notifier

import storage "github.com/stackrox/stackrox/generated/storage"

type Notifier interface {
	OnDelete(onDelete func(deployment *storage.Deployment))
	Deleted(deployment *storage.Deployment)
	OnUpsert(onUpsert func(deployment *storage.Deployment))
	Upserted(deployment *storage.Deployment)
}

func New() Notifier {
	return newNotifier()
}
