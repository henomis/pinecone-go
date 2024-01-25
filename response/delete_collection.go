package response

import (
	"io"
)

type IndexDeleteCollection struct {
	Response
}

func (r *IndexDeleteCollection) Decode(body io.Reader) error {
	return nil
}

func (r *IndexDeleteCollection) SetStatusCode(code int) error {
	r.Code = &code
	return nil
}

func (r *IndexDeleteCollection) AcceptContentType() string {
	return ""
}
