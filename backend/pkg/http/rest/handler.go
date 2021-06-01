package rest

import (
	"fmt"

	"restaurant-visualizer/pkg/list"
	"restaurant-visualizer/pkg/load"

	"restaurant-visualizer/pkg/middlewares"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func InitHandlers(loadService load.LoadService, listService list.ListService) chi.Router {
	fmt.Println("Initializing Handlers")
	router := chi.NewRouter()

	router.Use(
		middleware.Logger,
		middleware.RedirectSlashes,
		middlewares.SetJsonResponseContentType,
	)

	router.Get("/api", WelcomeHandler())

	router.Get("/api/buyer", ListBuyers(&listService))
	router.Post("/api/load", LoadData(&loadService))

	return router
}
