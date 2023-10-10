package request

import (
	"bytes"
	"encoding/json"
	"io"
)

type VectorDescribeIndexStats struct {
	Filter    map[string]any `json:"filter,omitempty"`
	IndexName string         `json:"-"`
	ProjectID string         `json:"-"`
}

func (r *VectorDescribeIndexStats) Path() (string, error) {
	return "/describe_index_stats", nil
}

func (r *VectorDescribeIndexStats) Encode() (io.Reader, error) {
	jsonBytes, err := json.Marshal(r)
	if err != nil {
		return nil, err
	}

	return bytes.NewReader(jsonBytes), nil
}

func (r *VectorDescribeIndexStats) ContentType() string {
	return "application/json"
}
