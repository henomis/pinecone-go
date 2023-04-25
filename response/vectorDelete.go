package response

import (
	"encoding/json"
	"io"
)

type VectorDelete struct {
	Response
}

func (r *VectorDelete) Decode(body io.Reader) error {
	return json.NewDecoder(body).Decode(r)
}

func (r *VectorDelete) SetStatusCode(code int) error {
	r.Code = &code
	return nil
}

func (r *VectorDelete) AcceptContentType() string {
	return "application/json"
}
