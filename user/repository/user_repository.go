package repository

import (
	"backend-test-chenxianhao/user-management/common"
	"backend-test-chenxianhao/user-management/domains"
	"context"
	"database/sql"
	"time"
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
		return err
	}
	var userId string
	err = stmt.QueryRowContext(ctx, u.Id, u.Name, u.Dob, u.Address, u.Description, time.Now()).Scan(&userId)
	return err
}

func (ur *userSQLRepository) Update(ctx context.Context, u *domains.User) error {
	query := "update t_accounts set name = $1, dob = $2, address = $3, description =$4, created_at = $5 where id = $6 returning id"
	stmt, err := ur.DB.PrepareContext(ctx, query)
	if err != nil {
		return err
	}
	var userId string
	err = stmt.QueryRowContext(ctx, u.Name, u.Dob, u.Address, u.Description, u.CreatedAt, u.Id).Scan(&userId)
	return err
}

func (ur *userSQLRepository) Delete(ctx context.Context, id string) error {
	query := "delete from t_accounts where id = $1"
	stmt, err := ur.DB.PrepareContext(ctx, query)
	if err != nil {
		return &common.BusinessError{Err: err, Code: common.DatabaseCode, Message: common.Database}
	}
	var userId string
	err = stmt.QueryRowContext(ctx, id).Scan(&userId)
	return err
}

func (ur *userSQLRepository) Get(ctx context.Context, limit int16, offset int16) ([]domains.User, error) {
	query := "select id, name, dob, address, description, created_at from t_accounts limit $1 offset $2"
	stmt, err := ur.DB.PrepareContext(ctx, query)
	results := []domains.User{}
	if err != nil {
		return nil, &common.BusinessError{Err: err, Code: common.DatabaseCode, Message: common.Param}
	}
	rows, err := stmt.QueryContext(ctx, limit, offset)
	if err != nil {
		return nil, &common.BusinessError{Err: err, Code: common.DatabaseCode, Message: common.Param}
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
			return nil, &common.BusinessError{Err: err, Code: common.DatabaseCode, Message: common.Param}
		}
		results = append(results, user)
	}

	return results, nil
}
