package request

import (
	"bytes"
	"encoding/json"
	"io"
)

type VectorSparseValues struct {
	Indices []int64   `json:"indices,omitempty"`
	Values  []float64 `json:"values,omitempty"`
}

type VectorUpdate struct {
	IndexHost    string                 `json:"-"`
	ID           string                 `json:"id"`
	Values       []float64              `json:"values,omitempty"`
	SparseValues *VectorSparseValues    `json:"sparseValues,omitempty"`
	SetMetadata  map[string]interface{} `json:"setMetadata,omitempty"`
	Namespace    *string                `json:"namespace,omitempty"`
}

func (r *VectorUpdate) Path() (string, error) {
	return "/vectors/update", nil
}

func (r *VectorUpdate) Encode() (io.Reader, error) {
	jsonBytes, err := json.Marshal(r)
	if err != nil {
		return nil, err
	}

	return bytes.NewReader(jsonBytes), nil
}

func (r *VectorUpdate) ContentType() string {
	return ContentTypeJSON
}
