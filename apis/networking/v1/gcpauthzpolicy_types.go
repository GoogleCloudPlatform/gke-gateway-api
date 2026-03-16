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

// GCPAuthzPolicyAction specifies the ACTION of the authorization policy.
// +kubebuilder:validation:Enum=ALLOW;DENY;CUSTOM;DENY_BY_DEFAULT
type GCPAuthzPolicyAction string

const (
	// Allow a request only if it matches the rules. This is the default type.
	Allow GCPAuthzPolicyAction = "ALLOW"
	// Deny a request if it matches any of the rules.
	Deny GCPAuthzPolicyAction = "DENY"
	// Custom action allows an extension to handle the user request if
	// the matching rules evaluate to true.
	Custom GCPAuthzPolicyAction = "CUSTOM"
	// DenyByDefault denies all requests in the ns and gets overridden by ALLOW rules.
	DenyByDefault GCPAuthzPolicyAction = "DENY_BY_DEFAULT"
)

// EnforcementLevel specifies the type of the authorization policy.
// +kubebuilder:validation:Enum=L7;L4
type EnforcementLevel string

const (
	// L7 specifies an application-level (Layer 7) policy.
	L7 EnforcementLevel = "L7"
	// L4 specifies a network-level (Layer 4) policy.
	L4 EnforcementLevel = "L4"
)

// StringMatchCriteriaType specifies the type of the string match criteria.
// +kubebuilder:validation:Enum=Exact;Prefix;Suffix;Contains
type StringMatchCriteriaType string

const (
	// StringExact is the default type. It matches the exact string.
	StringExact StringMatchCriteriaType = "Exact"
	// StringPrefix matches the prefix of the string.
	StringPrefix StringMatchCriteriaType = "Prefix"
	// StringSuffix matches the suffix of the string.
	StringSuffix StringMatchCriteriaType = "Suffix"
	// StringContains matches the string that contains the value.
	StringContains StringMatchCriteriaType = "Contains"
)

// HTTPMethod describes how to select a HTTP route by matching the HTTP
// method as defined by
// [RFC 7231](https://datatracker.ietf.org/doc/html/rfc7231#section-4) and
// [RFC 5789](https://datatracker.ietf.org/doc/html/rfc5789#section-2).
// The value is expected in upper case.
//
// Note that values may be added to this enum, implementations
// must ensure that unknown values will not cause a crash.
//
// Unknown values here must result in the implementation setting the
// Accepted Condition for the Route to `status: False`, with a
// Reason of `UnsupportedValue`.
//
// +kubebuilder:validation:Enum=GET;HEAD;POST;PUT;DELETE;CONNECT;OPTIONS;TRACE;PATCH
type HTTPMethod string

const (
	// HTTPMethodGet refers to the HTTP GET method.
	HTTPMethodGet HTTPMethod = "GET"
	// HTTPMethodHead refers to the HTTP HEAD method.
	HTTPMethodHead HTTPMethod = "HEAD"
	// HTTPMethodPost refers to the HTTP POST method.
	HTTPMethodPost HTTPMethod = "POST"
	// HTTPMethodPut refers to the HTTP PUT method.
	HTTPMethodPut HTTPMethod = "PUT"
	// HTTPMethodDelete refers to the HTTP DELETE method.
	HTTPMethodDelete HTTPMethod = "DELETE"
	// HTTPMethodConnect refers to the HTTP CONNECT method.
	HTTPMethodConnect HTTPMethod = "CONNECT"
	// HTTPMethodOptions refers to the HTTP OPTIONS method.
	HTTPMethodOptions HTTPMethod = "OPTIONS"
	// HTTPMethodTrace refers to the HTTP TRACE method.
	HTTPMethodTrace HTTPMethod = "TRACE"
	// HTTPMethodPatch refers to the HTTP PATCH method.
	HTTPMethodPatch HTTPMethod = "PATCH"
)

// +genclient
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:categories=gateway-api
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:name="Age",type=date,JSONPath=`.metadata.creationTimestamp`

