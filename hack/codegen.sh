#!/usr/bin/env bash
#
# Copyright 2024 Google LLC
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     https://www.apache.org/licenses/LICENSE-2.0
#
#     Unless required by applicable law or agreed to in writing, software
#     distributed under the License is distributed on an "AS IS" BASIS,
#     WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
#     See the License for the specific language governing permissions and
#     limitations under the License.
#

set -euo pipefail

readonly SCRIPT_DIR="$(readlink -f "$(dirname "${BASH_SOURCE[0]}")")"
readonly REPO_DIR="$(readlink -f $SCRIPT_DIR/..)"

export GOFLAGS="${GOFLAGS:-} -mod=readonly"

echo "Script dir: $SCRIPT_DIR"
echo "Repo dir: $REPO_DIR"
echo "GOFLAGS: $GOFLAGS"

readonly MODULE_PATH=github.com/GoogleCloudPlatform/gke-gateway-api
readonly CLIENT_PACKAGE_PATH="$MODULE_PATH/pkg/client"
readonly API_PACKAGE_PATH="$MODULE_PATH/apis/networking/v1"
readonly CLIENTSET_NAME=versioned
readonly CLIENTSET_PKG_NAME=clientset

# Work around tooling needing to operate on repos under GOPATH
if ! [ -d "$GOPATH/src/$MODULE_PATH" ]; then
  mkdir -p "$GOPATH/src/$(dirname $MODULE_PATH)"
  ln -s "${REPO_DIR}" "$GOPATH/src/$MODULE_PATH"
fi


GEN_FLAGS=()
GEN_FLAGS+=("--go-header-file ${REPO_DIR}/hack/boilerplate/boilerplate.go.txt")
if ${VERIFY:-false}; then
  echo "Running in verification mode"
  GEN_FLAGS+=("--verify-only")
fi

echo "Generating clientset at ${CLIENT_PACKAGE_PATH}/${CLIENTSET_PKG_NAME}"
go run k8s.io/code-generator/cmd/client-gen \
  --clientset-name "${CLIENTSET_NAME}" \
  --input-base "${MODULE_PATH}/apis" \
  --input "networking/v1" \
  --output-pkg "${CLIENT_PACKAGE_PATH}/${CLIENTSET_PKG_NAME}" \
  --output-dir "pkg/client/${CLIENTSET_PKG_NAME}" \
  ${GEN_FLAGS}

echo "Generating listers at ${CLIENT_PACKAGE_PATH}/listers"
go run k8s.io/code-generator/cmd/lister-gen \
  --output-dir "pkg/client/listers" \
  --output-pkg "${CLIENT_PACKAGE_PATH}/listers" \
  ${GEN_FLAGS} \
  ${MODULE_PATH}/apis/networking/v1

echo "Generating informers at ${CLIENT_PACKAGE_PATH}/informers"
go run k8s.io/code-generator/cmd/informer-gen \
  --versioned-clientset-package "${CLIENT_PACKAGE_PATH}/${CLIENTSET_PKG_NAME}/${CLIENTSET_NAME}" \
  --listers-package "${CLIENT_PACKAGE_PATH}/listers" \
  --output-dir "pkg/client/informers" \
  --output-pkg "${CLIENT_PACKAGE_PATH}/informers" \
  ${GEN_FLAGS} \
  ${MODULE_PATH}/apis/networking/v1

echo "Generating register for ${MODULE_PATH}/apis/networking/v1"
go run k8s.io/code-generator/cmd/register-gen \
  --output-file zz_generated.register.go \
  ${GEN_FLAGS} \
  ${MODULE_PATH}/apis/networking/v1

echo "Generating deepcopy for ${MODULE_PATH}/apis/networking/v1"
go run sigs.k8s.io/controller-tools/cmd/controller-gen \
  object:headerFile=${REPO_DIR}/hack/boilerplate/boilerplate.go.txt \
  paths="${MODULE_PATH}/apis/networking/v1"

echo "Generating crds for ${MODULE_PATH}/apis/networking/v1"
go run sigs.k8s.io/controller-tools/cmd/controller-gen \
  crd  \
  paths="${MODULE_PATH}/apis/networking/v1" \
  output:crd:artifacts:config="${REPO_DIR}/config/crd"

