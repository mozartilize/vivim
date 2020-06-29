package graphql

import (
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
)

func Init(app *echo.Echo, config *viper.Viper) {
	graphql := app.Group("/graphql")
	// graphql.Use(EchoContextToContextMiddleware)
	for _, r := range Handlers {
		graphql.Add(r.Method, r.Path, r.Handler)
	}
}
