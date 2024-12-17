package db

import (
	"database/sql"

	"log"

	_ "github.com/lib/pq"
)

func NewPostgresStorage(DataSource string) (*sql.DB, error) {
	db, err := sql.Open("postgres", DataSource)

	if err != nil {
		log.Fatal("Unable to connect to database", err)
	}

	return db, nil
}
