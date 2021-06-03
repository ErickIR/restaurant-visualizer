package models

type Transaction struct {
	Uid        string    `json:"uid,omitempty"`
	Id         string    `json:"id,omitempty"`
	BuyerId    string    `json:"buyerId,omitempty"`
	Buyer      Buyer     `json:"was_made_by,omitempty"`
	IpAddress  string    `json:"ipAddress,omitempty"`
	Device     string    `json:"device,omitempty"`
	ProductIds []string  `json:"productIds,omitempty"`
	Products   []Product `json:"bought,omitempty"`
	DType      []string  `json:"dgraph.type,omitempty"`
	Date       string    `json:"date,omitempty"`
}

func NewTransaction(id, buyerId, ipAddress, device, date string, productIds []string) (*Transaction, error) {
	return &Transaction{
		Id:         id,
		BuyerId:    buyerId,
		IpAddress:  ipAddress,
		Device:     device,
		ProductIds: productIds,
		DType:      []string{"Transaction"},
		Date:       date,
	}, nil
}
