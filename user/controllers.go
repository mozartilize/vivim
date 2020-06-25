package user

import (
	"net/http"
	"vivim/user/models"
	"vivim/utils"

	"github.com/labstack/echo/v4"
)

func getUser(c echo.Context) error {
	u := &models.User{
		ID:       "18446744073709551615",
		Username: "foo",
	}
	return c.JSON(http.StatusOK, u)
}

var ApiHandlers = []utils.RouteDef{
	{Method: http.MethodGet, Path: "/users/:id/", Handler: getUser},
}
