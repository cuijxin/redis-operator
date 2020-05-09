package v1alpha1

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// RedisSpec defines the desired state of Redis
type RedisSpec struct {
	Mode               string                     `json:"mode"`
	ImageName          string                     `json:"imageName,omitempty"`
	Size               *int32                     `json:"size,omitempty"`
	ImagePullPolicy    corev1.PullPolicy          `json:"imagePullPolicy,omitempty"`
	Master             RedisMaster                `json:"master,omitempty"`
	Slave              RedisSlave                 `json:"slave,omitempty"`
	RedisPassword      *string                    `json:"redisPassword,omitempty"`
	RedisExporter      bool                       `json:"exporter"`
	RedisExporterImage string                     `json:"redisExporterImage"`
	RedisConfig        map[string]string          `json:"redisConfig"`
	Resources          *Resources                 `json:"resources,omitempty"`
	Storage            *Storage                   `json:"storage,omitempty"`
	NodeSelector       map[string]string          `json:"nodeSelector,omitempty"`
	SecurityContext    *corev1.PodSecurityContext `json:"securityContext,omitempty"`
	PriorityClassName  string                     `json:"priorityClassName,omitempty"`
	Affinity           *corev1.Affinity           `json:"affinity,omitempty"`
}

// Storage is the inteface to add pvc and pv support in redis
type Storage struct {
	VolumeClaimTemplate corev1.PersistentVolumeClaim `json:"volumeClaimTemplate,omitempty"`
}

// RedisMaster interface will have the redis master configuration
type RedisMaster struct {
	Resources   Resources         `json:"resources,omitempty"`
	RedisConfig map[string]string `json:"redisConfig"`
}

// RedisSlave interface will have the redis slave configuration
type RedisSlave struct {
	Resources   Resources         `json:"resources,omitempty"`
	RedisConfig map[string]string `json:"redisConfig"`
}

// ResourceDescription describes CPU and memory resources defined for a cluster.
type ResourceDescription struct {
	CPU    string `json:"cpu"`
	Memory string `json:"memory"`
}

// Resources describes requests and limits for the cluster resouces.
type Resources struct {
	ResourceRequests ResourceDescription `json:"requests,omitempty"`
	ResourceLimits   ResourceDescription `json:"limits,omitempty"`
}

// RedisStatus defines the observed state of Redis
type RedisStatus struct {
	Cluster RedisSpec `json:"cluster,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Redis is the Schema for the redis API
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=redis,scope=Namespaced
type Redis struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   RedisSpec   `json:"spec,omitempty"`
	Status RedisStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// RedisList contains a list of Redis
type RedisList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Redis `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Redis{}, &RedisList{})
}
