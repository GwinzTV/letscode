package main

import (
	"log"
	"net/http"

	"letscode/project-01-url-shortener/internal/handlers"
	"letscode/project-01-url-shortener/internal/store"
)

func main() {
	st := store.NewFileStore("data/urls.json")
	h := handlers.NewHandler(st)

	addr := ":8080"
	log.Printf("starting server on %s", addr)
	if err := http.ListenAndServe(addr, h.Routes()); err != nil {
		log.Fatalf("server failed: %v", err)
	}
}
