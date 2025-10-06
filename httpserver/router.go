package httpserver

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/sonymuhamad/nexttalent-test/config"
	"net/http"
)

func InitRouter(_ config.EnvConfig, h *Handler) http.Handler {
	r := chi.NewRouter()

	r.Use(cors.AllowAll().Handler)

	r.Get("/GetCountry/{Name}", h.GetCountryByPersonName)

	r.Get("/GetCurrentTime", h.GetCurrentTimeWithTimezone)

	return r
}
