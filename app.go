package vivim

import (
	"net/http"
	"vivim/api"
	"vivim/chat"
	"vivim/db"
	"vivim/graphql"
	"vivim/static"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// CreateApp vivim create app factory
func CreateApp() *echo.Echo {
	app := echo.New()
	config := readConfig(".env", map[string]interface{}{
		"DATABASE_POOL_SIZE":       5,
		"DATABASE_POOL_RECYCLE":    10 * 60,
		"DATABASE_MAX_CONNECTIONS": 15,
		"SECRET_KEY":               "something secret",
	})

	app.Pre(middleware.AddTrailingSlash())
	app.Use(middleware.Logger())

	app.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set("config", config)
			return next(c)
		}
	})

	app.GET("/routes/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, app.Routes())
	})

	static.Init(app, config)
	api.Init(app, config)
	graphql.Init(app, config)
	chat.Init(app, config)

	db, err := db.GetDb(config)
	if err != nil {
		panic(err)
	}
	if r, err := db.Queryx("select 1"); err != nil {
		panic(err)
	} else {
		r.Close()
	}

	return app
}
