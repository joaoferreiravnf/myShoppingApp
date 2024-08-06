package main

import (
	"github.com/joaoferreiravnf/myShoppingApp.git/internal/config"
	"github.com/joaoferreiravnf/myShoppingApp.git/internal/models"
	"github.com/joaoferreiravnf/myShoppingApp.git/internal/repository"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"log"
	"time"
)

func main() {
	e := echo.New()

	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	dbConn, err := config.ConnectToDatabase()
	if err != nil {
		e.Logger.Fatal(err)
	}

	newRepo := repository.NewItemsRepository(dbConn)

	e.GET("/items",

	item := models.MarketItem{
		Name:     "Apple",
		Quantity: 5,
		Type:     "Fruit",
		Market:   "Local Market",
		AddedAt:  time.Now(),
		AddedBy:  "John Doe",
	}

	err = newRepo.CreateItem(item)
	if err != nil {
		log.Fatal(err)
	}
}
