package services

import (
	"github.com/joaoferreiravnf/myShoppingApp.git/internal/models"
	"github.com/joaoferreiravnf/myShoppingApp.git/internal/repository"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/pkg/errors"
	"html/template"
	"io"
	"net/http"
	"strconv"
)

type AppServer struct {
	*echo.Echo
	repo *repository.PostgresqlDb
}

func NewAppServer(repo *repository.PostgresqlDb) *AppServer {
	newEcho := echo.New()
	newEcho.Use(middleware.CORS())
	newServer := &AppServer{newEcho, repo}

	renderer := &TemplateRenderer{
		templates: template.Must(template.ParseGlob("views/*.html")),
	}
	newServer.Echo.Renderer = renderer

	newEcho.Use(middleware.Logger())

	itemGroup := newEcho.Group("/items")
	itemGroup.GET("", newServer.ListItems)
	itemGroup.POST("", newServer.CreateItem)

	return newServer
}

type TemplateRenderer struct {
	templates *template.Template
}

func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func (s *AppServer) StartServer() {
	s.Logger.Fatal(s.Start(":8080"))
}

func (s *AppServer) ListItems(c echo.Context) error {
	items, err := s.repo.ListItems(c.Request().Context())
	if err != nil {
		return errors.Wrap(err, "error listing items from database")
	}

	listData := models.ListItemsData{
		Items: items,
	}

	listData.GetTypes()
	listData.GetMarkets()

	return c.Render(http.StatusOK, "list_items.html", listData)
}

func (s *AppServer) CreateItem(c echo.Context) error {
	quantity, err := strconv.Atoi(c.FormValue("quantity"))
	if err != nil {
		return errors.Wrap(err, "error converting quantity field to int")
	}

	newItem := models.Item{
		Name:     c.FormValue("name"),
		Quantity: quantity,
		Type:     c.FormValue("type"),
		Market:   c.FormValue("market"),
		AddedBy:  c.FormValue("added_by"),
	}

	err = newItem.NormalizeFieldsForPersistence()
	if err != nil {
		return errors.Wrap(err, "error normalizing item fields for persistence")
	}

	err = s.repo.CreateItem(c.Request().Context(), newItem)
	if err != nil {
		return errors.Wrap(err, "error creating new item")
	}

	return c.Redirect(http.StatusSeeOther, "/items")
}
