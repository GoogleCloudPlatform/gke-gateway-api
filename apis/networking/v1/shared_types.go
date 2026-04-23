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

// ExtensionConditionType is a type of condition for a extension.
type ExtensionConditionType string

// ExtensionConditionReason is a reason for a extension condition.
type ExtensionConditionReason string

const (
	// ExtensionConditionAccepted indicates whether the extension or WasmPlugin
	// has been accepted or rejected and why.
	//
	// Possible reasons for this condition to be true are:
	//
	// * "Accepted"
	//
	// Possible reasons for this condition to be False are:
	//
	// * "Invalid"
	// * "Internal"
	// * "Required"
	// * "NotAllowed"
	// * "TooLarge"
	// * "InvalidExtensionService"
	// * "InvalidGCPWasmPlugin"
	// * "InvalidCELExpression"
	// * "NoMatchingTarget"
	// * "Conflicted"
	ExtensionConditionAccepted ExtensionConditionType = "Accepted"

	// ExtensionReasonAccepted reason is used with the "Accepted" condition
	// when the extension or WasmPlugin has been accepted and the ExtensionConditionAccepted is true.
	ExtensionReasonAccepted ExtensionConditionReason = "Accepted"

	// ExtensionReasonInvalid is used with the "Accepted" condition
	// when the extension or WasmPlugin is syntactically or semantically invalid and
	// the ExtensionConditionAccepted is false.
	ExtensionReasonInvalid ExtensionConditionReason = "Invalid"

	// ExtensionReasonInternal is used with the "Accepted" condition
	// when the server error happened when processing the request
	// and the ExtensionConditionAccepted is false.
	ExtensionReasonInternal ExtensionConditionReason = "Internal"

	// ExtensionReasonRequired is used with the "Accepted" condition
	// when the usage of the field or value is required and the ExtensionConditionAccepted is false.
	ExtensionReasonRequired ExtensionConditionReason = "Required"

	// ExtensionReasonNotAllowed is used with the "Accepted" condition
	// when the usage of the field or value is not allowed and the ExtensionConditionAccepted is false.
	ExtensionReasonNotAllowed ExtensionConditionReason = "NotAllowed"

	// ExtensionReasonTooLarge is used with the "Accepted" condition
	// when the size of the object is too large and the ExtensionConditionAccepted is false.
	ExtensionReasonTooLarge ExtensionConditionReason = "TooLarge"

	// ExtensionReasonInvalidExtensionService is used with the "Accepted" condition
	// when the extension service is invalid and the ExtensionConditionAccepted is false.
	ExtensionReasonInvalidExtensionService ExtensionConditionReason = "InvalidExtensionService"

	// ExtensionReasonInvalidGCPWasmPlugin is used with the "Accepted" condition
	// when the GCPWasmPlugin is invalid and the ExtensionConditionAccepted is false.
	ExtensionReasonInvalidGCPWasmPlugin ExtensionConditionReason = "InvalidGCPWasmPlugin"

	// ExtensionReasonInvalidCELExpression is used with the "Accepted" condition
	// when the CEL expression is invalid and the ExtensionConditionAccepted is false.
	ExtensionReasonInvalidCELExpression ExtensionConditionReason = "InvalidCELExpression"

	// ExtensionReasonNoMatchingTarget is used with the "Accepted" condition
	// when there are no matching TargetRef and the ExtensionConditionAccepted is false.
	ExtensionReasonNoMatchingTarget ExtensionConditionReason = "NoMatchingTarget"

	// ExtensionReasonConflicted is used with the "Accepted" condition
	// when the extension is conflicted with another extension and
	// the ExtensionConditionAccepted is false.
	ExtensionReasonConflicted ExtensionConditionReason = "Conflicted"

	// ExtensionConditionResolvedRefs indicates whether the controller was able
	// to resolve all the backendRefs.
	//
	// Possible reasons for this condition to be true are:
	//
	// * "ResolvedRefs"
	//
	// Possible reasons for this condition to be False are:
	//
	// * "ExtensionServiceNotFound"
	// * "GCPWasmPluginNotFound"
	ExtensionConditionResolvedRefs ExtensionConditionType = "ResolvedRefs"

	// ExtensionReasonResolvedRefs is used with the "ResolvedRefs" condition
	// when the controller was able to resolve all the backendRefs.
	ExtensionReasonResolvedRefs ExtensionConditionReason = "ResolvedRefs"

	// ExtensionReasonExtensionServiceNotFound is used with the "ResolvedRefs" condition
	// when the extension service is not found and the ExtensionConditionResolvedRefs is false.
	ExtensionReasonExtensionServiceNotFound ExtensionConditionReason = "ExtensionServiceNotFound"

	// ExtensionReasonGCPWasmPluginNotFound is used with the "ResolvedRefs" condition
	// when the GCPWasmPlugin is not found and the ExtensionConditionResolvedRefs is false.
	ExtensionReasonGCPWasmPluginNotFound ExtensionConditionReason = "GCPWasmPluginNotFound"
)

