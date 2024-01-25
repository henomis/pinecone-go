package request

import (
	"io"
)

type IndexDescribe struct {
	IndexName string `json:"-"`
}

func (r *IndexDescribe) Path() (string, error) {
	return "/indexes/" + r.IndexName, nil
}

func (r *IndexDescribe) Encode() (io.Reader, error) {
	return nil, nil
}

func (r *IndexDescribe) ContentType() string {
	return ""
}
