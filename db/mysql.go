package db

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func Connect() *sql.DB {
	db, err := sql.Open("mysql", "setia:setia@tcp(localhost:3306)/contacts")

	if err != nil {
		log.Fatal(err)
	}

	return db
}
