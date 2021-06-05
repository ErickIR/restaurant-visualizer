package dtos

import "restaurant-visualizer/pkg/models"

type BuyerInfo struct {
	Buyer            BuyerDto               `json:"buyer,omitempty"`
	Transactions     []TransactionInfo      `json:"transactions,omitempty"`
	BuyersWithSameIp []BuyersWithRelatedIps `json:"buyerWithSameIp,omitempty"`
	Products         []models.Product       `json:"products,omitempty"`
}

type BuyerDto struct {
	Id   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Age  int    `json:"age,omitempty"`
	Date string `json:"date,omitempty"`
}

type TransactionInfo struct {
	Id        string           `json:"id,omitempty"`
	IpAddress string           `json:"ipAddress,omitempty"`
	Device    string           `json:"device,omitempty"`
	Total     float64          `json:"total,omitempty"`
	Products  []models.Product `json:"products,omitempty"`
}

type BuyersWithRelatedIps struct {
	Id        string   `json:"id,omitempty"`
	Device    string   `json:"device,omitempty"`
	IpAddress string   `json:"ipAddress,omitempty"`
	BuyerInfo BuyerDto `json:"buyer,omitempty"`
}

func NewBuyerInformation(buyer BuyerDto, transactions []TransactionInfo, buyersSharingIp []BuyersWithRelatedIps, products []models.Product) *BuyerInfo {
	return &BuyerInfo{
		Buyer:            buyer,
		Transactions:     transactions,
		BuyersWithSameIp: buyersSharingIp,
		Products:         products,
	}
}
