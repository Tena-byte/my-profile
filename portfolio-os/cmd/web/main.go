package main

import (
	"encoding/json"
	"io/fs"
	"log"
	"net/http"

	"portfolio-os/internal/handlers"
	"portfolio-os/internal/renderer"
	"portfolio-os/internal/services"
	"portfolio-os/internal/web"
)

func main() {

	staticFS, err := fs.Sub(web.Files, "static")
	if err != nil {
		log.Fatal(err)
	}

	staticHandler := http.FileServer(http.FS(staticFS))

	renderer, err := renderer.New(web.Files)
	if err != nil {
		log.Fatal(err)
	}

	portfolioService, err := services.NewPortfolioService("data/portfolio.json")
	if err != nil {
		log.Fatal(err)
	}

	homeHandler := handlers.NewHomeHandler(
		renderer,
		portfolioService.GetPortfolio(),
	)
	mux := http.NewServeMux()

	mux.Handle(
		"GET /static/",
		http.StripPrefix("/static/", staticHandler),
	)

	mux.HandleFunc("GET /", homeHandler.Handle)

	mux.HandleFunc("GET /debug/portfolio", func(w http.ResponseWriter, r *http.Request) {
		portfolio := portfolioService.GetPortfolio()

		w.Header().Set("Content-Type", "application/json")

		if err := json.NewEncoder(w).Encode(portfolio); err != nil {
			http.Error(w, "Failed to encode portfolio", http.StatusInternalServerError)
		}
	})

	server := &http.Server{
		Addr:    ":3000",
		Handler: mux,
	}

	log.Println("Server running on http://localhost:3000")

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
