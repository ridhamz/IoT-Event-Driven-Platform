package api

import (
	"github.com/go-chi/chi/v5"
)

var router *chi.Mux

func InitRouter() *chi.Mux {
	if router == nil {
		router = chi.NewRouter()
	}

	SetupAuthRoutes()
	return router
}

func GetRouter() *chi.Mux {
	return router
}
