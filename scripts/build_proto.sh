#!/usr/bin/env bash

set -eu

# Get the directory of the script
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"

# Navigate to the parent directory
cd "$SCRIPT_DIR/.." || { echo "Directory does not exist, skipping build_proto"; exit 0; }

# Check if the contain any .proto files
if [ ! -d "proto" ] || [ -z "$(find proto -name '*.proto' -print -quit)" ]; then
    echo "No .proto files found in the 'proto' directory, skipping build_proto"
    exit 0
fi

# Recompile protobuf specification
protoc --go_out=. --go_opt=module=github.com/cloudnative-pg/cnpg-i-hello-world \
    --go-grpc_out=. --go-grpc_opt=module=github.com/cloudnative-pg/cnpg-i-hello-world \
    proto/*.proto
