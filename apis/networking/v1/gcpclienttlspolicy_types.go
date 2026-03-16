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

// TLSMode specifies the mode of the TLS policy.
// +kubebuilder:validation:Enum=MutualTLS;Disable
type TLSMode string

const (
	// MutualTLS is the default mode for TLS.
	MutualTLS TLSMode = "MutualTLS"
	// Disable disables TLS.
	Disable TLSMode = "Disable"
)

// SubjectAltName represents Subject Alternative Name.
type SubjectAltName struct {
	// URI contains Subject Alternative Name specified in a full URI format.
	// It MUST include both a scheme (e.g., "http" or "ftp") and a scheme-specific-part.
	// Common values include SPIFFE IDs like "spiffe://mycluster.example.com/ns/myns/sa/svc1sa".
	//
	// Support: Core
	URI v1.AbsoluteURI `json:"uri"`
}

// GCPClientTLSPolicySpec defines the desired state of GCPClientTLSPolicy.
// +kubebuilder:validation:XValidation:rule="!(self.tlsMode == 'Disable' && has(self.subjectAltNames) && size(self.subjectAltNames) > 0)",message="SubjectAltNames can only be set when TLSMode is not Disable (i.e., SubjectAltNames must be empty if TLSMode is Disable)"
// +kubebuilder:validation:XValidation:rule="self.targetRefs.all(t, !(t.kind == 'Namespace' && has(t.sectionName) && t.sectionName != \"\"))",message="SectionName cannot be set when targeting a Namespace"
// +kubebuilder:validation:XValidation:rule="self.targetRefs.all(t, t.kind == 'Service' || t.kind == 'Namespace')",message="TargetRefs can only target a Service or Namespace"
type GCPClientTLSPolicySpec struct {
	// SAN is the Subject Alternative Name (SAN) to use for this connection.
	// If unset, the SPIFFE prefix of the K8s service's namespace is set by default.
	// E.g. "spiffe://<fleet-project-id>.svc.id.goog/ns/<k8s-service-namespace>/sa/*".
	// If explicitly configured, then only exact string match semantics are supported.
	// If it’s not attached to a K8s service, then no SAN will be set by default.
	// +kubebuilder:validation:MaxItems=10
	// +optional
	SubjectAltNames []SubjectAltName `json:"subjectAltNames,omitempty"`

	// Defaults to MutualTLS if not set.
	// This is useful when all the services in a namespace are opted into MTLS.
	// But then certain services (and ports) need to be opted out.
	// +kubebuilder:default=MutualTLS
	// +optional
	TLSMode TLSMode `json:"tlsMode,omitempty"`

	// TargetRef is a reference to the object that this policy applies to.
	// Currently allowed targets are K8s service, service port (configured as a section name), and K8s namespace.
	//
	// When multiple ClientTLSPolicies apply to a specific outbound
	// connection the following precedence order is used to determine which
	// policy's settings are applied:
	//
	// 1. Service Port Specific Policy: A policy targeting a Service
	//    with a specific sectionName (port).
	// 2. Service Specific Policy: A policy targeting a Service
	//    without a sectionName.
	// 3. Namespace-wide Policy: A policy targeting all the Services
	//    in the Namespace. Must be the same namespace in which the policy
	//    is defined (i.e., cross namespace references are not allowed).
	//
	// If multiple policies exist at the same level of specificity,
	// the behavior is based on the documented conflict resolution rules
	// (i.e., the oldest policy wins, or it's an error condition).
	//
	// It is recommended that users avoid creating multiple conflicting
	// policies at the same specificity level.
	//
	// The most specific policy found will override the less specific ones.
	// This means the settings in the most specific policy will be used, and
	// settings from less specific policies will be ignored for that connection.
	// +kubebuilder:validation:MinItems=1
	// +kubebuilder:validation:MaxItems=10
	TargetRefs []v1.LocalPolicyTargetReferenceWithSectionName `json:"targetRefs"`
}

// GCPClientTLSPolicy is the Schema for the gcpclienttlspolicies API.
// +genclient
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Namespaced
type GCPClientTLSPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   GCPClientTLSPolicySpec `json:"spec,omitempty"`
	Status v1.PolicyStatus        `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// GCPClientTLSPolicyList contains a list of GCPClientTLSPolicy.
type GCPClientTLSPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GCPClientTLSPolicy `json:"items"`
}
