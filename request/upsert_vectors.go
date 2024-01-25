package request

import (
	"bytes"
	"encoding/json"
	"io"
)

type Vector struct {
	ID           string                 `json:"id"`
	Values       []float64              `json:"values"`
	SparseValues *VectorSparseValues    `json:"sparseValues,omitempty"`
	Metadata     map[string]interface{} `json:"metadata,omitempty"`
}

type VectorUpsert struct {
	IndexHost string   `json:"-"`
	Vectors   []Vector `json:"vectors"`
	Namespace string   `json:"namespace,omitempty"`
}

func (r *VectorUpsert) Path() (string, error) {
	return "/vectors/upsert", nil
}

func (r *VectorUpsert) Encode() (io.Reader, error) {
	jsonBytes, err := json.Marshal(r)
	if err != nil {
		return nil, err
	}

	return bytes.NewReader(jsonBytes), nil
}

func (r *VectorUpsert) ContentType() string {
	return "application/json"
}
