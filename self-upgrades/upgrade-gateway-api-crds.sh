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

is_upgrade_possible() {
    # Check if existing CRDs allow update (GKE 1.34+ check)
    local CRD_NAME="gatewayclasses.gateway.networking.k8s.io"

    if kubectl get crd "$CRD_NAME" >/dev/null 2>&1; then
        echo "Checking existing CRD management mode..."
        # Try to get the addonmanager mode label
        local EXISTING_MODE
        EXISTING_MODE=$(kubectl get crd "$CRD_NAME" -o jsonpath="{.metadata.labels.addonmanager\.kubernetes\.io/mode}" 2>/dev/null || true)
        
        if [ -z "$EXISTING_MODE" ]; then
            EXISTING_MODE=$(kubectl get crd "$CRD_NAME" -o jsonpath="{.metadata.labels['addonmanager.kubernetes.io/mode']}" 2>/dev/null || true)
        fi

        if [ -n "$EXISTING_MODE" ]; then
            echo "Existing addonmanager mode: $EXISTING_MODE"
            if [ "$EXISTING_MODE" != "NewerRevision" ]; then
                echo "Error: Updating Gateway API CRDs is not supported when they are managed by GKE Addon Manager in '$EXISTING_MODE' mode."
                echo "Updating is only supported from GKE 1.34+ where the mode is 'NewerRevision'."
                return 1
            fi
        else
            echo "Existing CRD is not managed by GKE Addon Manager (no mode label found). Proceeding."
        fi
    else
        echo "Gateway API CRDs not found in the cluster. Proceeding with fresh installation."
    fi
    return 0
}

if ! is_upgrade_possible; then
    exit 1
fi

# 2. Create a temporary directory for Kustomize context
TEMP_DIR=$(mktemp -d)
trap 'rm -rf "$TEMP_DIR"' EXIT

# Component version must strip the "v" prefix
COMP_VERSION=$(echo "$VERSION" | sed 's/^v//')

cat <<EOF > "$TEMP_DIR/kustomization.yaml"
apiVersion: kustomize.config.k8s.io/v1beta1
kind: Kustomization

resources:
  # Point to the upstream 'standard' channel specific to the requested tag
  - https://github.com/kubernetes-sigs/gateway-api/config/crd/?ref=${VERSION}

commonAnnotations:
  components.gke.io/component-name: gateway-api-crds
  components.gke.io/component-version: ${COMP_VERSION}
  components.gke.io/layer: addon
labels:
- includeSelectors: true
  pairs:
    addonmanager.kubernetes.io/mode: NewerRevision
EOF

# 4. Apply to the cluster
echo "Applying manifests to cluster..."
kubectl apply -k "$TEMP_DIR"

echo "---"
echo "✅ Success! Gateway API CRDs installed/updated to $VERSION."
