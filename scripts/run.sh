#!/usr/bin/env bash
##
## Copyright © contributors to CloudNativePG, established as
## CloudNativePG a Series of LF Projects, LLC.
##
## Licensed under the Apache License, Version 2.0 (the "License");
## you may not use this file except in compliance with the License.
## You may obtain a copy of the License at
##
##     http://www.apache.org/licenses/LICENSE-2.0
##
## Unless required by applicable law or agreed to in writing, software
## distributed under the License is distributed on an "AS IS" BASIS,
## WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
## See the License for the specific language governing permissions and
## limitations under the License.
##
## SPDX-License-Identifier: Apache-2.0
##

set -eu

cd "$(dirname "$0")/.." || exit

if [ -f .env ]; then
    source .env
fi

# The following script builds the plugin image and uploads it to the
# current kind cluster
# WARNING: This will fail with recent releases of kind due to https://github.com/kubernetes-sigs/kind/issues/3853
# See fix in CNPG https://github.com/cloudnative-pg/cloudnative-pg/pull/6770
# current_context=$(kubectl config view --raw -o json | jq -r '."current-context"' | sed "s/kind-//")
# kind load docker-image --name=${current_context} cnpg-i-hello-world:${VERSION:-latest}

# Constants
registry_name=registry.dev

load_image_registry() {
  local image=$1

  local image_reg_name=${registry_name}:5000/${image}
  local image_local_name=${image_reg_name/${registry_name}/127.0.0.1}
  docker tag "${image}" "${image_reg_name}"
  docker tag "${image}" "${image_local_name}"
  docker push -q "${image_local_name}"
}

# Now we deploy the plugin inside the `cnpg-system` workspace
kubectl apply -k kubernetes/

# We load the image into the registry (which is a prerequisite)
load_image_registry cnpg-i-hello-world:${VERSION:-latest}

# Patch the deployment to use the provided image
kubectl patch deployments.apps -n cnpg-system hello-world -p \
  "{\"spec\":{\"template\":{\"spec\":{\"containers\":[{\"name\":\"cnpg-i-hello-world\",\"image\":\"${registry_name}:5000/cnpg-i-hello-world:${VERSION:-latest}\"}]}}}}"
