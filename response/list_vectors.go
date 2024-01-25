package response

import (
	"encoding/json"
	"io"
)

type VectorList struct {
	Response
	Vectors    []VectorItem `json:"vectors"`
	Pagination Pagination   `json:"pagination"`
	Namespace  string       `json:"namespace"`
}

type VectorItem struct {
	ID string `json:"id"`
}

type Pagination struct {
	Next string `json:"next"`
}

func (r *VectorList) Decode(body io.Reader) error {
	return json.NewDecoder(body).Decode(r)
}

func (r *VectorList) SetStatusCode(code int) error {
	r.Code = &code
	return nil
}

func (r *VectorList) AcceptContentType() string {
	return "application/json"
}
