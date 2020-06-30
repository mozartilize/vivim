package api

import (
	"vivim/user"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/viper"
)

func Init(app *echo.Echo, config *viper.Viper) {
	ApiV1 := app.Group("/api/v1")

	ApiV1.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey:  []byte(config.GetString("secret_key")),
		TokenLookup: "query:token",
	}))

	for _, r := range user.ApiHandlers {
		ApiV1.Add(r.Method, r.Path, r.Handler)
	}
}
