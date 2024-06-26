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
	v1alpha1 "k8s.io/sample-controller/pkg/apis/samplecontroller/v1alpha1"
	scheme "k8s.io/sample-controller/pkg/generated/clientset/versioned/scheme"
)

// BookStoresGetter has a method to return a BookStoreInterface.
// A group's client should implement this interface.
type BookStoresGetter interface {
	BookStores(namespace string) BookStoreInterface
}

// BookStoreInterface has methods to work with BookStore resources.
type BookStoreInterface interface {
	Create(ctx context.Context, bookStore *v1alpha1.BookStore, opts v1.CreateOptions) (*v1alpha1.BookStore, error)
	Update(ctx context.Context, bookStore *v1alpha1.BookStore, opts v1.UpdateOptions) (*v1alpha1.BookStore, error)
	UpdateStatus(ctx context.Context, bookStore *v1alpha1.BookStore, opts v1.UpdateOptions) (*v1alpha1.BookStore, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.BookStore, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.BookStoreList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.BookStore, err error)
	BookStoreExpansion
}

// bookStores implements BookStoreInterface
type bookStores struct {
	client rest.Interface
	ns     string
}

// newBookStores returns a BookStores
func newBookStores(c *ShnV1alpha1Client, namespace string) *bookStores {
	return &bookStores{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the bookStore, and returns the corresponding bookStore object, and an error if there is any.
func (c *bookStores) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.BookStore, err error) {
	result = &v1alpha1.BookStore{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("bookstores").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of BookStores that match those selectors.
func (c *bookStores) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.BookStoreList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.BookStoreList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("bookstores").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested bookStores.
func (c *bookStores) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("bookstores").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a bookStore and creates it.  Returns the server's representation of the bookStore, and an error, if there is any.
func (c *bookStores) Create(ctx context.Context, bookStore *v1alpha1.BookStore, opts v1.CreateOptions) (result *v1alpha1.BookStore, err error) {
	result = &v1alpha1.BookStore{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("bookstores").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(bookStore).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a bookStore and updates it. Returns the server's representation of the bookStore, and an error, if there is any.
func (c *bookStores) Update(ctx context.Context, bookStore *v1alpha1.BookStore, opts v1.UpdateOptions) (result *v1alpha1.BookStore, err error) {
	result = &v1alpha1.BookStore{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("bookstores").
		Name(bookStore.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(bookStore).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *bookStores) UpdateStatus(ctx context.Context, bookStore *v1alpha1.BookStore, opts v1.UpdateOptions) (result *v1alpha1.BookStore, err error) {
	result = &v1alpha1.BookStore{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("bookstores").
		Name(bookStore.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(bookStore).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the bookStore and deletes it. Returns an error if one occurs.
func (c *bookStores) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("bookstores").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *bookStores) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("bookstores").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched bookStore.
func (c *bookStores) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.BookStore, err error) {
	result = &v1alpha1.BookStore{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("bookstores").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
