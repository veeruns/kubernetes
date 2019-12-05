package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"github.com/veeruns/kubernetes/pkg/apis/veeruresource"
)

var SchemeGroupVersion = schema.GroupVersion{
	Group:   VeeruResource.GroupName,
	Version: "v1",
}

var AddToScheme = runtime.NewSchemeBuilder(addKnownTypes).AddToScheme

func Resource(resource string) schema.GroupResource {
	return SchemeGroupVersion.WithResource(resource).GroupResource()
}

func addKnownTypes(scheme *runtime.Scheme) error {
	scheme.AddKnownTypes(
		SchemeGroupVersion,
		&VeeruResource{},
		&VeeruResourceList{},
	)

	// register the type in the scheme
	meta_v1.AddToGroupVersion(scheme, SchemeGroupVersion)
	return nil
}
