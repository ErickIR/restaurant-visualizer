package dgraph

import (
	"context"
	"fmt"
	"restaurant-visualizer/pkg/storage"

	"github.com/dgraph-io/dgo/v2/protos/api"
)

func LoadSchema(Db *storage.Storage) error {
	op := &api.Operation{}

	op.Schema = `
		date: string @index(exact) .
		id: string @index(exact, term) .
		name: string .
		age: int .
		price: int .
		device: string .
		buyerId: string @index(exact) .
		buyer: uid @reverse .
		ipAddress: string @index(exact) .
		productIds: [string] @index(exact) .
		product: [uid] @reverse .
		transaction: [uid] @reverse .

		type Transaction {
			id: string
			buyerId: string
			ipAddress: string
			device: string
			productIds: [string]
			buyer: uid
		}

		type Buyer {
			id: string
			name: string
			age: int
		}

		type Product {
			id: string
			name: string
			price: int
		}
	`

	ctx := context.Background()

	if err := Db.DbClient.Alter(ctx, op); err != nil {
		fmt.Println("Error altering Schema")
		return err
	}

	return nil
}
