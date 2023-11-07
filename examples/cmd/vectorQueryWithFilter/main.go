package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	pineconego "github.com/henomis/pinecone-go"
	"github.com/henomis/pinecone-go/request"
	"github.com/henomis/pinecone-go/response"
)

func main() {

	apiKey := os.Getenv("PINECONE_API_KEY")
	if apiKey == "" {
		panic("PINECONE_API_KEY is not set")
	}

	environment := os.Getenv("PINECONE_ENVIRONMENT")
	if environment == "" {
		panic("PINECONE_ENVIRONMENT is not set")
	}

	p := pineconego.New(environment, apiKey)

	id := "id3"
	req := &request.VectorQuery{
		IndexName: "test-index",
		ProjectID: "4ce27f9", // use Whoami() to get your project ID
		ID:        &id,
		Filter: request.Filter{
			"<your field>": map[string][]string{
				"$in": {"value1", "value2", "value3"},
			},
		},
		TopK: 10,
	}

	res := &response.VectorQuery{}
	err := p.VectorQuery(context.Background(), req, res)
	if err != nil {
		panic(err)
	}

	if !res.IsSuccess() {
		if res.Response.Code != nil {
			fmt.Printf("Error: %d -", *res.Response.Code)
		}

		if res.Response.Message != nil {
			fmt.Printf(" %s\n", *res.Response.Message)
		}

		if res.Response.RawBody != nil {
			fmt.Printf(" %s\n", *res.Response.RawBody)
		}

		return
	}

	b, err := json.MarshalIndent(res, "", "  ")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(b))

}
