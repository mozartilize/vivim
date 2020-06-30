package user

import (
	"net/http"
	"vivim/user/repository"
	"vivim/utils"

	"github.com/labstack/echo/v4"
)

func getUser(c echo.Context) error {
	u, err := repository.GetUserById(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusNotFound, nil)
	}
	return c.JSON(http.StatusOK, u)
}

var ApiHandlers = []utils.RouteDef{
	{Method: http.MethodGet, Path: "/users/:id/", Handler: getUser},
}
