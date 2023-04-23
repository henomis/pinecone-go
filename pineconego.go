package pineconego

import (
	"context"
	"fmt"
	"net/http"

	"github.com/henomis/pinecone-go/request"
	"github.com/henomis/pinecone-go/response"
	"github.com/henomis/restclientgo"
)

type PineconeGo struct {
	restClient  *restclientgo.RestClient
	environment string
	apiKey      string
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
		apiKey:      apiKey,
	}
}

func (p *PineconeGo) Whoami(ctx context.Context, req *request.Whoami, res *response.Whoami) error {
	p.restClient.SetEndpoint(fmt.Sprintf("https://controller.%s.pinecone.io", p.environment))
	return p.restClient.Get(ctx, req, res)
}

func (p *PineconeGo) DescribeIndexStats(
	ctx context.Context,
	req *request.VectorDescribeIndexStats,
	res *response.VectorDescribeIndexStats,
	indexName,
	projectID string,
) error {
	p.restClient.SetEndpoint(fmt.Sprintf("https://%s-%s.svc.%s.pinecone.io", indexName, projectID, p.environment))
	return p.restClient.Post(ctx, req, res)
}

func (p *PineconeGo) VectorQuery(
	ctx context.Context,
	req *request.VectorQuery,
	res *response.VectorQuery,
	indexName,
	projectID string,
) error {
	p.restClient.SetEndpoint(fmt.Sprintf("https://%s-%s.svc.%s.pinecone.io", indexName, projectID, p.environment))
	return p.restClient.Post(ctx, req, res)
}

func (p *PineconeGo) VectorDelete(
	ctx context.Context,
	req *request.IndexCreate,
	res *response.VectorDelete,
	indexName,
	projectID string,
) error {
	p.restClient.SetEndpoint(fmt.Sprintf("https://%s-%s.svc.%s.pinecone.io", indexName, projectID, p.environment))
	return p.restClient.Post(ctx, req, res)
}

func (p *PineconeGo) VectorFetch(
	ctx context.Context,
	req *request.VectorFetch,
	res *response.VectorFetch,
	indexName,
	projectID string,
) error {
	p.restClient.SetEndpoint(fmt.Sprintf("https://%s-%s.svc.%s.pinecone.io", indexName, projectID, p.environment))
	return p.restClient.Get(ctx, req, res)
}

func (p *PineconeGo) VectorUpdate(
	ctx context.Context,
	req *request.VectorUpdate,
	res *response.VectorUpdate,
	indexName,
	projectID string,
) error {
	p.restClient.SetEndpoint(fmt.Sprintf("https://%s-%s.svc.%s.pinecone.io", indexName, projectID, p.environment))
	return p.restClient.Post(ctx, req, res)
}

func (p *PineconeGo) VectorUpsert(
	ctx context.Context,
	req *request.VectorUpsert,
	res *response.VectorUpsert,
	indexName,
	projectID string,
) error {
	p.restClient.SetEndpoint(fmt.Sprintf("https://%s-%s.svc.%s.pinecone.io", indexName, projectID, p.environment))
	return p.restClient.Post(ctx, req, res)
}

func (p *PineconeGo) IndexCreate(ctx context.Context, req *request.IndexCreate, res *response.IndexCreate) error {
	p.restClient.SetEndpoint("https://controller.asia-northeast1-gcp.pinecone.io")
	return p.restClient.Post(ctx, req, res)
}
