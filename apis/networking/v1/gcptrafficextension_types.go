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

// +genclient
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
//
// +kubebuilder:metadata:labels="gateway.networking.k8s.io/policy=Direct"

// GCPTrafficExtension is the CRD for the Traffic Extension.
// It provides a way to add custom logic into Cloud Load Balancers by allowing
// the extension service to modify the headers and payloads of both requests
// and responses without impacting the choice of backend services or
// any other security policies associated with the backend service.
type GCPTrafficExtension struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// Spec defines the desired state of GCPTrafficExtension.
	Spec GCPTrafficExtensionSpec `json:"spec"`

	// Status defines the current state of GCPTrafficExtension.
	Status v1.PolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// GCPTrafficExtensionList contains a list of GCPTrafficExtensions.
type GCPTrafficExtensionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GCPTrafficExtension `json:"items"`
}

// GCPTrafficExtensionSpec defines the desired state of GCPTrafficExtension.
type GCPTrafficExtensionSpec struct {
	// TargetRefs is a list of API objects this extension applies to.
	// Valid Groups are:
	// - "gateway.networking.k8s.io"
	//
	// Valid Kinds are:
	// - "Gateway"
	//
	// +kubebuilder:validation:MinItems=1
	// +kubebuilder:validation:MaxItems=10
	TargetRefs []v1.LocalObjectReference `json:"targetRefs"`

	// ExtensionChains is a set of ordered extension chains that contain
	// the match conditions and extensions to execute. Match conditions for each
	// extension chain are evaluated in sequence for a given request. The first
	// extension chain that has conditions that match the request is executed.
	// Any subsequent extension chains do not execute.
	// Limited to 5 ExtensionChains.
	//
	// +kubebuilder:validation:MinItems=1
	// +kubebuilder:validation:MaxItems=5
	// +kubebuilder:validation:XValidation:message="supportedEvents must be set for GCPTrafficExtension",rule="self.all(ec, ec.extensions.all(e, size(e.supportedEvents) > 0))"
	ExtensionChains []ExtensionChain `json:"extensionChains"`
}
