package db

import (
	"database/sql"
	"log"

	"github.com/go-sql-driver/mysql"
)

func MySqlStorage(cfg mysql.Config) (*sql.DB, error) {

	db, err := sql.Open("mysql", cfg.FormatDSN())

	if err != nil {
		log.Fatal("Error occured while opening mysql connection")
	}

	return db, nil

}
