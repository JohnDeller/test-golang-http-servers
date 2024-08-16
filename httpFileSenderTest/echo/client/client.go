package echoClient

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"os"

	"github.com/go-resty/resty/v2"
	"github.com/valyala/fasthttp"
)

func UploadFile(filename string, targetUrl string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	part, err := writer.CreateFormFile("file", filename)
	if err != nil {
		return err
	}

	_, err = io.Copy(part, file)
	if err != nil {
		return err
	}

	err = writer.Close()
	if err != nil {
		return err
	}

	client := resty.New()

	resp, err := client.R().
		SetHeader("Content-Type", writer.FormDataContentType()).
		SetBody(body.Bytes()).
		Post(targetUrl)
	if err != nil {
		return err
	}

	if resp.StatusCode() != fasthttp.StatusOK {
		return fmt.Errorf("failed to upload file: %s", string(resp.Body()))
	}

	fmt.Println("File uploaded successfully")

	return nil
}
