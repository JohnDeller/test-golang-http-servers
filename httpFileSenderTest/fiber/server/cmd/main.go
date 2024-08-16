package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"

	fiberServer "httpFileSenderTest/fiber/server"
)

func main() {
	app := fiber.New(fiber.Config{
		BodyLimit: 20 << 30, // 20 GB limit
	})

	app.Post("/upload", fiberServer.UploadFileHandler)

	fmt.Println("Starting server at :8082")
	app.Listen(":8082")
}
