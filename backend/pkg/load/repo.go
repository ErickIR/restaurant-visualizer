package load

import (
	"context"
	"encoding/json"
	"fmt"
	"restaurant-visualizer/pkg/models"
	"restaurant-visualizer/pkg/storage"

	"github.com/dgraph-io/dgo/v2/protos/api"
)

type LoadRepo interface {
	SaveData(entity []byte) error
	FilterDuplicateBuyers(buyers []models.Buyer) ([]models.Buyer, error)
	FilterDuplicateProducts(products []models.Product) ([]models.Product, error)
	FilterDuplicateTransactions(transactions []models.Transaction) ([]models.Transaction, error)
	SetBuyersToTransactions(transactions []models.Transaction) error
}

type DgraphLoadRepo struct {
	db      storage.Storage
	context context.Context
}

type BuyersListResponse struct {
	Buyers []models.Buyer `json:"buyers,omitempty"`
}

type ProductsListResponse struct {
	Products []models.Product `json:"products,omitempty"`
}

type TransactionsListResponse struct {
	Transactions []models.Transaction `json:"transactions,omitempty"`
}

func NewDgraphLoadRepo(Db storage.Storage, context context.Context) *DgraphLoadRepo {
	return &DgraphLoadRepo{db: Db, context: context}
}

func (dgRepo *DgraphLoadRepo) SaveData(entity []byte) error {
	mutation := &api.Mutation{
		CommitNow: true,
	}

	mutation.SetJson = entity

	_, err := dgRepo.db.DbClient.NewTxn().Mutate(dgRepo.context, mutation)

	if err != nil {
		// log.Fatalf("Error running mutation: %v", err)
		return err
	}

	return nil
}

func (dgRepo *DgraphLoadRepo) FilterDuplicateBuyers(buyers []models.Buyer) ([]models.Buyer, error) {
	query := `
		query AllBuyersWithIds($idList: string) {
			buyers(func: anyofterms(id, $idList)) {
				uid
				id
				name
				age
			}
		}
	`
	var idList string

	for _, item := range buyers {
		idList += item.Id + " "
	}

	variables := map[string]string{"$idList": idList}

	resp, err := dgRepo.db.DbClient.NewReadOnlyTxn().QueryWithVars(dgRepo.context, query, variables)

	if err != nil {
		return nil, err
	}

	var dgraphResponse BuyersListResponse

	err = json.Unmarshal(resp.Json, &dgraphResponse)

	if err != nil {
		// log.Fatalf("Error parsing from JSON: %v", err)
		return nil, err
	}

	var filteredBuyers []models.Buyer
	buyerExist := make(map[string]bool)

	for _, buyer := range dgraphResponse.Buyers {
		buyerExist[buyer.Id] = true
	}

	for _, buyerToCheck := range buyers {
		exists := buyerExist[buyerToCheck.Id]

		if exists {
			continue
		}

		filteredBuyers = append(filteredBuyers, buyerToCheck)
	}

	return filteredBuyers, nil
}

func (dgRepo *DgraphLoadRepo) FilterDuplicateProducts(products []models.Product) ([]models.Product, error) {
	query := `
		query AllProductsWithIds($idList: string) {
			products(func: anyofterms(id, $idList)) {
				uid
				id
				name
				price
			}
		}
	`

	var idList string

	for _, item := range products {
		idList += item.Id + " "
	}

	variables := map[string]string{"$idList": idList}

	resp, err := dgRepo.db.DbClient.NewReadOnlyTxn().QueryWithVars(dgRepo.context, query, variables)

	if err != nil {
		return nil, err
	}

	var dgraphResponse ProductsListResponse

	err = json.Unmarshal(resp.Json, &dgraphResponse)

	if err != nil {
		return nil, err
	}

	var filteredProducts []models.Product
	productExist := make(map[string]bool)

	for _, product := range dgraphResponse.Products {
		productExist[product.Id] = true
	}

	for _, productToCheck := range products {
		exists := productExist[productToCheck.Id]

		if exists {
			continue
		}

		filteredProducts = append(filteredProducts, productToCheck)
	}

	return filteredProducts, nil
}

func (dgRepo *DgraphLoadRepo) FilterDuplicateTransactions(transactions []models.Transaction) ([]models.Transaction, error) {
	query := `
		query AllTransactionsWithIds($idList: string) {
			transactions(func: anyofterms(id, $idList)) {
				uid
				id
			}
		}
	`
	var idList string

	for _, item := range transactions {
		idList += item.Id + " "
	}

	variables := map[string]string{"$idList": idList}

	resp, err := dgRepo.db.DbClient.NewReadOnlyTxn().QueryWithVars(dgRepo.context, query, variables)

	if err != nil {
		return nil, err
	}

	var dgraphResponse TransactionsListResponse

	err = json.Unmarshal(resp.Json, &dgraphResponse)

	if err != nil {
		return nil, err
	}

	var filteredTransactions []models.Transaction
	transactionsExists := make(map[string]bool)

	for _, trans := range dgraphResponse.Transactions {
		transactionsExists[trans.Id] = true
	}

	for _, transToCheck := range transactions {
		exists := transactionsExists[transToCheck.Id]

		if exists {
			continue
		}

		filteredTransactions = append(filteredTransactions, transToCheck)
	}

	return filteredTransactions, nil
}

func (dgRepo *DgraphLoadRepo) SetBuyersToTransactions(transactions []models.Transaction) error {
	txn := dgRepo.db.DbClient.NewTxn()

	for _, transaction := range transactions {
		query := `
			query {
				buyer as var(func: eq(id, "%s"))
				transaction as var(func: eq(id, "%s"))
			}
		`

		formatQuery := fmt.Sprintf(query, transaction.BuyerId, transaction.Id)
		mutation := &api.Mutation{
			SetNquads: []byte(`uid(transaction) <buyer> uid(buyer) .`),
		}

		req := &api.Request{
			Query:     formatQuery,
			Mutations: []*api.Mutation{mutation},
		}

		_, err := txn.Do(dgRepo.context, req)

		if err != nil {
			return err
		}
	}

	err := txn.Commit(dgRepo.context)

	if err != nil {
		return err
	}

	return nil
}
