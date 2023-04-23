package response

import (
	"encoding/json"
	"io"
)

type IndexCreate struct {
	Response
}

func (r *IndexCreate) Decode(body io.Reader) error {
	return json.NewDecoder(body).Decode(r)
}

func (r *IndexCreate) SetBody(body io.Reader) error {
	return nil
}

func (r *IndexCreate) SetStatusCode(code int) error {
	r.Code = &code
	return nil
}

func (r *IndexCreate) AcceptContentType() string {
	return ""
}
