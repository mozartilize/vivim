package vivim

import (
	"net/http"
	"vivim/api"
	"vivim/graphql"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// var Config = readConfig("config", map[string]interface{}{})

func CreateApp() *echo.Echo {
	app := echo.New()
	app.Pre(middleware.AddTrailingSlash())

	app.GET("/routes/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, app.Routes())
	})

	api.Init(app)
	graphql.Init(app)
	return app
}

// if staticDir := Config.Get("STATIC_DIR"); staticDir == nil {
// 	log.Printf("Starting webpack dev server...")
// 	yarnServeCmd := exec.Command("yarn", "serve")
// 	yarnServeCmd.Dir = "static/vivim"
// 	yarnServeCmd.Start()

// 	static := App.Group("/*")
// 	feUrl, _ := url.Parse("http://localhost:8080")
// 	static.Use(
// 		middleware.Proxy(
// 			middleware.NewRoundRobinBalancer([]*middleware.ProxyTarget{
// 				{
// 					URL: feUrl,
// 				},
// 			})))
// } else {
// 	App.File("/favicon.ico", staticDir.(string)+"/favicon.ico")
// 	App.Static("/", staticDir.(string))
// 	App.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
// 		return func(c echo.Context) error {
// 			resp := next(c)
// 			var code int
// 			if he, ok := resp.(*echo.HTTPError); ok {
// 				code = he.Code
// 			}
// 			if code == 404 {
// 				return c.File(staticDir.(string) + "/index.html")
// 			}
// 			return resp
// 		}
// 	})
// }
