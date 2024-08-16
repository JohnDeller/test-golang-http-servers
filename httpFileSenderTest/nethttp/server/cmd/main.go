package main

import (
	"fmt"
	"net/http"

	nethttpServer "httpFileSenderTest/nethttp/server"
)

func main() {
	http.HandleFunc("/upload", nethttpServer.UploadFileHandler)
	fmt.Println("Starting server at :8081")
	http.ListenAndServe(":8081", nil)
}
