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
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="gateway.networking.k8s.io/policy=Direct"

// GCPBackendPolicy provides a way to apply LoadBalancer policy configuration with
// the GKE implementation of the Gateway API.
type GCPBackendPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// Spec defines the desired state of GCPBackendPolicy.
	Spec GCPBackendPolicySpec `json:"spec"`

	// Status defines the current state of GCPBackendPolicy.
	Status GCPBackendPolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// GCPBackendPolicyList contains a list of GCPBackendPolicy.
type GCPBackendPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GCPBackendPolicy `json:"items"`
}

// GCPBackendPolicySpec defines the desired state of GCPBackendPolicy.
type GCPBackendPolicySpec struct {
	// TargetRef identifies an API object to apply policy to.
	TargetRef v1alpha2.NamespacedPolicyTargetReference `json:"targetRef"`

	// Default defines default policy configuration for the targeted resource.
	// +optional
	Default *GCPBackendPolicyConfig `json:"default,omitempty"`
}

// GCPBackendPolicyConfig contains LoadBalancer policy configuration.
type GCPBackendPolicyConfig struct {
	Logging            *LoggingConfig         `json:"logging,omitempty"`
	SessionAffinity    *SessionAffinityConfig `json:"sessionAffinity,omitempty"`
	ConnectionDraining *ConnectionDraining    `json:"connectionDraining,omitempty"`
	// TimeoutSec is a BackendService parameter.
	// See https://cloud.google.com/compute/docs/reference/rest/v1/backendServices.
	// If the field is omitted, a default value (30s) will be used.
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=2147483647
	TimeoutSec *int64 `json:"timeoutSec,omitempty"`
	// SecurityPolicy is a reference to a GCP Cloud Armor SecurityPolicy resource.
	// +optional
	SecurityPolicy *string `json:"securityPolicy,omitempty"`
	// IAP contains the configurations for Identity-Aware Proxy.
	// See https://cloud.google.com/compute/docs/reference/rest/v1/backendServices
	// Identity-Aware Proxy manages access control policies for backend services associated with a HTTPRoute,
	// so they can be accessed only by authenticated users or applications with correct Identity and Access Management (IAM) role.
	// +optional
	IAP *IdentityAwareProxyConfig `json:"iap,omitempty"`
	// MaxRatePerEndpoint configures the target capacity for backends.
	// If the field is omitted, a default value (1e8) will be used.
	// In the future we may add selector based settings for MaxRatePerEndpoint but they will co-exist
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=1000000000
	// +optional
	MaxRatePerEndpoint *int64 `json:"maxRatePerEndpoint,omitempty"`
	// BackendPreference indicates whether the backend should be fully
	// utilized before sending traffic to backends with default preference.
	// Can only be configured for multi-cluster service backends when
	// GCPBackendPolicy targets ServiceExport.
	// The default value is DEFAULT.
	// +kubebuilder:validation:Enum=DEFAULT;PREFERRED
	// +optional
	BackendPreference *string `json:"backendPreference,omitempty"`
}

// ConnectionDraining contains configuration for connection draining
type ConnectionDraining struct {
	// DrainingTimeoutSec is a BackendService parameter.
	// It is used during removal of VMs from instance groups. This guarantees that for
	// the specified time all existing connections to a VM will remain untouched,
	// but no new connections will be accepted. Set timeout to zero to disable
	// connection draining. Enable the feature by specifying a timeout of up to
	// one hour. If the field is omitted, a default value (0s) will be used.
	// See https://cloud.google.com/compute/docs/reference/rest/v1/backendServices
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=3600
	DrainingTimeoutSec *int64 `json:"drainingTimeoutSec,omitempty"`
}

