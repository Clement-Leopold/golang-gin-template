package impl

import (
	"backend-test-chenxianhao/user-management/common"
	"backend-test-chenxianhao/user-management/domains"
	"backend-test-chenxianhao/user-management/util"
	"context"
	"errors"

	"github.com/op/go-logging"
)

var log = logging.MustGetLogger("user_impl")

type userImpl struct {
	userRepo domains.UserRepository
}

func UserMethodImpl(ur domains.UserRepository) domains.UserFunctions {
	return &userImpl{
		userRepo: ur,
	}
}

func (u *userImpl) GetByID(ctx context.Context, id string) (domains.User, error) {
	return u.userRepo.GetByID(ctx, id)
}

func (u *userImpl) Create(ctx context.Context, user *domains.User) error {
	return u.userRepo.Create(ctx, user)
}

func (u *userImpl) Delete(ctx context.Context, id string) error {
	origin, err := u.userRepo.GetByID(ctx, id)
	if err != nil {
		log.Error(err)
		return common.DatabaseError(err)
	}
	if origin.Id == "" {
		return common.ParamError(errors.New("id is not exist"))
	}
	return u.userRepo.Delete(ctx, id)
}

func (u *userImpl) Update(ctx context.Context, user *domains.User) error {
	origin, err := u.userRepo.GetByID(ctx, user.Id)
	if err != nil {
		log.Error(err)
		return common.DatabaseError(err)
	}
	if origin.Id == "" {
		return common.ParamError(errors.New("id is not exist"))
	}
	// use go routine for update followers distance. or use cron job.
	go func() {
		followers, err := u.userRepo.SelectAllFollowing(context.Background(), user.Id)
		if err != nil {
			log.Error(err)
			return
		}
		if len(followers) != 0 {
			for i := 0; i < len(followers); i++ {
				distance := util.GetDistance(user.XCoordinate, user.YCoordinate, followers[i].XCoordinate, followers[i].YCoordinate)
				followers[i].Distance = distance
			}
			err = u.userRepo.BatchUpdateForFollowerDistance(context.Background(), followers)
			if err != nil {
				log.Error(err)
				return
			}
		}
	}()
	return u.userRepo.Update(ctx, user)
}

func (u *userImpl) Get(ctx context.Context, limit int16, offset int16) ([]domains.User, error) {
	return u.userRepo.Get(ctx, limit, offset)
}

func (u *userImpl) Following(ctx context.Context, id string, followingId string) error {
	coordinates, err := u.userRepo.SelectTwoCoordinates(ctx, id, followingId)
	if err != nil {
		return err
	}
	distance := util.GetDistance(coordinates[0], coordinates[1], coordinates[2], coordinates[3])
	follower := &domains.Follower{
		Id:       followingId,
		Distance: distance,
	}
	return u.userRepo.Following(ctx, id, follower)
}

func (u *userImpl) UnFollowing(ctx context.Context, id string, followingID string) error {
	return u.userRepo.UnFollowing(ctx, id, followingID)
}

func (u *userImpl) GetFollowers(ctx context.Context, id string, limit int16, offset int16) ([]domains.Follower, error) {
	return u.userRepo.GetFollowers(ctx, id, limit, offset)
}

func (u *userImpl) GetFollowings(ctx context.Context, id string, limit int16, offset int16) ([]domains.Follower, error) {
	return u.userRepo.GetFollowings(ctx, id, limit, offset)
}

func (u *userImpl) GetMinimumDistanceForFollowing(ctx context.Context, name string) (domains.Follower, error) {
	return u.userRepo.GetMinimumDistanceForFollowing(ctx, name)
}
