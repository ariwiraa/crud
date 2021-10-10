package routes

import (
	"crud-echo/internal/handler/http"

	"github.com/labstack/echo/v4"
)

func Init() *echo.Echo {
	e := echo.New()

	e.GET("/", http.Status)

	return e
}
