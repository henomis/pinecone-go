package response

import (
	"encoding/json"
	"io"
)

type IndexCreateCollection struct {
	Response
}

func (r *IndexCreateCollection) Decode(body io.Reader) error {
	return json.NewDecoder(body).Decode(r)
}

func (r *IndexCreateCollection) SetBody(body io.Reader) error {
	return nil
}

func (r *IndexCreateCollection) SetStatusCode(code int) error {
	r.Code = &code
	return nil
}

func (r *IndexCreateCollection) AcceptContentType() string {
	return "application/json"
}
