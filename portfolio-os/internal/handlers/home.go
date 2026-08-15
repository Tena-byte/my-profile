package handlers

import (
	"net/http"

	"portfolio-os/internal/renderer"
)

type HomeHandler struct {
	renderer *renderer.Renderer
}

func NewHomeHandler(renderer *renderer.Renderer) *HomeHandler {
	return &HomeHandler{
		renderer: renderer,
	}
}

func (h *HomeHandler) Handle(w http.ResponseWriter, r *http.Request) {
	data := map[string]any{
		"Title": "Portfolio OS",
	}

	if err := h.renderer.Render(w, "base", data); err != nil {
		http.Error(w, "Failed to render page", http.StatusInternalServerError)
	}
}