// ExtensionServiceReference defines a reference to the Service
// within the namespace of the referrer.
// +kubebuilder:validation:XValidation:message="Group must be empty if kind is Service",rule="self.kind == 'Service' ? size(self.group) == 0 : true"
// +kubebuilder:validation:XValidation:message="Group must be set to `net.gke.io` if kind is ServiceImport",rule="self.kind == 'ServiceImport' ? self.group == 'net.gke.io' : true"
// +kubebuilder:validation:XValidation:message="Group must be set to `networking.gke.io` if kind is GCPWasmPlugin",rule="self.kind == 'GCPWasmPlugin' ? self.group == 'networking.gke.io' : true"
// +kubebuilder:validation:XValidation:message="Group must be set to `apim.googleapis.com` if kind is ApigeeBackendService",rule="self.kind == 'ApigeeBackendService' ? self.group == 'apim.googleapis.com' : true"
// +kubebuilder:validation:XValidation:message="Port has to be set if kind is Service",rule="self.kind == 'Service' ? has(self.port) : true"
// +kubebuilder:validation:XValidation:message="Port has to be set if kind is ServiceImport",rule="self.kind == 'ServiceImport' ? has(self.port) : true"
// +kubebuilder:validation:XValidation:message="Port has to be empty if kind is GCPWasmPlugin",rule="self.kind == 'GCPWasmPlugin' ? !has(self.port) : true"
type ExtensionServiceReference struct {
	// Group is the group of the referent.
	//
	// +kubebuilder:default=""
	// +kubebuilder:validation:Enum="";net.gke.io;networking.gke.io;apim.googleapis.com
	Group v1.Group `json:"group"`

	// Kind is kind of the referent.
	//
	// +kubebuilder:default=Service
	// +kubebuilder:validation:Enum=Service;ServiceImport;GCPWasmPlugin;ApigeeBackendService
	Kind v1.Kind `json:"kind"`

	// Name is the name of the referent.
	Name v1.ObjectName `json:"name"`

	// Port is the port of the referent.
	Port PortNumber `json:"port,omitempty"`
}

// PortNumber defines a network port.
//
// +kubebuilder:validation:Minimum=1
// +kubebuilder:validation:Maximum=65535
type PortNumber int32

// ExtensionChain is single extension chain wrapper that contains the match conditions
// and extensions to execute.
type ExtensionChain struct {
	// Name is the name for this extension chain.
	// The name is logged as part of the HTTP request logs.
	// The name must conform with RFC-1034, is restricted to lower-cased letters,
	// numbers and hyphens, and can have a maximum length of 63 characters.
	// Additionally, the first character must be a letter and the last a letter or a number.
	//
	// +kubebuilder:validation:MinLength=1
	// +kubebuilder:validation:MaxLength=63
	// +kubebuilder:validation:Pattern=`^[a-z]([a-z0-9-]{0,61}[a-z0-9])?$`
	Name string `json:"name"`

	// MatchCondition is the condition under which this chain is invoked for a request.
	//
	// +optional
	MatchCondition MatchCondition `json:"matchCondition,omitempty"`

	// Extensions is a set of extensions to execute for the matching request.
	// Up to 3 Extensions can be defined for each ExtensionChain
	// for the GCPTrafficExtension.
	// GCPRoutingExtension and GCPEdgeExtension chains are limited to 1 Extension
	// per ExtensionChain.
	//
	// +kubebuilder:validation:MinItems=1
	// +kubebuilder:validation:MaxItems=3
	Extensions []Extension `json:"extensions"`
}

// MatchCondition contains the conditions under which the extension chain
// is invoked for a request.
// The resulting MatchCondition is limited to 512 characters.
type MatchCondition struct {
	// CELExpressions are expressions that are used to match requests for which
	// the extension chain is executed.
	// Limited to 10 CELExpressions.
	//
	// Expressions are ORed together.
	//
	// +optional
	// +kubebuilder:validation:MaxItems=10
	CELExpressions []CELExpression `json:"celExpressions,omitempty"`
}

