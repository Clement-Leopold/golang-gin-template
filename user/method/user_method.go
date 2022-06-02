package method

import (
	"backend-test-chenxianhao/user-management/domains"
	"context"
	"time"

	"github.com/google/uuid"
)

type userMethod struct {
	userRepo domains.UserRepository
}

func UserMethodImpl(ur domains.UserRepository) domains.UserMethod {
	return &userMethod{
		userRepo: ur,
	}
}

func (u *userMethod) GetByID(ctx context.Context, id string) (domains.User, error) {
	return u.userRepo.GetByID(ctx, id)
}

func (u *userMethod) Create(ctx context.Context, user *domains.User) error {
	user.Id = uuid.New().String()
	user.CreatedAt = time.Now()
	return u.userRepo.Create(ctx, user)
}

func (u *userMethod) Delete(ctx context.Context, id string) error {
	return u.userRepo.Delete(ctx, id)
}

func (u *userMethod) Update(ctx context.Context, user *domains.User) error {
	user.CreatedAt = time.Now()
	return u.userRepo.Update(ctx, user)
}

func (u *userMethod) Get(ctx context.Context, limit int16, offset int16) ([]domains.User, error) {
	return u.userRepo.Get(ctx, limit, offset)
}

func (u *userMethod) Following(ctx context.Context, id string, followingId string) error {
	return u.userRepo.Following(ctx, id, followingId)
}

func (u *userMethod) UnFollowing(ctx context.Context, id string, followingID string) error {
	return u.userRepo.UnFollowing(ctx, id, followingID)
}

func (u *userMethod) GetFollowers(ctx context.Context, id string, limit int16, offset int16) ([]domains.Follower, error) {
	return u.userRepo.GetFollowers(ctx, id, limit, offset)
}

func (u *userMethod) GetFollowings(ctx context.Context, id string, limit int16, offset int16) ([]domains.Follower, error) {
	return u.userRepo.GetFollowings(ctx, id, limit, offset)
}
