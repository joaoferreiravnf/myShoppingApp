package services

import (
	"github.com/joaoferreiravnf/myShoppingApp.git/internal/repository"
	"github.com/labstack/echo/v4"
)

func ListItems(c echo.Context) error {
	items := repository.PostgresqlDb.ListItems()
}
