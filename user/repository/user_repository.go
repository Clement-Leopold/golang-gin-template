package repository

import (
	"backend-test-chenxianhao/user-management/common"
	"backend-test-chenxianhao/user-management/domains"
	"context"
	"database/sql"
)

type userSQLRepository struct {
	DB *sql.DB
}

func NewUserRepositoryImpl(db *sql.DB) domains.UserRepository {
	return &userSQLRepository{
		DB: db,
	}
}

func (ur *userSQLRepository) GetByID(ctx context.Context, id string) (res domains.User, err error) {
	query := "select id, name, dob, address, description, created_at from t_accounts where id = $1"
	stmt, err := ur.DB.PrepareContext(ctx, query)
	if err != nil {
		return domains.User{}, err
	}
	row := stmt.QueryRowContext(ctx, id)
	res = domains.User{}
	err = row.Scan(
		&res.Id,
		&res.Name,
		&res.Dob,
		&res.Address,
		&res.Description,
		&res.CreatedAt,
	)
	return
}

func (ur *userSQLRepository) Create(ctx context.Context, u *domains.User) error {
	query := "insert into t_accounts(id, name, dob, address, description, created_at) values ($1, $2, $3, $4, $5, $6) RETURNING id"
	stmt, err := ur.DB.PrepareContext(ctx, query)
	if err != nil {
		return common.DatabaseError(err)
	}
	var userId string
	err = stmt.QueryRowContext(ctx, u.Id, u.Name, u.Dob, u.Address, u.Description, u.CreatedAt).Scan(&userId)
	if err != nil || userId == "" {
		return common.DatabaseError(err)
	}
	return nil
}

func (ur *userSQLRepository) Update(ctx context.Context, u *domains.User) error {
	query := "update t_accounts set name = $1, dob = $2, address = $3, description =$4, created_at = $5 where id = $6 returning id"
	stmt, err := ur.DB.PrepareContext(ctx, query)
	if err != nil {
		return err
	}
	var userId string
	err = stmt.QueryRowContext(ctx, u.Name, u.Dob, u.Address, u.Description, u.CreatedAt, u.Id).Scan(&userId)
	if err != nil || userId == "" {
		return common.DatabaseError(err)
	}
	return nil
}

func (ur *userSQLRepository) Delete(ctx context.Context, id string) error {
	query := "delete from t_accounts where id = $1"
	stmt, err := ur.DB.PrepareContext(ctx, query)
	if err != nil {
		return common.DatabaseError(err)
	}
	_, err = stmt.ExecContext(ctx, id)
	if err != nil {
		return common.DatabaseError(err)
	}
	return nil
}

func (ur *userSQLRepository) Get(ctx context.Context, limit int16, offset int16) ([]domains.User, error) {
	query := "select id, name, dob, address, description, created_at from t_accounts limit $1 offset $2"
	stmt, err := ur.DB.PrepareContext(ctx, query)
	results := []domains.User{}
	if err != nil {
		return nil, common.DatabaseError(err)
	}
	rows, err := stmt.QueryContext(ctx, limit, offset)
	if err != nil {
		return nil, common.DatabaseError(err)
	}
	for rows.Next() {
		user := domains.User{}
		err = rows.Scan(
			&user.Id,
			&user.Name,
			&user.Dob,
			&user.Address,
			&user.Description,
			&user.CreatedAt,
		)
		if err != nil {
			return nil, common.DatabaseError(err)
		}
		results = append(results, user)
	}

	return results, nil
}

func (u *userSQLRepository) Following(ctx context.Context, id string, followingId string) error {
	query := "insert into t_followings (u_id, following_id) values (&1, &2)"
	stmt, err := u.DB.PrepareContext(ctx, query)
	if err != nil {
		return common.DatabaseError(err)
	}
	_, err = stmt.ExecContext(ctx, id, followingId)
	if err != nil {
		return common.DatabaseError(err)
	}
	return nil
}

func (u *userSQLRepository) UnFollowing(ctx context.Context, id string, followingId string) error {
	query := "delete from t_followings where u_id = &1 and following_id = &2"
	stmt, err := u.DB.PrepareContext(ctx, query)
	if err != nil {
		return common.DatabaseError(err)
	}
	_, err = stmt.ExecContext(ctx, id, followingId)
	if err != nil {
		return common.DatabaseError(err)
	}
	return nil
}

func (u *userSQLRepository) GetFollowers(ctx context.Context, id string, limit int16, offset int16) ([]domains.Follower, error) {
	query := "select u_id, name from t_following, t_accounts where u_id = id and following_id = &1 limit &2 offset &3"
	stmt, err := u.DB.PrepareContext(ctx, query)
	if err != nil {
		return nil, common.DatabaseError(err)
	}
	rows, err := stmt.QueryContext(ctx, id, limit, offset)
	if err != nil {
		return nil, common.DatabaseError(err)
	}
	results := []domains.Follower{}
	for rows.Next() {
		follower := domains.Follower{}
		err = rows.Scan(
			&follower.Id,
			&follower.Name,
		)
		if err != nil {
			return nil, common.DatabaseError(err)
		}
		results = append(results, follower)
	}
	return results, nil
}

func (u *userSQLRepository) GetFollowings(ctx context.Context, id string, limit int16, offset int16) ([]domains.Follower, error) {
	query := "select following_id, name from t_following, t_accounts where u_id = id and u_id = &1 limit &2 offset &3"
	stmt, err := u.DB.PrepareContext(ctx, query)
	if err != nil {
		return nil, common.DatabaseError(err)
	}
	rows, err := stmt.QueryContext(ctx, id, limit, offset)
	if err != nil {
		return nil, common.DatabaseError(err)
	}
	results := []domains.Follower{}
	for rows.Next() {
		follower := domains.Follower{}
		err = rows.Scan(
			&follower.Id,
			&follower.Name,
		)
		if err != nil {
			return nil, common.DatabaseError(err)
		}
		results = append(results, follower)
	}
	return results, nil
}
