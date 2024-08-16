package fasthttpServer

import (
	"fmt"
	"io"
	"os"

	"github.com/valyala/fasthttp"
)

func UploadFileHandler(ctx *fasthttp.RequestCtx) {
	if string(ctx.Method()) != "POST" {
		ctx.SetStatusCode(fasthttp.StatusMethodNotAllowed)
		ctx.SetBodyString("Invalid request method")
		return
	}

	_, err := ctx.Request.MultipartForm() // 20 GB limit
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		ctx.SetBodyString("Unable to parse form")
		return
	}

	//form.File

	handler, err := ctx.FormFile("file")
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		ctx.SetBodyString("Error retrieving the file")
		return
	}

	file, err := handler.Open()
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		ctx.SetBodyString("Error open the file")
		return
	}

	dst, err := os.Create("fasthttp/server/downloads/uploaded_" + handler.Filename)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		ctx.SetBodyString("Error saving the file")
		return
	}

	defer dst.Close()

	if _, err = io.Copy(dst, file); err != nil {
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		ctx.SetBodyString("Error saving the file")
		return
	}

	ctx.SetStatusCode(fasthttp.StatusOK)
	ctx.SetBodyString(fmt.Sprintf("File uploaded successfully: %s", handler.Filename))
}
