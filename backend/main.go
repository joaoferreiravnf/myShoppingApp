package main

import (
	"github.com/joaoferreiravnf/myShoppingApp.git/internal/config"
	"github.com/joaoferreiravnf/myShoppingApp.git/internal/repository"
	"github.com/joaoferreiravnf/myShoppingApp.git/internal/services"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	dbConn, err := config.ConnectToDatabase()
	if err != nil {
		log.Fatal(err)
	}

	dbSchema := os.Getenv("DB_SCHEMA")
	dbTable := os.Getenv("DB_TABLE")

	newRepo := repository.NewPostgresqlDb(dbConn, dbSchema, dbTable)
	newServer := services.NewAppServer(newRepo)
	newServer.StartServer()
}
