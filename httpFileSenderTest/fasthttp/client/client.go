package fasthttpClient

import (
	"bytes"
	"fmt"
	"github.com/valyala/fasthttp"
	"io"
	"mime/multipart"
	"os"
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

	req := fasthttp.AcquireRequest()
	resp := fasthttp.AcquireResponse()

	defer fasthttp.ReleaseRequest(req)
	defer fasthttp.ReleaseResponse(resp)

	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.SetRequestURI(targetUrl)
	req.Header.SetMethod(fasthttp.MethodPost)
	req.SetBody(body.Bytes())

	err = fasthttp.Do(req, resp)
	if err != nil {
		return err
	}

	if resp.StatusCode() != fasthttp.StatusOK {
		return fmt.Errorf("failed to upload file: %s", string(resp.Body()))
	}

	fmt.Println("File uploaded successfully")

	return nil
}
