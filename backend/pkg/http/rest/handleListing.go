package rest

import (
	"encoding/json"
	"fmt"
	"net/http"
	"restaurant-visualizer/pkg/http/response"
	"restaurant-visualizer/pkg/list"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func WelcomeHandler() func(rw http.ResponseWriter, r *http.Request) {
	return func(rw http.ResponseWriter, r *http.Request) {
		json.NewEncoder(rw).Encode("Restaurant Visualizer API")
	}
}

func ListBuyers(s list.DataEnlister) func(rw http.ResponseWriter, r *http.Request) {
	return func(rw http.ResponseWriter, r *http.Request) {
		var page = 1
		var size = 10

		pageParam := r.URL.Query().Get("page")
		sizeParam := r.URL.Query().Get("size")

		if pageParam != "" {
			if intPage, err := strconv.Atoi(pageParam); err == nil {
				page = intPage
			}
		}

		if sizeParam != "" {
			if intSize, err := strconv.Atoi(sizeParam); err == nil {
				size = intSize
			}
		}

		buyers, err := s.GetAllBuyers(page, size)

		if err != nil {
			response := response.NewFailedResponse(err.Error())
			json.NewEncoder(rw).Encode(response)
			return
		}

		count, err := s.GetBuyersCount()

		if err != nil {
			response := response.NewFailedResponse(err.Error())
			json.NewEncoder(rw).Encode(response)
			return
		}

		totalPages := count / size

		var next string
		var previous string

		if page < totalPages {
			next = fmt.Sprintf("/api/buyer?page=%v&size=%v", page+1, size)
		} else {
			next = ""
		}

		if page > 1 {
			previous = fmt.Sprintf("/api/buyer?page=%v&size=%v", page-1, size)
		} else {
			previous = ""
		}

		response := response.NewPaginatedResponse(buyers, page, size, totalPages, count, "", next, previous)
		json.NewEncoder(rw).Encode(response)
	}

}

func GetBuyerInformation(s list.DataEnlister) func(rw http.ResponseWriter, r *http.Request) {
	return func(rw http.ResponseWriter, r *http.Request) {
		buyerId := chi.URLParam(r, "buyerId")

		buyerInfo, err := s.GetBuyerInformation(buyerId)

		if err != nil {
			response := response.NewFailedResponse(err.Error())
			json.NewEncoder(rw).Encode(response)
		}

		response := response.NewSuccessResponse(buyerInfo, "Data fetched.")
		json.NewEncoder(rw).Encode(response)
	}
}
