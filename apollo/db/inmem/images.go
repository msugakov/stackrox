package inmem

import (
	"fmt"
	"sort"
	"sync"

	"bitbucket.org/stack-rox/apollo/apollo/db"
	"bitbucket.org/stack-rox/apollo/pkg/api/generated/api/v1"
	"github.com/golang/protobuf/proto"
)

type imageStore struct {
	images     map[string]*v1.Image
	imageMutex sync.Mutex

	persistent db.ImageStorage
}

func newImageStore(persistent db.ImageStorage) *imageStore {
	return &imageStore{
		images:     make(map[string]*v1.Image),
		persistent: persistent,
	}
}

func (s *imageStore) clone(image *v1.Image) *v1.Image {
	return proto.Clone(image).(*v1.Image)
}

func (s *imageStore) loadFromPersistent() error {
	s.imageMutex.Lock()
	defer s.imageMutex.Unlock()
	images, err := s.persistent.GetImages(&v1.GetImagesRequest{})
	if err != nil {
		return err
	}
	for _, image := range images {
		s.images[image.Sha] = image
	}
	return nil
}

// GetImages returns all images
func (s *imageStore) GetImages(request *v1.GetImagesRequest) ([]*v1.Image, error) {
	s.imageMutex.Lock()
	defer s.imageMutex.Unlock()
	images := make([]*v1.Image, 0, len(s.images))
	for _, image := range s.images {
		images = append(images, s.clone(image))
	}
	sort.SliceStable(images, func(i, j int) bool { return images[i].Sha < images[j].Sha })
	return images, nil
}

func (s *imageStore) insertImage(image *v1.Image) {
	s.images[image.Sha] = s.clone(image)
}

// AddImage adds an image to the database
func (s *imageStore) AddImage(image *v1.Image) error {
	s.imageMutex.Lock()
	defer s.imageMutex.Unlock()
	if _, ok := s.images[image.Sha]; ok {
		return fmt.Errorf("Cannot add image %v because it already exists", image.Sha)
	}
	if err := s.persistent.AddImage(image); err != nil {
		return err
	}
	s.insertImage(image)
	return nil
}

// UpdateImage updates an image
func (s *imageStore) UpdateImage(image *v1.Image) error {
	s.imageMutex.Lock()
	defer s.imageMutex.Unlock()
	if err := s.persistent.UpdateImage(image); err != nil {
		return err
	}
	s.insertImage(image)
	return nil
}

// RemoveImage removes a specific image specified by it's SHA
func (s *imageStore) RemoveImage(sha string) error {
	s.imageMutex.Lock()
	defer s.imageMutex.Unlock()
	if err := s.persistent.RemoveImage(sha); err != nil {
		return err
	}
	delete(s.images, sha)
	return nil
}
