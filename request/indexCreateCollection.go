package request

import (
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
	return nil, nil
}

func (r *IndexCreateCollection) ContentType() string {
	return ""
}
