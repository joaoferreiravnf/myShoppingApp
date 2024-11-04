package services

import (
	"github.com/joaoferreiravnf/myShoppingApp.git/internal/repository"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/pkg/errors"
	"net/http"
)

type AppServer struct {
	*echo.Echo
	repo *repository.PostgresqlDb
}

func NewAppServer(repo *repository.PostgresqlDb) *AppServer {
	newEcho := echo.New()
	newEcho.Use(middleware.CORS())
	newServer := &AppServer{newEcho, repo}

	itemGroup := newEcho.Group("/items")
	itemGroup.GET("", newServer.ListItems)

	return newServer
}

func (s *AppServer) StartServer() {
	s.Logger.Fatal(s.Start(":8080"))
}

func (s *AppServer) ListItems(c echo.Context) error {
	items, err := s.repo.ListItems(c.Request().Context())
	if err != nil {
		return errors.Wrap(err, "error listing items from database")
	}

	return c.JSON(http.StatusOK, items)
}
