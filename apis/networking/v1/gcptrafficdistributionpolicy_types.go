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

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

// +genclient
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="gateway.networking.k8s.io/policy=Direct"

// GCPTrafficDistributionPolicy contains settings that configure how traffic should
// be distributed to its targeting service(s).
type GCPTrafficDistributionPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// Spec defines the desired state for traffic distribution policy settings.
	Spec GCPTrafficDistributionPolicySpec `json:"spec"`

	// Status provides the current state of GCPTrafficDistributionPolicy.
	Status PolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// GCPTrafficDistributionPolicyList contains a list of GCPTrafficDistributionPolicy.
type GCPTrafficDistributionPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GCPTrafficDistributionPolicy `json:"items"`
}

// GCPTrafficDistributionPolicySpec defines the desired state of GCPTrafficDistributionPolicy.
type GCPTrafficDistributionPolicySpec struct {
	// TargetRefs identifies an API object to apply policy to.
	// A GCPTrafficDistributionPolicy can only target Service
	// local namespace.

	// +kubebuilder:validation:XValidation:message="TargetRefs must reference Service",rule="self.all(x, x.kind == 'Service' && x.group == '')"
	// +kubebuilder:validation:MinItems=1
	// +kubebuilder:validation:MaxItems=16
	TargetRefs []LocalPolicyTargetReference `json:"targetRefs"`

	// Default defines default policy configuration for the targeted resource.
	// +optional
	Default *GCPTrafficDistributionPolicyConfig `json:"default,omitempty"`
}

// GCPTrafficDistributionPolicyConfig defines the settings of GCPTrafficDistributionPolicy.
type GCPTrafficDistributionPolicyConfig struct {
	// The load balancing algorithm used to determine traffic distribution weighting at
	// cluster/zone level.
	// ServiceLbAlgorithm works together with LocalityLbAlgorithm.
	// Refer to https://cloud.google.com/load-balancing/docs/service-lb-policy for a
	// more detailed explanation of how they work together.
	// Supported values: SPRAY_TO_REGION / WATERFALL_BY_ZONE / WATERFALL_BY_REGION
	// Refer to https://cloud.google.com/load-balancing/docs/service-lb-policy#lb-algos
	// explanation of the algorithms.
	// Default to WATERFALL_BY_REGION.
	// +kubebuilder:validation:Enum=SPRAY_TO_REGION;WATERFALL_BY_ZONE;WATERFALL_BY_REGION
	ServiceLbAlgorithm *string `json:"serviceLbAlgorithm,omitempty"`

	// The load balancing algorithm used within the scope of the locality. This algorithm
	// affects how an individual endpoint is selected for a particular request.
	// LocalityLbAlgorithm works together with ServiceLbAlgorithm.
	// Refer to https://cloud.google.com/load-balancing/docs/service-lb-policy for a
	// more detailed explanation of how they work together.
	// Default to ROUND_ROBIN.
	// +kubebuilder:validation:Enum=ROUND_ROBIN;LEAST_REQUEST;RING_HASH;RANDOM;ORIGINAL_DESTINATION;MAGLEV;WEIGHTED_ROUND_ROBIN
	LocalityLbAlgorithm *string `json:"localityLbAlgorithm,omitempty"`

	// AutoCapacityDrain contains configurations for auto draining.
	//
	// +optional
	AutoCapacityDrain *AutoCapacityDrain `json:"autoCapacityDrain,omitempty"`

	// FailoverConfig contains configurations for failover behaviors.
	//
	// +optional
	FailoverConfig *FailoverConfig `json:"failoverConfig,omitempty"`
}

// AutoCapacityDrain contains configurations for auto draining.
type AutoCapacityDrain struct {
	// If set to 'True', backends in a certain (cluster, zone) will be
	// drained(considered to have 0 capacity) when less than 25% of the endpoints
	// there are healthy. Default to false.
	EnableAutoCapacityDrain *bool `json:"enableAutoCapacityDrain,omitempty"`
}

// FailoverConfig contains configurations for failover behaviors.
type FailoverConfig struct {
	// The percentage threshold that a load balancer will begin to send traffic
	// to failover backends. When not specified, the dataplane uses its own
	// builtin default value. For Envoy the default value is 70. Proxyless gRPC
	// defaults to 50.
	//
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=100
	FailoverHealthThreshold *int32 `json:"failoverHealthThreshold,omitempty"`
}
