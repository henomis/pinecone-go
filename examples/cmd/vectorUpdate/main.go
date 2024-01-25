package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	pineconego "github.com/henomis/pinecone-go/v2"
	"github.com/henomis/pinecone-go/v2/request"
	"github.com/henomis/pinecone-go/v2/response"
)

func main() {

	apiKey := os.Getenv("PINECONE_API_KEY")
	if apiKey == "" {
		panic("PINECONE_API_KEY is not set")
	}

	p := pineconego.New(apiKey)

	req := &request.VectorUpdate{
		IndexHost: "https://test-index-xxxxxx.xxx.gcp-starter.pinecone.io",
		ID:        "id3",
		Values:    []float64{1.1, 2.2, 3, 4, 5, 6, 7, 8, 9, 10},
	}

	res := &response.VectorUpdate{}
	err := p.VectorUpdate(context.Background(), req, res)
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
