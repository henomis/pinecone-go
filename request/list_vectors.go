package request

import (
	"io"

	"github.com/henomis/restclientgo"
)

type VectorList struct {
	IndexHost       string  `json:"-"`
	Prefix          *string `json:"-"`
	Limit           *int    `json:"-"`
	PaginationToken *string `json:"-"`
	Namespace       *string `json:"-"`
}

func (r *VectorList) Path() (string, error) {
	urlValues := restclientgo.NewURLValues()

	if r.Prefix != nil {
		urlValues.Add("prefix", r.Prefix)
	}

	if r.Limit != nil {
		urlValues.AddInt("limit", r.Limit)
	}

	if r.PaginationToken != nil {
		urlValues.Add("paginationToken", r.PaginationToken)
	}

	if r.Namespace != nil {
		urlValues.Add("namespace", r.Namespace)
	}

	parameters := ""
	if encodedURLValues := urlValues.Encode(); encodedURLValues != "" {
		parameters = "?" + encodedURLValues
	}

	return "/vectors/list" + parameters, nil
}

func (r *VectorList) Encode() (io.Reader, error) {
	return nil, nil
}

func (r *VectorList) ContentType() string {
	return ""
}
