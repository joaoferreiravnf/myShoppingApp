// config/config.go
package config

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
	"os"
)

// ConnectToDatabase opens a new connection to the database
func ConnectToDatabase() (*sql.DB, error) {
	dbHost, dbUser, dbPort, dbName := getEnvVariables()

	connStr := fmt.Sprintf("host=%s user=%s port=%s dbname=%s sslmode=disable", dbHost, dbUser, dbPort, dbName)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return &sql.DB{}, errors.Wrap(err, "error opening connection to db")
	}
	return db, nil
}

func getEnvVariables() (string, string, string, string) {
	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	return dbHost, dbUser, dbPort, dbName
}
