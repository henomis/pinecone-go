package request

import (
	"io"
)

type IndexList struct {
}

func (r *IndexList) Path() (string, error) {
	return "/databases", nil
}

func (r *IndexList) Encode() (io.Reader, error) {
	return nil, nil
}

func (r *IndexList) ContentType() string {
	return ""
}