// GCPAuthzPolicy is the CRD for Authorization Policy.
// This policy enables access control on workloads.
type GCPAuthzPolicy struct {
	metav1.TypeMeta `json:",inline"`
	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// Spec defines the implementation of this definition.
	// +required
	Spec GCPAuthzPolicySpec `json:"spec,omitempty"`

	// Status defines the current state of GCPAuthzPolicy.
	// +optional
	Status v1.PolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// GCPAuthzPolicyList contains a list of GCPAuthzPolicy.
type GCPAuthzPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GCPAuthzPolicy `json:"items"`
}

// GCPAuthzPolicySpec is the spec for Authorization Policy.
// +kubebuilder:validation:XValidation:message="At least one rule is required when the action is not CUSTOM or DENY_BY_DEFAULT",rule="(self.action == 'ALLOW' || self.action == 'DENY') ? size(self.rules) > 0 : true"
// +kubebuilder:validation:XValidation:message="CustomProviders are required when the action is CUSTOM",rule="!(self.action == 'CUSTOM' && !has(self.customProviders)) && !(self.action != 'CUSTOM' && has(self.customProviders))"
// +kubebuilder:validation:XValidation:message="When Action is DENY_BY_DEFAULT, Rules and CustomProviders must be empty",rule="self.action != 'DENY_BY_DEFAULT' || (!has(self.rules) && !has(self.customProviders))"
// +kubebuilder:validation:XValidation:message="When Action is CUSTOM, EnforcementLevel must be L7",rule="self.action != 'CUSTOM' || self.enforcementLevel == 'L7'"
// +kubebuilder:validation:XValidation:message="When EnforcementLevel is L4, only principals are allowed in sources and notSources, and no operations are allowed",rule="self.enforcementLevel != 'L4' || self.rules.all(r, (!has(r.from) || ((!has(r.from.sources) || r.from.sources.all(s, has(s.principals) && s.principals.size() > 0 && !has(s.resources))) && (!has(r.from.notSources) || r.from.notSources.all(s, has(s.principals) && s.principals.size() > 0 && !has(s.resources))))) && !has(r.to))"
// +kubebuilder:validation:XValidation:message="When Resources is set in GCPAuthzPolicySource, at least one TargetRef must have Kind=Gateway",rule="(!has(self.rules) || !self.rules.exists(r, has(r.from) && ((has(r.from.sources) && r.from.sources.exists(s, has(s.resources))) || (has(r.from.notSources) && r.from.notSources.exists(s, has(s.resources)))))) || self.targetRefs.exists(t, t.kind == 'Gateway')"
// +kubebuilder:validation:XValidation:message="Only one TargetRef of kind=Pod is allowed",rule="self.targetRefs.filter(t, t.kind == 'Pod').size() <= 1"
// +kubebuilder:validation:XValidation:message="principalSelector must be CLIENT_CERT_URI_SAN when TargetRef kind is Pod.",rule="!self.targetRefs.exists(t, t.kind == 'Pod') || !has(self.rules) || self.rules.all(r, !has(r.from) || ((!has(r.from.sources) || r.from.sources.all(s, !has(s.principals) || s.principals.all(p, !has(p.principalSelector) || p.principalSelector == 'CLIENT_CERT_URI_SAN'))) && (!has(r.from.notSources) || r.from.notSources.all(s, !has(s.principals) || s.principals.all(p, !has(p.principalSelector) || p.principalSelector == 'CLIENT_CERT_URI_SAN')))))"
type GCPAuthzPolicySpec struct {
	// Type specifies the type of the authorization policy.
	// +required
	EnforcementLevel EnforcementLevel `json:"enforcementLevel,omitempty"`
	// A list of rules to match the request.
	// +kubebuilder:validation:MaxItems=10
	// +optional
	Rules []GCPAuthPolicyRule `json:"rules,omitempty"`
	// The action to take if the request is matched with the rules. Default is ALLOW if not specified.
	// +kubebuilder:validation:Enum=ALLOW;DENY;CUSTOM;DENY_BY_DEFAULT
	// +kubebuilder:default=ALLOW
	// +optional
	Action *GCPAuthzPolicyAction `json:"action,omitempty"`
	// CustomProviders defines the extension providers for authorization policy.
	// +optional
	CustomProviders *GCPAuthzPolicyCustomProviders `json:"customProviders,omitempty"`
	// TargetRefs identifies a list of API objects to apply policy to.
	// Limited to 10 TargetRef, can not be empty.
	// +kubebuilder:validation:MinItems=1
	// +kubebuilder:validation:MaxItems=10
	// +optional
	TargetRefs []LocalObjectReference `json:"targetRefs,omitempty"`
}

