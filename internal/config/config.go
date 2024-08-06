// config/config.go
package config

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"os"
)

// ConnectToDatabase opens a new connection to the database
func ConnectToDatabase() (*sql.DB, error) {
	dbHost, dbUser, dbPassword, dbName := getEnvVariables()

	connStr := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", dbHost, dbUser, dbPassword, dbName)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return &sql.DB{}, err
	}
	return db, nil
}

func getEnvVariables() (string, string, string, string) {
	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	return dbHost, dbUser, dbPassword, dbName
}
