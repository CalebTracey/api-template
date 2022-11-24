package main

import (
	"context"
	"fmt"
	"github.com/calebtracey/api-template/pkg/openapi3"
	"log"
)

func main() {
	client, err := openapi3.NewClientWithResponses("http://0.0.0.0:6080")
	if err != nil {
		log.Fatalf("Couldn't instantiate client: %s", err)
	}

	reqType := "psql"
	table := "example"
	id := "123"
	respC, err := client.AddToDatabaseWithResponse(context.Background(),
		openapi3.AddToDatabaseJSONRequestBody{
			RequestType: &reqType,
			Table:       &table,
			Id:          &id,
		})
	if err != nil {
		log.Fatalf("Couldn't get competition %s", err)
	}

	fmt.Printf("\tRows Affected: %s\n", *respC.JSON201.RowsAffected)
	fmt.Printf("\tLast Inserted ID: %v\n", *respC.JSON201.LastInsertID)
	fmt.Printf("\tMessage: %v\n", *respC.JSON201.Message)
}
