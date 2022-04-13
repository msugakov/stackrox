package service

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	dsMocks "github.com/stackrox/stackrox/central/user/datastore/mocks"
	v1 "github.com/stackrox/stackrox/generated/api/v1"
	"github.com/stackrox/stackrox/generated/storage"
	"github.com/stretchr/testify/suite"
)

func TestUserService(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(UserServiceTestSuite))
}

type UserServiceTestSuite struct {
	suite.Suite

	mockCtrl *gomock.Controller

	usersMock *dsMocks.MockDataStore
	ser       Service
}

func (suite *UserServiceTestSuite) SetupSuite() {
	suite.mockCtrl = gomock.NewController(suite.T())

	suite.usersMock = dsMocks.NewMockDataStore(suite.mockCtrl)
	suite.ser = New(suite.usersMock)
}

func (suite *UserServiceTestSuite) TestGetUsersAttributes() {
	expectedContext := context.Background()
	users := []*storage.User{
		{
			Id:             "user1",
			AuthProviderId: "ap1",
			Attributes: []*storage.UserAttribute{
				{
					Key:   "name",
					Value: "user1",
				},
				{
					Key:   "email",
					Value: "user@derp.com",
				},
			},
		},
		{
			Id:             "user2",
			AuthProviderId: "ap1",
			Attributes: []*storage.UserAttribute{
				{
					Key:   "name",
					Value: "user2",
				},
			},
		},
		{
			Id:             "user3",
			AuthProviderId: "ap2",
			Attributes: []*storage.UserAttribute{
				{
					Key:   "name",
					Value: "user1",
				},
			},
		},
		{
			Id:             "user4",
			AuthProviderId: "ap1",
			Attributes: []*storage.UserAttribute{
				{
					Key:   "name",
					Value: "user1",
				},
				{
					Key:   "email",
					Value: "user@derp.com",
				},
			},
		},
	}

	expectedAttributes := []*v1.UserAttributeTuple{
		{
			AuthProviderId: "ap1",
			Key:            "name",
			Value:          "user1",
		},
		{
			AuthProviderId: "ap1",
			Key:            "email",
			Value:          "user@derp.com",
		},
		{
			AuthProviderId: "ap1",
			Key:            "name",
			Value:          "user2",
		},
		{
			AuthProviderId: "ap2",
			Key:            "name",
			Value:          "user1",
		},
	}

	suite.usersMock.EXPECT().GetAllUsers(expectedContext).Return(users, nil)

	resp, err := suite.ser.GetUsersAttributes(expectedContext, nil)
	suite.NoError(err, "request should not fail with valid user data")

	suite.Equal(len(expectedAttributes), len(resp.GetUsersAttributes()))
}
