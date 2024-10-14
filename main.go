package main

import (
	"database/sql"
	"github.com/tanishqtrivedi27/ecom/api"
	"log"

	"github.com/tanishqtrivedi27/ecom/config"
	"github.com/tanishqtrivedi27/ecom/db"
)

func main() {
	dbConfig := db.PostgresConfig{
		Host:     config.Envs.PublicHost,
		User:     config.Envs.DBUser,
		Password: config.Envs.DBPassword,
		DBName:   config.Envs.DBName,
		Port:     5432,
	}

	postgreSQLStorage, err := db.NewPostgreSQLStorage(dbConfig)
	if err != nil {
		log.Fatal(err)
	}
	defer func(postgreSQLStorage *sql.DB) {
		err := postgreSQLStorage.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(postgreSQLStorage)

	server := api.NewAPIServer(":8080", postgreSQLStorage)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
