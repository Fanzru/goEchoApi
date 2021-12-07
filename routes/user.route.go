package routes

import (
	"github.com/labstack/echo/v4"
	"goEchoApi/controllers"
)



func InitUser(e *echo.Echo) *echo.Echo{
	e.GET("/", controllers.GetUser)
	e.POST("/post",controllers.Create)
	return e
}


