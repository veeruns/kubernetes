package v1

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// https://medium.com/@trstringer/create-kubernetes-controllers-for-core-and-custom-resources-62fc35ad64a3

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type VeeruResource struct {
	meta_v1.TypeMeta   `json:",inline"`
	meta_v1.ObjectMeta `json:"metadata,omitempty"`
	Spec               VeeruResourceSpec `json:"spec"`
}

type VeeruResourceSpec struct {
	// Message and SomeValue are example custom spec fields
	//
	// this is where you would put your custom resource data
	Message   string `json:"message"`
	SomeValue string `json:"someValue"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// VeeruResourceList is a list of VeeruResource resources
type VeeruResourceList struct {
	meta_v1.TypeMeta `json:",inline"`
	meta_v1.ListMeta `json:"metadata"`

	Items []VeeruResource `json:"items"`
}
