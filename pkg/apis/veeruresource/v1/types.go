package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

//VeeruResource some object with spec
type VeeruResource struct {
	meta_v1.TypeMeta   `json:",inline"`
	meta_v1.ObjectMeta `json:"metadata,omitempty"`

	Spec VeeruResourceSpec `json:"spec"`
}

// VeeruResourceSpec is the what all the crd contains.
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

//VeeruResourceStatus is the status
type VeeruResourceStatus struct {
	RandomString int32 `json:"randomstring"`
}
