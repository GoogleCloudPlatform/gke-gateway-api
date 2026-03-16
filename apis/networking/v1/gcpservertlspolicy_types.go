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

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	v1 "sigs.k8s.io/gateway-api/apis/v1"
)

// MTLSMode defines the mTLS mode for inbound connections on the targeted workload(s).
// +kubebuilder:validation:Enum=Disabled;Permissive;Strict
type MTLSMode string

const (
	// Permissive allows connection to be plaintext or mTLS.
	Permissive MTLSMode = "Permissive"
	// Strict allows connection to be only mTLS.
	Strict MTLSMode = "Strict"
	// Disabled allows connection to be plaintext only.
	Disabled MTLSMode = "Disabled"
)

// PortOverride defines port-specific mTLS settings.
type PortOverride struct {
	// Port specifies the port number to which this override applies.
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=65535
	// +required
	Port int32 `json:"port,omitempty"`

	// MtlsMode specifies the mTLS mode for the specified port.
	// Overrides the default MTLSMode set in the policy specification.
	// +required
	MtlsMode MTLSMode `json:"mtlsMode,omitempty"`
}

// PolicyTargetReferenceWithLabelSelectors specifies a reference to a set of Kubernetes
// objects by Group and Kind, with an optional label selector to narrow down the matching
// objects.
//
// Currently, we only support label selectors when targeting Pods.
// This restriction is intentional to limit the complexity and potential
// ambiguity of supporting label selectors for arbitrary Kubernetes kinds.
// Unless there is a very strong justification in the future, we plan to keep this
// functionality limited to selecting Pods only.
//
// This is currently experimental in the Gateway API and should only be used
// for policies implemented within Gateway API. It is currently not intended for general-purpose
// use outside of Gateway API resources.
// +kubebuilder:validation:XValidation:rule="!has(self.selector.matchLabels) || self.kind == 'Pod'",message="selector.matchLabels can only be used when targeting pods."
// +kubebuilder:validation:XValidation:rule="!has(self.selector.matchExpressions) || size(self.selector.matchExpressions) == 0",message="selector.matchExpressions are not supported."
type PolicyTargetReferenceWithLabelSelectors struct {
	// Group is the group of the target object.
	// +optional
	Group v1.Group `json:"group,omitempty"`

	// Kind is the kind of the target object.
	// +required
	Kind v1.Kind `json:"kind,omitempty"`

	// Selector is the label selector of target objects of the specified kind.
	// +required
	Selector metav1.LabelSelector `json:"selector,omitempty"`
}

// GCPServerTLSPolicySpec defines the desired state of GCPServerTLSPolicy.
// +kubebuilder:validation:XValidation:rule="!(self.targetRefs.exists(ref, (!has(ref.selector.matchLabels) || size(ref.selector.matchLabels) == 0)) && has(self.portOverrides))",message="portOverrides cannot be set when targeting a whole namespace (i.e., when the TargetRefs selector.matchLabels is absent or empty)."
type GCPServerTLSPolicySpec struct {
	// MTLSMode defines the default mutual TLS settings for the inbound connections on the targeted workload(s).
	// Can be one of Disabled | Permissive | Strict.
	// Defaults to Strict if not set.
	// This applies to all ports unless overridden by PortOverride.
	// +kubebuilder:default=Strict
	// +optional
	MTLSMode *MTLSMode `json:"mtlsMode,omitempty"`

	// PortOverrides allows specifying different mTLS settings for individual ports
	// on the targeted workload(s). If a port is not listed, it inherits the
	// default MTLSMode.
	// +kubebuilder:validation:MaxItems=5
	// +listType=map
	// +listMapKey=port
	// +optional
	PortOverrides []PortOverride `json:"portOverrides,omitempty"`

	// Policy selection works as follows with precedence from highest to lowest:
	// 1. Workload port policy (workload matched by label selector, and matching port in `portOverrides`)
	// 2. Workload policy (workload matched by label selector)
	// 3. Namespace policy (label selector is absent or empty)

	// At most one policy can be picked for a given workload port. If there's no workload port-specific policy, the system checks for a workload policy. If that doesn't exist, the system checks for a namespace policy. If no namespace policy exists, then MTLS is disabled and workload will accept clear text traffic across all of its ports.
	// If multiple policies are defined at the same level of precedence for the same target, the one with the oldest `creationTimestamp` is chosen. It is highly recommended to avoid configuring multiple policies for the same target.
	// +kubebuilder:validation:MinItems=1
	// +kubebuilder:validation:MaxItems=1
	// +required
	TargetRefs []PolicyTargetReferenceWithLabelSelectors `json:"targetRefs,omitempty"`
}

// +genclient
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:categories=gateway-api
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:name="Age",type=date,JSONPath=`.metadata.creationTimestamp`

// GCPServerTLSPolicy is the Schema for the GCPServerTLSPolicy API.
// It configures server-side inbound authentication settings (mTLS mode) for workloads
// within a specific namespace, based either on workload labels (selector), or
// applying to the entire namespace.
type GCPServerTLSPolicy struct {
	// +optional
	metav1.TypeMeta `json:",inline,omitempty"`
	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// Spec defines the implementation of this definition.
	// +required
	Spec GCPServerTLSPolicySpec `json:"spec,omitempty"`

	// +optional
	Status v1.PolicyStatus `json:"status,omitempty"`
}

// GCPServerTLSPolicyList contains a list of GCPServerTLSPolicy resources.
// +kubebuilder:object:root=true
type GCPServerTLSPolicyList struct {
	// +optional
	metav1.TypeMeta `json:",inline,omitempty"`
	// +optional
	metav1.ListMeta `json:"metadata,omitempty"`
	// +optional
	Items []GCPServerTLSPolicy `json:"items,omitempty"`
}
