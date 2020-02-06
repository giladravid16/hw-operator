package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// GiladHelloWorldSpec defines the desired state of GiladHelloWorld
type GiladHelloWorldSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html
	Message string `json:"message"`
}

// GiladHelloWorldStatus defines the observed state of GiladHelloWorld
type GiladHelloWorldStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html
	Message string `json:"message"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// GiladHelloWorld is the Schema for the giladhelloworlds API
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=giladhelloworlds,scope=Namespaced
type GiladHelloWorld struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   GiladHelloWorldSpec   `json:"spec,omitempty"`
	Status GiladHelloWorldStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// GiladHelloWorldList contains a list of GiladHelloWorld
type GiladHelloWorldList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GiladHelloWorld `json:"items"`
}

func init() {
	SchemeBuilder.Register(&GiladHelloWorld{}, &GiladHelloWorldList{})
}
