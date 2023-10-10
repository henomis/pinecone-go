package request

import (
	"bytes"
	"encoding/json"
	"io"
)

type VectorQuery struct {
	ID              *string        `json:"id,omitempty"`
	Namespace       *string        `json:"namespace,omitempty"`
	TopK            int32          `json:"topK"`
	Filter          map[string]any `json:"filter,omitempty"`
	IncludeValues   *bool          `json:"includeValues,omitempty"`
	IncludeMetadata *bool          `json:"includeMetadata,omitempty"`
	Vector          []float64      `json:"vector,omitempty"`
	SparseVector    []SparseVector `json:"sparseVector,omitempty"`

	IndexName string `json:"-"`
	ProjectID string `json:"-"`
}

type SparseVector struct {
	Indices []int64   `json:"indices,omitempty"`
	Values  []float64 `json:"values,omitempty"`
}

func (r *VectorQuery) Path() (string, error) {
	return "/query", nil
}

func (r *VectorQuery) Encode() (io.Reader, error) {
	jsonBytes, err := json.Marshal(r)
	if err != nil {
		return nil, err
	}

	return bytes.NewReader(jsonBytes), nil
}

func (r *VectorQuery) ContentType() string {
	return "application/json"
}
