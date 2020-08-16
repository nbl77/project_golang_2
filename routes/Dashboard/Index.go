package Dashboard

import (
	"Inventory_Project/config"
	"github.com/labstack/echo/v4"
	"net"
	"net/http"
)

func Home(ctx echo.Context) error {
	data := &config.M{
		"title":"Dashboard",
		"path": "dashboard",
	}
	ctx.Render(http.StatusOK,"header",data)
	ctx.Render(http.StatusOK,"sidenav",data)
	return ctx.Render(http.StatusOK,"index.html",net.Interface{})
}