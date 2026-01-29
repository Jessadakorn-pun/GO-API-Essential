package database

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func PostgresConnect(ConnectionString string) {

	// Connect to SLQ database
	sdb, err := sql.Open("postgres", ConnectionString)
	if err != nil {
		log.Fatal(err)
	}

	// Check Connection
	err = sdb.Ping()
	if err != nil {
		log.Fatal(err)
	}

	DB = sdb

}
