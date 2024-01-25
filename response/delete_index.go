package response

import (
	"io"
)

type IndexDelete struct {
	Response
}

func (r *IndexDelete) Decode(_ io.Reader) error {
	return nil
}

func (r *IndexDelete) SetStatusCode(code int) error {
	r.Code = &code
	return nil
}

func (r *IndexDelete) AcceptContentType() string {
	return ""
}
