package Barang

import (
	"Inventory_Project/config"
	"Inventory_Project/cookie_conf"
	"github.com/labstack/echo/v4"
	"net/http"
)

func Keluar(ctx echo.Context) error {
	data := &config.M{
		"title": "Barang Keluar",
		"path": "barang-keluar",
		"alert":      cookie_conf.CookieExist(ctx, "alert"),
		"alert_data": cookie_conf.GetCookieAlert(ctx),
	}
	ctx.Render(http.StatusOK,"header",data)
	ctx.Render(http.StatusOK,"sidenav",data)
	return ctx.Render(http.StatusOK,"barang_keluar.html",data)
}
