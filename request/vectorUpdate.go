package request

import (
	"bytes"
	"encoding/json"
	"io"
)

type VectorUpdate struct {
	ID           string `json:"id"`
	Values       []int  `json:"values,omitempty"`
	SparseValues struct {
		Indices []int `json:"indices"`
		Values  []int `json:"values"`
	} `json:"sparseValues,omitempty"`
	SetMetadata struct {
		NewKey string `json:"newKey"`
	} `json:"setMetadata,omitempty"`
	Namespace string `json:"namespace,omitempty"`
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
	return "application/json"
}
