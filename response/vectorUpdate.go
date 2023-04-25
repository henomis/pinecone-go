package response

import (
	"encoding/json"
	"io"
)

type VectorUpdate struct {
	Response
}

func (r *VectorUpdate) Decode(body io.Reader) error {
	return json.NewDecoder(body).Decode(r)
}

func (r *VectorUpdate) SetStatusCode(code int) error {
	r.Code = &code
	return nil
}

func (r *VectorUpdate) AcceptContentType() string {
	return "application/json"
}
