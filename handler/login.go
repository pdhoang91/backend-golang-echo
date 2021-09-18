package handler

import (
	"back-end-echo/models"
	"net/http"

	"github.com/labstack/echo"
)

func Login(c echo.Context) error {

	response := &models.LoginResponse{Token: "123456789"}
	return c.JSON(http.StatusOK, response)
}
