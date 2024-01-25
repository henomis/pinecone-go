package request

import (
	"bytes"
	"encoding/json"
	"io"
)

type IndexCreateCollection struct {
	Name   string `json:"name"`
	Source string `json:"source"`
}

func (r *IndexCreateCollection) Path() (string, error) {
	return "/collections", nil
}

func (r *IndexCreateCollection) Encode() (io.Reader, error) {
	jsonBytes, err := json.Marshal(r)
	if err != nil {
		return nil, err
	}

	return bytes.NewReader(jsonBytes), nil
}

func (r *IndexCreateCollection) ContentType() string {
	return ContentTypeJSON
}
