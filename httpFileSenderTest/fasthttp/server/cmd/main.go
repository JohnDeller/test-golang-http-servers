package main

import (
	"fmt"

	"github.com/valyala/fasthttp"

	fasthttpServer "httpFileSenderTest/fasthttp/server"
)

func main() {
	server := &fasthttp.Server{
		Handler: func(ctx *fasthttp.RequestCtx) {
			switch string(ctx.Path()) {
			case "/upload":
				fasthttpServer.UploadFileHandler(ctx)
			default:
				ctx.SetStatusCode(fasthttp.StatusNotFound)
				ctx.SetBodyString("Not Found")
			}
		},
		MaxRequestBodySize: 20 << 30, // 20 GB limit
	}

	fmt.Println("Starting server at :8083")
	if err := server.ListenAndServe(":8083"); err != nil {
		fmt.Printf("Error starting server: %v\n", err)
	}
}
