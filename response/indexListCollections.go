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
	return nil
}

func (r *IndexListCollections) SetBody(body io.Reader) error {
	data, err := io.ReadAll(body)
	if err != nil {
		return err
	}

	return json.Unmarshal(data, &r.Collections)

}

func (r *IndexListCollections) SetStatusCode(code int) error {
	r.Code = &code
	return nil
}

func (r *IndexListCollections) AcceptContentType() string {
	return "application/json"
}
