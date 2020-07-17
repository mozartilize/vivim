package chat

import (
	"fmt"
	"log"
	"net/http"
	"vivim/utils"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
)

var upgrader = websocket.Upgrader{}

// serveWs handles websocket requests from the peer.
func serveWs(c echo.Context) error {
	conn, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	cc := c.(*ChatContext)
	if err != nil {
		log.Println(err)
		return err
	}

	client := &Client{hub: cc.hub, conn: conn, send: make(chan []byte, 256)}
	client.hub.register <- client

	// Allow collection of memory referenced by the caller by doing all work in
	// new goroutines.
	r := make(chan error)
	go client.writePump(r)
	go client.readPump()
	select {
	case err := <-r:
		return err
	}
}

func hello(c echo.Context) error {
	ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		return err
	}
	defer ws.Close()

	for {
		// Write
		err := ws.WriteMessage(websocket.TextMessage, []byte("Hello, Client!"))
		if err != nil {
			c.Logger().Error(err)
			return err
		}

		// Read
		_, msg, err := ws.ReadMessage()
		if err != nil {
			c.Logger().Error(err)
			return ws.Close()
		}
		fmt.Printf("%s\n", msg)
	}
}

var WsHandlers = []utils.RouteDef{
	{Method: http.MethodGet, Path: "/", Handler: serveWs},
	{Method: http.MethodGet, Path: "/echo/", Handler: hello},
}
