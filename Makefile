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

# CRD docs generated using crd-ref-docs
# https://github.com/elastic/crd-ref-docs
#
# Pages site generated using mkdocs-material
# https://squidfunk.github.io/mkdocs-material/

.PHONY: docs
docs: docs-crd-ref-docs-gen
	mkdocs build

.PHONY: docs-serve
docs-serve: docs-crd-ref-docs-gen
	mkdocs serve

.PHONY: docs-deploy
docs-deploy: docs-crd-ref-docs-gen
	mkdocs gh-deploy --force

.PHONY: docs-crd-ref-docs-gen
docs-crd-ref-docs-gen:
	crd-ref-docs \
	  --config config.yaml \
	  --renderer markdown \
	  --source-path apis \
	  --output-path docs/index.md

