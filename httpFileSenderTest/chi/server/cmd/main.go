package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	chiServer "httpFileSenderTest/chi/server"
)

func main() {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestSize(20 << 30)) // 20 GB limit

	r.Post("/upload", chiServer.UploadFileHandler)

	fmt.Println("Starting server at :8085")
	http.ListenAndServe(":8085", r)
}
