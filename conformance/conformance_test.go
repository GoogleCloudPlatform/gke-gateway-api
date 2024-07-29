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

package conformance_test

import (
	"testing"
	"time"

	"sigs.k8s.io/gateway-api/conformance"
	"sigs.k8s.io/gateway-api/conformance/utils/suite"
)

func TestConformance(t *testing.T) {
	options := conformance.DefaultOptions(t)

	// Configure skip tests, supported features and exempt features
	options.SkipTests = suite.ParseSkipTests("HTTPRouteHostnameIntersection")
	options.SupportedFeatures = suite.ParseSupportedFeatures("Gateway,GatewayPort8080,HTTPRoute,HTTPRouteResponseHeaderModification,HTTPRouteSchemeRedirect,HTTPRoutePathRedirect,HTTPRouteHostRewrite,HTTPRouteRequestMirror")
	options.ExemptFeatures = suite.ParseSupportedFeatures("GatewayStaticAddresses,GatewayHTTPListenerIsolation,HTTPRouteBackendRequestHeaderModification,HTTPRouteQueryParamMatching,HTTPRouteMethodMatching,HTTPRoutePortRedirect,HTTPRoutePathRewrite,HTTPRouteRequestMultipleMirrors,HTTPRouteRequestTimeout,HTTPRouteBackendTimeout,HTTPRouteParentRefPort")

	// Configure timeout config
	options.TimeoutConfig.DefaultTestTimeout = 300 * time.Second
	options.TimeoutConfig.GatewayListenersMustHaveConditions = 300 * time.Second
	options.TimeoutConfig.GatewayStatusMustHaveListeners = 300 * time.Second
	options.TimeoutConfig.HTTPRouteMustHaveCondition = 300 * time.Second
	options.TimeoutConfig.LatestObservedGenerationSet = 300 * time.Second
	options.TimeoutConfig.MaxTimeToConsistency = 600 * time.Second
	options.TimeoutConfig.NamespacesMustBeReady = 600 * time.Second
	options.TimeoutConfig.RequestTimeout = 20 * time.Second
	options.TimeoutConfig.RouteMustHaveParents = 600 * time.Second

	conformance.RunConformanceWithOptions(t, options)
}
