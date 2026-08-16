package middleware

import (
	"net/http"

	"portfolio-os/internal/services"
)

func Analytics(
	analyticsService *services.AnalyticsService,
) func(http.Handler) http.Handler {

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			next.ServeHTTP(w, r)
		})
	}
}