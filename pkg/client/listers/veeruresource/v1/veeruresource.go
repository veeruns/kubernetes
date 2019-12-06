/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/veeruns/kubernetes/pkg/apis/veeruresource/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// VeeruResourceLister helps list VeeruResources.
type VeeruResourceLister interface {
	// List lists all VeeruResources in the indexer.
	List(selector labels.Selector) (ret []*v1.VeeruResource, err error)
	// VeeruResources returns an object that can list and get VeeruResources.
	VeeruResources(namespace string) VeeruResourceNamespaceLister
	VeeruResourceListerExpansion
}

// veeruResourceLister implements the VeeruResourceLister interface.
type veeruResourceLister struct {
	indexer cache.Indexer
}

// NewVeeruResourceLister returns a new VeeruResourceLister.
func NewVeeruResourceLister(indexer cache.Indexer) VeeruResourceLister {
	return &veeruResourceLister{indexer: indexer}
}

// List lists all VeeruResources in the indexer.
func (s *veeruResourceLister) List(selector labels.Selector) (ret []*v1.VeeruResource, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.VeeruResource))
	})
	return ret, err
}

// VeeruResources returns an object that can list and get VeeruResources.
func (s *veeruResourceLister) VeeruResources(namespace string) VeeruResourceNamespaceLister {
	return veeruResourceNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// VeeruResourceNamespaceLister helps list and get VeeruResources.
type VeeruResourceNamespaceLister interface {
	// List lists all VeeruResources in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1.VeeruResource, err error)
	// Get retrieves the VeeruResource from the indexer for a given namespace and name.
	Get(name string) (*v1.VeeruResource, error)
	VeeruResourceNamespaceListerExpansion
}

// veeruResourceNamespaceLister implements the VeeruResourceNamespaceLister
// interface.
type veeruResourceNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all VeeruResources in the indexer for a given namespace.
func (s veeruResourceNamespaceLister) List(selector labels.Selector) (ret []*v1.VeeruResource, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.VeeruResource))
	})
	return ret, err
}

// Get retrieves the VeeruResource from the indexer for a given namespace and name.
func (s veeruResourceNamespaceLister) Get(name string) (*v1.VeeruResource, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("veeruresource"), name)
	}
	return obj.(*v1.VeeruResource), nil
}
