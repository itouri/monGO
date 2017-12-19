package main

import (

	//"fmt"

	"./handlers"
	"./interceptor"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// これを入れないと OPTION のメソッドがさばけずエラーになる
	// TODO これがない時のエラー原因を理解する
	e.Use(middleware.CORS())

	e.GET("/api/spot/:id", handlers.GetSpot)
	e.POST("/api/spot", handlers.PostSpot)

	e.GET("/calc", handlers.GetCalc)

	// basic auth
	e.GET("/secret", handlers.GetSecret, interceptor.BasicAuth())

	e.Static("/", "public")

	// Start server
	//e.Run(standard.New(":1323"))
	e.Start(":1323")
}
