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
		date: string @index(exact) .
		was_made_by: uid .
		made: [uid] .
		ipAddress: string @index(exact) .
		productIds: [string] @index(exact) .
		bought: [uid] .
		was_bought: [uid] .

		type Transaction {
			id: string
			buyerId: string
			was_made_by: uid
			ipAddress: string
			device: string
			productIds: [string]
			bought: [uid]
			date: string
		}

		type Buyer {
			id: string
			name: string
			age: int
			made: [uid]
			date: string
		}

		type Product {
			id: string
			name: string
			price: int
			was_bought: [uid]
			date: string
		}
	`

	ctx := context.Background()

	if err := Db.DbClient.Alter(ctx, op); err != nil {
		fmt.Println("Error altering Schema")
		return err
	}

	return nil
}
