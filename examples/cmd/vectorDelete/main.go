package main

import (
	"context"

	pineconego "github.com/henomis/pinecone-go"
	"github.com/henomis/pinecone-go/request"
	"github.com/henomis/pinecone-go/response"
)

func main() {
	p := pineconego.New("pippo", "461e1ec3-470f-4d85-b694-ddf9462ea369")

	r := &request.IndexCreate{
		Name:      "test",
		Dimension: 1,
	}
	rr := &response.IndexCreate{}
	p.IndexCreate(context.Background(), r, rr)

	println(rr.Body)
}
