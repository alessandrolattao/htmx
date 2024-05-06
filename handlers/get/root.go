package get

import (
	"strconv"

	"github.com/alessandrolattao/htmx/helpers"
	"github.com/alessandrolattao/htmx/views"
	"github.com/labstack/echo/v4"
)

// Root is the main handler for the root path.
func Root(c echo.Context) error {
	if !checkLoginCookie(c) {
		Login := views.Login()
		return Login.Render(c.Request().Context(), c.Response().Writer)
	}

	Dashboard := views.Dashboard(c)
	return Dashboard.Render(c.Request().Context(), c.Response().Writer)
}

// checkLoginCookie checks if the user is logged in by checking the login cookie.
func checkLoginCookie(c echo.Context) bool {

	cookie := helpers.GetCookieValue(c, "login")
	if cookie == "" {
		helpers.UnsetCookie(c, "login")
		helpers.UnsetCookie(c, "email")
		return false
	}

	loginStatus, err := strconv.ParseBool(cookie)
	if err != nil {
		helpers.UnsetCookie(c, "login")
		helpers.UnsetCookie(c, "email")
		return false
	}

	return loginStatus
}
