package main

import (
	"github.com/joaoferreiravnf/myShoppingApp.git/internal/config"
	"github.com/joaoferreiravnf/myShoppingApp.git/internal/repository"
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

	/*	item := models.Item{
			Name:     "aPPle PIT",
			Quantity: 5,
			Type:     "Fruit",
			Market:   "Local Market",
			AddedAt:  time.Now(),
			AddedBy:  "John",
		}

		item.NormalizeNameForPersistence()

		err = newRepo.CreateItem(item)
		if err != nil {
			c.Logger.Fatal(err)
		}*/

	err = newRepo.DeleteItem(10)
	if err != nil {
		log.Fatal(err)
	}
}
