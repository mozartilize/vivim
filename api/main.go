package api

import (
	"vivim/auth"
	"vivim/user"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/viper"
)

func Init(app *echo.Echo, config *viper.Viper) {
	ApiV1 := app.Group("/api/v1")

	AuthenticatedApiV1 := ApiV1.Group("")

	AuthenticatedApiV1.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey:  []byte(config.GetString("secret_key")),
		TokenLookup: "query:token",
	}))

	for _, r := range auth.UnauthApiHandlers {
		ApiV1.Add(r.Method, r.Path, r.Handler)
	}
	for _, r := range user.ApiHandlers {
		AuthenticatedApiV1.Add(r.Method, r.Path, r.Handler)
	}
}
