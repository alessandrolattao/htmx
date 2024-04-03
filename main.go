package main

import (
	"github.com/alessandrolattao/htmx/handlers/get"
	"github.com/alessandrolattao/htmx/handlers/post"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Main View
	e.GET("/", get.Root)

	// Login
	e.POST("/login", post.Login)

	// Logout
	e.POST("/logout", post.Logout)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
