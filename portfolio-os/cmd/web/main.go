package main

import (
	"log"
	"net/http"

	"portfolio-os/internal/handlers"
	"portfolio-os/internal/renderer"
)

func main() {
	renderer, err := renderer.New()
	if err != nil {
		log.Fatal(err)
	}

	homeHandler := handlers.NewHomeHandler(renderer)

	mux := http.NewServeMux()

	mux.HandleFunc("GET /", homeHandler.Handle)

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	log.Println("Server running on http://localhost:8080")

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}