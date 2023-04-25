package response

import (
	"encoding/json"
	"io"
)

type IndexConfig struct {
	KBits  *int  `json:"k_bits,omitempty"`
	Hybrid *bool `json:"hybrid,omitempty"`
}

type IndexStatus struct {
	Ready *bool   `json:"ready,omitempty"`
	State *string `json:"state,omitempty"`
}

type IndexDescribe struct {
	Response
	Database struct {
		Name           *string                `json:"name,omitempty"`
		Dimension      *int                   `json:"dimension,omitempty"`
		Metric         *string                `json:"metric,omitempty"`
		Pods           *int                   `json:"pods,omitempty"`
		Replicas       *int                   `json:"replicas,omitempty"`
		Shards         *int                   `json:"shards,omitempty"`
		PodType        *string                `json:"pod_type,omitempty"`
		IndexConfig    *IndexConfig           `json:"index_config,omitempty"`
		MetadataConfig map[string]interface{} `json:"metadata_config,omitempty"`
		Status         *IndexStatus           `json:"status"`
	} `json:"database,omitempty"`
}

func (r *IndexDescribe) Decode(body io.Reader) error {
	return json.NewDecoder(body).Decode(r)
}

func (r *IndexDescribe) SetStatusCode(code int) error {
	r.Code = &code
	return nil
}

func (r *IndexDescribe) AcceptContentType() string {
	return "application/json"
}
