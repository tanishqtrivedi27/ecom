package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"

	_ "github.com/lib/pq"
)

func main() {
	// PostgreSQL connection parameters
	const (
		host     = "localhost" // or the Docker container IP
		port     = 5432        // or the port your Docker container is exposed on
		user     = "postgres"
		password = "postgres"
		dbname   = "ecom"
	)

	// Connection string
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	// Connect to PostgreSQL database
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Ensure connection is valid
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfully connected to the database!")


	folderPath := "cmd/migrate/migrations"
	err = filepath.Walk(folderPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			file, err := os.Open(path)
			if err != nil {
				log.Fatal(err)
			}
			defer file.Close()

			var queries string
			scanner := bufio.NewScanner(file)
			for scanner.Scan() {
				queries += scanner.Text() + "\n"
			}

			if err := scanner.Err(); err != nil {
				log.Fatal(err)
			}

			_, err = db.Exec(queries)
			if err != nil {
				log.Fatal(err)
			}
		}
		return nil
	})

	// Handle errors, if any
	if err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Println("Successfully executed query!")
}
