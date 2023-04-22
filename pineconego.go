package pineconego

import (
	"context"
	"net/http"

	"github.com/henomis/pinecone-go/request"
	"github.com/henomis/pinecone-go/response"
	"github.com/henomis/restclientgo"
)

type PineconeGo struct {
	restClient  *restclientgo.RestClient
	environment string
}

func New(environment, apiKey string) *PineconeGo {
	endpoint := "https://" + environment + ".api.pinecone.io"
	restClient := restclientgo.New(endpoint)
	restClient.SetRequestModifier(func(req *http.Request) *http.Request {
		req.Header.Set("Api-Key", apiKey)
		return req
	})
	return &PineconeGo{
		restClient:  restClient,
		environment: environment,
	}
}

func (p *PineconeGo) DescribeIndexStats(ctx context.Context, req *request.VectorDescribeIndexStats, res *response.VectorDescribeIndexStats) error {
	return p.restClient.Post(ctx, req, res)
}

func (p *PineconeGo) VectorQuery(ctx context.Context, req *request.VectorQuery, res *response.VectorQuery) error {
	return p.restClient.Post(ctx, req, res)
}

func (p *PineconeGo) VectorDelete(ctx context.Context, req *request.VectorDelete, res *response.VectorDelete) error {
	return p.restClient.Post(ctx, req, res)
}
