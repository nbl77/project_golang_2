package Barang

import (
	"Inventory_Project/config"
	"Inventory_Project/cookie_conf"
	"github.com/labstack/echo/v4"
	"net/http"
)

func Masuk(ctx echo.Context) error {
	data := &config.M{
		"title": "Barang Masuk",
		"path": "barang-masuk",
		"alert":      cookie_conf.CookieExist(ctx, "alert"),
		"alert_data": cookie_conf.GetCookieAlert(ctx),
	}
	ctx.Render(http.StatusOK,"header",data)
	ctx.Render(http.StatusOK,"sidenav",data)
	return ctx.Render(http.StatusOK,"barang_masuk.html",data)
}
