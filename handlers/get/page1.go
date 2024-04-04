package get

import (
	"github.com/alessandrolattao/htmx/pages"
	"github.com/labstack/echo/v4"
)

func Page1(c echo.Context) error {
	Dashboard := pages.Page1()
	return Dashboard.Render(c.Request().Context(), c.Response().Writer)
}
