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
	v1alpha2 "sigs.k8s.io/gateway-api/apis/v1alpha2"
)

// +genclient
// +kubebuilder:object:root=true
// +kubebuilder:resource:categories=gateway-api
// +kubebuilder:storageversion
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Age",type=date,JSONPath=`.metadata.creationTimestamp`

// GCPSessionAffinityPolicy provides a way to apply session affinity policy configuration.
type GCPSessionAffinityPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// Spec defines the desired state of GCPSessionAffinityPolicy.
	Spec GCPSessionAffinityPolicySpec `json:"spec"`

	// Status defines the current state of GCPSessionAffinityPolicy.
	Status GCPSessionAffinityPolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// GCPSessionAffinityPolicyList contains a list of GCPSessionAffinityPolicy.
type GCPSessionAffinityPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GCPSessionAffinityPolicy `json:"items"`
}

// GCPSessionAffinityPolicySpec defines the desired state of GCPSessionAffinityPolicy.
type GCPSessionAffinityPolicySpec struct {
	// GCPSessionAffinitySpec is shared with GCPSessionAffinityFilter
	GCPSessionAffinitySpec `json:",inline"`

	// TargetRef identifies an API object to apply policy to.
	TargetRef v1alpha2.NamespacedPolicyTargetReference `json:"targetRef"`
}

// GCPSessionAffinityPolicyStatus defines the observed state of GCPSessionAffinityPolicy.
type GCPSessionAffinityPolicyStatus struct {
	// Conditions describe the current conditions of the GCPSessionAffinityPolicy.
	//
	// +optional
	// +listType=map
	// +listMapKey=type
	// +kubebuilder:validation:MaxItems=8
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
