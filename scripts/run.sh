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
kubectl apply -f kubernetes/
kubectl rollout restart deployment/hello-world -n cnpg-system