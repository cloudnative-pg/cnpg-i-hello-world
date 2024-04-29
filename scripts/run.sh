#!/usr/bin/env bash

set -eu

cd "$(dirname "$0")/.." || exit

if [ -f .env ]; then
    source .env
fi

# The following script deployes this plugin in a locally running CNPG installation
# in Kind
current_context=$(kubectl config view --raw -o json | jq -r '."current-context"' | sed "s/kind-//")
kind load docker-image --name=${current_context} cnpg-i-hello-world:${VERSION:-latest}

kubectl patch deployment -n cnpg-system  cnpg-controller-manager --patch-file  kubernetes/deployment-patch.json
kubectl rollout restart deployment -n cnpg-system  cnpg-controller-manager
kubectl rollout status deployment -n cnpg-system  cnpg-controller-manager
