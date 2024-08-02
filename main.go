package main

import (
	"github.com/joaoferreiravnf/myShoppingApp.git/internal/config"
	"github.com/joaoferreiravnf/myShoppingApp.git/internal/models"
	"github.com/joaoferreiravnf/myShoppingApp.git/internal/repository"
	"github.com/joho/godotenv"
	"log"
	"time"
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

	defer dbConn.Close()

	item := models.MarketItem{
		Name:     "Apple",
		Quantity: 5,
		Type:     "Fruit",
		Market:   "Local Market",
		AddedAt:  time.Now(),
		AddedBy:  "John Doe",
	}

	newRepo := repository.NewRepository(dbConn)

	err = newRepo.CreateUser(item)
	if err != nil {
		log.Fatal(err)
	}
}
