package post

import (
	"net/http"

	"github.com/alessandrolattao/htmx/helpers"
	"github.com/labstack/echo/v4"
)

func Login(c echo.Context) error {
	helpers.SetCookie(c, "login", "true")
	helpers.SetCookie(c, "email", c.FormValue("email"))
	return c.Redirect(http.StatusSeeOther, "/")
}
