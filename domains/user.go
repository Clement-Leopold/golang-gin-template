package domains

import (
	"context"
	"time"
)

// User: in database called account;
// To simplify the spatial question I used X-Y coordinates.
// and use async func to do the nearest math when following.
// And the alternative:
// 1. POSTGIS extensions for postgres to do the spatial work.
// 2. Cache the queue for following and distance to do the math.
type User struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	// date of birth  use string to avoid BindJson problem.
	Dob     string `json:"dob" `
	Address string `json:"address"`
	// for now it may be appropriate to just let coordinates in User structure.
	XCoordinate int64  `json:"x_coordinate"`
	YCoordinate int64  `json:"y_coordinate"`
	Description string `json:"description"`
	// database create time
	CreatedAt time.Time `json:"-"`
}

// Follower: several attributes of User used in Follower API.
type Follower struct {
	AutoId      int     `json:"-"`
	Id          string  `json:"id"`
	Name        string  `json:"name"`
	XCoordinate int64   `json:"-"`
	YCoordinate int64   `json:"-"`
	Distance    float64 `json:"-"`
}

// UserFunctions user's use case, standard CRUD
type UserFunctions interface {
	// get user by id
	GetByID(ctx context.Context, id string) (User, error)
	// Update: update user info, and use go routine to update the distance of followers.
	Update(ctx context.Context, u *User) error
	Create(ctx context.Context, u *User) error
	Delete(ctx context.Context, id string) error
	// get a partial of users
	Get(ctx context.Context, limit int16, offset int16) ([]User, error)
	//Following: following
	Following(ctx context.Context, id string, followerId string) error
	//UnFollowing: not following
	UnFollowing(ctx context.Context, id string, followingId string) error
	//GetFollowers: get followers
	GetFollowers(ctx context.Context, id string, limit int16, offset int16) ([]Follower, error)
	//GetFollowings: get followings
	GetFollowings(ctx context.Context, id string, limit int16, offset int16) ([]Follower, error)
	// GetMinimumDistanceForFollowing: find the following user with minimum distance.
	GetMinimumDistanceForFollowing(ctx context.Context, name string) (Follower, error)
}

// UserRepository represent the user's repository method, standard CRUD
type UserRepository interface {
	GetByID(ctx context.Context, id string) (User, error)
	Update(ctx context.Context, u *User) error
	Create(ctx context.Context, u *User) error
	Delete(ctx context.Context, id string) error
	Get(ctx context.Context, numbers int16, offset int16) ([]User, error)
	// search two users coordinates.
	SelectTwoCoordinates(ctx context.Context, idOne string, idTwo string) ([]int64, error)
	//Following: following
	Following(ctx context.Context, id string, follower *Follower) error
	//UnFollowing: not following
	UnFollowing(ctx context.Context, id string, followingId string) error
	//GetFollowers: get followers
	GetFollowers(ctx context.Context, id string, limit int16, offset int16) ([]Follower, error)
	//GetFollowings: get followings
	GetFollowings(ctx context.Context, id string, limit int16, offset int16) ([]Follower, error)
	//BatchUpdateForFollowerDistance: batch update all distance.
	BatchUpdateForFollowerDistance(ctx context.Context, followers []Follower) error
	//SelectAllFollowing: select all following the id.
	SelectAllFollowing(ctx context.Context, uId string) ([]Follower, error)
	// GetMinimumDistanceForFollowing: find the following user with minimum distance.
	GetMinimumDistanceForFollowing(ctx context.Context, name string) (Follower, error)
}
