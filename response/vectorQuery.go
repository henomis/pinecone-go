package response

import (
	"encoding/json"
	"io"
)

type QueryMatch struct {
	ID           string                 `json:"id"`
	Score        float64                `json:"score"`
	Values       []float64              `json:"values,omitempty"`
	SparseValues *VectorSparseValues    `json:"sparseValues,omitempty"`
	Metadata     map[string]interface{} `json:"metadata,omitempty"`
}

type VectorQuery struct {
	Response
	Matches   []QueryMatch `json:"matches"`
	Namespace *string      `json:"namespace,omitempty"`
}

func (r *VectorQuery) Decode(body io.Reader) error {
	return json.NewDecoder(body).Decode(r)
}

func (r *VectorQuery) SetStatusCode(code int) error {
	r.Code = &code
	return nil
}

func (r *VectorQuery) AcceptContentType() string {
	return "application/json"
}
