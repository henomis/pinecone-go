package request

import (
	"bytes"
	"encoding/json"
	"io"
)

type VectorUpsert struct {
	Vectors []struct {
		ID           string `json:"id"`
		Values       []int  `json:"values"`
		SparseValues struct {
			Indices []int `json:"indices"`
			Values  []int `json:"values"`
		} `json:"sparseValues,omitempty"`
		Metadata struct {
			NewKey string `json:"newKey"`
		} `json:"metadata,omitempty"`
	} `json:"vectors,omitempty"`
	Namespace string `json:"namespace,omitempty"`
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
