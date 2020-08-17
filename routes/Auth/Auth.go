package Auth

import (
	"Inventory_Project/config"
	"Inventory_Project/cookie_conf"
	"Inventory_Project/db"
	"crypto/md5"
	"encoding/hex"
	"github.com/labstack/echo/v4"
	"net/http"
)
type M map[string]interface {}
func Login(ctx echo.Context) error {
	data := &M{
		"title":"Login",
		"alert":      cookie_conf.CookieExist(ctx, "alert"),
		"alert_data": cookie_conf.GetCookieAlert(ctx),
	}
	ctx.Render(http.StatusOK,"header",data)
	return ctx.Render(http.StatusOK,"login.html",data)
}
func LoginPost(ctx echo.Context) error {
	username := ctx.FormValue("username")
	password := toMD5(ctx.FormValue("password"))
	row := db.SelectParam("SELECT count(*) FROM admin WHERE username=? AND password=?",username,password)
	var result int
	row.Scan(&result)
	if result > 0 {
		cookie_conf.SetCookie(ctx,config.COOKIE_LOGIN_KEY,config.COOKIE_LOGIN_VAL)
	}else {
		cookie_conf.SetCookieAlert(ctx,"danger","Username Atau Password Salah!")
	}
	return ctx.Redirect(http.StatusFound,"/login")
}
func ChangePass(ctx echo.Context) error {
	oldPass := toMD5(ctx.FormValue("old-pass"))
	newPass := toMD5(ctx.FormValue("new-pass"))
	row := db.SelectParam("SELECT count(*) FROM admin WHERE password=?",oldPass)
	var result int
	row.Scan(&result)
	if result > 0 {
		db.Execute("UPDATE admin SET password=?",newPass)
		cookie_conf.SetCookieAlert(ctx,"success","Berhasil Merubah Password!")
		return ctx.Redirect(http.StatusFound,"/logout")
	}
	cookie_conf.SetCookieAlert(ctx,"danger","Password Yang Anda Masukan Salah!")
	return ctx.Redirect(http.StatusFound,"/")
}
func toMD5(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}