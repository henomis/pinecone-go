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
	Name             string                 `json:"name"`
	Dimension        int                    `json:"dimension"`
	Metric           *Metric                `json:"metric,omitempty"`
	Pods             *int                   `json:"pods,omitempty"`
	Replicas         *int                   `json:"replicas,omitempty"`
	PodType          *string                `json:"pod_type,omitempty"`
	MetadataConfig   map[string]interface{} `json:"metadata_config,omitempty"`
	SourceCollection *string                `json:"source_collection,omitempty"`
}

func (r *IndexCreate) Path() (string, error) {
	return "/databases", nil
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
