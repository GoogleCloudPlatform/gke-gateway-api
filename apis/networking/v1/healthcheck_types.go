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

// HealthCheckType is the HealthCheck protocol type.
type HealthCheckType string

const (
	// TCP is ProtocolType of TCP
	TCP HealthCheckType = "TCP"
	// HTTP is ProtocolType of HTTP
	HTTP HealthCheckType = "HTTP"
	// HTTPS is ProtocolType of HTTPS
	HTTPS HealthCheckType = "HTTPS"
	// HTTP2 is ProtocolType of HTTP2
	HTTP2 HealthCheckType = "HTTP2"
	// GRPC is ProtocolType of GRPC
	GRPC HealthCheckType = "GRPC"
)

// PortSpecificationType is the PortSpecification type.
type PortSpecificationType string

const (
	// UseFixedPort is PortSpecificationType of USE_FIXED_PORT
	UseFixedPort PortSpecificationType = "USE_FIXED_PORT"
	// UseNamedPort is PortSpecificationType of USE_NAMED_PORT
	UseNamedPort PortSpecificationType = "USE_NAMED_PORT"
	// UseServingPort is PortSpecificationType of USE_SERVING_PORT
	UseServingPort PortSpecificationType = "USE_SERVING_PORT"
)

// ProxyHeaderType is the ProxyHeader type.
type ProxyHeaderType string

const (
	// None is ProxyHeaderType of NONE
	None ProxyHeaderType = "NONE"
	// ProxyV1 is ProxyHeaderType of PROXY_V1
	ProxyV1 ProxyHeaderType = "PROXY_V1"
)

// +genclient
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="gateway.networking.k8s.io/policy=Direct"

// HealthCheckPolicy provides a way to create and attach a HealthCheck to a BackendService with
// the GKE implementation of the Gateway API. This policy can only be attached to a BackendService.
type HealthCheckPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// Spec defines the desired state of HealthCheckPolicy.
	Spec HealthCheckPolicySpec `json:"spec"`

	// Status defines the current state of HealthCheckPolicy.
	Status HealthCheckPolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// HealthCheckPolicyList contains a list of HealthCheckPolicy.
type HealthCheckPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []HealthCheckPolicy `json:"items"`
}

// HealthCheckPolicySpec defines the desired state of HealthCheckPolicy.
type HealthCheckPolicySpec struct {
	// TargetRef identifies an API object to apply policy to.
	TargetRef v1alpha2.NamespacedPolicyTargetReference `json:"targetRef"`

	// Default defines default policy configuration for the targeted resource.
	// +optional
	Default *HealthCheckPolicyConfig `json:"default,omitempty"`
}

// HealthCheckPolicyConfig contains HealthCheck policy configuration.
// +kubebuilder:validation:XValidation:rule="has(self.checkIntervalSec) && has(self.timeoutSec) ? self.checkIntervalSec >= self.timeoutSec : true",message="timeOutSec cannot exceed checkIntervalSec"
// +kubebuilder:validation:XValidation:rule="!has(self.checkIntervalSec) && has(self.timeoutSec) ? 5 >= self.timeoutSec : true",message="when checkIntervalSec is unspecified, timeOutSec cannot exceed 5, which is the default value of checkIntervalSec"
// +kubebuilder:validation:XValidation:rule="has(self.checkIntervalSec) && !has(self.timeoutSec) ? self.checkIntervalSec >= 5 : true",message="when timeoutSec is unspecified, checkIntervalSec must be at least 5, which is the default value of timeoutSec"
type HealthCheckPolicyConfig struct {
	// How often (in seconds) to send a health check.
	// If not specified, a default value of 5 seconds will be used.
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=300
	CheckIntervalSec *int64 `json:"checkIntervalSec,omitempty"`
	// How long (in seconds) to wait before claiming failure.
	// If not specified, a default value of 5 seconds will be used.
	// It is invalid for timeoutSec to have greater value than checkIntervalSec.
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=300
	TimeoutSec *int64 `json:"timeoutSec,omitempty"`
	// A so-far healthy instance will be marked unhealthy after this many consecutive failures.
	// If not specified, a default value of 2 will be used.
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=10
	UnhealthyThreshold *int64 `json:"unhealthyThreshold,omitempty"`
	// A so-far unhealthy instance will be marked healthy after this many consecutive successes.
	// If not specified, a default value of 2 will be used.
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=10
	HealthyThreshold *int64 `json:"healthyThreshold,omitempty"`
	// Specifies the type of the healthCheck, either TCP, HTTP, HTTPS, HTTP2 or GRPC.
	// Exactly one of the protocol-specific health check field must be specified,
	// which must match type field.
	// Config contains per protocol (i.e. HTTP, HTTPS, HTTP2, TCP, GRPC) configuration.
	// If not specified, health check type defaults to HTTP.
	Config *HealthCheck `json:"config,omitempty"`
	// LogConfig configures logging on this health check.
	LogConfig *LogConfig `json:"logConfig,omitempty"`
}