// CELExpression contains the conditions under which the extension chain is invoked
// for a request.
//
// CELMatcher and BackendRef are ANDed together.
type CELExpression struct {
	// CELMatcher is a Common Expression Language (CEL) expression that is used
	// to match requests for which the extension chain is executed.
	//
	// For more information, see [CEL matcher language
	// reference](https://cloud.google.com/service-extensions/docs/cel-matcher-language-reference).
	//
	// +optional
	// +kubebuilder:validation:MaxLength=512
	// +kubebuilder:validation:Pattern=`^(( )|(request.headers)|(request.method)|(request.host)|(request.path)|(request.query)|(request.scheme)|(request.backend_service_num_endpoints)|(response.code)|(response.grpc_status)|(response.headers)|(source.address)|(source.port)|(connection.requested_server_name)|(connection.tls_version)|(connection.sha256_peer_certificate_digest)|(endsWith)|(startsWith)|(matches)|(contains)|(lowerAscii)|(upperAscii)|(int)|(==)|(!=)|(&&)|(>=)|(<=)|(>)|(<)|(\|\|)|(!)|(\+)|(\.)|(/)|('[a-zA-Z0-9\/\-._~%!$&'()*+,;=:\"\\]*')|("[a-zA-Z0-9\/\-._~%!$&'()*+,;=:\"\\]*")|(“[a-zA-Z0-9\/\-._~%!$&'()*+,;=:\"\\]*”)|(R"[a-zA-Z0-9\/\-._~%!$&'()*+,;=:\\]*")|[0-9]|(\()|(\))|(\[)|(\]))+$`
	CELMatcher string `json:"celMatcher,omitempty"`

	// BackendRefs are Kubernetes Services that are used to match requests
	// for which the extension chain is executed.
	//
	// +optional
	// +kubebuilder:validation:MaxItems=1
	// +kubebuilder:validation:XValidation:message="Only backendRefs of kind Service or ServiceImport are supported in CEL expression",rule="self.all(ref, ref.kind == 'Service' || ref.kind == 'ServiceImport')"
	BackendRefs []ExtensionServiceReference `json:"backendRefs,omitempty"`
}

