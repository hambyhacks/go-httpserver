package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func fsHandler() http.Handler {
	var w http.ResponseWriter
	var rw *http.Request
	r := chi.NewRouter()

	// Use logging middleware
	r.Use(middleware.Logger)
	r.Handle("/*", serveFiles(w, rw))

	return r
}

func serveFiles(w http.ResponseWriter, r *http.Request) http.Handler {
	// Get current working directory
	dir, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}

	files := http.FileServer(http.Dir(dir))
	return files
}
