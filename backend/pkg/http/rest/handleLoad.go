package rest

import (
	"encoding/json"
	"net/http"
	"restaurant-visualizer/pkg/http/response"
	"restaurant-visualizer/pkg/load"
)

func LoadData(loadService load.DataLoader) func(rw http.ResponseWriter, r *http.Request) {
	return func(rw http.ResponseWriter, r *http.Request) {

		date := r.URL.Query().Get("date")

		err := loadService.LoadData(date)

		if err != nil {
			response := response.NewFailedResponse(err.Error())
			json.NewEncoder(rw).Encode(response)
		} else {
			response := response.NewSuccessResponse(nil, "Data Loaded Successfully.")
			json.NewEncoder(rw).Encode(response)
		}
	}
}
