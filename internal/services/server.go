package services

import (
	"github.com/joaoferreiravnf/myShoppingApp.git/internal/repository"
	"github.com/labstack/echo/v4"
)

type AppServer struct {
	*echo.Echo
	repo *repository.PostgresqlDb
}

func NewAppServer(repo *repository.PostgresqlDb) *AppServer {
	e := echo.New()

	s := &AppServer{e, repo}

	return s
}

func StartServer() {
	e := echo.New()

	Register(e)

	e.Logger.Fatal(e.Start(":8080"))
}
