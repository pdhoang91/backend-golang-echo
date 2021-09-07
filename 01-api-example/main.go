package main

import (
	"log"
	"net/http"

	"github.com/labstack/echo"
)

func main() {
	// Echo instance
	sever := echo.New()

	// Routes
	sever.GET("/", hello)

	// Routes
	sever.POST("/login", login)

	// Start server
	sever.Logger.Fatal(sever.Start(":8080"))
}

// Handler
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello world")
}

func login(c echo.Context) error {

	rep := new(LoginRequest)
	c.Bind(rep)
	log.Printf("req data %+v", rep)

	if rep.UserName != "admin" && rep.PassWord != "123123" {
		return c.JSON(http.StatusUnauthorized, "userName/passWord invalid")
	}

	return c.JSON(http.StatusOK, &LoginRespone{
		Token: "12121",
	})
}

type TestStruct struct {
	name string
	age  int
}

type LoginRespone struct {
	Token string `json:"token"`
}

type LoginRequest struct {
	UserName string `json:"userName" form:"userName" xml:"userName" query:"userName"`
	PassWord string `json:"passWord" form:"passWord" xml:"passWord" query:"passWord"`
}