// WorkloadSelector defines the selector for the workloads to which the policy is applied.
// MatchExpressions are not supported. Only MatchLabels is supported at the moment because AuthzPolicy
// already has complex merging scenarios across pod labels and namespace.
type WorkloadSelector struct {
	// MatchLabels is a map of {key, value} pairs which defines the pod labels to match.
	// +optional
	MatchLabels map[string]string `json:"matchLabels,omitempty"`
}

// LocalObjectReference identifies an API object within the namespace of the
// referrer.
// The API object must be valid in the cluster; the Group and Kind must
// be registered in the cluster for this reference to be valid.
//
// References to objects with invalid Group and Kind are not valid, and must
// be rejected by the implementation, with appropriate Conditions set
// on the containing object.
// +kubebuilder:validation:XValidation:message="Kind must be either 'Gateway' or 'Pod'",rule="self.kind == 'Gateway' || self.kind == 'Pod'"
// +kubebuilder:validation:XValidation:message="If Kind is Gateway, Name must be set and Selector must be empty.",rule="self.kind != 'Gateway' || (has(self.name) && !has(self.selector))"
// +kubebuilder:validation:XValidation:message="If Kind is Pod, Name must be empty and Selector must be set.",rule="self.kind != 'Pod' || (!has(self.name) && has(self.selector))"
// +kubebuilder:validation:XValidation:message="If Kind is Gateway, Group must be gateway.networking.k8s.io.",rule="self.kind != 'Gateway' || self.group == 'gateway.networking.k8s.io'"
type LocalObjectReference struct {
	// Group is the group of the referent. For example, "gateway.networking.k8s.io".
	// When unspecified or empty string, core API group is inferred.
	// +required
	Group v1.Group `json:"group"`

	// Kind is kind of the referent. For example "Gateway" or "Pod".
	// +required
	Kind v1.Kind `json:"kind,omitempty"`

	// Name is the name of the referent.
	// +optional
	Name v1.ObjectName `json:"name,omitempty"`

	// Selector is the label selector of target objects of the specified kind.
	// +optional
	Selector *WorkloadSelector `json:"selector,omitempty"`
}

// GCPAuthPolicyRule matches requests from a list of sources that perform a list of operations subject to a
// list of conditions. A match occurs when at least one source, one operation and all conditions
// matches the request. An empty rule is always matched.
type GCPAuthPolicyRule struct {
	// From specifies the source of a request.
	// If not set, any source is allowed.
	// +optional
	From *GCPAuthzPolicyFrom `json:"from,omitempty"`
	// To specifies the operations of a request.
	// If not set, any operation is allowed.
	// +optional
	To *GCPAuthzPolicyTo `json:"to,omitempty"`
	// when specifies a list of additional conditions of a request.
	// If not set, any condition is allowed.
	// It is supported via CEL expression.
	// +kubebuilder:validation:MinLength=1
	// +kubebuilder:validation:MaxLength=512
	// +optional
	When *string `json:"when,omitempty"`
}

// GCPAuthzPolicyFrom includes a list of sources.
type GCPAuthzPolicyFrom struct {
	// Sources specifies the source of a request.
	// +kubebuilder:validation:MaxItems=5
	// +optional
	Sources []GCPAuthzPolicySource `json:"sources,omitempty"`
	// NotSources specifies the not sources of a request.
	// +kubebuilder:validation:MaxItems=5
	// +optional
	NotSources []GCPAuthzPolicySource `json:"notSources,omitempty"`
}

