package request

import (
	"io"
)

type IndexListCollections struct {
}

func (r *IndexListCollections) Path() (string, error) {
	return "/collections", nil
}

func (r *IndexListCollections) Encode() (io.Reader, error) {
	return nil, nil
}

func (r *IndexListCollections) ContentType() string {
	return ""
}
