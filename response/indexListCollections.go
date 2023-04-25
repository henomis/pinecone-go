package response

import (
	"encoding/json"
	"io"
)

type IndexListCollections struct {
	Response
	Collections []string
}

func (r *IndexListCollections) Decode(body io.Reader) error {
	b, err := io.ReadAll(body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(b, &r.Collections)
	if err != nil {
		err = json.Unmarshal(b, &r.Response)
		if err != nil {
			return err
		}
	}

	return nil
}

func (r *IndexListCollections) SetStatusCode(code int) error {
	r.Code = &code
	return nil
}

func (r *IndexListCollections) AcceptContentType() string {
	return "application/json"
}
