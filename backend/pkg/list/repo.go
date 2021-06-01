package list

import (
	"context"
	"encoding/json"
	"fmt"
	dtos "restaurant-visualizer/pkg/dtos/out"
	"restaurant-visualizer/pkg/models"
	"restaurant-visualizer/pkg/storage"
	"strconv"
)

type ListRepo interface {
	GetAllBuyers(page, size int) ([]models.Buyer, error)
	GetTotalBuyersCount() (int, error)
	GetBuyerInformation(buyerId string) (*dtos.BuyerInfo, error)
}

type DgraphListRepo struct {
	db      storage.Storage
	context context.Context
}

type BuyerInfoResponse struct {
	BuyerInfo        []dtos.BuyerInfo            `json:"buyer,omitempty"`
	BuyersWithSameIp []dtos.BuyersWithRelatedIps `json:"buyersWithSameIp,omitempty"`
}

type BuyersListResponse struct {
	Buyers []models.Buyer `json:"buyers,omitempty"`
}

func NewDgraphListRepo(Db storage.Storage, context context.Context) *DgraphListRepo {
	return &DgraphListRepo{db: Db, context: context}
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

func (dgRepo *DgraphListRepo) GetBuyerInformation(buyerId string) (*dtos.BuyerInfo, error) {
	query := `
		query getBuyerInformation($buyerId: string) {
			buyer(func: eq(id, $buyerId)){
				id
				name
				age
				transactions: made {
					id
					ipAddress as ipAddress
					device
					products: bought {
						id
						name
						price as price
					}
					total:  sum(val(price))
				}
			}
				
			buyersWithSameIp(func: eq(ipAddress, val(ipAddress)), first: 10) @filter(NOT uid(ipAddress)){
				device
				ipAddress
				buyer: was_made_by {
					id
					name
					age
				}
			}
		}
	`

	variables := map[string]string{"$buyerId": buyerId}

	resp, err := dgRepo.db.DbClient.NewReadOnlyTxn().QueryWithVars(dgRepo.context, query, variables)

	if err != nil {
		return nil, err
	}

	var dgraphResponse BuyerInfoResponse

	err = json.Unmarshal(resp.Json, &dgraphResponse)

	if err != nil {
		return nil, err
	}

	if len(dgraphResponse.BuyerInfo) == 0 {
		return nil, nil
	}

	result := dgraphResponse.BuyerInfo[0]

	fmt.Println(len(dgraphResponse.BuyersWithSameIp))

	buyerInfo := dtos.NewBuyerInformation(result.Id, result.Name, result.Age, result.Transactions, dgraphResponse.BuyersWithSameIp)

	return buyerInfo, nil
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
