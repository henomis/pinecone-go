package request

import (
	"bytes"
	"encoding/json"
	"io"
)

type IndexConfigure struct {
	IndexName string             `json:"-"`
	Spec      IndexConfigureSpec `json:"spec"`
}

type IndexConfigureSpec struct {
	Pod IndexConfigureSpecPod `json:"pod"`
}

type IndexConfigureSpecPod struct {
	Replicas *int    `json:"replicas,omitempty"`
	PodType  *string `json:"pod_type,omitempty"`
}

func (r *IndexConfigure) Path() (string, error) {
	return "/indexes/" + r.IndexName, nil
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
