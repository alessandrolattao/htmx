package get

import (
	"github.com/alessandrolattao/htmx/pages"
	"github.com/labstack/echo/v4"
)

func Dashboard(c echo.Context) error {
	Dashboard := pages.Dashboard()
	return Dashboard.Render(c.Request().Context(), c.Response().Writer)
}
