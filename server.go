package main

import (
	"github.com/iamsangil/chat/app/controller"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Static("/", "./")
	e.GET("/ws", controller.ChatController)
	e.Logger.Fatal(e.Start(":1323"))
}
