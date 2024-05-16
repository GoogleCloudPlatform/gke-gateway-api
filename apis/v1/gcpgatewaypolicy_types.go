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

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// GCPGatewayPolicy provides a way to apply SSL policy and other configuration to
// the GKE Gateway.
type GCPGatewayPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// Spec defines the desired state of GCPGatewayPolicy.
	Spec GCPGatewayPolicySpec `json:"spec"`

	// Status defines the current state of GCPGatewayPolicy.
	Status GCPGatewayPolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// GCPGatewayPolicyList contains a list of GCPGatewayPolicies.
type GCPGatewayPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []*GCPGatewayPolicy `json:"items"`
}

// GCPGatewayPolicySpec defines the desired state of GCPGatewayPolicy.
type GCPGatewayPolicySpec struct {
	// TargetRef identifies an API object to apply policy to.
	TargetRef v1alpha2.PolicyTargetReference `json:"targetRef"`

	// Default defines default gateway policy configuration for the targeted resource.
	// +optional
	Default *GCPGatewayPolicyConfig `json:"default,omitempty"`
}

// GCPGatewayPolicyConfig contains gateway policy configuration.
type GCPGatewayPolicyConfig struct {
	// SslPolicy can be a raw name or a path. If it is set to a raw name, the region of the SslPolicy
	// will derive from the attached load balancer

	// Examples:
	// "[name-of-ssl-policy]"
	// "projects/[projectID]/global/targetHttpsProxies/[name-of-ssl-policy]"
	// "projects/[projectID]/regions/[region]/targetHttpsProxies/[name-of-ssl-policy]"

	// +optional
	SslPolicy string `json:"sslPolicy,omitempty"`
	// +optional
	AllowGlobalAccess bool `json:"allowGlobalAccess,omitempty"`
	// Region allows to specify load balancer's region for Multi-cluster Gateway.
	// +optional
	Region string `json:"region,omitempty"`
}

// GCPGatewayPolicyStatus defines the observed state of GCPGatewayPolicy.
type GCPGatewayPolicyStatus struct {
	// Conditions describe the current conditions of the GatewayPolicy.
	//
	// +optional
	// +listType=map
	// +listMapKey=type
	// +kubebuilder:validation:MaxItems=8
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
