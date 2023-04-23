package response

import (
	"encoding/json"
	"io"
)

type VectorFetch struct {
	Response
	Vectors struct {
		AdditionalProp struct {
			ID           string    `json:"id"`
			Values       []float64 `json:"values"`
			SparseValues struct {
				Indices []int     `json:"indices"`
				Values  []float64 `json:"values"`
			} `json:"sparseValues"`
			Metadata struct {
				Genre string `json:"genre"`
				Year  int    `json:"year"`
			} `json:"metadata"`
		} `json:"additionalProp"`
	} `json:"vectors"`
	Namespace string `json:"namespace"`
}

func (r *VectorFetch) Decode(body io.Reader) error {
	return json.NewDecoder(body).Decode(r)
}

func (r *VectorFetch) SetBody(body io.Reader) error {
	return nil
}

func (r *VectorFetch) SetStatusCode(code int) error {
	r.Code = &code
	return nil
}

func (r *VectorFetch) AcceptContentType() string {
	return "application/json"
}
