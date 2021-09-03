package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func main() {
	// Echo instance
	e := echo.New()

	// Routes
	e.GET("/", hello)

	e.GET("/home", home)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}

// Handler
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func home(c echo.Context) error {
	return c.String(http.StatusOK, "This is Home")
}
