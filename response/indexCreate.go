package response

import (
	"encoding/json"
	"io"
)

type IndexCreate struct {
	Response
	Body string
}

func (r *IndexCreate) Decode(body io.Reader) error {
	return json.NewDecoder(body).Decode(r)
}

func (r *IndexCreate) SetBody(body io.Reader) error {

	b, err := io.ReadAll(body)
	if err != nil {
		return err
	}

	r.Body = string(b)
	return nil

}

func (r *IndexCreate) SetStatusCode(code int) error {
	r.Code = &code
	return nil
}

func (r *IndexCreate) AcceptContentType() string {
	return "application/json"
}
