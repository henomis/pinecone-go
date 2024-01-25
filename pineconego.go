package pineconego

import (
	"context"
	"net/http"

	"github.com/henomis/pinecone-go/request"
	"github.com/henomis/pinecone-go/response"
	"github.com/henomis/restclientgo"
)

type PineconeGo struct {
	restClient *restclientgo.RestClient
	apiKey     string
}

const (
	// indexEndpointTemplate  = "https://controller.%s.pinecone.io"
	// vectorEndpointTemplate = "https://%s-%s.svc.%s.pinecone.io"

	controlPlaneEndpoint = "https://api.pinecone.io"
)

func New(apiKey string) *PineconeGo {

	restClient := restclientgo.New("")
	restClient.SetRequestModifier(func(req *http.Request) *http.Request {
		req.Header.Set("Api-Key", apiKey)
		return req
	})
	return &PineconeGo{
		restClient: restClient,
		apiKey:     apiKey,
	}
}

func (p *PineconeGo) Whoami(ctx context.Context, req *request.Whoami, res *response.Whoami) error {
	// p.restClient.SetEndpoint(fmt.Sprintf(indexEndpointTemplate, p.environment))
	return p.restClient.Get(ctx, req, res)
}

func (p *PineconeGo) VectorDescribeIndexStats(
	ctx context.Context,
	req *request.VectorDescribeIndexStats,
	res *response.VectorDescribeIndexStats,

) error {
	// p.restClient.SetEndpoint(fmt.Sprintf(vectorEndpointTemplate, req.IndexName, req.ProjectID, p.environment))
	return p.restClient.Post(ctx, req, res)
}

func (p *PineconeGo) VectorQuery(
	ctx context.Context,
	req *request.VectorQuery,
	res *response.VectorQuery,
) error {
	// p.restClient.SetEndpoint(fmt.Sprintf(vectorEndpointTemplate, req.IndexName, req.ProjectID, p.environment))
	return p.restClient.Post(ctx, req, res)
}

func (p *PineconeGo) VectorDelete(
	ctx context.Context,
	req *request.VectorDelete,
	res *response.VectorDelete,
) error {
	// p.restClient.SetEndpoint(fmt.Sprintf(vectorEndpointTemplate, req.IndexName, req.ProjectID, p.environment))
	return p.restClient.Post(ctx, req, res)
}

func (p *PineconeGo) VectorFetch(
	ctx context.Context,
	req *request.VectorFetch,
	res *response.VectorFetch,
) error {
	// p.restClient.SetEndpoint(fmt.Sprintf(vectorEndpointTemplate, req.IndexName, req.ProjectID, p.environment))
	return p.restClient.Get(ctx, req, res)
}

func (p *PineconeGo) VectorUpdate(
	ctx context.Context,
	req *request.VectorUpdate,
	res *response.VectorUpdate,
) error {
	// p.restClient.SetEndpoint(fmt.Sprintf(vectorEndpointTemplate, req.IndexName, req.ProjectID, p.environment))
	return p.restClient.Post(ctx, req, res)
}

func (p *PineconeGo) VectorUpsert(ctx context.Context, req *request.VectorUpsert, res *response.VectorUpsert) error {
	// p.restClient.SetEndpoint(fmt.Sprintf(vectorEndpointTemplate, req.IndexName, req.ProjectID, p.environment))
	return p.restClient.Post(ctx, req, res)
}

func (p *PineconeGo) IndexListCollections(
	ctx context.Context,
	req *request.IndexListCollections,
	res *response.IndexListCollections,
) error {
	p.restClient.SetEndpoint(controlPlaneEndpoint)
	return p.restClient.Get(ctx, req, res)
}

func (p *PineconeGo) IndexCreateCollection(
	ctx context.Context,
	req *request.IndexCreateCollection,
	res *response.IndexCreateCollection,
) error {
	// p.restClient.SetEndpoint(fmt.Sprintf(indexEndpointTemplate, p.environment))
	return p.restClient.Post(ctx, req, res)
}

func (p *PineconeGo) IndexDescribeCollection(
	ctx context.Context,
	req *request.IndexDescribeCollection,
	res *response.IndexDescribeCollection,
) error {
	// p.restClient.SetEndpoint(fmt.Sprintf(indexEndpointTemplate, p.environment))
	return p.restClient.Get(ctx, req, res)
}

func (p *PineconeGo) IndexDeleteCollection(
	ctx context.Context,
	req *request.IndexDeleteCollection,
	res *response.IndexDeleteCollection,
) error {
	// p.restClient.SetEndpoint(fmt.Sprintf(indexEndpointTemplate, p.environment))
	return p.restClient.Delete(ctx, req, res)
}

func (p *PineconeGo) IndexList(
	ctx context.Context,
	req *request.IndexList,
	res *response.IndexList,
) error {
	p.restClient.SetEndpoint(controlPlaneEndpoint)
	return p.restClient.Get(ctx, req, res)
}

func (p *PineconeGo) IndexCreate(
	ctx context.Context,
	req *request.IndexCreate,
	res *response.IndexCreate,
) error {
	p.restClient.SetEndpoint(controlPlaneEndpoint)
	return p.restClient.Post(ctx, req, res)
}

func (p *PineconeGo) IndexDescribe(
	ctx context.Context,
	req *request.IndexDescribe,
	res *response.IndexDescribe,
) error {
	p.restClient.SetEndpoint(controlPlaneEndpoint)
	return p.restClient.Get(ctx, req, res)
}

func (p *PineconeGo) IndexDelete(
	ctx context.Context,
	req *request.IndexDelete,
	res *response.IndexDelete,
) error {
	p.restClient.SetEndpoint(controlPlaneEndpoint)
	return p.restClient.Delete(ctx, req, res)
}

func (p *PineconeGo) IndexConfigure(
	ctx context.Context,
	req *request.IndexConfigure,
	res *response.IndexConfigure,
) error {
	p.restClient.SetEndpoint(controlPlaneEndpoint)
	return p.restClient.Patch(ctx, req, res)
}
