package handlers

import (
	"net/http"

	"portfolio-os/internal/models"
	"portfolio-os/internal/renderer"
)

type HomeHandler struct {
	renderer  *renderer.Renderer
	portfolio *models.Portfolio
}

func NewHomeHandler(
	renderer *renderer.Renderer,
	portfolio *models.Portfolio,
) *HomeHandler {
	return &HomeHandler{
		renderer:  renderer,
		portfolio: portfolio,
	}
}

func (h *HomeHandler) Handle(w http.ResponseWriter, r *http.Request) {
	data := map[string]any{
		"Title":     h.portfolio.Profile.Name + " — Portfolio",
		"Portfolio": h.portfolio,
	}

	if err := h.renderer.Render(w, "base", data); err != nil {
		http.Error(w, "Failed to render page", http.StatusInternalServerError)
	}
}