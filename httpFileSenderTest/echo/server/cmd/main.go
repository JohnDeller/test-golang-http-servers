package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	echoServer "httpFileSenderTest/echo/server"
)

func main() {
	e := echo.New()
	e.Use(middleware.BodyLimit("20GI")) // 20 GB limit

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.POST("/upload", echoServer.UploadFileHandler)

	e.Logger.Fatal(e.Start(":8084"))
}
