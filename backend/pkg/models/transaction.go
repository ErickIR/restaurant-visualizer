package models

type Transaction struct {
	Uid        string   `json:"uid,omitempty"`
	Id         string   `json:"id,omitempty"`
	BuyerId    string   `json:"buyerId,omitempty"`
	Buyer      Buyer    `json:"buyer,omitempty"`
	IpAddress  string   `json:"ipAddress,omitempty"`
	Device     string   `json:"device,omitempty"`
	ProductIds []string `json:"productIds,omitempty"`
	DType      []string `json:"dgraph.type,omitempty"`
}

func NewTransaction(id string, buyerId string, ipAddress string, device string, productIds []string) (*Transaction, error) {
	return &Transaction{
		Id:         id,
		BuyerId:    buyerId,
		IpAddress:  ipAddress,
		Device:     device,
		ProductIds: productIds,
		DType:      []string{"Transaction"},
	}, nil
}
