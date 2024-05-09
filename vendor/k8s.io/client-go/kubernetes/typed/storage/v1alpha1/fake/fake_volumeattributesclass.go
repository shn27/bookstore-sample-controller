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

package fake

import (
	"context"
	json "encoding/json"
	"fmt"

	v1alpha1 "k8s.io/api/storage/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	storagev1alpha1 "k8s.io/client-go/applyconfigurations/storage/v1alpha1"
	testing "k8s.io/client-go/testing"
)

// FakeVolumeAttributesClasses implements VolumeAttributesClassInterface
type FakeVolumeAttributesClasses struct {
	Fake *FakeStorageV1alpha1
}

var volumeattributesclassesResource = v1alpha1.SchemeGroupVersion.WithResource("volumeattributesclasses")

var volumeattributesclassesKind = v1alpha1.SchemeGroupVersion.WithKind("VolumeAttributesClass")

// Get takes name of the volumeAttributesClass, and returns the corresponding volumeAttributesClass object, and an error if there is any.
func (c *FakeVolumeAttributesClasses) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.VolumeAttributesClass, err error) {
	emptyResult := &v1alpha1.VolumeAttributesClass{}
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(volumeattributesclassesResource, name), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.VolumeAttributesClass), err
}

// List takes label and field selectors, and returns the list of VolumeAttributesClasses that match those selectors.
func (c *FakeVolumeAttributesClasses) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.VolumeAttributesClassList, err error) {
	emptyResult := &v1alpha1.VolumeAttributesClassList{}
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(volumeattributesclassesResource, volumeattributesclassesKind, opts), emptyResult)
	if obj == nil {
		return emptyResult, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.VolumeAttributesClassList{ListMeta: obj.(*v1alpha1.VolumeAttributesClassList).ListMeta}
	for _, item := range obj.(*v1alpha1.VolumeAttributesClassList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested volumeAttributesClasses.
func (c *FakeVolumeAttributesClasses) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(volumeattributesclassesResource, opts))
}

// Create takes the representation of a volumeAttributesClass and creates it.  Returns the server's representation of the volumeAttributesClass, and an error, if there is any.
func (c *FakeVolumeAttributesClasses) Create(ctx context.Context, volumeAttributesClass *v1alpha1.VolumeAttributesClass, opts v1.CreateOptions) (result *v1alpha1.VolumeAttributesClass, err error) {
	emptyResult := &v1alpha1.VolumeAttributesClass{}
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(volumeattributesclassesResource, volumeAttributesClass), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.VolumeAttributesClass), err
}

// Update takes the representation of a volumeAttributesClass and updates it. Returns the server's representation of the volumeAttributesClass, and an error, if there is any.
func (c *FakeVolumeAttributesClasses) Update(ctx context.Context, volumeAttributesClass *v1alpha1.VolumeAttributesClass, opts v1.UpdateOptions) (result *v1alpha1.VolumeAttributesClass, err error) {
	emptyResult := &v1alpha1.VolumeAttributesClass{}
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(volumeattributesclassesResource, volumeAttributesClass), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.VolumeAttributesClass), err
}

// Delete takes name of the volumeAttributesClass and deletes it. Returns an error if one occurs.
func (c *FakeVolumeAttributesClasses) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteActionWithOptions(volumeattributesclassesResource, name, opts), &v1alpha1.VolumeAttributesClass{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeVolumeAttributesClasses) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(volumeattributesclassesResource, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.VolumeAttributesClassList{})
	return err
}

// Patch applies the patch and returns the patched volumeAttributesClass.
func (c *FakeVolumeAttributesClasses) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.VolumeAttributesClass, err error) {
	emptyResult := &v1alpha1.VolumeAttributesClass{}
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(volumeattributesclassesResource, name, pt, data, subresources...), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.VolumeAttributesClass), err
}

// Apply takes the given apply declarative configuration, applies it and returns the applied volumeAttributesClass.
func (c *FakeVolumeAttributesClasses) Apply(ctx context.Context, volumeAttributesClass *storagev1alpha1.VolumeAttributesClassApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha1.VolumeAttributesClass, err error) {
	if volumeAttributesClass == nil {
		return nil, fmt.Errorf("volumeAttributesClass provided to Apply must not be nil")
	}
	data, err := json.Marshal(volumeAttributesClass)
	if err != nil {
		return nil, err
	}
	name := volumeAttributesClass.Name
	if name == nil {
		return nil, fmt.Errorf("volumeAttributesClass.Name must be provided to Apply")
	}
	emptyResult := &v1alpha1.VolumeAttributesClass{}
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(volumeattributesclassesResource, *name, types.ApplyPatchType, data), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.VolumeAttributesClass), err
}