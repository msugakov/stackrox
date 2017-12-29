package inmem

import (
	"fmt"
	"sort"
	"sync"

	"bitbucket.org/stack-rox/apollo/apollo/db"
	"bitbucket.org/stack-rox/apollo/pkg/api/generated/api/v1"
	"github.com/golang/protobuf/proto"
)

type imagePolicyStore struct {
	imagePolicies      map[string]*v1.ImagePolicy
	imagePoliciesMutex sync.Mutex

	persistent db.ImagePolicyStorage
}

func newImagePolicyStore(persistent db.ImagePolicyStorage) *imagePolicyStore {
	return &imagePolicyStore{
		imagePolicies: make(map[string]*v1.ImagePolicy),
		persistent:    persistent,
	}
}

func (s *imagePolicyStore) clone(policy *v1.ImagePolicy) *v1.ImagePolicy {
	return proto.Clone(policy).(*v1.ImagePolicy)
}

func (s *imagePolicyStore) loadFromPersistent() error {
	s.imagePoliciesMutex.Lock()
	defer s.imagePoliciesMutex.Unlock()
	policies, err := s.persistent.GetImagePolicies(&v1.GetImagePoliciesRequest{})
	if err != nil {
		return err
	}
	for _, p := range policies {
		s.imagePolicies[p.Name] = p
	}
	return nil
}

// GetImagePolicies returns image policies according to request.
func (s *imagePolicyStore) GetImagePolicies(request *v1.GetImagePoliciesRequest) ([]*v1.ImagePolicy, error) {
	s.imagePoliciesMutex.Lock()
	defer s.imagePoliciesMutex.Unlock()
	policies := make([]*v1.ImagePolicy, 0, len(s.imagePolicies))
	for _, v := range s.imagePolicies {
		policies = append(policies, s.clone(v))
	}
	if request.Name != "" {
		val, ok := s.imagePolicies[request.Name]
		if ok {
			return []*v1.ImagePolicy{s.clone(val)}, nil
		}
		return policies, nil
	}
	sort.SliceStable(policies, func(i, j int) bool { return policies[i].Name < policies[j].Name })
	return policies, nil
}

// AddImagePolicy adds the image policy to the database.
func (s *imagePolicyStore) AddImagePolicy(policy *v1.ImagePolicy) error {
	s.imagePoliciesMutex.Lock()
	defer s.imagePoliciesMutex.Unlock()
	if _, ok := s.imagePolicies[policy.Name]; ok {
		return fmt.Errorf("policy with name %v already exists and cannot be added again", policy.Name)
	}
	if err := s.persistent.AddImagePolicy(policy); err != nil {
		return err
	}
	s.upsertImagePolicy(policy)
	return nil
}

// UpdateImagePolicy updates the image policy.
func (s *imagePolicyStore) UpdateImagePolicy(policy *v1.ImagePolicy) error {
	s.imagePoliciesMutex.Lock()
	defer s.imagePoliciesMutex.Unlock()
	if err := s.persistent.UpdateImagePolicy(policy); err != nil {
		return err
	}
	s.upsertImagePolicy(policy)
	return nil
}

func (s *imagePolicyStore) upsertImagePolicy(policy *v1.ImagePolicy) {
	s.imagePolicies[policy.Name] = s.clone(policy)
}

// RemoveImagePolicy removes the image policy.
func (s *imagePolicyStore) RemoveImagePolicy(name string) error {
	s.imagePoliciesMutex.Lock()
	defer s.imagePoliciesMutex.Unlock()
	if err := s.persistent.RemoveImagePolicy(name); err != nil {
		return err
	}
	delete(s.imagePolicies, name)
	return nil
}
