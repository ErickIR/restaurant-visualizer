package rest

import (
	"encoding/json"
	"net/http"
)

func WelcomeHandler() func(rw http.ResponseWriter, r *http.Request) {
	return func(rw http.ResponseWriter, r *http.Request) {
		json.NewEncoder(rw).Encode("Restaurant Visualizer API")
	}
}
