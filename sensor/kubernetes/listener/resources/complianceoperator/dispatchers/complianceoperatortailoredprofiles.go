package dispatchers

import (
	"github.com/stackrox/stackrox/generated/internalapi/central"
	"github.com/stackrox/stackrox/generated/storage"
	"github.com/stackrox/stackrox/pkg/complianceoperator/api/v1alpha1"
	"github.com/stackrox/stackrox/pkg/set"
	"github.com/stackrox/stackrox/pkg/stringutils"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/tools/cache"
)

// TailoredProfileDispatcher handles compliance operator tailored profile objects
type TailoredProfileDispatcher struct {
	profileLister cache.GenericLister
}

// NewTailoredProfileDispatcher creates and returns a new tailored profile dispatcher
func NewTailoredProfileDispatcher(profileLister cache.GenericLister) *TailoredProfileDispatcher {
	return &TailoredProfileDispatcher{
		profileLister: profileLister,
	}
}

// ProcessEvent processes a compliance operator tailored profile
func (c *TailoredProfileDispatcher) ProcessEvent(obj, _ interface{}, action central.ResourceAction) []*central.SensorEvent {
	var tailoredProfile v1alpha1.TailoredProfile

	unstructuredObject, ok := obj.(*unstructured.Unstructured)
	if !ok {
		log.Errorf("Not of type 'unstructured': %T", obj)
		return nil
	}

	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(unstructuredObject.Object, &tailoredProfile); err != nil {
		log.Errorf("error converting unstructured to tailored compliance profile: %v", err)
		return nil
	}

	if tailoredProfile.Status.ID == "" {
		log.Warnf("Tailored profile %s does not have an ID. Skipping...", tailoredProfile.Name)
		return nil
	}

	profileObj, err := c.profileLister.ByNamespace(tailoredProfile.GetNamespace()).Get(tailoredProfile.Spec.Extends)
	if err != nil {
		log.Errorf("error getting profile %s: %v", tailoredProfile.Spec.Extends, err)
		return nil
	}
	unstructuredObject, ok = profileObj.(*unstructured.Unstructured)
	if !ok {
		log.Errorf("Fetched profile not of type 'unstructured': %T", obj)
		return nil
	}

	var complianceProfile v1alpha1.Profile
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(unstructuredObject.Object, &complianceProfile); err != nil {
		log.Errorf("error converting unstructured to compliance profile: %v", err)
		return nil
	}

	protoProfile := &storage.ComplianceOperatorProfile{
		Id:        string(tailoredProfile.UID),
		ProfileId: tailoredProfile.Status.ID,
		Name:      tailoredProfile.Name,
		// We want to use the original compliance profiles labels and annotations as they hold data about the type of profile
		Labels:      complianceProfile.Labels,
		Annotations: complianceProfile.Annotations,
		Description: stringutils.FirstNonEmpty(tailoredProfile.Spec.Description, complianceProfile.Description),
	}
	removedRules := set.NewStringSet()
	for _, rule := range tailoredProfile.Spec.DisableRules {
		removedRules.Add(rule.Name)
	}

	for _, r := range complianceProfile.Rules {
		if removedRules.Contains(string(r)) {
			continue
		}
		protoProfile.Rules = append(protoProfile.Rules, &storage.ComplianceOperatorProfile_Rule{
			Name: string(r),
		})
	}
	for _, rule := range tailoredProfile.Spec.EnableRules {
		protoProfile.Rules = append(protoProfile.Rules, &storage.ComplianceOperatorProfile_Rule{
			Name: rule.Name,
		})
	}

	return []*central.SensorEvent{
		{
			Id:     protoProfile.GetId(),
			Action: action,
			Resource: &central.SensorEvent_ComplianceOperatorProfile{
				ComplianceOperatorProfile: protoProfile,
			},
		},
	}
}
