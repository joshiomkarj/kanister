package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// BlueprintSpec defines the desired state of Blueprint
// +k8s:openapi-gen=true
type BlueprintSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html
}

// BlueprintStatus defines the observed state of Blueprint
// +k8s:openapi-gen=true
type BlueprintStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html
}

// These names are used to query Blueprint API objects.
const (
	BlueprintResourceName       = "blueprint"
	BlueprintResourceNamePlural = "blueprints"
)

var _ runtime.Object = (*Blueprint)(nil)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Blueprint is the Schema for the blueprints API
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status
type Blueprint struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec    BlueprintSpec               `json:"spec,omitempty"`
	Status  BlueprintStatus             `json:"status,omitempty"`
	Actions map[string]*BlueprintAction `json:"actions"`
}

// BlueprintAction describes the set of phases that constitute an action.
type BlueprintAction struct {
	Name               string              `json:"name"`
	Kind               string              `json:"kind"`
	ConfigMapNames     []string            `json:"configMapNames"`
	SecretNames        []string            `json:"secretNames"`
	InputArtifactNames []string            `json:"inputArtifactNames"`
	OutputArtifacts    map[string]Artifact `json:"outputArtifacts"`
	Phases             []BlueprintPhase    `json:"phases"`
}

// BlueprintPhase is a an individual unit of execution.
type BlueprintPhase struct {
	Func       string                     `json:"func"`
	Name       string                     `json:"name"`
	ObjectRefs map[string]ObjectReference `json:"objects"`
	Args       map[string]ArgsInterface     `json:"args"`
}

type ArgsInterface interface{}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// BlueprintList contains a list of Blueprint
type BlueprintList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Blueprint `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Blueprint{}, &BlueprintList{})
}
