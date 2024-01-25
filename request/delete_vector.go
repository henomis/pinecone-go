package request

import (
	"bytes"
	"encoding/json"
	"io"
)

type VectorDelete struct {
	IndexHost string   `json:"-"`
	IDs       []string `json:"ids,omitempty"`
	DeleteAll *bool    `json:"deleteAll,omitempty"`
	Namespace *string  `json:"namespace,omitempty"`
	Filter    Filter   `json:"filter,omitempty"`
}

func (r *VectorDelete) Path() (string, error) {
	return "/vectors/delete", nil
}

func (r *VectorDelete) Encode() (io.Reader, error) {
	jsonBytes, err := json.Marshal(r)
	if err != nil {
		return nil, err
	}

	return bytes.NewReader(jsonBytes), nil
}

func (r *VectorDelete) ContentType() string {
	return ContentTypeJSON
}
