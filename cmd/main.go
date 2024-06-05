package main

import (
	"log"

	"github.com/tanishqtrivedi27/ecom/cmd/api"
	"github.com/tanishqtrivedi27/ecom/config"
	"github.com/tanishqtrivedi27/ecom/db"
)

func main() {
	dbconfig := db.PostgresConfig{
		User:     config.Envs.DBUser,
		Password: config.Envs.DBPasword,
		DBName:   config.Envs.DBName,
	}

	db, err := db.NewPostgreSQLStorage(dbconfig)
	if err != nil {
		log.Fatal(err)
	}

	server := api.NewAPIServer("localhost:8080", db)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
