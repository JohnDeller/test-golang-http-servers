package fiberServer

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/gofiber/fiber/v2"
)

func UploadFileHandler(c *fiber.Ctx) error {
	if c.Method() != http.MethodPost {
		return c.Status(fiber.StatusMethodNotAllowed).SendString("Invalid request method")
	}

	file, err := c.FormFile("file")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Error retrieving the file")
	}

	src, err := file.Open()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Error opening the file")
	}

	defer src.Close()

	dst, err := os.Create("fiber/server/downloads/uploaded_" + file.Filename)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Error saving the file")
	}

	defer dst.Close()

	if _, err = io.Copy(dst, src); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Error saving the file")
	}

	return c.SendString(fmt.Sprintf("File uploaded successfully: %s", file.Filename))
}
