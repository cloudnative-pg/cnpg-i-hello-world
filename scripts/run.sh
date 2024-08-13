#!/usr/bin/env bash

set -eu

cd "$(dirname "$0")/.." || exit

if [ -f .env ]; then
    source .env
fi

# The following script build the plugin image and uploads it to the
# current kind cluster
current_context=$(kubectl config view --raw -o json | jq -r '."current-context"' | sed "s/kind-//")
kind load docker-image --name=${current_context} cnpg-i-hello-world:${VERSION:-latest}

# Now we deploy the plugin inside the `cnpg-system` workspace
kubectl apply -k kubernetes/
# Patch the deployment to use the provided image and disable imagePullPolicy, since the image is already loaded
kubectl patch deployments.apps -n cnpg-system hello-world -p \
  "{\"spec\":{\"template\":{\"spec\":{\"containers\":[{\"name\":\"cnpg-i-hello-world\",\"image\":\"cnpg-i-hello-world:${VERSION:-latest}\",\"imagePullPolicy\":\"Never\"}]}}}}"
