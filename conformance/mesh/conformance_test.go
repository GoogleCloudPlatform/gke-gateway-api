/*
Copyright 2026 Google LLC

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package mesh_test

import (
	"testing"
	"time"

	"sigs.k8s.io/gateway-api/conformance"
	"sigs.k8s.io/gateway-api/conformance/utils/suite"
)

func TestMeshConformance(t *testing.T) {
	options := conformance.DefaultOptions(t)

	// Configure profile, supported features and exempt features
	options.ConformanceProfiles = suite.ParseConformanceProfilesSlice("MESH-HTTP")
	options.EnableAllSupportedFeatures = true
	options.ExemptFeatures = suite.ParseSupportedFeaturesSlice("Gateway,GRPCRoute,GRPCRouteNamedRouteRule")
	options.AllowCRDsMismatch = true
	options.CleanupBaseResources = true
	options.Debug = true

	// Inject CSM Sidecar Proxy Labels into test namespaces
	if options.NamespaceLabels == nil {
		options.NamespaceLabels = make(map[string]string)
	}
	options.NamespaceLabels["mesh.cloud.google.com/csm-injection"] = "sidecar"
	options.NamespaceLabels["istio-injection"] = "enabled"

	// Configure timeout config
	options.TimeoutConfig.DefaultTestTimeout = 300 * time.Second
	options.TimeoutConfig.HTTPRouteMustHaveCondition = 300 * time.Second
	options.TimeoutConfig.LatestObservedGenerationSet = 300 * time.Second
	options.TimeoutConfig.MaxTimeToConsistency = 1200 * time.Second
	options.TimeoutConfig.NamespacesMustBeReady = 600 * time.Second
	options.TimeoutConfig.RequestTimeout = 20 * time.Second
	options.TimeoutConfig.RouteMustHaveParents = 600 * time.Second

	conformance.RunConformanceWithOptions(t, options)
}
