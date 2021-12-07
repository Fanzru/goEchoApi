package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Init() *echo.Echo{
	e := echo.New()
	e = InitUser(e)
	e.Debug = true
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	return e
}