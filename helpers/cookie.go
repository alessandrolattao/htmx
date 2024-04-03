package helpers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func SetCookie(c echo.Context, name string, value string) {
	cookie := new(http.Cookie)
	cookie.Name = name
	cookie.Value = value
	c.SetCookie(cookie)
}

func GetCookieValue(c echo.Context, name string) string {
	cookie, err := c.Cookie(name)
	if err != nil {
		return ""
	}
	return cookie.Value
}

func UnsetCookie(c echo.Context, name string) {
	cookie := new(http.Cookie)
	cookie.Name = name
	cookie.Value = ""
	cookie.Expires = cookie.Expires.AddDate(0, 0, -1)
	c.SetCookie(cookie)
}
