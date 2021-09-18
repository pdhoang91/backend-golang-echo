package main

import (
	"back-end-echo/handler"
	"back-end-echo/mdw"
	"fmt"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	// Echo instance
	sever := echo.New()

	// Routes
	sever.GET("/", handler.Hello)

	fmt.Println("Test")

	// Routes
	sever.POST("/login", handler.Login, middleware.BasicAuth(mdw.BasicAuth))

	// Start server
	sever.Logger.Fatal(sever.Start(":8080"))
}
