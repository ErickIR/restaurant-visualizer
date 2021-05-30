package rest

import (
	"fmt"

	"restaurant-visualizer/pkg/load"

	"restaurant-visualizer/pkg/middlewares"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func InitHandlers(ls load.LoadService) chi.Router {
	fmt.Println("Initializing Handlers")
	router := chi.NewRouter()

	router.Use(
		middleware.Logger,
		middleware.RedirectSlashes,
		middlewares.SetJsonResponseContentType,
	)

	router.Get("/api", WelcomeHandler())

	router.Post("/api/load", LoadData(&ls))

	return router
}
