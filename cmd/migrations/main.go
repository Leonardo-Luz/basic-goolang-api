package main

import (
	"database/sql"
	"log"

	"fmt"

	"github.com/leonardo-luz/basic-goolang-api/cmd/api"
	"github.com/leonardo-luz/basic-goolang-api/config"
	"github.com/leonardo-luz/basic-goolang-api/db"
)

func main() {
	var DataSource = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", config.Envs.POSTGRES_HOST, config.Envs.POSTGRES_PORT, config.Envs.POSTGRES_USER, config.Envs.POSTGRES_PASSWORD, config.Envs.POSTGRES_DATABASE)

	db, err := db.NewPostgresStorage(DataSource)

	if err != nil {
		log.Fatal(err)
	}

	initStorage(db)

	server := api.NewApiServer(":8080", db)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}

func initStorage(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Database successfully connected!")
}
