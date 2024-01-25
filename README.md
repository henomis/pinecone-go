# Unofficial Pinecone Go SDK

[![GoDoc](https://godoc.org/github.com/henomis/pinecone-go/v2?status.svg)](https://godoc.org/github.com/henomis/pinecone-go/v2) [![Go Report Card](https://goreportcard.com/badge/github.com/henomis/pinecone-go/v2)](https://goreportcard.com/report/github.com/henomis/pinecone-go/v2) [![GitHub release](https://img.shields.io/github/release/henomis/pinecone-go.svg)](https://github.com/henomis/pinecone-go/v2/releases)

This is [Pinecone](https://pinecone.io)'s **unofficial** Go client, designed to enable you to use Pinecone's services easily from your own applications.

## Pinecone

[Pinecone](https://pinecone.io) is a managed, cloud-native vector database that allows you to build high-performance vector search applications.

## API support v2

| **Index Operations** | **Status** | **Vector Operations** | **Status** |
| -------------------- | ---------- | --------------------- | ---------- |
| List Collections     | 🟢         | DescribeIndexStats    | 🟢         |
| Create Collection    | 🟢         | Query                 | 🟢         |
| Describe Collection  | 🟢         | Delete                | 🟢         |
| Delete Collection    | 🟢         | Fetch                 | 🟢         |
| List Indexes         | 🟢         | Update                | 🟢         |
| Create Index         | 🟢         | Upsert                | 🟢         |
| Describe Index       | 🟢         | List                  | 🟢         |
| Delete Index         | 🟢         |                       |            |
| Configure Index      | 🟢         |                       |            |

## Getting started

### Installation

You can load pinecone-go into your project by using:

```
go get github.com/henomis/pinecone-go/v2@v2.0.0
```

### Configuration

The only thing you need to start using Pinecone's APIs is the developer API key and related environment. Copy and paste them in the corresponding place in the code, select the API and the parameters you want to use, and that's it.

### Usage

Please refer to the [examples folder](examples/cmd/) to see how to use the SDK.

Here below a simple usage example:

```go
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
```

## Who uses pinecone-go?

- [LinGoose](https://github.com/henomis/lingoose) Go framework for building awesome LLM apps