// LoggingConfig contains configuration for logging.
type LoggingConfig struct {
	// Enabled denotes whether to enable logging for the load balancer traffic
	// served by this backend service. If not specified, this defaults to false,
	// which means logging is disabled by default.
	Enabled *bool `json:"enabled,omitempty"`
	// This field can only be specified if logging is enabled for this backend
	// service. The value of the field must be in range [0, 1e6]. This is
	// converted to a floating point value in the range [0, 1] by dividing by 1e6
	// for use with the GCE api and interpreted as the proportion of requests that
	// will be logged. By default all requests will be logged.
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=1000000
	SampleRate *int32 `json:"sampleRate,omitempty"`
}

// SessionAffinityConfig contains configuration for stickiness parameters.
type SessionAffinityConfig struct {
	// Type specifies the type of session affinity to use. If not specified, this
	// defaults to NONE.
	// +kubebuilder:validation:Enum=CLIENT_IP;CLIENT_IP_PORT_PROTO;CLIENT_IP_PROTO;GENERATED_COOKIE;HEADER_FIELD;HTTP_COOKIE;NONE
	Type *string `json:"type,omitempty"`
	// CookieTTLSec specifies the lifetime of cookies in seconds. This setting
	// requires GENERATED_COOKIE or HTTP_COOKIE session affinity. If set to 0, the
	// cookie is non-persistent and lasts only until the end of the browser
	// session (or equivalent). The maximum allowed value is two weeks
	// (1,209,600).
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=1209600
	CookieTTLSec *int64 `json:"cookieTtlSec,omitempty"`
}

// IdentityAwareProxyConfig contains the configurations for Identity-Aware Proxy.
// Identity-Aware Proxy manages access control policies for backend services associated with a HTTPRoute,
// so they can be accessed only by authenticated users or applications with correct Identity and Access Management (IAM) role.
// See https://cloud.google.com/compute/docs/reference/rest/v1/backendServices
type IdentityAwareProxyConfig struct {
	// Enabled denotes whether the serving infrastructure will authenticate and authorize all incoming requests.
	// If true, the ClientID and Oauth2ClientSecret fields must be non-empty.
	// If not specified, this defaults to false, which means Identity-Aware Proxy is disabled by default.
	Enabled *bool `json:"enabled,omitempty"`
	// Oauth2ClientSecret contains the OAuth2 client secret to use for the authentication flow.
	// To use a custom OAuth client, provide both ClientID and Oauth2ClientSecret. If neither is provided, Google managed OAuth is the default.
	// +optional
	Oauth2ClientSecret *Oauth2ClientSecret `json:"oauth2ClientSecret,omitempty"`
	// ClientID is the OAuth2 client ID to use for the authentication flow.
	// See iap.oauth2ClientId in https://cloud.google.com/compute/docs/reference/rest/v1/backendServices
	// To use a custom OAuth client, provide both ClientID and Oauth2ClientSecret. If neither is provided, Google managed OAuth is the default.
	// +optional
	ClientID *string `json:"clientID,omitempty"`
}

// Oauth2ClientSecret contains the OAuth2 client secret to use for the authentication flow.
// See https://cloud.google.com/compute/docs/reference/rest/v1/backendServices
type Oauth2ClientSecret struct {
	// Name is the reference to the secret resource.
	Name *string `json:"name,omitempty"`
	// Namespace will be supported in the future if people ask for cross namespace IAP reference grant support.
	// Namespace *string `json:"namespace,omitempty"`
}

// GCPBackendPolicyStatus defines the observed state of GCPBackendPolicy.
type GCPBackendPolicyStatus struct {
	// Ancestors is a list of ancestor resources (usually Gateways) that are
	// associated with the policy, and the status of the policy with respect to
	// each ancestor.
	//
	// +optional
	// +kubebuilder:validation:MaxItems=16
	Ancestors []PolicyAncestorStatus `json:"ancestors,omitempty"`

	// Conditions describe the current conditions of the GCPBackendPolicy.
	//
	// +optional
	// +listType=map
	// +listMapKey=type
	// +kubebuilder:validation:MaxItems=8
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
