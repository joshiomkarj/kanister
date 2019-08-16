package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// ProfileSpec defines the desired state of Profile
// +k8s:openapi-gen=true
type ProfileSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html
}

// ProfileStatus defines the observed state of Profile
// +k8s:openapi-gen=true
type ProfileStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html
}

// These names are used to query Profile API objects.
const (
	ProfileResourceName       = "profile"
	ProfileResourceNamePlural = "profiles"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Profile is the Schema for the profiles API
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status
type Profile struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Location      Location   `json:"location"`
	Credential    Credential `json:"credential"`
	SkipSSLVerify bool       `json:"skipSSLVerify"`
}

// LocationType
type LocationType string

const (
	LocationTypeGCS         LocationType = "gcs"
	LocationTypeS3Compliant LocationType = "s3Compliant"
	LocationTypeAzure       LocationType = "azure"
)

// Location
type Location struct {
	Type     LocationType `json:"type"`
	Bucket   string       `json:"bucket"`
	Endpoint string       `json:"endpoint"`
	Prefix   string       `json:"prefix"`
	Region   string       `json:"region"`
}

// CredentialType
type CredentialType string

const (
	CredentialTypeKeyPair CredentialType = "keyPair"
)

// Credential
type Credential struct {
	Type    CredentialType `json:"type"`
	KeyPair *KeyPair       `json:"keyPair"`
}

// KeyPair
type KeyPair struct {
	IDField     string          `json:"idField"`
	SecretField string          `json:"secretField"`
	Secret      ObjectReference `json:"secret"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ProfileList contains a list of Profile
type ProfileList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Profile `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Profile{}, &ProfileList{})
}
