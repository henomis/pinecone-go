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

	b, err := io.ReadAll(body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(b, &r.Indexes)
	if err != nil {
		err = json.Unmarshal(b, &r.Response)
		if err != nil {
			return err
		}
	}

	return nil
}

func (r *IndexList) SetStatusCode(code int) error {
	r.Code = &code
	return nil
}

func (r *IndexList) AcceptContentType() string {
	return "application/json"
}
