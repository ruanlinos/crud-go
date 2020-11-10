package db

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql" //Driver for connection
)

// Connect open the connection with database.
func Connect() (*sql.DB, error) {

	urlConnection := "docker:docker@/course?charset=utf8&parseTime=True&loc=Local"

	db, err := sql.Open("mysql", urlConnection)

	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil

}
