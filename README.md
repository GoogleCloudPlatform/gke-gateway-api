# GKE Gateway API types

This repo contains the type definitions for GKE Gateway Service Policy types
as described in
[Configure Gateway resources using Policies](https://cloud.google.com/kubernetes-engine/docs/how-to/configure-gateway-resources)


# Conformance Testing
This repo also contains the custom setup for GKE Gateway conformance testing.

Refer to https://gateway-api.sigs.k8s.io/concepts/conformance/ for more information.

## Run a single test case
Example:

```
go test ./conformance -run TestConformance -v -args \
    --gateway-class=gke-l7-rilb \
    --run-test=HTTPRouteRequestMirror
```

## Obtain a conformance report
Example:

```
go test ./conformance -run TestConformance -v -timeout=2h -args \
    --gateway-class=gke-l7-global-external-managed \
    --conformance-profiles=GATEWAY-HTTP \
    --organization=GKE \
    --project=gke-gateway \
    --url=https://cloud.google.com/kubernetes-engine/docs/concepts/gateway-api \
    --version=1.30.3-gke.1211000 \
    --contact=gke-gateway-dev@google.com \
    --report-output="/path/to/report"
```