// HealthCheck is a union struct that contains per protocol (i.e. HTTP, HTTPS, HTTP2, TCP, GRPC)
// configuration.
// +union
// +kubebuilder:validation:MaxProperties=2
// +kubebuilder:validation:MinProperties=2
type HealthCheck struct {
	// Specifies the type of the healthCheck, either TCP, HTTP, HTTPS, HTTP2 or GRPC.
	// Exactly one of the protocol-specific health check field must be specified,
	// which must match type field.
	// +unionDiscriminator
	// +kubebuilder:validation:Enum=TCP;HTTP;HTTPS;HTTP2;GRPC
	Type HealthCheckType `json:"type,omitempty"`
	// TCP is the health check configuration of type TCP.
	// +optional
	TCP *TCPHealthCheck `json:"tcpHealthCheck,omitempty"`
	// HTTP is the health check configuration of type HTTP.
	// +optional
	HTTP *HTTPHealthCheck `json:"httpHealthCheck,omitempty"`
	// HTTPS is the health check configuration of type HTTPS.
	// +optional
	HTTPS *HTTPSHealthCheck `json:"httpsHealthCheck,omitempty"`
	// HTTP2 is the health check configuration of type HTTP2.
	// +optional
	HTTP2 *HTTP2HealthCheck `json:"http2HealthCheck,omitempty"`
	// GRPC is the health check configuration of type GRPC.
	// +optional
	GRPC *GRPCHealthCheck `json:"grpcHealthCheck,omitempty"`
}

// CommonHealthCheck holds all the fields that are common across all protocol health checks.
// +union
type CommonHealthCheck struct {
	// Specifies how port is selected for health checking, can be one of following values:
	//
	// USE_FIXED_PORT: The port number in port is used for health checking.
	// USE_NAMED_PORT: The portName is used for health checking.
	// USE_SERVING_PORT: For NetworkEndpointGroup, the port specified for each network endpoint
	// is used for health checking. For other backends, the port or named port specified in the
	// Backend Service is used for health checking.
	//
	// If not specified, Protocol health check follows behavior specified in port and portName fields.
	// If neither Port nor PortName is specified, this defaults to USE_SERVING_PORT.
	// +unionDiscriminator
	// +kubebuilder:validation:Enum=USE_FIXED_PORT;USE_NAMED_PORT;USE_SERVING_PORT
	PortSpecification *PortSpecificationType `json:"portSpecification,omitempty"`
	// The TCP port number for the health check request. Valid values are 1 through 65535.
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=65535
	// +optional
	Port *int64 `json:"port,omitempty"`
	// Port name as defined in InstanceGroup#NamedPort#name.
	// If both port and portName are defined, port takes precedence.
	// +optional
	// +kubebuilder:validation:MaxLength=63
	// +kubebuilder:validation:Pattern=[a-z]([-a-z0-9]*[a-z0-9])?
	PortName *string `json:"portName,omitempty"`
}

