package static

import (
	"log"
	"net/url"
	"os/exec"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/viper"
)

// Init static app
func Init(app *echo.Echo, config *viper.Viper) {
	if staticDir := config.Get("STATIC_DIR"); staticDir == nil {
		log.Printf("Starting webpack dev server...")
		yarnServeCmd := exec.Command("yarn", "serve")
		yarnServeCmd.Dir = "static/vivim"
		yarnServeCmd.Start()

		feURL, _ := url.Parse("http://localhost:8080")

		static := app.Group("/*")
		static.Use(
			middleware.Proxy(
				middleware.NewRoundRobinBalancer([]*middleware.ProxyTarget{
					{
						URL: feURL,
					},
				})))
	} else {
		static := app.Group("/")
		static.Static("assets", staticDir.(string))
		static.GET("*", func(c echo.Context) error {
			return c.File(staticDir.(string) + "/index.html")
		})
	}
}
