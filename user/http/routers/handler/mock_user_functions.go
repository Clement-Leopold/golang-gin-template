package handler

import (
	"backend-test-chenxianhao/user-management/domains"
	"context"

	"github.com/stretchr/testify/mock"
)

// MockUserFunctions is a mock of interface
type MockUserFunctions struct {
	mock.Mock
}

func (u *MockUserFunctions) GetByID(ctx context.Context, id string) (domains.User, error) {
	ret := u.Called(ctx, id)

	var r0 domains.User
	if ret.Get(0) != nil {
		r0 = ret.Get(0).(domains.User)
	}

	var r1 error

	if ret.Get(1) != nil {
		r1 = ret.Get(1).(error)
	}

	return r0, r1
}

func (u *MockUserFunctions) Update(context.Context, *domains.User) error {
	return nil
}
func (u *MockUserFunctions) Get(ctx context.Context, limit int16, offset int16) ([]domains.User, error) {
	return nil, nil
}
func (u *MockUserFunctions) Create(context.Context, *domains.User) error {
	return nil
}
func (u *MockUserFunctions) Delete(ctx context.Context, id string) error {
	return nil
}

func (u *MockUserFunctions) Following(ctx context.Context, id string, followerId string) error {
	return nil
}
func (u *MockUserFunctions) UnFollowing(ctx context.Context, id string, followingId string) error {
	return nil
}
func (u *MockUserFunctions) GetFollowers(ctx context.Context, id string, limit int16, offset int16) ([]domains.Follower, error) {
	return nil, nil
}

func (u *MockUserFunctions) GetFollowings(ctx context.Context, id string, limit int16, offset int16) ([]domains.Follower, error) {
	return nil, nil
}

func (u *MockUserFunctions) GetMinimumDistanceForFollowing(ctx context.Context, name string) (domains.Follower, error) {
	return domains.Follower{}, nil
}
