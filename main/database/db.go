package database

import (
	"github.com/jmoiron/sqlx"
	"log"
)

var (
	SQL *sqlx.DB
)

func Connect() () {
	var err error

	if SQL, err = sqlx.Connect("postgres",
		"user=db_api password=MTIzNHF3ZXI= dbname=TodoAppDB sslmode=disable"); err != nil {
		log.Println("SQL Driver Error", err)
	}

	// Check if is alive
	if err = SQL.Ping(); err != nil {
		log.Println("Database Error", err)
	}
}
