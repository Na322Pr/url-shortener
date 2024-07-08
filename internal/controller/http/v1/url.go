package v1

import (
	"net/http"
	"net/url"
	"url-shortener/internal/service"

	"github.com/labstack/echo/v4"
)

type urlRoutes struct {
	urlService service.Url
}

func NewUrlRoutes(g *echo.Group, urlService service.Url) {
	r := &urlRoutes{urlService: urlService}

	g.GET("", r.getURL)
	g.POST("/create", r.CreateURL)
	g.DELETE("/delete", r.DeleteURL)
}

type getURLRequest struct {
	Alias string `query:"alias"`
}

func (r *urlRoutes) getURL(c echo.Context) error {
	var input getURLRequest

	if err := c.Bind(&input); err != nil {
		return err
	}

	urlToRedirect, err := r.urlService.GetURL(input.Alias)
	if err != nil {
		return err
	}

	return c.Redirect(http.StatusMovedPermanently, urlToRedirect)
}

type CreateURLRequest struct {
	Url   string `json:"url" form:"url"`
	Alias string `json:"alias" form:"alias"`
}

type CreateURLResponse struct {
	Id    int    `json:"id,omitempty"`
	Error string `json:"error,omitempty"`
}

func (r *urlRoutes) CreateURL(c echo.Context) error {
	var input CreateURLRequest

	if err := c.Bind(&input); err != nil {
		return err
	}

	if input.Alias == "" {
		return c.JSON(http.StatusOK, CreateURLResponse{
			Error: "Empty Alias",
		})
	}

	if input.Url == "" {
		return c.JSON(http.StatusOK, CreateURLResponse{
			Error: "Empty URL",
		})
	}

	if _, err := url.ParseRequestURI(input.Url); err != nil {
		return c.JSON(http.StatusOK, CreateURLResponse{
			Error: "Invalid URL",
		})
	}

	id, err := r.urlService.CreateURL(input.Url, input.Alias)

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, CreateURLResponse{
		Id: id,
	})
}

type deleteURLRequest struct {
	Id int `json:"url_id" form:"url_id"`
}

func (r *urlRoutes) DeleteURL(c echo.Context) error {
	var input deleteURLRequest

	if err := c.Bind(&input); err != nil {
		return err
	}

	println(input.Id)

	err := r.urlService.DeleteURLbyID(input.Id)
	if err != nil {
		return err
	}

	return c.NoContent(http.StatusOK)
}
