package response

import (
	"encoding/json"
	"io"
)

type VectorDescribeIndexStats struct {
	Response
	Dimension        *int32               `json:"dimension,omitempty"`
	IndexFullness    *float32             `json:"index_fullness,omitempty"`
	TotalVectorCount *int64               `json:"totalVectorCount,omitempty"`
	Namespaces       map[string]Namespace `json:"namespaces,omitempty"`
}

type Namespace struct {
	VectorCount *int32 `json:"vectorCount,omitempty"`
}

func (r *VectorDescribeIndexStats) Decode(body io.Reader) error {
	return json.NewDecoder(body).Decode(r)
}

func (r *VectorDescribeIndexStats) SetStatusCode(code int) error {
	r.Code = &code
	return nil
}

func (r *VectorDescribeIndexStats) AcceptContentType() string {
	return "application/json"
}
