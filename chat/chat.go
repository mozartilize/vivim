package chat

import (
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
)

type ChatContext struct {
	echo.Context
	hub *Hub
}

func Init(app *echo.Echo, config *viper.Viper) {
	hub := newHub()
	go hub.run()

	chat := app.Group("/chat")
	chat.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cc := &ChatContext{c, hub}
			return next(cc)
		}
	})
	for _, r := range WsHandlers {
		chat.Add(r.Method, r.Path, r.Handler)
	}
}
