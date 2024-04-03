package post

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

func Logout(c echo.Context) error {
	cookie := new(http.Cookie)
	cookie.Name = "login"
	cookie.Value = ""
	cookie.Expires = time.Unix(0, 0)
	c.SetCookie(cookie)
	return c.Redirect(http.StatusSeeOther, "/")
}
