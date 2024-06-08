package main

import (
	"log"

	"github.com/tanishqtrivedi27/ecom/cmd/api"
	"github.com/tanishqtrivedi27/ecom/config"
	"github.com/tanishqtrivedi27/ecom/db"
)

func main() {
	dbconfig := db.PostgresConfig{
		Host:     config.Envs.PublicHost,
		User:     config.Envs.DBUser,
		Password: config.Envs.DBPasword,
		DBName:   config.Envs.DBName,
		Port:     5432,
	}

	db, err := db.NewPostgreSQLStorage(dbconfig)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	server := api.NewAPIServer(":8080", db)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
