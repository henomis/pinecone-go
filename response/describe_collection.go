package response

import (
	"encoding/json"
	"io"
)

type IndexDescribeCollection struct {
	Response
	Collection
}

func (r *IndexDescribeCollection) Decode(body io.Reader) error {
	return json.NewDecoder(body).Decode(r)
}

func (r *IndexDescribeCollection) SetStatusCode(code int) error {
	r.Code = &code
	return nil
}

func (r *IndexDescribeCollection) AcceptContentType() string {
	return "application/json"
}
