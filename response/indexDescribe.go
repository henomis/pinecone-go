package response

import (
	"encoding/json"
	"io"
)

type IndexDescribe struct {
	Response
	Database struct {
		Name        string `json:"name"`
		Dimension   int    `json:"dimension"`
		Metric      string `json:"metric"`
		Pods        int    `json:"pods"`
		Replicas    int    `json:"replicas"`
		Shards      int    `json:"shards"`
		PodType     string `json:"pod_type"`
		IndexConfig struct {
			KBits  int  `json:"k_bits"`
			Hybrid bool `json:"hybrid"`
		} `json:"index_config"`
		MetadataConfig map[string]interface{} `json:"metadata_config,omitempty"`
		Status         struct {
			Ready bool   `json:"ready"`
			State string `json:"state"`
		} `json:"status"`
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
