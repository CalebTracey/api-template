package main

import (
	"context"
	"fmt"
	"github.com/calebtracey/rugby-data-api/pkg/openapi3"
	"log"
)

func main() {
	client, err := openapi3.NewClientWithResponses("http://0.0.0.0:6080")
	if err != nil {
		log.Fatalf("Couldn't instantiate client: %s", err)
	}

	source := "psql_db"
	compId := "180659"
	respC, err := client.GetCompetitionWithResponse(context.Background(),
		openapi3.GetCompetitionJSONRequestBody{
			Source:        &source,
			CompetitionID: &compId,
		})
	if err != nil {
		log.Fatalf("Couldn't get competition %s", err)
	}

	fmt.Printf("Competition\n\tID: %s\n", *respC.JSON201.Id)
	fmt.Printf("\tName: %s\n", *respC.JSON201.Name)
	fmt.Printf("\tTeams: %v\n", *respC.JSON201.Teams)
	fmt.Printf("\tMessage: %v\n", *respC.JSON201.Message)
}
