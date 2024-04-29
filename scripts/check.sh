#!/usr/bin/env bash

cd "$(dirname "$0")/.." || exit

if [ ! -f "bin/golangci-lint" ]; then
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s v1.55.2
fi

bin/golangci-lint run
