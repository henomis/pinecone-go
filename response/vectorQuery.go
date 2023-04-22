package response

import (
	"encoding/json"
	"io"
)

type VectorQuery struct {
	Response
	Matches []struct {
		ID           string    `json:"id"`
		Score        *float64  `json:"score,omitempty"`
		Values       []float64 `json:"values,omitempty"`
		SparseValues struct {
			Indices []int     `json:"indices"`
			Values  []float64 `json:"values"`
		} `json:"sparseValues,omitempty"`
		Metadata *struct {
			Genre string `json:"genre"`
			Year  int    `json:"year"`
		} `json:"metadata,omitempty"`
	} `json:"matches"`
	Namespace *string `json:"namespace,omitempty"`
}

func (r *VectorQuery) Decode(body io.Reader) error {
	return json.NewDecoder(body).Decode(r)
}

func (r *VectorQuery) SetBody(body io.Reader) error {
	return nil
}

func (r *VectorQuery) SetStatusCode(code int) error {
	r.Code = &code
	return nil
}

func (r *VectorQuery) AcceptContentType() string {
	return "application/json"
}
