/*
* Copyright 2024 Google LLC
*
* Licensed under the Apache License, Version 2.0 (the "License");
* you may not use this file except in compliance with the License.
* You may obtain a copy of the License at
*
*     https://www.apache.org/licenses/LICENSE-2.0
*
*     Unless required by applicable law or agreed to in writing, software
*     distributed under the License is distributed on an "AS IS" BASIS,
*     WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
*     See the License for the specific language governing permissions and
*     limitations under the License.
 */

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1 "github.com/GoogleCloudPlatform/gke-gateway-api/apis/networking/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeGCPSessionAffinityFilters implements GCPSessionAffinityFilterInterface
type FakeGCPSessionAffinityFilters struct {
	Fake *FakeNetworkingV1
	ns   string
}

var gcpsessionaffinityfiltersResource = v1.SchemeGroupVersion.WithResource("gcpsessionaffinityfilters")

var gcpsessionaffinityfiltersKind = v1.SchemeGroupVersion.WithKind("GCPSessionAffinityFilter")

// Get takes name of the gCPSessionAffinityFilter, and returns the corresponding gCPSessionAffinityFilter object, and an error if there is any.
func (c *FakeGCPSessionAffinityFilters) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.GCPSessionAffinityFilter, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(gcpsessionaffinityfiltersResource, c.ns, name), &v1.GCPSessionAffinityFilter{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.GCPSessionAffinityFilter), err
}

// List takes label and field selectors, and returns the list of GCPSessionAffinityFilters that match those selectors.
func (c *FakeGCPSessionAffinityFilters) List(ctx context.Context, opts metav1.ListOptions) (result *v1.GCPSessionAffinityFilterList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(gcpsessionaffinityfiltersResource, gcpsessionaffinityfiltersKind, c.ns, opts), &v1.GCPSessionAffinityFilterList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1.GCPSessionAffinityFilterList{ListMeta: obj.(*v1.GCPSessionAffinityFilterList).ListMeta}
	for _, item := range obj.(*v1.GCPSessionAffinityFilterList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested gCPSessionAffinityFilters.
func (c *FakeGCPSessionAffinityFilters) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(gcpsessionaffinityfiltersResource, c.ns, opts))

}

// Create takes the representation of a gCPSessionAffinityFilter and creates it.  Returns the server's representation of the gCPSessionAffinityFilter, and an error, if there is any.
func (c *FakeGCPSessionAffinityFilters) Create(ctx context.Context, gCPSessionAffinityFilter *v1.GCPSessionAffinityFilter, opts metav1.CreateOptions) (result *v1.GCPSessionAffinityFilter, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(gcpsessionaffinityfiltersResource, c.ns, gCPSessionAffinityFilter), &v1.GCPSessionAffinityFilter{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.GCPSessionAffinityFilter), err
}

// Update takes the representation of a gCPSessionAffinityFilter and updates it. Returns the server's representation of the gCPSessionAffinityFilter, and an error, if there is any.
func (c *FakeGCPSessionAffinityFilters) Update(ctx context.Context, gCPSessionAffinityFilter *v1.GCPSessionAffinityFilter, opts metav1.UpdateOptions) (result *v1.GCPSessionAffinityFilter, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(gcpsessionaffinityfiltersResource, c.ns, gCPSessionAffinityFilter), &v1.GCPSessionAffinityFilter{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.GCPSessionAffinityFilter), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeGCPSessionAffinityFilters) UpdateStatus(ctx context.Context, gCPSessionAffinityFilter *v1.GCPSessionAffinityFilter, opts metav1.UpdateOptions) (*v1.GCPSessionAffinityFilter, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(gcpsessionaffinityfiltersResource, "status", c.ns, gCPSessionAffinityFilter), &v1.GCPSessionAffinityFilter{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.GCPSessionAffinityFilter), err
}

// Delete takes name of the gCPSessionAffinityFilter and deletes it. Returns an error if one occurs.
func (c *FakeGCPSessionAffinityFilters) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(gcpsessionaffinityfiltersResource, c.ns, name, opts), &v1.GCPSessionAffinityFilter{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeGCPSessionAffinityFilters) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(gcpsessionaffinityfiltersResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1.GCPSessionAffinityFilterList{})
	return err
}

// Patch applies the patch and returns the patched gCPSessionAffinityFilter.
func (c *FakeGCPSessionAffinityFilters) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.GCPSessionAffinityFilter, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(gcpsessionaffinityfiltersResource, c.ns, name, pt, data, subresources...), &v1.GCPSessionAffinityFilter{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.GCPSessionAffinityFilter), err
}
