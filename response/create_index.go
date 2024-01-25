package response

import (
	"encoding/json"
	"io"
)

type IndexCreate struct {
	Response
	Name      string            `json:"name"`
	Dimension int               `json:"dimension"`
	Metric    *Metric           `json:"metric,omitempty"`
	Host      string            `json:"host"`
	Spec      Spec              `json:"spec"`
	Status    IndexCreateStatus `json:"status"`
}

type Metric string

const (
	MetricCosine     Metric = "cosine"
	MetricEuclidean  Metric = "euclidean"
	MetricDotProduct Metric = "dotproduct"
)

type Spec struct {
	Serverless *ServerlessSpec `json:"serverless,omitempty"`
	Pod        *PodSpec        `json:"pod,omitempty"`
}

type ServerlessSpec struct {
	Cloud  ServerlessSpecCloud `json:"cloud"`
	Region string              `json:"region"`
}

type ServerlessSpecCloud string

const (
	ServerlessSpecCloudAWS   ServerlessSpecCloud = "aws"
	ServerlessSpecCloudGCP   ServerlessSpecCloud = "gcp"
	ServerlessSpecCloudAzure ServerlessSpecCloud = "azure"
)

type PodSpec struct {
	Environment      string          `json:"environment"`
	Replicas         *int            `json:"replicas,omitempty"`
	PodType          string          `json:"pod_type"`
	Pods             *int            `json:"pods,omitempty"`
	Shards           *int            `json:"shards,omitempty"`
	MetadataConfig   *MetadataConfig `json:"metadata_config,omitempty"`
	SourceCollection *string         `json:"source_collection,omitempty"`
}

type MetadataConfig struct {
	Indexed []string `json:"indexed"`
}

type IndexCreateStatus struct {
	Ready bool   `json:"ready"`
	State string `json:"state"`
}

type IndexCreateState string

const (
	IndexCreateStateInitializing         IndexCreateState = "Initializing"
	IndexCreateStateInitializationFailed IndexCreateState = "InitializationFailed"
	IndexCreateStateScalingUp            IndexCreateState = "ScalingUp"
	IndexCreateStateScalingDown          IndexCreateState = "ScalingDown"
	IndexCreateStateScalingUpPodSize     IndexCreateState = "ScalingUpPodSize"
	IndexCreateStateScalingDownPodSize   IndexCreateState = "ScalingDownPodSize"
	IndexCreateStateTerminating          IndexCreateState = "Terminating"
	IndexCreateStateReady                IndexCreateState = "Ready"
)

func (r *IndexCreate) Decode(body io.Reader) error {
	return json.NewDecoder(body).Decode(r)
}

func (r *IndexCreate) SetStatusCode(code int) error {
	r.Code = &code
	return nil
}

func (r *IndexCreate) AcceptContentType() string {
	return "application/json"
}
