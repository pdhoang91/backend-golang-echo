package mdw

import "github.com/labstack/echo"

func BasicAuth(username, password string, c echo.Context) (bool, error) {
	if username != "admin" || password != "123123" {
		return false, nil
		//return false, nil
	}
	return true, nil
}
