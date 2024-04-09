#!/usr/bin/env bash

cd "$(dirname "$0")/.." || exit

if [ -z "$(which docker)" ]; then
    echo "Please install docker"
    exit 1
fi

docker build -t cnpg-i-hello-world:${VERSION:-latest} .
