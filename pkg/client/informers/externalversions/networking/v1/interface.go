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

// Code generated by informer-gen. DO NOT EDIT.

package v1

import (
	internalinterfaces "github.com/GoogleCloudPlatform/gke-gateway-api/pkg/client/informers/externalversions/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// GCPBackendPolicies returns a GCPBackendPolicyInformer.
	GCPBackendPolicies() GCPBackendPolicyInformer
	// GCPGatewayPolicies returns a GCPGatewayPolicyInformer.
	GCPGatewayPolicies() GCPGatewayPolicyInformer
	// GCPSessionAffinityFilters returns a GCPSessionAffinityFilterInformer.
	GCPSessionAffinityFilters() GCPSessionAffinityFilterInformer
	// GCPSessionAffinityPolicies returns a GCPSessionAffinityPolicyInformer.
	GCPSessionAffinityPolicies() GCPSessionAffinityPolicyInformer
	// HealthCheckPolicies returns a HealthCheckPolicyInformer.
	HealthCheckPolicies() HealthCheckPolicyInformer
}

type version struct {
	factory          internalinterfaces.SharedInformerFactory
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory, namespace string, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &version{factory: f, namespace: namespace, tweakListOptions: tweakListOptions}
}

// GCPBackendPolicies returns a GCPBackendPolicyInformer.
func (v *version) GCPBackendPolicies() GCPBackendPolicyInformer {
	return &gCPBackendPolicyInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// GCPGatewayPolicies returns a GCPGatewayPolicyInformer.
func (v *version) GCPGatewayPolicies() GCPGatewayPolicyInformer {
	return &gCPGatewayPolicyInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// GCPSessionAffinityFilters returns a GCPSessionAffinityFilterInformer.
func (v *version) GCPSessionAffinityFilters() GCPSessionAffinityFilterInformer {
	return &gCPSessionAffinityFilterInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// GCPSessionAffinityPolicies returns a GCPSessionAffinityPolicyInformer.
func (v *version) GCPSessionAffinityPolicies() GCPSessionAffinityPolicyInformer {
	return &gCPSessionAffinityPolicyInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// HealthCheckPolicies returns a HealthCheckPolicyInformer.
func (v *version) HealthCheckPolicies() HealthCheckPolicyInformer {
	return &healthCheckPolicyInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}