// Extension is a single extension in the chain to execute for the matching request.
// +kubebuilder:validation:XValidation:message="timeout must be between 10-10000 milliseconds",rule="has(self.timeout) ? duration(self.timeout) >= duration('10ms') && duration(self.timeout) <= duration('10000ms') : true"
// +kubebuilder:validation:XValidation:message="Extensions with backendRef kind GCPWasmPlugin do not support timeout",rule="has(self.backendRef) && self.backendRef.kind == 'GCPWasmPlugin' ? !has(self.timeout) : true"
// +kubebuilder:validation:XValidation:message="authority must be set if backendRef kind is set to Service or ServiceImport",rule="has(self.backendRef) && (self.backendRef.kind == 'Service' || self.backendRef.kind == 'ServiceImport') ? has(self.authority) && size(self.authority) != 0 : true"
// +kubebuilder:validation:XValidation:message="authority must not be set if the backendRef kind is ApigeeBackendService",rule="has(self.backendRef) && self.backendRef.kind == 'ApigeeBackendService' ? !has(self.authority) : true"
// +kubebuilder:validation:XValidation:message="Extension with googleAPIServiceName do not support authority",rule="has(self.googleAPIServiceName) ? !has(self.authority) : true"
// +kubebuilder:validation:XValidation:message="Extensions with backendRef kind GCPWasmPlugin do not support authority",rule="has(self.backendRef) && self.backendRef.kind == 'GCPWasmPlugin' ? !has(self.authority) : true"
// +kubebuilder:validation:XValidation:message="Exactly one of backendRef or googleAPIServiceName should be set",rule="!(has(self.backendRef) && has(self.googleAPIServiceName)) && (has(self.backendRef) || has(self.googleAPIServiceName))"
// +kubebuilder:validation:XValidation:message="Extensions with backendRef kind GCPWasmPlugin support only RequestHeaders, RequestBody, ResponseHeaders and ResponseBody events",rule="has(self.backendRef) && self.backendRef.kind == 'GCPWasmPlugin' ? self.supportedEvents.all(e, e in ['RequestHeaders', 'RequestBody', 'ResponseHeaders', 'ResponseBody']) : true"
// +kubebuilder:validation:XValidation:message="Extension with backendRef kind GCPWasmPlugin do not support metadata",rule="has(self.backendRef) && self.backendRef.kind == 'GCPWasmPlugin' ? !has(self.metadata) : true"
// +kubebuilder:validation:XValidation:message="requestBodySendMode can be configured only for extensions using backendRef with kind Service, ApigeeBackendService or ServiceImport",rule="has(self.requestBodySendMode) ? has(self.backendRef) && (self.backendRef.kind == 'Service' || self.backendRef.kind == 'ApigeeBackendService' || self.backendRef.kind == 'ServiceImport') : true"
// +kubebuilder:validation:XValidation:message="responseBodySendMode can be configured only for extensions using backendRef with kind Service, ApigeeBackendService or ServiceImport",rule="has(self.responseBodySendMode) ? has(self.backendRef) && (self.backendRef.kind == 'Service' || self.backendRef.kind == 'ApigeeBackendService' || self.backendRef.kind == 'ServiceImport') : true"
// +kubebuilder:validation:XValidation:message="If requestBodySendMode is set to `Streamed`, then the `supportedEvents` list must contain `RequestBody` event",rule="has(self.requestBodySendMode) && self.requestBodySendMode == 'Streamed' ? self.supportedEvents.exists(e, e == 'RequestBody') : true"
// +kubebuilder:validation:XValidation:message="If responseBodySendMode is set to `Streamed`, then the `supportedEvents` list must contain `ResponseBody` event",rule="has(self.responseBodySendMode) && self.responseBodySendMode == 'Streamed' ? self.supportedEvents.exists(e, e == 'ResponseBody') : true"
// +kubebuilder:validation:XValidation:message="If requestBodySendMode is set to `FullDuplexStreamed`, then the `supportedEvents` list must contain at least both: `RequestBody` and `RequestTrailers“ events",rule="has(self.requestBodySendMode) && self.requestBodySendMode == 'FullDuplexStreamed' ? self.supportedEvents.exists(e, e == 'RequestBody') && self.supportedEvents.exists(e, e == 'RequestTrailers') : true"
// +kubebuilder:validation:XValidation:message="If responseBodySendMode is set to `FullDuplexStreamed`, then the `supportedEvents` list must contain at least both: `ResponseBody` and `ResponseTrailers“ events",rule="has(self.responseBodySendMode) && self.responseBodySendMode == 'FullDuplexStreamed' ? self.supportedEvents.exists(e, e == 'ResponseBody') && self.supportedEvents.exists(e, e == 'ResponseTrailers') : true"
// +kubebuilder:validation:XValidation:message="If observabilityMode is set to `TRUE`, then the `responseBodySendMode` and `requestBodySendMode` must be not set or set to `Streamed`",rule="has(self.observabilityMode) && self.observabilityMode ? (!has(self.responseBodySendMode) || self.responseBodySendMode == 'Streamed') && (!has(self.requestBodySendMode) || self.requestBodySendMode == 'Streamed') : true"
type Extension struct {
	// Name is the name for this chain.
	// The name is logged as part of the HTTP request logs.
	// The name must conform with RFC-1034, is restricted to lower-cased
	// letters, numbers and hyphens, and can have a maximum length of 63
	// characters. Additionally, the first character must be a letter and the
	// last a letter or a number.
	//
	// +kubebuilder:validation:MinLength=1
	// +kubebuilder:validation:MaxLength=63
	// +kubebuilder:validation:Pattern=`^[a-z]([a-z0-9-]{0,61}[a-z0-9])?$`
	Name string `json:"name"`

	// BackendRef identifies an API object that runs the extension.
	// Exactly one of BackendRef or GoogleAPIServiceName should be set.
	// Valid Kinds are:
	// - "Service"
	// - "ServiceImport"
	// - "GCPWasmPlugin"
	//
	// +optional
	BackendRef *ExtensionServiceReference `json:"backendRef,omitempty"`

	// LINT.IfChange
	// We have to keep the validation in sync with the DEP CLH.

	// GoogleAPIServiceName is the name of the Google API service that runs the extension.
	// The name must be in one of the following formats:
	// - <serviceName>.<region>.rep.googleapis.com in regional case.
	// Exactly one of BackendRef or GoogleAPIServiceName should be set.
	//
	// +optional
	// +kubebuilder:validation:MaxLength=100
	// +kubebuilder:validation:Pattern=`^[a-z0-9.-]+\.(googleapis|sandbox\.googleapis)\.com$`
	GoogleAPIServiceName string `json:"googleAPIServiceName,omitempty"`

	// LINT.ThenChange(//depot/google3/cloud/hosted/networkservices/dep/clh/consumerservice/config/startup.pi)

	// Authority is the `:authority` header in the gRPC request sent from Envoy
	// to the callout extension backend service.
	// Required for extensions with callout backend service.
	//
	// +optional
	// +kubebuilder:validation:MaxLength=1000
	// +kubebuilder:validation:Pattern=`^[A-Za-z0-9-_:%\.\[\]]*$`
	Authority string `json:"authority,omitempty"`

	// SupportedEvents is a set of events during request or response
	// processing for which this extension is called.
	//
	// This field is required for the `GCPTrafficExtension` resource.
	// This field is optional for the `GCPRoutingExtension` resource.
	// This field is required for the `GCPEdgeExtension` resource and must only
	// contain `RequestHeaders`.
	//
	// If requestBodySendMode is set for the `GCPRoutingExtension` resource,
	// then the `supportedEvents` list can only contain `RequestHeaders` events.
	// If requestBodySendMode is set to `FullDuplexStreamed` for the `GCPRoutingExtension` resource,
	// then the `supportedEvents` list can only contain `RequestHeaders`, `RequestBody`
	// and `RequestTrailers` events.
	//
	// If unspecified, `RequestHeaders` event is assumed as supported.
	// Limited to 6 events.
	//
	// +optional
	// +kubebuilder:validation:MaxItems=6
	SupportedEvents []EventType `json:"supportedEvents"`

	// Timeout specifies the timeout for each individual message on the stream.
	// The timeout must be between 10-10000 milliseconds.
	// If omitted, the default timeout is 1000 milliseconds.
	//
	// +optional
	Timeout *v1.Duration `json:"timeout,omitempty"`

	// FailOpen determines how the proxy behaves if the call to the extension
	// fails or times out.
	//
	// When set to `TRUE`, request or response processing continues without
	// error. Any subsequent extensions in the extension chain are also
	// executed. When set to `FALSE` or the default setting of `FALSE` is used,
	// one of the following happens:
	//
	// * If response headers have not been delivered to the downstream client,
	// a generic 500 error is returned to the client. The error response can be
	// tailored by configuring a custom error response in the load balancer.
	//
	// * If response headers have been delivered, then the HTTP stream to the
	// downstream client is reset.
	//
	// +optional
	FailOpen bool `json:"failOpen,omitempty"`

	// ForwardHeaders is a list of the HTTP headers to forward to the extension
	// (from the client or backend). If omitted, all headers are sent.
	// Each element indicates the header name.
	//
	// +optional
	// +kubebuilder:validation:MaxItems=50
	ForwardHeaders []HTTPHeaderName `json:"forwardHeaders,omitempty"`

	// Metadata provided here is included as part of the
	// `metadata_context` (of type `google.protobuf.Struct`) in the
	// `ProcessingRequest` message sent to the extension
	// server. The metadata is available under the namespace
	// `com.google.<extension_type>.<resource_name>.<extension_chain_name>.<extension_name>`.
	// For example:
	// `com.google.lb_traffic_extension.lbtrafficextension1.chain1.ext1`.
	//
	// The following variables are supported in the metadata:
	//
	// `{forwarding_rule_id}` - substituted with the forwarding rule's fully
	//   qualified resource name.
	//
	// +optional
	// +kubebuilder:validation:MaxProperties=16
	// +kubebuilder:validation:XValidation:message="Metadata keys must only contain valid characters (matching ^([A-Za-z0-9\\/\\-._~%!$&'()*+,;=:\\s\\[\\]\\{\\}]{1,63})$) and must be up to 63 characters long.",rule="self.all(key, key.matches(r\"\"\"^([A-Za-z0-9\\/\\-._~%!$&'()*+,;=:\\s\\{\\}\\[\\]]{1,63})$\"\"\"))"
	Metadata map[MetadataKey]MetadataValue `json:"metadata,omitempty"`

	// RequestBodySendMode configures processing mode for request processing. If not specified
	// and RequestBody event is supported, the default value STREAMED is used.
	// Supported by both `GCPTrafficExtension` and `GCPRoutingExtension` resources
	// and only for extensions with `backendRef`.
	// When this field is set to `FullDuplexStreamed`, `supportedEvents`
	// must include both `RequestBody` and `RequestTrailers`.
	//
	// +optional
	RequestBodySendMode BodySendMode `json:"requestBodySendMode,omitempty"`

	// ResponseBodySendMode configures processing mode for response processing.
	// If not specified and RequestBody event is supported, the default value STREAMED is used.
	// Supported only by `GCPTrafficExtension` resource
	// and only for extensions with `backendRef`.
	// When this field is set to `FullDuplexStreamed`, `supportedEvents`
	// must include both `ResponseBody` and `ResponseTrailers`.
	//
	// +optional
	ResponseBodySendMode BodySendMode `json:"responseBodySendMode,omitempty"`

	// ObservabilityMode configures the observability mode for the extension.
	// This field is helpful when you want to try out the extension in async
	// log-only mode.
	// Supported by `GCPTrafficExtension` and `GCPRoutingExtension` resources
	// attached to gateways with regional gateway classes.
	// ObservabilityMode - set to true - is only supported with not set or `STREAMED`
	// `requestBodySendMode` and `responseBodySendMode`.
	//
	// +optional
	ObservabilityMode bool `json:"observabilityMode,omitempty"`
}

