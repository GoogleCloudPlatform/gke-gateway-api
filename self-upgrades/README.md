# Self-Managed Gateway API CRD Upgrades on GKE

This directory contains tooling to allow GKE customers to manually upgrade their Kubernetes Gateway API Custom Resource Definitions (CRDs) to newer upstream versions.

## Overview

By default, GKE manages the installation and lifecycle of Gateway API CRDs via the GKE release cycle. However, you may wish to access newer features available in the upstream Gateway API project before they are backported to your specific GKE version.

This process allows you to apply a **strictly newer** version of the **Standard** Channel CRDs.

## Prerequisites

* A GKE cluster with Gateway API enabled (--`gateway-api=standard`).  
* `kubectl` installed and configured for your cluster.  
* `kustomize` (required by the installation script).

## Usage

We provide a helper script (`upgrade-gateway-api-crds.sh`) that automates the upgrade process. This script fetches the specified version from the upstream project and injects the specific annotations required for GKE to recognize the manual upgrade.

### 1. Identify Target Version

Determine the upstream version you wish to install (e.g., `v1.4.0`). You can find available releases on the [Kubernetes Gateway API Releases page](https://github.com/kubernetes-sigs/gateway-api/releases).

### 2. Run the Script

Execute the script with the desired version tag as the argument:

```shell
# Make the script executable
chmod +x upgrade-gateway-api-crds.sh

# Install specific version (e.g., v1.4.0)
./upgrade-gateway-api-crds.sh v1.4.0
```

## How it Works

When you run this script:

1. It downloads the **Standard Channel** CRDs for the requested version.  
2. It injects the `components.gke.io/component-version` annotation.  
3. It applies the manifests to your cluster.

If the version you apply is higher than what GKE bundles, GKE will respect your manual installation and pause its own reconciliation.

## Important Limitations

* **Standard Channel Only:** You may only install CRDs from the `standard` channel. GKE policy restricts the use of `experimental` CRDs; attempting to install them will result in a validation error.  
* **Strictly Newer Versions:** You can only upgrade to a version newer than what is currently on your cluster. You cannot downgrade below the version GKE provides.  
* **Automatic Updates:** If the GKE managed version eventually exceeds your manually installed version (e.g., during a cluster upgrade), GKE will resume management of the CRDs and update them automatically.  
* **GCP-Specific CRDs:** This script applies to OSS Kubernetes CRDs (e.g., `Gateway`, `HTTPRoute`).
