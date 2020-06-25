package graphql

import (
	"context"
	"fmt"
	"net/http"
	"vivim/graphql/generated"
	"vivim/utils"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/labstack/echo/v4"
)

type GqlgenContext struct {
	echo.Context
	ctx context.Context
}

func EchoContextFromContext(ctx context.Context) (*echo.Context, error) {
	echoContext := ctx.Value("EchoContextKey")
	if echoContext == nil {
		err := fmt.Errorf("could not retrieve echo.Context")
		return nil, err
	}

	ec, ok := echoContext.(*echo.Context)
	if !ok {
		err := fmt.Errorf("echo.Context has wrong type")
		return nil, err
	}
	return ec, nil
}

func EchoContextToContextMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := context.WithValue(c.Request().Context(), "EchoContextKey", c)
		c.SetRequest(c.Request().WithContext(ctx))

		cc := &GqlgenContext{c, ctx}

		return next(cc)
	}
}

func playgroundHandler() echo.HandlerFunc {
	handler := playground.Handler("GraphQL Playground", "/graphql/query/")
	return func(c echo.Context) error {
		handler.ServeHTTP(c.Response(), c.Request())
		return nil
	}
}

func graphqlHandler() echo.HandlerFunc {
	handler := handler.NewDefaultServer(
		generated.NewExecutableSchema(
			generated.Config{Resolvers: &Resolver{}}))
	return func(c echo.Context) error {
		handler.ServeHTTP(c.Response(), c.Request())
		return nil
	}
}

var Handlers = []utils.RouteDef{
	{Method: http.MethodGet, Path: "/", Handler: playgroundHandler()},
	{Method: http.MethodPost, Path: "/query/", Handler: graphqlHandler()},
}
