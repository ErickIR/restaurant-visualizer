package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"restaurant-visualizer/pkg/http/rest"
	"restaurant-visualizer/pkg/integration"
	"restaurant-visualizer/pkg/load"
	"restaurant-visualizer/pkg/storage"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("../../.env")

	if err != nil {
		log.Fatal(err.Error())
	}

	port := os.Getenv("PORT")

	db, err := storage.NewClient()

	if err != nil {
		log.Fatalf("Error creating a new DGraph Client: %v", err)
	}

	client := http.Client{}
	context := context.Background()

	loadRepo := load.NewDgraphLoadRepo(*db, context)

	externalService := integration.NewExternalService(client)
	loadService := load.NewService(loadRepo, externalService)

	router := rest.InitHandlers(*loadService)

	log.Printf("Starting server on http://localhost%s/api/\n", port)
	log.Fatal(http.ListenAndServe(port, router))
}
