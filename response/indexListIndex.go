package response

import (
	"encoding/json"
	"io"
)

type IndexList struct {
	Response
	Indexes []string
}

func (r *IndexList) Decode(body io.Reader) error {
	return nil
}

func (r *IndexList) SetBody(body io.Reader) error {
	data, err := io.ReadAll(body)
	if err != nil {
		return err
	}

	return json.Unmarshal(data, &r.Indexes)

}

func (r *IndexList) SetStatusCode(code int) error {
	r.Code = &code
	return nil
}

func (r *IndexList) AcceptContentType() string {
	return "application/json"
}
