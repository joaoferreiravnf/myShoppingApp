package services

import (
	"database/sql"
	"github.com/joaoferreiravnf/myShoppingApp.git/internal/repository"
	"github.com/labstack/echo/v4"
)

type AppServer struct {
	*echo.Echo
	repo *repository.postgresqlDb
}

func NewAppServer(repo *repository.postgresqlDb, db *sql.DB) *AppServer {
	e := echo.New()

	s := &AppServer{e, repo}

	e.GET("/items", s.listItems)

	return s
}

func (s *AppServer) listItems(c echo.Context) error {
	items, err := s.repo.ListItems()
	if err != nil {
		return c.JSON(500, "An error occurred while fetching the items")
	}

	return c.JSON(200, items)
}
