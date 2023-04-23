package response

import (
	"io"
)

type IndexConfigure struct {
	Response
}

func (r *IndexConfigure) Decode(body io.Reader) error {
	return nil
}

func (r *IndexConfigure) SetBody(body io.Reader) error {
	return nil
}

func (r *IndexConfigure) SetStatusCode(code int) error {
	r.Code = &code
	return nil
}

func (r *IndexConfigure) AcceptContentType() string {
	return ""
}