// GCPAuthzPolicyTo includes a list of operations.
type GCPAuthzPolicyTo struct {
	// Operation specifies the operation of a request.
	// +kubebuilder:validation:MaxItems=5
	// +optional
	Operations []GCPAuthzPolicyOperation `json:"operations,omitempty"`
	// NotOperations specifies the not operation of a request.
	// +kubebuilder:validation:MaxItems=5
	// +optional
	NotOperations []GCPAuthzPolicyOperation `json:"notOperations,omitempty"`
}

// PrincipalSelector specifies subsets of authentication attributes, such as
// URI SANs in the validated client's certificate.
// +kubebuilder:validation:Enum=CLIENT_CERT_URI_SAN;CLIENT_CERT_DNS_NAME_SAN;CLIENT_CERT_COMMON_NAME
type PrincipalSelector string

const (
	// ClientCertURISAN means the principal rule is matched against a list of URI SANs in the
	// validated client's certificate. A match happens when there is any
	// exact URI SAN value match. This is the default principal selector.
	ClientCertURISAN PrincipalSelector = "CLIENT_CERT_URI_SAN"
	// ClientCertDNSNameSAN means the principal rule is matched against a list of DNS Name SANs in the
	// validated client's certificate. A match happens when there is any
	// exact DNS Name SAN value match.
	// This is only applicable for Application Load Balancers
	// except for classic Global External Application load balancer.
	// CLIENT_CERT_DNS_NAME_SAN is not supported for INTERNAL_SELF_MANAGED
	// load balancing scheme.
	ClientCertDNSNameSAN PrincipalSelector = "CLIENT_CERT_DNS_NAME_SAN"
	// ClientCertCommonName means the principal rule is matched against the common name in the client's
	// certificate. Authorization against multiple common names in the
	// client certificate is not supported. Requests with multiple common
	// names in the client certificate will be rejected if
	// CLIENT_CERT_COMMON_NAME is set as the principal selector. A match
	// happens when there is an exact common name value match.
	// This is only applicable for Application Load Balancers
	// except for global external Application Load Balancer and
	// classic Application Load Balancer.
	// CLIENT_CERT_COMMON_NAME is not supported for INTERNAL_SELF_MANAGED
	// load balancing scheme.
	ClientCertCommonName PrincipalSelector = "CLIENT_CERT_COMMON_NAME"
)

// Principal describes the properties of a principal to be matched against.
type Principal struct {
	// PrincipalSelector is an enum to decide what principal value the principal rule
	// will match against. If not specified, the PrincipalSelector is
	// CLIENT_CERT_URI_SAN.
	// One of CLIENT_CERT_URI_SAN, CLIENT_CERT_DNS_NAME_SAN, or CLIENT_CERT_COMMON_NAME.
	// +optional
	PrincipalSelector *PrincipalSelector `json:"principalSelector,omitempty"`
	// Principal is a non-empty string whose value is matched against the
	// principal value based on the principal_selector. Only exact match can
	// be applied for CLIENT_CERT_URI_SAN, CLIENT_CERT_DNS_NAME_SAN,
	// CLIENT_CERT_COMMON_NAME selectors.
	// +kubebuilder:validation:XValidation:message="Only Exact is allowed for StringMatchCriteria Type in Principals",rule="self.type == 'Exact'"
	Principal StringMatchCriteria `json:"principal"`
}

// GCPAuthzPolicySource specifies the source identities of a request.
// Fields in the AuthzPolicySource are ANDed together.
type GCPAuthzPolicySource struct {
	// Principals includes the list of peer identities derived from the peer certificate.
	// The peer identity is in the format of `"<TRUST_DOMAIN>/ns/<NAMESPACE>/sa/<SERVICE_ACCOUNT>"`,
	// for example, `"{projectid}.svc.id.goog/ns/foo/sa/productpage"`.
	// This field requires mTLS enabled.
	// +kubebuilder:validation:MaxItems=10
	// +optional
	Principals []Principal `json:"principals,omitempty"`
	// Resources describes the properties of a client VM
	// resource accessing the internal application load balancers.
	// +kubebuilder:validation:MaxItems=10
	// +optional
	Resources []Resource `json:"resources,omitempty"`
}

