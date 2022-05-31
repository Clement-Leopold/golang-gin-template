package common

import (
	"database/sql"
	"fmt"

	"github.com/spf13/viper"
)

var DB *sql.DB

// Opening a database and save the reference to `Database` struct.
// singleton
func Init() {
	host := viper.GetString(`database.host`)
	port := viper.GetString(`database.port`)
	pass := viper.GetString(`database.pass`)
	user := viper.GetString(`database.user`)
	name := viper.GetString(`database.name`)
	connStr := fmt.Sprintf("user=%s dbname=%s password=%s  host=%s port=%s sslmode=disable", user, name, pass, host, port)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	DB = db
}
