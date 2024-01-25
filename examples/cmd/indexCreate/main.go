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
	replicas := 1
	shards := 1
	metric := request.MetricCosine
	req := &request.IndexCreate{
		Name:      "test-index",
		Dimension: 10,
		Metric:    &metric,
		Spec: request.Spec{
			Pod: &request.PodSpec{
				Replicas:    &replicas,
				Shards:      &shards,
				PodType:     "s1.x1",
				Environment: "gcp-starter",
			},
		},
	}
	res := &response.IndexCreate{}
	err := p.IndexCreate(context.Background(), req, res)
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
