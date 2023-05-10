package response

import (
	"encoding/json"
	"io"
)

type Database struct {
	Name      string `json:"name"`
	Metric    string `json:"metric"`
	Dimension int    `json:"dimension"`
	Replicas  int    `json:"replicas"`
	Shards    int    `json:"shards"`
	Pods      int    `json:"pods"`
	PodType   string `json:"pod_type"`
}
type Status struct {
	Waiting []interface{} `json:"waiting"` // not ducomented
	Crashed []interface{} `json:"crashed"` // not ducomented
	Host    string        `json:"host"`
	Port    int           `json:"port"`
	State   string        `json:"state"`
	Ready   bool          `json:"ready"`
}

type IndexDescribe struct {
	Response
	Database Database `json:"database"`
	Status   Status   `json:"status"`
}

func (r *IndexDescribe) Decode(body io.Reader) error {
	return json.NewDecoder(body).Decode(r)
}

func (r *IndexDescribe) SetStatusCode(code int) error {
	r.Code = &code
	return nil
}

func (r *IndexDescribe) AcceptContentType() string {
	return "application/json"
}
