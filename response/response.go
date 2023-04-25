package response

import "io"

type Response struct {
	Code          *int                 `json:"code,omitempty"`
	Message       *string              `json:"message,omitempty"`
	Details       []Detail             `json:"details,omitempty"`
	Dimension     *int32               `json:"dimension,omitempty"`
	IndexFullness *string              `json:"index_fullness,omitempty"`
	Namespaces    map[string]Namespace `json:"namespaces,omitempty"`
	RawBody       *string              `json:"body,omitempty"`
}

type Namespace struct {
	VectorCount *int32 `json:"vectorCount,omitempty"`
}

type Detail struct {
	TypeURL *string `json:"typeUrl,omitempty"`
	Value   *string `json:"value,omitempty"`
}

func (r *Response) IsSuccess() bool {
	return (r.Code != nil) && (*r.Code >= 200 && *r.Code < 300)
}

func (r *Response) SetBody(body io.Reader) error {

	b, err := io.ReadAll(body)
	if err != nil {
		return err
	}

	s := string(b)
	r.RawBody = &s

	return nil
}
