package utils

import "github.com/labstack/echo/v4"

type RouteDef struct {
	Method  string
	Path    string
	Handler echo.HandlerFunc
}
