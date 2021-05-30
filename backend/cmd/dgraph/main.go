package main

import (
	"log"
	"restaurant-visualizer/pkg/dgraph"
	"restaurant-visualizer/pkg/storage"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	Db, err := storage.NewClient()

	if err != nil {
		log.Fatalf("Error connecting to DGraph: %v", err)
	}

	defer Db.Cancel()

	err = dgraph.LoadSchema(Db)

	if err != nil {
		log.Fatalf("Error loading schema: %v", err)
	}

	log.Println("Schema successfully loaded .")
}
