package v1alpha1

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// VaultSecretSpec defines the desired state of VaultSecret
// +k8s:openapi-gen=true
type VaultSecretSpec struct {
	// Keys is an array of Keys, which should be included in the Kubernetes
	// secret. If the Keys field is ommitted all keys from the Vault secret will
	// be included in the Kubernetes secret.
	Keys []string `json:"keys,omitempty"`
	// Path is the path of the corresponding secret in Vault.
	Path string `json:"path"`
	// Type is the type of the Kubernetes secret, which will be created by the
	// Vault Secrets Operator.
	Type corev1.SecretType `json:"type"`
}

// VaultSecretStatus defines the observed state of VaultSecret
// +k8s:openapi-gen=true
type VaultSecretStatus struct {
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// VaultSecret is the Schema for the vaultsecrets API
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status
type VaultSecret struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   VaultSecretSpec   `json:"spec,omitempty"`
	Status VaultSecretStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// VaultSecretList contains a list of VaultSecret
type VaultSecretList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VaultSecret `json:"items"`
}

func init() {
	SchemeBuilder.Register(&VaultSecret{}, &VaultSecretList{})
}