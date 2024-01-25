package response

import (
	"encoding/json"
	"io"
)

type IndexList struct {
	Response
	Indexes []Index `json:"indexes"`
}

type Index struct {
	Name      string            `json:"name"`
	Dimension int               `json:"dimension"`
	Metric    *Metric           `json:"metric,omitempty"`
	Host      string            `json:"host"`
	Spec      Spec              `json:"spec"`
	Status    IndexCreateStatus `json:"status"`
}

func (r *IndexList) Decode(body io.Reader) error {
	return json.NewDecoder(body).Decode(r)
}

func (r *IndexList) SetStatusCode(code int) error {
	r.Code = &code
	return nil
}

func (r *IndexList) AcceptContentType() string {
	return ContentTypeJSON
}
