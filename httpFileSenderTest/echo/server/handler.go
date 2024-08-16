package echoServer

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

func UploadFileHandler(c echo.Context) error {
	if c.Request().Method != http.MethodPost {
		return c.String(http.StatusMethodNotAllowed, "Invalid request method")
	}

	file, err := c.FormFile("file")
	if err != nil {
		return c.String(http.StatusBadRequest, "Error retrieving the file")
	}

	src, err := file.Open()
	if err != nil {
		return c.String(http.StatusInternalServerError, "Error opening the file")
	}

	defer src.Close()

	dst, err := os.Create("echo/server/downloads/uploaded_" + file.Filename)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Error saving the file")
	}

	defer dst.Close()

	if _, err = io.Copy(dst, src); err != nil {
		return c.String(http.StatusInternalServerError, "Error saving the file")
	}

	return c.String(http.StatusOK, fmt.Sprintf("File uploaded successfully: %s", file.Filename))
}
