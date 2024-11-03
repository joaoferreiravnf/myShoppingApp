package services

import "github.com/labstack/echo/v4"

func Register(e *echo.Echo) {
	itemGroup := e.Group("/items")

	itemGroup.GET("", ListItems)
}
