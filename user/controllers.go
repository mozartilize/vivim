package user

import (
	"net/http"
	"vivim/user/repository"
	"vivim/utils"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

func getUser(c echo.Context) error {
	u, err := repository.GetUserById(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusNotFound, nil)
	}
	return c.JSON(http.StatusOK, u)
}

func foo(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["name"].(string)
	return c.JSON(http.StatusOK, map[string]string{"username": name})
}

var ApiHandlers = []utils.RouteDef{
	{Method: http.MethodGet, Path: "/users/:id/", Handler: getUser},
	{Method: http.MethodGet, Path: "/foo/", Handler: foo},
}
