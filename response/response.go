package response

type Response struct {
	Code          *int                 `json:"code,omitempty"`
	Message       *string              `json:"message,omitempty"`
	Details       []Detail             `json:"details,omitempty"`
	Dimension     *int32               `json:"dimension,omitempty"`
	IndexFullness *string              `json:"index_fullness,omitempty"`
	Namespaces    map[string]Namespace `json:"namespaces,omitempty"`
}

type Namespace struct {
	VectorCount *int32 `json:"vectorCount,omitempty"`
}

type Detail struct {
	TypeURL *string `json:"typeUrl,omitempty"`
	Value   *string `json:"value,omitempty"`
}
