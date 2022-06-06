package repository

import (
	"backend-test-chenxianhao/user-management/common"
	"backend-test-chenxianhao/user-management/domains"
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/op/go-logging"
)

var log = logging.MustGetLogger("user_repo")

type userSQLRepository struct {
	DB *sql.DB
}

func NewUserRepositoryImpl(db *sql.DB) domains.UserRepository {
	return &userSQLRepository{
		DB: db,
	}
}

func (ur *userSQLRepository) GetByID(ctx context.Context, id string) (res domains.User, err error) {
	query := "select id, name, dob, address, x_coordinate, y_coordinate, description from t_accounts where id = $1"
	stmt, err := ur.DB.PrepareContext(ctx, query)
	if err != nil {
		log.Error(err)
		return domains.User{}, common.DatabaseError(err)
	}
	row := stmt.QueryRowContext(ctx, id)
	res = domains.User{}
	err = row.Scan(
		&res.Id,
		&res.Name,
		&res.Dob,
		&res.Address,
		&res.XCoordinate,
		&res.YCoordinate,
		&res.Description,
	)
	return
}

func (ur *userSQLRepository) Create(ctx context.Context, u *domains.User) error {
	query := "insert into t_accounts(id, name, dob, address, x_coordinate, y_coordinate, description, created_at) values ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id"
	stmt, err := ur.DB.PrepareContext(ctx, query)
	if err != nil {
		return common.DatabaseError(err)
	}
	var userId string
	err = stmt.QueryRowContext(ctx, u.Id, u.Name, u.Dob, u.Address, u.XCoordinate, u.YCoordinate, u.Description, u.CreatedAt).Scan(&userId)
	if err != nil || userId == "" {
		return common.DatabaseError(err)
	}
	return nil
}

// to Simplify the situation, I just assume whole update.
func (ur *userSQLRepository) Update(ctx context.Context, u *domains.User) error {
	query := "update t_accounts set name = $1, dob = $2, address = $3, x_coordinate = $4, y_coordinate = $5,description =$6, created_at = $7 where id = $8 returning id"
	stmt, err := ur.DB.PrepareContext(ctx, query)
	if err != nil {
		log.Error(err)
		return common.DatabaseError(err)
	}
	var userId string
	err = stmt.QueryRowContext(ctx, u.Name, u.Dob, u.Address, u.XCoordinate, u.YCoordinate, u.Description, u.CreatedAt, u.Id).Scan(&userId)
	if err != nil || userId == "" {
		log.Error(err)
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
	query := "select id, name, dob, address, x_coordinate, y_coordinate,description from t_accounts limit $1 offset $2"
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
			&user.XCoordinate,
			&user.YCoordinate,
			&user.Description,
		)
		if err != nil {
			return nil, common.DatabaseError(err)
		}
		results = append(results, user)
	}

	return results, nil
}

func (u *userSQLRepository) Following(ctx context.Context, id string, follower *domains.Follower) error {
	query := "insert into t_followings(u_id, following_id, distance) values ($1, $2, $3)"

	stmt, err := u.DB.PrepareContext(ctx, query)
	if err != nil {
		return common.DatabaseError(err)
	}

	_, err = stmt.ExecContext(ctx, id, follower.Id, follower.Distance)
	if err != nil {
		return common.DatabaseError(err)
	}

	return nil
}

func (u *userSQLRepository) UnFollowing(ctx context.Context, id string, followingId string) error {
	query := "delete from t_followings where u_id = $1 and following_id = $2"
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
	query := "select id, name, distance from t_followings, t_accounts where following_id = id and following_id = $1 limit $2 offset $3"
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
			&follower.Distance,
		)
		if err != nil {
			return nil, common.DatabaseError(err)
		}
		results = append(results, follower)
	}
	return results, nil
}

func (u *userSQLRepository) GetFollowings(ctx context.Context, id string, limit int16, offset int16) ([]domains.Follower, error) {
	query := "select following_id, name, distance from t_followings, t_accounts where u_id = id and u_id = $1 limit $2 offset $3"
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
			&follower.Distance,
		)
		if err != nil {
			return nil, common.DatabaseError(err)
		}
		results = append(results, follower)
	}
	return results, nil
}

func (u *userSQLRepository) SelectTwoCoordinates(ctx context.Context, uIdOne string, uIdTwo string) ([]int64, error) {
	query := "select x_coordinate, y_coordinate from t_accounts where id = $1 or id = $2"
	stmt, err := u.DB.PrepareContext(ctx, query)
	if err != nil {
		return nil, common.DatabaseError(err)
	}
	rows, err := stmt.QueryContext(ctx, uIdOne, uIdTwo)
	if err != nil {
		return nil, common.DatabaseError(err)
	}
	results := []int64{}
	for rows.Next() {
		var xCoordinate int64
		var yCoordinate int64
		rows.Scan(&xCoordinate, &yCoordinate)
		results = append(results, xCoordinate, yCoordinate)
	}
	return results, nil
}

func (u *userSQLRepository) SelectAllFollowing(ctx context.Context, uId string) ([]domains.Follower, error) {
	query := `select t_followings.auto_id, 
					t_followings.following_id,
					t_accounts.x_coordinate, 
					t_accounts.y_coordinate
					from 
						t_followings 
					left join 
						t_accounts 
					on  following_id = id where t_followings.u_id = $1`
	stmt, err := u.DB.PrepareContext(ctx, query)
	if err != nil {
		log.Error(err)
		return nil, common.DatabaseError(err)
	}
	rows, err := stmt.QueryContext(ctx, uId)
	if err != nil {
		log.Error(err)
		return nil, common.DatabaseError(err)
	}
	results := []domains.Follower{}
	for rows.Next() {
		follower := domains.Follower{}
		rows.Scan(
			&follower.AutoId,
			&follower.Id,
			&follower.XCoordinate,
			&follower.YCoordinate)
		results = append(results, follower)
	}
	return results, nil

}

func (u *userSQLRepository) BatchUpdateForFollowerDistance(ctx context.Context, followers []domains.Follower) error {
	query := "update t_followings set distance = tmp.distance from (values %s) as tmp(auto_id, distance) where tmp.auto_id = t_followings.auto_id"
	values := []string{}
	for _, v := range followers {
		value := fmt.Sprintf("(%d,%f)", v.AutoId, v.Distance)
		values = append(values, value)
	}
	query = fmt.Sprintf(query, strings.Join(values, ","))

	stmt, err := u.DB.PrepareContext(ctx, query)
	if err != nil {
		log.Error(err)
		return common.DatabaseError(err)
	}
	_, err = stmt.ExecContext(ctx)

	if err != nil {
		log.Error(err)
		return common.DatabaseError(err)
	}

	return nil
}

func (u *userSQLRepository) GetMinimumDistanceForFollowing(ctx context.Context, name string) (domains.Follower, error) {
	query := "select id, name from t_followings, t_accounts where u_id = id and name = $1 order by distance limit 1"
	stmt, err := u.DB.PrepareContext(ctx, query)
	if err != nil {
		log.Error(err)
		return domains.Follower{}, common.DatabaseError(err)
	}
	result := stmt.QueryRowContext(ctx, name)
	follower := domains.Follower{}
	result.Scan(
		&follower.Id,
		&follower.Name,
	)
	return follower, nil
}
