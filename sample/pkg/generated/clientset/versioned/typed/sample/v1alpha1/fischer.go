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

// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	"time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
	v1alpha1 "github.com/pluralsh/apiserver-runtime/sample/pkg/apis/sample/v1alpha1"
	scheme "github.com/pluralsh/apiserver-runtime/sample/pkg/generated/clientset/versioned/scheme"
)

// FischersGetter has a method to return a FischerInterface.
// A group's client should implement this interface.
type FischersGetter interface {
	Fischers() FischerInterface
}

// FischerInterface has methods to work with Fischer resources.
type FischerInterface interface {
	Create(ctx context.Context, fischer *v1alpha1.Fischer, opts v1.CreateOptions) (*v1alpha1.Fischer, error)
	Update(ctx context.Context, fischer *v1alpha1.Fischer, opts v1.UpdateOptions) (*v1alpha1.Fischer, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.Fischer, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.FischerList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Fischer, err error)
	FischerExpansion
}

// fischers implements FischerInterface
type fischers struct {
	client rest.Interface
}

// newFischers returns a Fischers
func newFischers(c *SampleV1alpha1Client) *fischers {
	return &fischers{
		client: c.RESTClient(),
	}
}

// Get takes name of the fischer, and returns the corresponding fischer object, and an error if there is any.
func (c *fischers) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.Fischer, err error) {
	result = &v1alpha1.Fischer{}
	err = c.client.Get().
		Resource("fischers").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Fischers that match those selectors.
func (c *fischers) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.FischerList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.FischerList{}
	err = c.client.Get().
		Resource("fischers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested fischers.
func (c *fischers) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("fischers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a fischer and creates it.  Returns the server's representation of the fischer, and an error, if there is any.
func (c *fischers) Create(ctx context.Context, fischer *v1alpha1.Fischer, opts v1.CreateOptions) (result *v1alpha1.Fischer, err error) {
	result = &v1alpha1.Fischer{}
	err = c.client.Post().
		Resource("fischers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(fischer).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a fischer and updates it. Returns the server's representation of the fischer, and an error, if there is any.
func (c *fischers) Update(ctx context.Context, fischer *v1alpha1.Fischer, opts v1.UpdateOptions) (result *v1alpha1.Fischer, err error) {
	result = &v1alpha1.Fischer{}
	err = c.client.Put().
		Resource("fischers").
		Name(fischer.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(fischer).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the fischer and deletes it. Returns an error if one occurs.
func (c *fischers) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("fischers").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *fischers) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("fischers").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched fischer.
func (c *fischers) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Fischer, err error) {
	result = &v1alpha1.Fischer{}
	err = c.client.Patch(pt).
		Resource("fischers").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
