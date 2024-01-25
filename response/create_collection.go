package response

import (
	"encoding/json"
	"io"
)

type IndexCreateCollection struct {
	Response
	Name        string           `json:"name"`
	Size        int              `json:"size"`
	Status      CollectionStatus `json:"status"`
	Dimension   int32            `json:"dimension"`
	RecordCount int32            `json:"record_count"`
	Environment *string          `json:"environment,omitempty"`
}

func (r *IndexCreateCollection) Decode(body io.Reader) error {
	return json.NewDecoder(body).Decode(r)
}

func (r *IndexCreateCollection) SetStatusCode(code int) error {
	r.Code = &code
	return nil
}

func (r *IndexCreateCollection) AcceptContentType() string {
	return "application/json"
}
