package response

import (
	"encoding/json"
	"io"
)

type IndexDescribe struct {
	Response
	Index
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
