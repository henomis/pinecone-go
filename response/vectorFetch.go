package response

import (
	"encoding/json"
	"io"
)

type VectorSparseValues struct {
	Indices []int64   `json:"indices,omitempty"`
	Values  []float32 `json:"values,omitempty"`
}

type Vector struct {
	ID           *string                `json:"id,omitempty"`
	Values       []float32              `json:"values,omitempty"`
	SparseValues *VectorSparseValues    `json:"sparseValues,omitempty"`
	Metadata     map[string]interface{} `json:"metadata,omitempty"`
}

type VectorFetch struct {
	Response
	Vectors   map[string]Vector `json:"vectors,omitempty"`
	Namespace *string           `json:"namespace,omitempty"`
}

func (r *VectorFetch) Decode(body io.Reader) error {
	return json.NewDecoder(body).Decode(r)
}

func (r *VectorFetch) SetStatusCode(code int) error {
	r.Code = &code
	return nil
}

func (r *VectorFetch) AcceptContentType() string {
	return "application/json"
}
