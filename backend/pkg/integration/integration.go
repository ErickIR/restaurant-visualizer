package integration

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"restaurant-visualizer/pkg/models"
	"strconv"
	"strings"
)

type ExternalGetter interface {
	GetBuyers(date string) ([]models.Buyer, error)
	GetProducts(date string) ([]models.Product, error)
	GetTransactions(date string) ([]models.Transaction, error)
}

type ExternalService struct {
	client http.Client
}

func NewExternalService(client http.Client) *ExternalService {
	return &ExternalService{client}
}

func (es *ExternalService) GetBuyers(date string) ([]models.Buyer, error) {
	buyersUrl := os.Getenv("BUYERS_URL")

	buyersUrl = fmt.Sprintf("%s?date=%s", buyersUrl, date)

	resp, err := es.client.Get(buyersUrl)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var buyersList []models.Buyer

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &buyersList)

	if err != nil {
		return nil, err
	}

	var validBuyersToLoad []models.Buyer
	duplicate := make(map[string]bool)

	for _, item := range buyersList {
		buyer, err := models.NewBuyer(item.Id, item.Name, item.Age)

		if err != nil {
			log.Println(err)
		}

		exist := duplicate[item.Id]

		if exist {
			continue
		} else {
			duplicate[item.Id] = true
		}

		validBuyersToLoad = append(validBuyersToLoad, *buyer)
	}

	fmt.Println("Data retrieved")

	return validBuyersToLoad, nil
}

func (es *ExternalService) GetProducts(date string) ([]models.Product, error) {
	productsUrl := os.Getenv("PRODUCTS_URL")

	productsUrl = fmt.Sprintf("%s?date=%s", productsUrl, date)

	resp, err := es.client.Get(productsUrl)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	// var productsList []models.Product

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	products, err := convertByteArrayCsvToProductList(body)

	if err != nil {
		return nil, err
	}

	return products, nil
}

func (es *ExternalService) GetTransactions(date string) ([]models.Transaction, error) {
	transactionsUrl := os.Getenv("TRANSACTIONS_URL")

	transactionsUrl = fmt.Sprintf("%s?date=%s", transactionsUrl, date)

	resp, err := es.client.Get(transactionsUrl)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	transactions, err := convertByteArrayNoStandardToTransactionsList(body, date)

	if err != nil {
		return nil, err
	}

	return transactions, nil
}

func convertByteArrayNoStandardToTransactionsList(bytes []byte, date string) ([]models.Transaction, error) {
	var transactions []models.Transaction

	bodyString := string(bytes)

	lines := strings.Split(bodyString, "#")

	for _, line := range lines {
		if len(line) == 0 {
			continue
		}

		transactionData := strings.Split(line, "\x00")

		if len(transactionData) < 5 {
			continue
		}

		productIds := strings.Split(transactionData[4][1:len(transactionData[4])-1], ",")

		transaction, err := models.NewTransaction(transactionData[0], transactionData[1], transactionData[2], transactionData[3], date, productIds)

		if err != nil {
			continue
		}

		transactions = append(transactions, *transaction)
	}

	return transactions, nil
}

func convertByteArrayCsvToProductList(bytes []byte) ([]models.Product, error) {
	var products []models.Product

	csvString := string(bytes)

	csvLines := strings.Split(csvString, "\n")

	for _, line := range csvLines {
		if len(line) == 0 {
			continue
		}

		productData := strings.Split(line, "'")

		if len(productData) != 3 {
			continue
		}

		price, err := strconv.ParseFloat(productData[2], 32)

		if err != nil {
			continue
		}

		product, err := models.NewProduct(productData[0], productData[1], float32(price))

		if err != nil {
			continue
		}

		products = append(products, *product)

	}

	return products, nil
}
