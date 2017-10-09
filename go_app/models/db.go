package models

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

var DB *sqlx.DB

func init() {
	var err error
	driver_name := "sqlite3"
	if driver_name == "" {
		log.Fatal("Invalid driver name")
	}
	dsn := "../db/development.sqlite3"
	if dsn == "" {
		log.Fatal("Invalid DSN")
	}
	DB, err = sqlx.Connect(driver_name, dsn)
	if err != nil {
		log.Fatal(err)
	}
}
