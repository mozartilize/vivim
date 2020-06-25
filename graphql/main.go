package graphql

import "github.com/labstack/echo/v4"

func Init(app *echo.Echo) {
	graphql := app.Group("/graphql")
	// graphql.Use(EchoContextToContextMiddleware)
	for _, r := range Handlers {
		graphql.Add(r.Method, r.Path, r.Handler)
	}
}
