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

// FakeGCPGatewayPolicies implements GCPGatewayPolicyInterface
type FakeGCPGatewayPolicies struct {
	Fake *FakeNetworkingV1
	ns   string
}

var gcpgatewaypoliciesResource = v1.SchemeGroupVersion.WithResource("gcpgatewaypolicies")

var gcpgatewaypoliciesKind = v1.SchemeGroupVersion.WithKind("GCPGatewayPolicy")

// Get takes name of the gCPGatewayPolicy, and returns the corresponding gCPGatewayPolicy object, and an error if there is any.
func (c *FakeGCPGatewayPolicies) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.GCPGatewayPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(gcpgatewaypoliciesResource, c.ns, name), &v1.GCPGatewayPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.GCPGatewayPolicy), err
}

// List takes label and field selectors, and returns the list of GCPGatewayPolicies that match those selectors.
func (c *FakeGCPGatewayPolicies) List(ctx context.Context, opts metav1.ListOptions) (result *v1.GCPGatewayPolicyList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(gcpgatewaypoliciesResource, gcpgatewaypoliciesKind, c.ns, opts), &v1.GCPGatewayPolicyList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1.GCPGatewayPolicyList{ListMeta: obj.(*v1.GCPGatewayPolicyList).ListMeta}
	for _, item := range obj.(*v1.GCPGatewayPolicyList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested gCPGatewayPolicies.
func (c *FakeGCPGatewayPolicies) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(gcpgatewaypoliciesResource, c.ns, opts))

}

// Create takes the representation of a gCPGatewayPolicy and creates it.  Returns the server's representation of the gCPGatewayPolicy, and an error, if there is any.
func (c *FakeGCPGatewayPolicies) Create(ctx context.Context, gCPGatewayPolicy *v1.GCPGatewayPolicy, opts metav1.CreateOptions) (result *v1.GCPGatewayPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(gcpgatewaypoliciesResource, c.ns, gCPGatewayPolicy), &v1.GCPGatewayPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.GCPGatewayPolicy), err
}

// Update takes the representation of a gCPGatewayPolicy and updates it. Returns the server's representation of the gCPGatewayPolicy, and an error, if there is any.
func (c *FakeGCPGatewayPolicies) Update(ctx context.Context, gCPGatewayPolicy *v1.GCPGatewayPolicy, opts metav1.UpdateOptions) (result *v1.GCPGatewayPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(gcpgatewaypoliciesResource, c.ns, gCPGatewayPolicy), &v1.GCPGatewayPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.GCPGatewayPolicy), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeGCPGatewayPolicies) UpdateStatus(ctx context.Context, gCPGatewayPolicy *v1.GCPGatewayPolicy, opts metav1.UpdateOptions) (*v1.GCPGatewayPolicy, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(gcpgatewaypoliciesResource, "status", c.ns, gCPGatewayPolicy), &v1.GCPGatewayPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.GCPGatewayPolicy), err
}

// Delete takes name of the gCPGatewayPolicy and deletes it. Returns an error if one occurs.
func (c *FakeGCPGatewayPolicies) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(gcpgatewaypoliciesResource, c.ns, name, opts), &v1.GCPGatewayPolicy{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeGCPGatewayPolicies) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(gcpgatewaypoliciesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1.GCPGatewayPolicyList{})
	return err
}

// Patch applies the patch and returns the patched gCPGatewayPolicy.
func (c *FakeGCPGatewayPolicies) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.GCPGatewayPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(gcpgatewaypoliciesResource, c.ns, name, pt, data, subresources...), &v1.GCPGatewayPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.GCPGatewayPolicy), err
}