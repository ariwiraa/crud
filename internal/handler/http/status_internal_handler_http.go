package http

import (
	"crud-echo/entity"
	"net/http"

	"github.com/labstack/echo/v4"
)

//Menampilkan response ketika berhasil dijalankan
func Status(echoCtx echo.Context) error {
	var res = entity.NewResponse(http.StatusOK, "Berhasil dijalankan", nil)
	return echoCtx.JSON(res.Status, res)
}
