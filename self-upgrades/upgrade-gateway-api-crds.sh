#!/bin/bash
set -e

usage() {
    echo "Usage: $0 <VERSION>"
    echo "Example: $0 v1.4.0"
    echo "Available versions: https://github.com/kubernetes-sigs/gateway-api/releases"
    exit 1
}

VERSION=$1
if [ -z "$VERSION" ]; then
    echo "Error: No version specified."
    usage
fi

# Ensure the version starts with 'v' (standard convention for Gateway API tags)
if [[ ! "$VERSION" =~ ^v ]]; then
    echo "Warning: Version '$VERSION' does not start with 'v'. Upstream tags usually require it (e.g., v1.4.0)."
    read -p "Continue anyway? (y/n) " -n 1 -r
    echo
    if [[ ! $REPLY =~ ^[Yy]$ ]]; then
        exit 1
    fi
fi

echo "---"
echo "Preparing to install Kubernetes Gateway API Standard CRDs"
echo "Target Version: $VERSION"
echo "---"

# 2. Create a temporary directory for Kustomize context
TEMP_DIR=$(mktemp -d)
trap 'rm -rf "$TEMP_DIR"' EXIT

cat <<EOF > "$TEMP_DIR/kustomization.yaml"
apiVersion: kustomize.config.k8s.io/v1beta1
kind: Kustomization

resources:
  # Point to the upstream 'standard' channel specific to the requested tag
  - https://github.com/kubernetes-sigs/gateway-api/config/crd?ref=${VERSION}

commonAnnotations:
  components.gke.io/component-version: "${VERSION}"
EOF

# 4. Apply to the cluster
echo "Applying manifests to cluster..."
kubectl apply -k "$TEMP_DIR"

echo "---"
echo "✅ Success! Gateway API CRDs installed/updated to $VERSION."