// GCPAuthzPolicyOperation is the spec for the operation of the rule.
// Fields in the AuthzPolicyOperation are ANDed together.
type GCPAuthzPolicyOperation struct {
	// Headers defines a set of headers to match against for a given request.
	// All headers in the set must match. Limited to 10 matches.
	// +kubebuilder:validation:MaxItems=10
	// +optional
	Headers []HTTPHeaderMatch `json:"headers,omitempty"`
	// Hosts is a list of HTTP Hosts to match against.
	// Limited to 10 matches.
	// +kubebuilder:validation:MaxItems=10
	// +optional
	Hosts []StringMatchCriteria `json:"hosts,omitempty"`
	// Methods is a list of HTTP methods to match against.
	// Each entry must be a valid HTTP method name (GET, PUT, POST, HEAD, PATCH, DELETE, OPTIONS, CONNECT, TRACE).
	// Limited to 9 matches.
	// +kubebuilder:validation:MaxItems=9
	// +optional
	Methods []HTTPMethod `json:"methods,omitempty"`
	// Paths is a list paths to match against.
	// Limited to 10 matches.
	// +kubebuilder:validation:MaxItems=10
	// +optional
	Paths []StringMatchCriteria `json:"paths,omitempty"`
}

// HTTPHeaderMatch builds the header match criteria.
type HTTPHeaderMatch struct {
	// Type specifies how to match against the value of the header.
	// +required
	Type StringMatchCriteriaType `json:"type,omitempty"`
	// Name is the name of the header.
	// +kubebuilder:validation:MinLength=1
	// +kubebuilder:validation:MaxLength=256
	// +kubebuilder:validation:Pattern=`^[A-Za-z0-9!#$%&'*+\-.^_\x60|~]+$`
	// +required
	Name string `json:"name,omitempty"`
	// Value is the value of the header.
	// +kubebuilder:validation:MaxLength=256
	// +required
	Value string `json:"value,omitempty"`
	// IgnoreCase is true then the matching should be case insensitive.
	// +optional
	IgnoreCase bool `json:"ignoreCase,omitempty"`
}

// StringMatchCriteria defines the match criteria for a string.
type StringMatchCriteria struct {
	// Type is the type of the string match criteria.
	// +required
	Type StringMatchCriteriaType `json:"type,omitempty"`
	// Value is the match.
	// +required
	Value string `json:"value,omitempty"`
	// IgnoreCase is true then the matching should be case insensitive.
	// +optional
	IgnoreCase bool `json:"ignoreCase,omitempty"`
}

// GCPAuthzPolicyCustomProviders defines the custom providers for authorization policy.
type GCPAuthzPolicyCustomProviders struct {
	// ExtensionRefs identifies a list of Authz Extensions.
	// Limited to 2 ExtensionRefs.
	// +kubebuilder:validation:MinItems=1
	// +kubebuilder:validation:MaxItems=2
	// +required
	ExtensionRefs []v1.LocalObjectReference `json:"extensionRefs,omitempty"`
}

// Resource defines the Andromeda credentials.
// It is only applicable internal L7 LBs.
type Resource struct {
	// TagValueIDSet is a list of resource tag value permanent IDs to match against
	// the resource manager tag values associated with the source VM of a request.
	// All IDs must match.
	// +kubebuilder:validation:MinItems=1
	// +kubebuilder:validation:MaxItems=10
	// +optional
	TagValueIDSet []int64 `json:"tagValueIdSet,omitempty"`

	// IAMServiceAccount to match against the GCP IAM service account
	// associated with the source VM of a request.
	// +optional
	IAMServiceAccount *StringMatchCriteria `json:"iamServiceAccount,omitempty"`
}
