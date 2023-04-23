package request

import (
	"bytes"
	"encoding/json"
	"io"
)

type VectorFetch struct {
	IDs       []string `json:"ids"`
	Namespace string   `json:"namespace,omitempty"`
}

func (r *VectorFetch) Path() (string, error) {
	return "/vectors/fetch", nil
}

func (r *VectorFetch) Encode() (io.Reader, error) {
	jsonBytes, err := json.Marshal(r)
	if err != nil {
		return nil, err
	}

	return bytes.NewReader(jsonBytes), nil
}

func (r *VectorFetch) ContentType() string {
	return "application/json"
}
