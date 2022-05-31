package domains

import (
	"context"
	"time"
)

type User struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	// date of birth  use string to avoid BindJson problem.
	Dob         string `json:"dob" `
	Address     string `json:"address"`
	Description string `json:"description"`
	// database create time
	CreatedAt time.Time
}

// UserMethod user's use case, standard CRUD
type UseMethod interface {
	// get user by id
	GetByID(ctx context.Context, id string) (User, error)
	Update(ctx context.Context, u *User) error
	Create(ctx context.Context, u *User) error
	Delete(ctx context.Context, id string) error
	// get a partial of users
	Get(ctx context.Context, limit int16, offset int16) ([]User, error)
}

// UserRepository represent the user's repository method, standard CRUD
type UserRepository interface {
	GetByID(ctx context.Context, id string) (User, error)
	Update(ctx context.Context, u *User) error
	Create(ctx context.Context, u *User) error
	Delete(ctx context.Context, id string) error
	Get(ctx context.Context, numbers int16, offset int16) ([]User, error)
}
