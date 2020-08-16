package cookie_conf

import (
	"Inventory_Project/config"
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
	"time"
)

func SetCookie(response echo.Context, key, val string) bool {
	cookie := &http.Cookie{
		Name: key,
		Value: val,
		Expires: time.Now().Add(config.COOKIE_LOGIN_EXPIRES * time.Hour),
	}
	response.SetCookie(cookie)
	return true
}
func SetCookieAlert(response echo.Context, status, message string) bool {
	cookie := &http.Cookie{
		Name: "alert",
		Value: status +","+ message,
		Path: "/",
		Expires: time.Now().Add(config.COOKIE_LOGIN_EXPIRES * time.Hour),
	}
	response.SetCookie(cookie)
	return true
}
func GetCookieAlert(response echo.Context) (map[string]string) {
	key := "alert"
	cookie, err := response.Cookie(key)
	if err != nil {
		return nil
	}
	if !CookieExist(response,key) {
		return nil
	}
	res := strings.Split(cookie.Value,",")
	data := map[string]string{
		"status":res[0],
		"message":res[1],
	}
	DeleteCookie(response,key)
	return data
}
func GetCookie(response echo.Context, key string) (string,error) {
	cookie, err := response.Cookie(key)
	if err != nil {
		return "",err
	}
	return cookie.Value,nil
}
func DeleteCookie(response echo.Context, key string) bool {
	cookie := &http.Cookie{
		Name:    key,
		Value:   "",
		Expires: time.Unix(0, 0),
		MaxAge:   -1,
	}
	response.SetCookie(cookie)
	return true
}
func CookieExist(response echo.Context, key string) bool {
	_, err := response.Cookie(key)
	if err != nil {
		return false
	}
	return true
}