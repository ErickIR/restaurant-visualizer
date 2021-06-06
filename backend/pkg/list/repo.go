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
	GetBuyersByDate(date string) ([]models.Buyer, error)
}

type DgraphListRepo struct {
	db      storage.Storage
	context context.Context
}

type BuyerInfoResponse struct {
	BuyerInfo        []dtos.BuyerDto             `json:"buyer,omitempty"`
	Transactions     []dtos.TransactionInfo      `json:"transactions,omitempty"`
	BuyersWithSameIp []dtos.BuyersWithRelatedIps `json:"buyersWithSameIp,omitempty"`
	Products         []models.Product            `json:"top10Products,omitempty"`
}

type BuyersListResponse struct {
	Buyers []models.Buyer `json:"buyers,omitempty"`
}

func NewDgraphListRepo(Db storage.Storage, context context.Context) *DgraphListRepo {
	return &DgraphListRepo{db: Db, context: context}
}

func (dgRepo *DgraphListRepo) GetBuyersByDate(date string) ([]models.Buyer, error) {
	query := `
		query GetBuyersByDate($date: string) {
			buyers(func: eq(date, $date)) @filter(type(Buyer)){
				id
				name
				age
				date
			}
		}
	`

	variables := map[string]string{"$date": date}

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

func (dgRepo *DgraphListRepo) GetAllBuyers(offset, size int) ([]models.Buyer, error) {
	query := `
		query GetAllBuyers($offset: int, $size: int) {
			buyers(func: type(Buyer), offset: $offset, first: $size, orderdesc: date) {
				id
				name
				age
				date
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
				date
			}
				
			transactions(func: eq(buyerId, $buyerId), first: 10){
				id
				ipAddress as ipAddress
				device
				date
				products: bought {
					id
					name
					price as price
				}
				total: sum(val(price))
			}
				
			buyersWithSameIp(func: eq(ipAddress, val(ipAddress)), first: 10) @filter(NOT uid(ipAddress)) 
			{
				device
				ipAddress
				buyer: was_made_by  {
					id
					name
					age
				}
			}

			var(func: eq(id, $buyerId)){
				made {
					bought {
						productsBought as id
					}
				}
			} 
		
			var(func: eq(id, val(productsBought))){
				id
				name
				price
				was_bought {
					id
					bought @filter(NOT uid(productsBought)) {
						productsToBeRecommended as id
					}
				}
			}
				
			var(func: eq(id, val(productsToBeRecommended))){
				id
				total as count(was_bought)
			}
				
			top10Products(func: uid(total), orderdesc: val(total), first: 10){
					id
					name
					price
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

	buyerInfo := dtos.NewBuyerInformation(result, dgraphResponse.Transactions, dgraphResponse.BuyersWithSameIp, dgraphResponse.Products)

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
