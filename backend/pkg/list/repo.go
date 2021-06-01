package list

import (
	"context"
	"encoding/json"
	"fmt"
	"restaurant-visualizer/pkg/models"
	"restaurant-visualizer/pkg/storage"
	"strconv"
)

type ListRepo interface {
	// Query(query string, variables map[string]string) (interface{}, error)
	GetAllBuyers(page, size int) ([]models.Buyer, error)
	GetTotalBuyersCount() (int, error)
}

type DgraphListRepo struct {
	db      storage.Storage
	context context.Context
}

type BuyersListResponse struct {
	Buyers []models.Buyer `json:"buyers,omitempty"`
}

func NewDgraphListRepo(Db storage.Storage, context context.Context) *DgraphListRepo {
	return &DgraphListRepo{db: Db, context: context}
}

func (dgRepo *DgraphListRepo) Query(query string, variables map[string]string) (interface{}, error) {
	resp, err := dgRepo.db.DbClient.NewReadOnlyTxn().QueryWithVars(dgRepo.context, query, variables)

	if err != nil {
		return nil, err
	}

	var result interface{}

	err = json.Unmarshal(resp.Json, &result)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (dgRepo *DgraphListRepo) GetAllBuyers(offset, size int) ([]models.Buyer, error) {
	query := `
		query GetAllBuyers($offset: int, $size: int) {
			buyers(func: type(Buyer), offset: $offset, first: $size) {
				id
				name
				age
			}
		}
	`

	variables := map[string]string{"$offset": fmt.Sprint(offset), "$size": fmt.Sprint(size)}

	resp, err := dgRepo.db.DbClient.NewReadOnlyTxn().QueryWithVars(dgRepo.context, query, variables)

	if err != nil {
		return nil, err
	}

	var dgraphResponse BuyersListResponse

	err = json.Unmarshal(resp.Json, &dgraphResponse)

	if err != nil {
		return nil, err
	}

	return dgraphResponse.Buyers, nil
}

func (dgRepo *DgraphListRepo) GetTotalBuyersCount() (int, error) {
	query := `
		query {
			total(func: type(Buyer))  {
				count: count(uid)
			}
		}
	`

	resp, err := dgRepo.db.DbClient.NewReadOnlyTxn().Query(dgRepo.context, query)

	if err != nil {
		return 0, err
	}

	json := string(resp.Json)

	countStr := json[19 : len(json)-3]

	count, err := strconv.Atoi(countStr)

	if err != nil {
		return 0, err
	}

	return count, nil
}
