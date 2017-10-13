package models

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var DB *sqlx.DB

func init() {
	var err error
	driver_name := "mysql"
	if driver_name == "" {
		log.Fatal("Invalid driver name")
	}
	dsn := "root:@tcp(localhost:3306)/example_read_rails_session_development?charset=utf8&parseTime=True&loc=Local"
	if dsn == "" {
		log.Fatal("Invalid DSN")
	}
	DB, err = sqlx.Connect(driver_name, dsn)
	if err != nil {
		log.Fatal(err)
	}
}
