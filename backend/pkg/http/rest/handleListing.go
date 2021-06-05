package rest

import (
	"encoding/json"
	"errors"
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

func ListBuyersByDate(s list.DataEnlister) func(rw http.ResponseWriter, r *http.Request) {
	return func(rw http.ResponseWriter, r *http.Request) {
		date := r.URL.Query().Get("date")

		if date == "" {
			response := response.NewFailedResponse(errors.New("date cannot be empty").Error())
			json.NewEncoder(rw).Encode(response)
		}

		buyers, err := s.GetBuyersByDate(date)

		if err != nil {
			response := response.NewFailedResponse(err.Error())
			json.NewEncoder(rw).Encode(response)
			return
		}

		var message string
		if len(buyers) == 0 {
			message = "There is no data available."
		} else {
			message = "Data fetched."
		}

		response := response.NewSuccessResponse(buyers, message)
		json.NewEncoder(rw).Encode(response)
	}
}

func ListBuyers(s list.DataEnlister) func(rw http.ResponseWriter, r *http.Request) {
	return func(rw http.ResponseWriter, r *http.Request) {
		var page = 1
		var size = 30

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
			next = fmt.Sprintf("/buyer?page=%v&size=%v", page+1, size)
		} else {
			next = ""
		}

		if page > 1 {
			previous = fmt.Sprintf("/buyer?page=%v&size=%v", page-1, size)
		} else {
			previous = ""
		}

		var message string
		if len(buyers) == 0 {
			message = "There is no data available."
		} else {
			message = "Data fetched."
		}

		response := response.NewPaginatedResponse(buyers, page, size, totalPages, count, message, next, previous)
		json.NewEncoder(rw).Encode(response)
	}

}

func GetBuyerInformation(s list.DataEnlister) func(rw http.ResponseWriter, r *http.Request) {
	return func(rw http.ResponseWriter, r *http.Request) {
		buyerId := chi.URLParam(r, "buyerId")

		if buyerId == "" {
			response := response.NewFailedResponse(errors.New("buyerId cannot be empty").Error())
			json.NewEncoder(rw).Encode(response)
		}

		buyerInfo, err := s.GetBuyerInformation(buyerId)

		if err != nil {
			response := response.NewFailedResponse(err.Error())
			json.NewEncoder(rw).Encode(response)
			return
		}

		response := response.NewSuccessResponse(buyerInfo, "Data fetched.")
		json.NewEncoder(rw).Encode(response)
	}
}
