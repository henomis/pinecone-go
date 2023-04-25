package request

import (
	"io"

	"github.com/henomis/restclientgo"
)

type VectorFetch struct {
	IDs       []string `json:"ids"`
	Namespace string   `json:"namespace,omitempty"`
	IndexName string   `json:"-"`
	ProjectID string   `json:"-"`
}

func (r *VectorFetch) Path() (string, error) {

	newString := func(s string) *string {
		return &s
	}

	urlValues := restclientgo.NewURLValues()

	// OMG! This is so ugly. Hey Pinecone, can you fix this?
	for _, id := range r.IDs {
		urlValues.Add("ids", newString(id))
	}

	if r.Namespace != "" {
		urlValues.Add("namespace", newString(r.Namespace))
	}

	return "/vectors/fetch?" + urlValues.Encode(), nil
}

func (r *VectorFetch) Encode() (io.Reader, error) {
	return nil, nil
}

func (r *VectorFetch) ContentType() string {
	return ""
}
