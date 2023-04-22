package response

import (
	"encoding/json"
	"io"
)

type VectorDescribeIndexStats struct {
	Response
}

func (r *VectorDescribeIndexStats) Decode(body io.Reader) error {
	return json.NewDecoder(body).Decode(r)
}

func (r *VectorDescribeIndexStats) SetBody(body io.Reader) error {
	return nil
}

func (r *VectorDescribeIndexStats) SetStatusCode(code int) error {
	r.Code = &code
	return nil
}

func (r *VectorDescribeIndexStats) AcceptContentType() string {
	return "application/json"
}
