package v1

import (
	"url-shortener/internal/service"

	"github.com/labstack/echo/v4"
)

func NewRouter(handler *echo.Echo, service *service.Service) {
	v1 := handler.Group("/api/v1")
	{
		NewUrlRoutes(v1.Group("/urls"), service.Url)
	}
}
