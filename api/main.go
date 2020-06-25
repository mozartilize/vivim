package api

import (
	"vivim/user"

	"github.com/labstack/echo/v4"
)

func Init(app *echo.Echo) {
	ApiV1 := app.Group("/api/v1")
	for _, r := range user.ApiHandlers {
		ApiV1.Add(r.Method, r.Path, r.Handler)
	}
}
