package response

import (
	"encoding/json"
	"io"
)

type IndexListCollections struct {
	Response
	Collections []Collection `json:"collections"`
}

type CollectionStatus string

const (
	CollectionStatusInitializing CollectionStatus = "Initializing"
	CollectionStatusReady        CollectionStatus = "Ready"
	CollectionStatusTerminating  CollectionStatus = "Terminating"
)

type Collection struct {
	Name        string           `json:"name"`
	Size        int64            `json:"size"`
	Status      CollectionStatus `json:"status"`
	Dimension   int32            `json:"dimension"`
	RecordCount int32            `json:"record_count"`
	Environment *string          `json:"environment,omitempty"`
}

func (r *IndexListCollections) Decode(body io.Reader) error {
	b, err := io.ReadAll(body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(b, &r.Collections)
	if err != nil {
		err = json.Unmarshal(b, &r.Response)
		if err != nil {
			return err
		}
	}

	return nil
}

func (r *IndexListCollections) SetStatusCode(code int) error {
	r.Code = &code
	return nil
}

func (r *IndexListCollections) AcceptContentType() string {
	return ContentTypeJSON
}
