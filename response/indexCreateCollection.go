package response

import (
	"io"
)

type IndexCreateCollection struct {
	Response
}

func (r *IndexCreateCollection) Decode(body io.Reader) error {
	return nil
}

func (r *IndexCreateCollection) SetStatusCode(code int) error {
	r.Code = &code
	return nil
}

func (r *IndexCreateCollection) AcceptContentType() string {
	return ""
}