// EventType identifies the part of the request or response for which the extension is called.
// +kubebuilder:validation:Enum=RequestHeaders;RequestBody;ResponseHeaders;ResponseBody;RequestTrailers;ResponseTrailers
type EventType string

const (
	// EventTypeRequestHeaders can be used to call the extension
	// when the HTTP request headers arrive.
	EventTypeRequestHeaders EventType = "RequestHeaders"

	// EventTypeRequestBody can be used to call the extension
	// when the HTTP request body arrives.
	EventTypeRequestBody EventType = "RequestBody"

	// EventTypeResponseHeaders can be used to call the extension
	// when the HTTP response headers arrive.
	EventTypeResponseHeaders EventType = "ResponseHeaders"

	// EventTypeResponseBody can be used to call the extension
	// when the HTTP response body arrives.
	EventTypeResponseBody EventType = "ResponseBody"

	// EventTypeRequestTrailers can be used to call the extension
	// when the HTTP request trailers arrive.
	EventTypeRequestTrailers EventType = "RequestTrailers"

	// EventTypeResponseTrailers can be used to call the extension
	// when the HTTP response trailers arrive.
	EventTypeResponseTrailers EventType = "ResponseTrailers"
)

// BodySendMode defines the send mode for the body processing.
// +kubebuilder:validation:Enum=Streamed;FullDuplexStreamed
type BodySendMode string

