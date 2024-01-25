package request

import (
	"bytes"
	"encoding/json"
	"io"
)

type Metric string

const (
	MetricCosine     Metric = "cosine"
	MetricEuclidean  Metric = "euclidean"
	MetricDotProduct Metric = "dotproduct"
)

type IndexCreate struct {
	Name      string  `json:"name"`
	Dimension int     `json:"dimension"`
	Metric    *Metric `json:"metric,omitempty"`
	Spec      Spec    `json:"spec"`
}

type Spec struct {
	Serverless *ServerlessSpec `json:"serverless,omitempty"`
	Pod        *PodSpec        `json:"pod,omitempty"`
}

type ServerlessSpec struct {
	Cloud  ServerlessSpecCloud `json:"cloud"`
	Region string              `json:"region"`
}

type ServerlessSpecCloud string

const (
	ServerlessSpecCloudAWS   ServerlessSpecCloud = "aws"
	ServerlessSpecCloudGCP   ServerlessSpecCloud = "gcp"
	ServerlessSpecCloudAzure ServerlessSpecCloud = "azure"
)

type PodSpec struct {
	Environment      string          `json:"environment"`
	Replicas         *int            `json:"replicas,omitempty"`
	PodType          string          `json:"pod_type"`
	Pods             *int            `json:"pods,omitempty"`
	Shards           *int            `json:"shards,omitempty"`
	MetadataConfig   *MetadataConfig `json:"metadata_config,omitempty"`
	SourceCollection *string         `json:"source_collection,omitempty"`
}

type MetadataConfig struct {
	Indexed []string `json:"indexed"`
}

func (r *IndexCreate) Path() (string, error) {
	return "/indexes", nil
}

func (r *IndexCreate) Encode() (io.Reader, error) {
	jsonBytes, err := json.Marshal(r)
	if err != nil {
		return nil, err
	}

	return bytes.NewReader(jsonBytes), nil
}

func (r *IndexCreate) ContentType() string {
	return "application/json"
}
