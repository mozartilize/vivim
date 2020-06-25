package main

import (
	"vivim"
)

func main() {
	app := vivim.CreateApp()
	app.Logger.Fatal(app.Start(":1323"))
}
