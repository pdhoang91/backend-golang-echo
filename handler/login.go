package handler

import (
	"net/http"
	"test3/models"

	"github.com/labstack/echo"
)

func Login(c echo.Context) error {

	respone := &models.LoginRespone{"123456789"}
	return c.JSON(http.StatusOK, respone)
}
