# Step 1: build image
FROM golang:1.23 AS builder

# Cache the dependencies
WORKDIR /app
COPY go.mod go.sum /app/
RUN go mod download

# Compile the application
COPY . /app
RUN --mount=type=cache,target=/root/.cache/go-build ./scripts/build.sh

# Step 2: build the image to be actually run
FROM golang:1-alpine
USER 10001:10001
COPY --from=builder /app/bin/cnpg-i-hello-world /app/bin/cnpg-i-hello-world
ENTRYPOINT ["/app/bin/cnpg-i-hello-world"]