// CommonHTTPHealthCheck holds all the fields that are common across all HTTP health checks.
type CommonHTTPHealthCheck struct {
	// Host is the value of the host header in the HTTP health check request. This
	// matches the RFC 1123 definition of a hostname with 1 notable exception that
	// numeric IP addresses are not allowed.
	// If not specified or left empty, the IP on behalf of which this health check is
	// performed will be used.
	// +kubebuilder:validation:MaxLength=2048
	// +kubebuilder:validation:Pattern=`^[a-z0-9]([-a-z0-9]*[a-z0-9])?(\.[a-z0-9]([-a-z0-9]*[a-z0-9])?)*$`
	Host *string `json:"host,omitempty"`
	// The request path of the HTTP health check request.
	// If not specified or left empty, a default value of "/" is used.
	// +kubebuilder:validation:MaxLength=2048
	// +kubebuilder:validation:Pattern=`\/[A-Za-z0-9\/\-._~%!?$&'()*+,;=:]*$`
	RequestPath *string `json:"requestPath,omitempty"`
	// Specifies the type of proxy header to append before sending data to the backend,
	// either NONE or PROXY_V1. If not specified, this defaults to NONE.
	// +kubebuilder:validation:Enum=NONE;PROXY_V1
	ProxyHeader *ProxyHeaderType `json:"proxyHeader,omitempty"`
	// The string to match anywhere in the first 1024 bytes of the response body.
	// If not specified or left empty, the status code determines health.
	// The response data can only be ASCII.
	// +kubebuilder:validation:MaxLength=1024
	// +kubebuilder:validation:Pattern=[\x00-\xFF]+
	Response *string `json:"response,omitempty"`
}

// TCPHealthCheck is the health check configuration of type TCP
type TCPHealthCheck struct {
	CommonHealthCheck `json:",inline"`
	// The application data to send once the TCP connection has been established. If not specified,
	// this defaults to empty. If both request and response are empty, the connection establishment
	// alone will indicate health. The request data can only be ASCII.
	// +kubebuilder:validation:MaxLength=1024
	// +kubebuilder:validation:Pattern=[\x00-\xFF]+
	Request *string `json:"request,omitempty"`
	// The bytes to match against the beginning of the response data.
	// If not specified or left empty, any response will indicate health.
	// The response data can only be ASCII.
	// +kubebuilder:validation:MaxLength=1024
	// +kubebuilder:validation:Pattern=[\x00-\xFF]+
	Response *string `json:"response,omitempty"`
	// Specifies the type of proxy header to append before sending data to the backend,
	// either NONE or PROXY_V1. If not specified, this defaults to NONE.
	// +kubebuilder:validation:Enum=NONE;PROXY_V1
	ProxyHeader *ProxyHeaderType `json:"proxyHeader,omitempty"`
}

// HTTPHealthCheck is the health check configuration of type HTTP
type HTTPHealthCheck struct {
	CommonHealthCheck     `json:",inline"`
	CommonHTTPHealthCheck `json:",inline"`
}

// HTTPSHealthCheck is the health check configuration of type HTTPS
type HTTPSHealthCheck struct {
	CommonHealthCheck     `json:",inline"`
	CommonHTTPHealthCheck `json:",inline"`
}

// HTTP2HealthCheck is the health check configuration of type HTTP2
type HTTP2HealthCheck struct {
	CommonHealthCheck     `json:",inline"`
	CommonHTTPHealthCheck `json:",inline"`
}

// GRPCHealthCheck is the health check configuration of type GRPC
type GRPCHealthCheck struct {
	CommonHealthCheck `json:",inline"`
	// The gRPC service name for the health check. This field is optional.
	// The value of grpcServiceName has the following meanings by convention:
	// - Empty serviceName means the overall status of all services at the backend.
	// - Non-empty serviceName means the health of that gRPC service, as defined by
	//   the owner of the service.
	// The grpcServiceName can only be ASCII.
	// +kubebuilder:validation:MaxLength=1024
	// +kubebuilder:validation:Pattern=[\x00-\xFF]+
	GRPCServiceName *string `json:"grpcServiceName,omitempty"`
}

// LogConfig configures logging on this health check.
type LogConfig struct {
	// Enabled indicates whether or not to export health check logs. If not
	// specified, this defaults to false, which means health check logging will be
	// disabled.
	Enabled *bool `json:"enabled,omitempty"`
}

// HealthCheckPolicyStatus defines the observed state of HealthCheckPolicy.
type HealthCheckPolicyStatus struct {
	// Conditions describe the current conditions of the HealthCheckPolicy.
	//
	// +optional
	// +listType=map
	// +listMapKey=type
	// +kubebuilder:validation:MaxItems=8
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}
