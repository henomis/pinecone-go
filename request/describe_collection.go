package request

import (
	"io"
)

type IndexDeleteCollection struct {
	CollectionName string `json:"-"`
}

func (r *IndexDeleteCollection) Path() (string, error) {
	return "/collections/" + r.CollectionName, nil
}

func (r *IndexDeleteCollection) Encode() (io.Reader, error) {
	return nil, nil
}

func (r *IndexDeleteCollection) ContentType() string {
	return ""
}
