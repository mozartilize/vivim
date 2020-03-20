package main

import (
	"log"
	"net/http"
	"net/url"
	"os/exec"

	"context"

	"vivim/graph"
	"vivim/graph/generated"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

type person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
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
			generated.Config{Resolvers: &graph.Resolver{}}))
	return func(c echo.Context) error {
		handler.ServeHTTP(c.Response(), c.Request())
		return nil
	}
}

type GqlgenContext struct {
	echo.Context
	ctx context.Context
}

func EchoContextToContextMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := context.WithValue(c.Request().Context(), "EchoContextKey", c)
		c.SetRequest(c.Request().WithContext(ctx))

		cc := &GqlgenContext{c, ctx}

		return next(cc)
	}
}

func main() {
	config := readConfig("config", map[string]interface{}{})

	e := echo.New()
	e.Pre(middleware.AddTrailingSlash())

	graphql := e.Group("/graphql")
	// graphql.Use(EchoContextToContextMiddleware)
	graphql.GET("/", playgroundHandler())
	graphql.POST("/query/", graphqlHandler())

	api := e.Group("/api/v1")
	api.GET("/", func(c echo.Context) error {
		u := &person{
			Name: "world",
			Age:  10000,
		}
		return c.JSON(http.StatusOK, u)
	})

	if staticDir := config.Get("STATIC_DIR"); staticDir == nil {
		log.Printf("Starting webpack dev server...")
		yarnServeCmd := exec.Command("yarn", "serve")
		yarnServeCmd.Dir = "static/vivim"
		yarnServeCmd.Start()

		static := e.Group("/*")
		feUrl, _ := url.Parse("http://localhost:8080")
		static.Use(
			middleware.Proxy(
				middleware.NewRoundRobinBalancer([]*middleware.ProxyTarget{
					{
						URL: feUrl,
					},
				})))
	} else {
		e.File("/favicon.ico", staticDir.(string)+"/favicon.ico")
		e.Static("/", staticDir.(string))
		e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
			return func(c echo.Context) error {
				resp := next(c)
				var code int
				if he, ok := resp.(*echo.HTTPError); ok {
					code = he.Code
				}
				if code == 404 {
					return c.File(staticDir.(string) + "/index.html")
				}
				return resp
			}
		})
	}

	e.Logger.Fatal(e.Start(":1323"))
}
