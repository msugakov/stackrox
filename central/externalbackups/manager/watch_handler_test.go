package manager

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stackrox/stackrox/central/externalbackups/manager/mocks"
	"github.com/stackrox/stackrox/generated/storage"
	"github.com/stretchr/testify/suite"
)

func TestWatchHandler(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(WatchHandlerTestSuite))
}

type WatchHandlerTestSuite struct {
	suite.Suite

	mockCtrl *gomock.Controller
	mgr      *mocks.MockManager
}

func (s *WatchHandlerTestSuite) SetupTest() {
	s.mockCtrl = gomock.NewController(s.T())
	s.mgr = mocks.NewMockManager(s.mockCtrl)
}

func (s *WatchHandlerTestSuite) TearDownTest() {
	s.mockCtrl.Finish()
}

func (s *WatchHandlerTestSuite) TestOnStableUpdate() {
	watch := &watchHandler{
		mgr:  s.mgr,
		id:   "wooo",
		file: "test_data",
	}
	testGcsConfig := &storage.ExternalBackup{
		Name: "test",
		Type: "gcs",
	}

	s.mgr.EXPECT().Upsert(gomock.Any(), testGcsConfig).Return(nil)
	watch.OnStableUpdate(testGcsConfig, nil)
	s.NotEmpty(testGcsConfig.GetId())
}

func (s *WatchHandlerTestSuite) TestOnChange() {
	watch := &watchHandler{
		mgr:  s.mgr,
		id:   "wooo",
		file: "test_data",
	}

	testConfig, err := watch.OnChange("./testdata")
	s.NoError(err)
	s.NotNil(testConfig)
}
