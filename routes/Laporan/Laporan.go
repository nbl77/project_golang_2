package Laporan

import (
	"Inventory_Project/config"
	"github.com/labstack/echo/v4"
	"net/http"
)

func Laporan(ctx echo.Context) error {
	data := &config.M{
		"title": "Laporan",
		"path": "laporan",
	}
	ctx.Render(http.StatusOK,"header",data)
	ctx.Render(http.StatusOK,"sidenav",data)
	return ctx.Render(http.StatusOK,"laporan.html",data)
}
