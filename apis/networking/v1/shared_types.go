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

// PolicyConditionType is a type of condition for a service policy.
type PolicyConditionType string

// PolicyConditionReason is a reason for a service policy condition.
type PolicyConditionReason string

const (
	// PolicyConditionAttached indicates whether the policy has been accepted or rejected
	// by a targeted resource, and why.
	//
	// Possible reasons for this condition to be true are:
	//
	// * "Attached"
	//
	// Possible reasons for this condition to be False are:
	//
	// * "Conflicted"
	//
	PolicyConditionAttached PolicyConditionType = "Attached"

	// PolicyReasonAttached is used with the "Attached" condition when the policy has been
	// accepted by the targeted resource.
	PolicyReasonAttached PolicyConditionReason = "Attached"

	// PolicyReasonConflicted is used with the "Attached" condition when the policy has not
	// been accepted by a targeted resource because there is another policy that targets the same
	// resource and has higher precedence.
	PolicyReasonConflicted PolicyConditionReason = "Conflicted"

	// PolicyReasonInvalid is used with the "Attached" condition when the policy is syntactically
	// or semantically invalid.
	PolicyReasonInvalid PolicyConditionReason = "Invalid"

	// PolicyReasonTargetNotFound is used with the "Attached" condition when the policy is attached to
	// an invalid target resource
	PolicyReasonTargetNotFound PolicyConditionReason = "TargetNotFound"
)


// LocalPolicyTargetReference identifies an API object to apply a direct or
// inherited policy to. This should be used as part of Policy resources
// that can target Gateway API resources. For more information on how this
// policy attachment model works, and a sample Policy resource, refer to
// the policy attachment documentation for Gateway API.
type LocalPolicyTargetReference struct {
	// Group is the group of the target resource.
	Group v1.Group `json:"group"`

	// Kind is kind of the target resource.
	Kind v1.Kind `json:"kind"`

	// Name is the name of the target resource.
	Name v1.ObjectName `json:"name"`
}

// PolicyAncestorStatus describes the status of a route with respect to an
// associated Ancestor.
//
// Ancestors refer to objects that are either the Target of a policy or above it
// in terms of object hierarchy. For example, if a policy targets a Service, the
// Policy's Ancestors are, in order, the Service, the HTTPRoute, the Gateway, and
// the GatewayClass. Almost always, in this hierarchy, the Gateway will be the most
// useful object to place Policy status on, so we recommend that implementations
// SHOULD use Gateway as the PolicyAncestorStatus object unless the designers
// have a _very_ good reason otherwise.
//
// In the context of policy attachment, the Ancestor is used to distinguish which
// resource results in a distinct application of this policy. For example, if a policy
// targets a Service, it may have a distinct result per attached Gateway.
//
// Policies targeting the same resource may have different effects depending on the
// ancestors of those resources. For example, different Gateways targeting the same
// Service may have different capabilities, especially if they have different underlying
// implementations.
//
// For example, in BackendTLSPolicy, the Policy attaches to a Service that is
// used as a backend in a HTTPRoute that is itself attached to a Gateway.
// In this case, the relevant object for status is the Gateway, and that is the
// ancestor object referred to in this status.
//
// Note that a parent is also an ancestor, so for objects where the parent is the
// relevant object for status, this struct SHOULD still be used.
//
// This struct is intended to be used in a slice that's effectively a map,
// with a composite key made up of the AncestorRef and the ControllerName.
type PolicyAncestorStatus struct {
	// AncestorRef corresponds with a ParentRef in the spec that this
	// PolicyAncestorStatus struct describes the status of.
	AncestorRef v1.ParentReference `json:"ancestorRef"`

	// ControllerName is a domain/path string that indicates the name of the
	// controller that wrote this status. This corresponds with the
	// controllerName field on GatewayClass.
	//
	// Example: "example.net/gateway-controller".
	//
	// The format of this field is DOMAIN "/" PATH, where DOMAIN and PATH are
	// valid Kubernetes names
	// (https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names).
	//
	// Controllers MUST populate this field when writing status. Controllers should ensure that
	// entries to status populated with their ControllerName are cleaned up when they are no
	// longer necessary.
	ControllerName v1.GatewayController `json:"controllerName"`

	// Conditions describes the status of the Policy with respect to the given Ancestor.
	//
	// +listType=map
	// +listMapKey=type
	// +kubebuilder:validation:MinItems=1
	// +kubebuilder:validation:MaxItems=8
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}

// PolicyStatus defines the common attributes that all Policies should include within
// their status.
type PolicyStatus struct {
	// Ancestors is a list of ancestor resources (usually Gateways) that are
	// associated with the policy, and the status of the policy with respect to
	// each ancestor. When this policy attaches to a parent, the controller that
	// manages the parent and the ancestors MUST add an entry to this list when
	// the controller first sees the policy and SHOULD update the entry as
	// appropriate when the relevant ancestor is modified.
	//
	// Note that choosing the relevant ancestor is left to the Policy designers;
	// an important part of Policy design is designing the right object level at
	// which to namespace this status.
	//
	// Note also that implementations MUST ONLY populate ancestor status for
	// the Ancestor resources they are responsible for. Implementations MUST
	// use the ControllerName field to uniquely identify the entries in this list
	// that they are responsible for.
	//
	// Note that to achieve this, the list of PolicyAncestorStatus structs
	// MUST be treated as a map with a composite key, made up of the AncestorRef
	// and ControllerName fields combined.
	//
	// A maximum of 16 ancestors will be represented in this list. An empty list
	// means the Policy is not relevant for any ancestors.
	//
	// If this slice is full, implementations MUST NOT add further entries.
	// Instead they MUST consider the policy unimplementable and signal that
	// on any related resources such as the ancestor that would be referenced
	// here. For example, if this list was full on BackendTLSPolicy, no
	// additional Gateways would be able to reference the Service targeted by
	// the BackendTLSPolicy.
	//
	// +kubebuilder:validation:MaxItems=16
	Ancestors []PolicyAncestorStatus `json:"ancestors"`
}
