package response

import (
	"encoding/json"
	"io"
)

type Whoami struct {
	Response
	Body string
}

func (r *Whoami) Decode(body io.Reader) error {
	return json.NewDecoder(body).Decode(r)
}

func (r *Whoami) SetBody(body io.Reader) error {

	b, err := io.ReadAll(body)
	if err != nil {
		return err
	}

	r.Body = string(b)
	return nil
}

func (r *Whoami) SetStatusCode(code int) error {
	r.Code = &code
	return nil
}

func (r *Whoami) AcceptContentType() string {
	return "application/json"
}
