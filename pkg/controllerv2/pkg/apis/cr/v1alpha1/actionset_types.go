package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

const (
	// SchemeVersion is the API version of objects in this package.
	SchemeVersion = "v1alpha1"
	// ResourceGroup is the API group of resources in this package.
	ResourceGroup = "cr.kanister.io"
)

// SchemeGroupVersion is group version used to register these objects
var SchemeGroupVersion = schema.GroupVersion{Group: ResourceGroup, Version: SchemeVersion}

// These names are used to query ActionSet API objects.
const (
	ActionSetResourceName       = "actionset"
	ActionSetResourceNamePlural = "actionsets"
)

var _ runtime.Object = (*ActionSet)(nil)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// ActionSetSpec defines the desired state of ActionSet
// +k8s:openapi-gen=true
type ActionSetSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html
	Actions []ActionSpec `json:"actions"`
}

// ActionSetStatus defines the observed state of ActionSet
// +k8s:openapi-gen=true
type ActionSetStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html
	State   State          `json:"state"`
	Actions []ActionStatus `json:"actions"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ActionSet is the Schema for the actionsets API
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status
type ActionSet struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ActionSetSpec   `json:"spec,omitempty"`
	Status ActionSetStatus `json:"status,omitempty"`
}

// ObjectReference refers to a kubernetes object.
type ObjectReference struct {
	// API version of the referent.
	APIVersion string `json:"apiVersion"`
	// API Group of the referent.
	Group string `json:"group"`
	// Resource name of the referent.
	Resource string `json:"resource"`
	// Kind of the referent.
	// More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	Kind string `json:"kind"`
	// Name of the referent.
	// More info: http://kubernetes.io/docs/user-guide/identifiers#names
	Name string `json:"name"`
	// Namespace of the referent.
	// More info: http://kubernetes.io/docs/user-guide/namespaces
	Namespace string `json:"namespace,omitempty"`
}

// ActionSpec is the specification for a single Action.
type ActionSpec struct {
	// Name is the action we'll perform. For example: `backup` or `restore`.
	Name string `json:"name"`
	// Object refers to the thing we'll perform this action on.
	Object ObjectReference `json:"object"`
	// Blueprint with instructions on how to execute this action.
	Blueprint string `json:"blueprint,omitempty"`
	// Artifacts will be passed as inputs into this phase.
	Artifacts map[string]Artifact `json:"artifacts,omitempty"`
	// ConfigMaps that we'll get and pass into the blueprint.
	ConfigMaps map[string]ObjectReference `json:"configMaps"`
	// Secrets that we'll get and pass into the blueprint.
	Secrets map[string]ObjectReference `json:"secrets"`
	// Profile is use to specify the location where store artifacts and the
	// credentials authorized to access them.
	Profile *ObjectReference `json:"profile"`
	// Options will be used to specify additional values
	// to be used in the Blueprint.
	Options map[string]string `json:"options"`
}

// ActionStatus is updated as we execute phases.
type ActionStatus struct {
	// Name is the action we'll perform. For example: `backup` or `restore`.
	Name string `json:"name"`
	// Object refers to the thing we'll perform this action on.
	Object ObjectReference `json:"object"`
	// Blueprint with instructions on how to execute this action.
	Blueprint string `json:"blueprint"`
	// Phases are sub-actions an are executed sequentially.
	Phases []Phase `json:"phases"`
	// Artifacts created by this phase.
	Artifacts map[string]Artifact `json:"artifacts"`
}

// State is the current state of a phase of execution.
type State string

const (
	// StatePending mean this action or phase has yet to be executed.
	StatePending State = "pending"
	// StateRunning means this action or phase is currently executing.
	StateRunning State = "running"
	// StateFailed means this action or phase was unsuccessful.
	StateFailed State = "failed"
	// StateComplete means this action or phase finished successfully.
	StateComplete State = "complete"
)

// Phase is subcomponent of an action.
type Phase struct {
	Name   string                    `json:"name"`
	State  State                     `json:"state"`
	Output map[string]PhaseInterface `json:"output"`
}

type PhaseInterface interface{}

// k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Artifact tracks objects produced by an action.
type Artifact struct {
	KeyValue map[string]string `json:"keyValue"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ActionSetList contains a list of ActionSet
type ActionSetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ActionSet `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ActionSet{}, &ActionSetList{})
}
