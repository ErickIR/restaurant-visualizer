package storage

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/dgraph-io/dgo/v2"
	"github.com/dgraph-io/dgo/v2/protos/api"
	"google.golang.org/grpc"
)

type CancelFunc func()

type Storage struct {
	DbClient *dgo.Dgraph
	Cancel   CancelFunc
}

func NewClient() (*Storage, error) {
	dgraphUrl := os.Getenv("DGRAPH_URL")
	conn, err := grpc.Dial(dgraphUrl, grpc.WithInsecure())

	if err != nil {
		fmt.Println("Error Connecting to Dgraph")
		return nil, err
	}

	dgClient := dgo.NewDgraphClient(
		api.NewDgraphClient(conn),
	)

	cancelFunc := func() {
		if err := conn.Close(); err != nil {
			log.Printf("Error while closing connection: %v", err)
		}
	}
	return &Storage{DbClient: dgClient, Cancel: cancelFunc}, nil
}

func Example_setObject() {
	Db, err := NewClient()

	if err != nil {
		log.Fatalf("Error connecting to the DB: %v", err)
	}

	defer Db.Cancel()

	// date := time.Date(1980, 01, 01, 23, 0, 0, 0, time.UTC)

	op := &api.Operation{}
	op.Schema = `
		id: string @index(exact) .
		name: string .
		age: int .
		price: float .
		date: datetime .
		Buyer: uid .
		Bought: [uid] .
		buyerId: string .
		ipAddress: string @index(exact) .
		device: string @index(exact) .
		productIds: [string] .
		type Transaction {
			id: string
			Buyer: Buyer
			buyerId: string
			ipAddress: string
			device: string
			productIds: [string]
			Bought: [Product]
		}
		type Buyer {
			id: string
			name: string
			age: int
		}
		type Product {
			id: string
			name: string
			price: float
		}
	`

	ctx := context.Background()
	if err := Db.DbClient.Alter(ctx, op); err != nil {
		log.Fatalf("Error altering schema: %v", err)
	}

	// mutation := &api.Mutation{CommitNow: true}

	// pb, err := json.Marshal(transaction)

	// if err != nil {
	// 	log.Fatalf("Error parsing to JSON: %v", err)
	// }

	// mutation.SetJson = pb
	// response, err := dg.NewTxn().Mutate(ctx, mutation)

	// if err != nil {
	// 	log.Fatalf("Error running mutation: %v", err)
	// }

	// uid := response.Uids["trans"]

	// fmt.Println(uid)
	// variables := map[string]string{"$id1": uid}
	// query := `query Me($id1: string) {
	// 	me(func: uid($id1)) {
	// 		uid
	// 		id
	// 		ipAddress
	// 		device
	// 		dgraph.type
	// 		buyer {
	// 			id
	// 			name
	// 			age
	// 		}
	// 		bought {
	// 			id
	// 			name
	// 			price
	// 		}
	// 	}
	// }`

	// resp, err := dg.NewTxn().QueryWithVars(ctx, query, variables)

	// query := `
	// 	{
	// 		total(func: has(_predicate_)) {count(uid)}
	// 	}
	// `

	// resp, err := dg.NewTxn().Query(ctx, query)

	// if err != nil {
	// 	log.Fatalf("Error running query: %v", err)
	// }

	// type Root struct {
	// 	Me []Transaction `json:"me"`
	// }

	// fmt.Println(resp.Json)

	// var r Root
	// err = json.Unmarshal(resp.Json, &r)

	// if err != nil {
	// 	log.Fatalf("Error parsing from JSON: %v", err)
	// }

	// out, _ := json.MarshalIndent(r, "", "\t")
	// fmt.Printf("%s\n", out)
}