const (
	// BodySendModeStreamed can be used to indicate that calls are executed in the streaming mode.
	//
	// This is the default mode.
	BodySendModeStreamed BodySendMode = "Streamed"

	// BodySendModeFullDuplexStreamed can be used to indicate that calls are executed
	// in the full duplex streaming mode. Subsequent chunks will be sent
	// for processing without waiting for the response for the previous event.
	BodySendModeFullDuplexStreamed BodySendMode = "FullDuplexStreamed"
)

// HTTPHeaderName is the name of the HTTP header.
type HTTPHeaderName v1.HTTPHeaderName

// MetadataKey is the key of an metadata in GCP Extensions.
// The CEL validation may be removed from this field in the future because
// in conjunction with the Metadata map object it is ignored and performed at the map level.
//
// +kubebuilder:validation:MinLength=1
// +kubebuilder:validation:MaxLength=63
// +kubebuilder:validation:Pattern=`^([A-Za-z0-9\/\-._~%!$&'()*+,;=:"\s\[\]\{\}])*$`
type MetadataKey string

// MetadataValue is the value of an metadata in GCP Extensions.
//
// +kubebuilder:validation:MinLength=1
// +kubebuilder:validation:MaxLength=1023
// +kubebuilder:validation:Pattern=`^([A-Za-z0-9\/\-._~%!$&'()*+,;=:"\s\[\]\{\}])*$`
type MetadataValue string

// WireFormat is the format of communication supported by the callout extension.
//
// +kubebuilder:validation:Enum=ExtProcGRPC
type WireFormat string

const (
	// WireFormatExtProcGRPC can be used to indicate that the callout backend service
	// uses ExtProc GRPC API.
	WireFormatExtProcGRPC WireFormat = "ExtProcGRPC"
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
