package main

import (
	"test3/handler"
	"test3/mdw"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	// Echo instance
	sever := echo.New()

	// Routes
	sever.GET("/", handler.Hello)

	// Routes
	sever.POST("/login", handler.Login, middleware.BasicAuth(mdw.BasicAuth))

	// Start server
	sever.Logger.Fatal(sever.Start(":8080"))
}
