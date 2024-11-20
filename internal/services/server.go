package services

import (
	"crypto/subtle"
	"github.com/joaoferreiravnf/myShoppingApp.git/internal/config"
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
	echo   *echo.Echo
	repo   *repository.PostgresqlDb
	config *config.Config
}

func NewAppServer(repo *repository.PostgresqlDb, configs *config.Config) *AppServer {
	newEcho := echo.New()

	newServer := &AppServer{
		newEcho,
		repo,
		configs,
	}

	renderer := &TemplateRenderer{
		templates: template.Must(template.ParseGlob("views/*.html")),
	}
	newServer.echo.Renderer = renderer

	newServer.echo.HTTPErrorHandler = customErrorHandler

	newServer.echo.Use(middleware.Logger())
	basicAuth := newServer.basicAuth()

	itemGroup := newServer.echo.Group("/items", basicAuth...)

	itemGroup.GET("", newServer.ListItems)
	itemGroup.POST("/create", newServer.CreateItem)
	itemGroup.POST("/delete/:id", newServer.DeleteItem)

	return newServer
}

func (s *AppServer) basicAuth() []echo.MiddlewareFunc {
	return []echo.MiddlewareFunc{
		middleware.BasicAuth(func(username, password string, _ echo.Context) (bool, error) {
			if subtle.ConstantTimeCompare([]byte(username), []byte(s.config.AppAuth.Username)) == 1 &&
				subtle.ConstantTimeCompare([]byte(password), []byte(s.config.AppAuth.Password)) == 1 {
				return true, nil
			}
			return false, nil
		}),
	}
}

func customErrorHandler(err error, c echo.Context) {
	code := http.StatusInternalServerError
	message := err.Error()

	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
	}

	if code == http.StatusUnauthorized {
		c.Response().Header().Set(echo.HeaderWWWAuthenticate, `Basic realm="Restricted"`)
	}

	err = c.Render(code, "error.html", map[string]interface{}{
		"status":  "error",
		"message": message,
	})
	if err != nil {
		return
	}
}

type TemplateRenderer struct {
	templates *template.Template
}

func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func (s *AppServer) StartServer() {
	s.echo.Logger.Fatal(s.echo.Start(":8080"))
}

func (s *AppServer) ListItems(c echo.Context) error {
	items, err := s.repo.ListItems(c.Request().Context())
	if err != nil {
		return errors.Wrap(err, "error listing items")
	}

	listData := models.ListItemsData{
		Items: items,
	}

	listData.GetTypes()
	listData.GetMarkets()
	listData.GetQuantities()

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

func (s *AppServer) DeleteItem(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return errors.Wrap(err, "error converting id field to int")
	}

	item := models.Item{
		ID:   id,
		Name: c.FormValue("name"),
	}

	err = s.repo.DeleteItem(c.Request().Context(), item)
	if err != nil {
		return errors.Wrap(err, "error deleting item")
	}

	return c.Redirect(http.StatusSeeOther, "/items")
}
