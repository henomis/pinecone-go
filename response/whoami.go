package response

import (
	"encoding/json"
	"io"
)

type Whoami struct {
	Response
	ProjectID string `json:"project_name"`
	UserLabel string `json:"user_label"`
	UserName  string `json:"user_name"`
}

func (r *Whoami) Decode(body io.Reader) error {
	return json.NewDecoder(body).Decode(r)
}

func (r *Whoami) SetStatusCode(code int) error {
	r.Code = &code
	return nil
}

func (r *Whoami) AcceptContentType() string {
	return "application/json"
}
