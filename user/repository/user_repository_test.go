package repository

import (
	"backend-test-chenxianhao/user-management/domains"
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/spf13/viper"
)

// test basic CRUD functions using go-sqlmock
func TestCRUD(t *testing.T) {
	db, mock, err := sqlmock.NewWithDSN(InitDSN())
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	user := prepareUserMockData()
	userSQLRepository := NewUserRepositoryImpl(db)
	defer db.Close()
	// test Create
	mock.ExpectPrepare("insert into t_accounts").
		ExpectQuery().
		WithArgs(user.Id, user.Name, user.Dob, user.Address, user.Description, user.CreatedAt).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(user.Id))

	if err := userSQLRepository.Create(context.Background(), user); err != nil {
		t.Errorf("was expecting no error, but there was one: %s", err)
	}
	// test Update
	user.Address = "test Update"
	user.CreatedAt = time.Now()

	mock.ExpectPrepare("update t_accounts").
		ExpectQuery().
		WithArgs(user.Name, user.Dob, user.Address, user.Description, user.CreatedAt, user.Id).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(user.Id))

	if err := userSQLRepository.Update(context.Background(), user); err != nil {
		t.Errorf("was expecting no error, but there was one: %s", err)
	}

	// test Select
	mock.ExpectPrepare("select id, name").
		ExpectQuery().
		WithArgs(user.Id).
		WillReturnRows(sqlmock.
			NewRows([]string{"id", "name", "dob", "address", "description", "create_at"}).
			AddRow(user.Id, user.Name, user.Dob, user.Address, user.Description, user.CreatedAt))

	_, err = userSQLRepository.GetByID(context.Background(), user.Id)
	if err != nil {
		t.Errorf("was expecting no error, but there was one: %s", err)
	}
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
		CreatedAt:   time.Now(),
	}
}
