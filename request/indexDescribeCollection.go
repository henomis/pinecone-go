package request

import (
	"io"
)

type IndexDescribeCollection struct {
	CollectionName string `json:"-"`
}

func (r *IndexDescribeCollection) Path() (string, error) {
	return "/collections/" + r.CollectionName, nil
}

func (r *IndexDescribeCollection) Encode() (io.Reader, error) {
	return nil, nil
}

func (r *IndexDescribeCollection) ContentType() string {
	return ""
}
