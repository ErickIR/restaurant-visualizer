package models

type Product struct {
	Uid          string        `json:"uid,omitempty"`
	Id           string        `json:"id,omitempty"`
	Name         string        `json:"name,omitempty"`
	Price        float32       `json:"price,omitempty"`
	Transactions []Transaction `json:"was_bought,omitempty"`
	DType        []string      `json:"dgraph.type,omitempty"`
}

func NewProduct(id string, name string, price float32) (*Product, error) {
	return &Product{
		Id:    id,
		Name:  name,
		Price: price,
		DType: []string{"Product"},
	}, nil
}
