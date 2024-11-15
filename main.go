package main

import (
	"github.com/joaoferreiravnf/myShoppingApp.git/internal/config"
	"github.com/joaoferreiravnf/myShoppingApp.git/internal/repository"
	"github.com/joaoferreiravnf/myShoppingApp.git/internal/services"
	"log"
)

func main() {
	configs, err := config.LoadConfigs("secrets.yaml")
	if err != nil {
		log.Fatalf("error instanciating database config: %s", err)
	}

	dbConn, err := config.ConnectToDatabase(configs.DatabaseConfig)
	if err != nil {
		log.Fatalf("error connecting to database: %s", err)
	}

	newRepo := repository.NewPostgresqlDb(dbConn, configs.DatabaseConfig.Schema, configs.DatabaseConfig.Table)
	newServer := services.NewAppServer(newRepo, configs)
	newServer.StartServer()
}
