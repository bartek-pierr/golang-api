package application

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/bartek-pierr/golang-api.git/handler"
)

func loadRoutes() *chi.Mux {
	router := chi.NewRouter()
	router.Use(middleware.Logger)

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	router.Route("/logos", loadLogoRoutes)

	return router
}

func loadLogoRoutes(router chi.Router) {
	logoHandler := &handler.Logo{}

	router.Post("/", logoHandler.Create)
	router.Get("/", logoHandler.List)
	router.Get("/{id}", logoHandler.GetByID)
	router.Put("/{id}", logoHandler.UpdateByID)
	router.Delete("/{id}", logoHandler.DeleteByID)
}
