package ErrorPage

import (
	"github.com/labstack/echo/v4"
	"net"
)

func Err404(ctx echo.Context) error {
	return ctx.Render(200,"404.html",net.Interface{})
}