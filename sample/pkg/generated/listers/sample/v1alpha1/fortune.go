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

package v1alpha1

import (
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
	v1alpha1 "github.com/pluralsh/apiserver-runtime/sample/pkg/apis/sample/v1alpha1"
)

// FortuneLister helps list Fortunes.
// All objects returned here must be treated as read-only.
type FortuneLister interface {
	// List lists all Fortunes in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.Fortune, err error)
	// Fortunes returns an object that can list and get Fortunes.
	Fortunes(namespace string) FortuneNamespaceLister
	FortuneListerExpansion
}

// fortuneLister implements the FortuneLister interface.
type fortuneLister struct {
	indexer cache.Indexer
}

// NewFortuneLister returns a new FortuneLister.
func NewFortuneLister(indexer cache.Indexer) FortuneLister {
	return &fortuneLister{indexer: indexer}
}

// List lists all Fortunes in the indexer.
func (s *fortuneLister) List(selector labels.Selector) (ret []*v1alpha1.Fortune, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Fortune))
	})
	return ret, err
}

// Fortunes returns an object that can list and get Fortunes.
func (s *fortuneLister) Fortunes(namespace string) FortuneNamespaceLister {
	return fortuneNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// FortuneNamespaceLister helps list and get Fortunes.
// All objects returned here must be treated as read-only.
type FortuneNamespaceLister interface {
	// List lists all Fortunes in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.Fortune, err error)
	// Get retrieves the Fortune from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.Fortune, error)
	FortuneNamespaceListerExpansion
}

// fortuneNamespaceLister implements the FortuneNamespaceLister
// interface.
type fortuneNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Fortunes in the indexer for a given namespace.
func (s fortuneNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.Fortune, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Fortune))
	})
	return ret, err
}

// Get retrieves the Fortune from the indexer for a given namespace and name.
func (s fortuneNamespaceLister) Get(name string) (*v1alpha1.Fortune, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("fortune"), name)
	}
	return obj.(*v1alpha1.Fortune), nil
}
