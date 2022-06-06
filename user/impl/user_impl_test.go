package impl

import (
	"backend-test-chenxianhao/user-management/domains"
	"backend-test-chenxianhao/user-management/user/repository"
	"backend-test-chenxianhao/user-management/util"
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

// test basic CRUD functions using go-sqlmock
func TestCRUD(t *testing.T) {
	db, mock, err := sqlmock.NewWithDSN(InitDSN())
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	user := prepareUserMockData()
	userSQLRepository := repository.NewUserRepositoryImpl(db)
	userImpl := UserMethodImpl(userSQLRepository)
	defer db.Close()

	mock.ExpectPrepare("insert into t_accounts").
		ExpectQuery().
		WithArgs(user.Id, user.Name, user.Dob, user.Address, user.XCoordinate, user.YCoordinate, user.Description, user.CreatedAt).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(user.Id))

	if err := userImpl.Create(context.Background(), user); err != nil {
		t.Errorf("was expecting no error, but there was one: %s", err)
	}

	mock.ExpectPrepare("update t_accounts").
		ExpectQuery().
		WithArgs(user.Name, user.Dob, user.Address, user.XCoordinate, user.YCoordinate, user.Description, user.CreatedAt, user.Id).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(user.Id))

	// for test go routine run.
	mock.ExpectPrepare("select auto_id, following_id, x_coordinate").
		ExpectQuery().
		WithArgs(user.Id).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(user.Id))

	if err := userImpl.Update(context.Background(), user); err != nil {
		t.Errorf("was expecting no error, but there was one: %s", err)
	}
	time.Sleep(time.Second * 1)
	// test Select
	mock.ExpectPrepare("select id, name").
		ExpectQuery().
		WithArgs(user.Id).
		WillReturnRows(sqlmock.
			NewRows([]string{"id", "name", "dob", "address", "x_coordinate", "y_coordinate", "description"}).
			AddRow(user.Id, user.Name, user.Dob, user.Address, user.XCoordinate, user.YCoordinate, user.Description))

	_, err = userImpl.GetByID(context.Background(), user.Id)
	if err != nil {
		t.Errorf("was expecting no error, but there was one: %s", err)
	}
}

// test follower sql parameter match.
func TestFollowing(t *testing.T) {
	db, mock, err := sqlmock.NewWithDSN(InitDSN())
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	user := prepareFollower()
	following := prepareFollowing()
	userSQLRepository := repository.NewUserRepositoryImpl(db)
	userImpl := UserMethodImpl(userSQLRepository)
	followRelation := &domains.Follower{
		Id:       following.Id,
		Distance: util.GetDistance(user.XCoordinate, user.YCoordinate, following.XCoordinate, following.YCoordinate),
	}

	// test following
	mock.ExpectPrepare("select x_coordinate, y_coordinate from").
		ExpectQuery().
		WithArgs(user.Id, following.Id).
		WillReturnRows(sqlmock.
			NewRows([]string{"x_coordinate", "y_coordinate"}).
			AddRow(user.XCoordinate, user.YCoordinate).
			AddRow(following.XCoordinate, following.YCoordinate))

	mock.ExpectPrepare("insert into t_followings").
		ExpectExec().
		WithArgs(user.Id, following.Id, followRelation.Distance).
		WillReturnResult(sqlmock.NewResult(1, 1))

	if err = userImpl.Following(context.Background(), user.Id, followRelation.Id); err != nil {
		assert.Error(t, err)
	}
	// test followers
	mock.ExpectPrepare("select id, name, distance from t_followings").
		ExpectQuery().
		// limit offset is not important
		WithArgs(following.Id, 5, 5).
		WillReturnRows(sqlmock.
			NewRows([]string{"u_id", "name", "distance"}).
			AddRow(user.Id, user.Name, 8))

	if _, err = userImpl.GetFollowers(context.Background(), following.Id, 5, 5); err != nil {
		assert.Error(t, err)
	}

	// test followings
	mock.ExpectPrepare("select following_id, name, distance from").
		ExpectQuery().
		// limit offset is not important
		WithArgs(user.Id, 5, 5).
		WillReturnRows(sqlmock.
			NewRows([]string{"u_id", "name", "distance"}).
			AddRow(following.Id, following.Name, 8))

	followings, err := userImpl.GetFollowings(context.Background(), user.Id, 5, 5)
	if err != nil {
		assert.Error(t, err)
	}
	assert.Equal(t, len(followings), 1)

}

func InitDSN() string {
	viper.SetConfigFile(`config.json`)
	host := viper.GetString(`database.host`)
	port := viper.GetString(`database.port`)
	pass := viper.GetString(`database.pass`)
	user := viper.GetString(`database.user`)
	name := viper.GetString(`database.name`)
	return fmt.Sprintf("user=%s dbname=%s password=%s  host=%s port=%s sslmode=disable", user, name, pass, host, port)

}

func prepareUserMockData() *domains.User {
	return &domains.User{
		Id:          uuid.NewString(),
		Name:        "test",
		Dob:         "0000-00-00",
		Address:     "test",
		Description: "test",
		XCoordinate: 2,
		YCoordinate: 15,
		CreatedAt:   time.Now(),
	}
}

func prepareFollower() *domains.User {
	return &domains.User{
		Id:          uuid.NewString(),
		Name:        "follower",
		Dob:         "0000-00-00",
		Address:     "followerAddress",
		Description: "follower",
		XCoordinate: 1,
		YCoordinate: 1,
		CreatedAt:   time.Now(),
	}
}

func prepareFollowing() *domains.User {
	return &domains.User{
		Id:          uuid.NewString(),
		Name:        "following",
		Dob:         "0000-00-00",
		Address:     "following",
		Description: "following",
		XCoordinate: 3,
		YCoordinate: 3,
		CreatedAt:   time.Now(),
	}
}
