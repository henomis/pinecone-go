package request

import (
	"io"
)

type IndexDelete struct {
	IndexName string `json:"-"`
}

func (r *IndexDelete) Path() (string, error) {
	return "/databases/" + r.IndexName, nil
}

func (r *IndexDelete) Encode() (io.Reader, error) {
	return nil, nil
}

func (r *IndexDelete) ContentType() string {
	return ""
}
