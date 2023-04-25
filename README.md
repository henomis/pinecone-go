# Unofficial Pinecode Go SDK


[![Build Status](https://github.com/henomis/pinecone-go/actions/workflows/test.yml/badge.svg?branch=main)](https://github.com/henomis/pinecone-go/actions/workflows/test.yml?query=branch%3Amain) [![GoDoc](https://godoc.org/github.com/henomis/pinecone-go?status.svg)](https://godoc.org/github.com/henomis/pinecone-go) [![Go Report Card](https://goreportcard.com/badge/github.com/henomis/pinecone-go)](https://goreportcard.com/report/github.com/henomis/pinecone-go) [![GitHub release](https://img.shields.io/github/release/henomis/pinecone-go.svg)](https://github.com/henomis/pinecone-go/releases)

This is [Pinecone](https://pinecone.io)'s **unofficial** Go client, designed to enable you to use Pinecone's services easily from your own applications.

## Pinecone

[Pinecone](https://pinecone.io) is a managed, cloud-native vector database that allows you to build high-performance vector search applications.


## Getting started

### Installation

You can load pinecone-go into your project by using:
```
go get github.com/henomis/pinecone-go
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

	metric := request.MetricCosine
	pods := 1
	replicas := 1
	podType := "s1.x1"
	req := &request.IndexCreate{
		Name:      "test-index",
		Dimension: 10,
		Metric:    &metric,
		Pods:      &pods,
		Replicas:  &replicas,
		PodType:   &podType,
		MetadataConfig: map[string]interface{}{
			"key1": "value1",
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

