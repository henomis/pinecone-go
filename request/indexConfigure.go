package request

import (
	"bytes"
	"encoding/json"
	"io"
)

type IndexConfigure struct {
	IndexName string `json:"-"`

	Replicas int    `json:"replicas"`
	PodType  string `json:"pod_type"`
}

func (r *IndexConfigure) Path() (string, error) {
	return "/databases/" + r.IndexName, nil
}

func (r *IndexConfigure) Encode() (io.Reader, error) {
	jsonBytes, err := json.Marshal(r)
	if err != nil {
		return nil, err
	}

	return bytes.NewReader(jsonBytes), nil
}

func (r *IndexConfigure) ContentType() string {
	return "application/json"
}
