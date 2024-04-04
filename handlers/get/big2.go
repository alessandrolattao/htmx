package get

import (
	"github.com/alessandrolattao/htmx/components"
	"github.com/labstack/echo/v4"
)

func Big2(c echo.Context) error {
	component := components.Big2()
	return component.Render(c.Request().Context(), c.Response().Writer)
}
