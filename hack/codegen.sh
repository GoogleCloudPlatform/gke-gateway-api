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

readonly MODULE_NAME=github.com/GoogleCloudPlatform/gke-gateway-api
readonly CLIENT_PKG="$MODULE_NAME/pkg/client"
readonly API_DIR="$MODULE_NAME/apis/networking/v1"

# Work around tooling needing to operate on repos under GOPATH
if ! [ -d "$GOPATH/src/$MODULE_NAME" ]; then
  mkdir -p "$GOPATH/src/$(dirname $MODULE_NAME)"
  ln -s "${REPO_DIR}" "$GOPATH/src/$MODULE_NAME"
fi

echo "Generating deepcopy for ${MODULE_NAME}/apis/v1"
go run sigs.k8s.io/controller-tools/cmd/controller-gen \
  object:headerFile=${REPO_DIR}/hack/boilerplate/boilerplate.go.txt \
  paths="${MODULE_NAME}/apis/networking/v1"

