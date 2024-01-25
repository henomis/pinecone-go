package response

import (
	"encoding/json"
	"io"
)

type IndexConfigure struct {
	Response
	Index
}

func (r *IndexConfigure) Decode(body io.Reader) error {
	return json.NewDecoder(body).Decode(r)
}

func (r *IndexConfigure) SetStatusCode(code int) error {
	r.Code = &code
	return nil
}

func (r *IndexConfigure) AcceptContentType() string {
	return ContentTypeJSON
}
