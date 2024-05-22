# Implementation Workflow

## Identity

1. Define a struct inside the `internal/identity` package that implements
   the `pluginhelper.IdentityServer` interface.

2. Implement the following methods:

    - `GetPluginMetadata`: return human-readable information about the plugin.
    - `GetPluginCapabilities`: specify the features supported by the plugin.
      In the hello-world example, the
      `PluginCapability_Service_TYPE_LIFECYCLE_SERVICE` is defined
      in the corresponding Go [file](../internal/lifecycle/lifecycle.go).
    - `Probe`: indicate whether the plugin is ready to serve requests; this
      example is stateless, so it will always be ready.

## Implement the supported features

This example implements the lifecycle service capabilities,
and the `OperatorLifecycleServer` interface is implemented inside the
`internal/lifecycle` package.

The `OperatorLifecycleServer` interface requires several methods:

- `GetCapabilities`: describe the resources and operations the plugin
  should be notified for

- `LifecycleHook`: is invoked for every operation against the
  Kubernetes API server that matches the specifications
  returned by `GetCapabilities`

  In this function, the plugin is expected to do pattern matching
  using the `Kind` and the operation `Type` and proceed with the
  proper logic.

## Register the plugin webhooks on cluster creation and mutation

The operator interface offers a way for the plugin to interact with
the Cluster resource webhooks.

Do that, the plugin should implement the [operator](https://github.com/cloudnative-pg/cnpg-i/blob/main/proto/operator.proto)
interface, specifically the `MutateCluster`, `ValidateClusterCreate`,
and `ValidateClusterChange` rpc calls.

- `MutateCluster`: enriches the plugin defaulting webhook

- `ValidateClusterCreate` and `ValidateClusterChange`: enriches
  the plugin validation logic.

The package `internal/operator` implements this interface.

## Startup Command

This example plugin is installed in a sidecar of the operator, and its
main command is implemented in the `main.go` file.

This function uses the plugin helper library to create a GRPC server.

Plugin developers are expected to use the `pluginhelper.CreateMainCmd`
to implement the `main` function, passing an implemented `Identity`
struct.

Further implementations can be registered within the callback function.

In the example we propose, that's done for  **operator** and for
the **lifecycle** services in `cmd/plugin/plugin.go`:

```
operator.RegisterOperatorServer(server, operatorImpl.Implementation{})
lifecycle.RegisterOperatorLifecycleServer(server, lifecycleImpl.Implementation{})
```

# Testing the plugin locally

The repository provides a `Makefile` that contains several helpful
commands to test the plugin in a CNPG development environment.

By executing `make run`, a Docker image containing the executable
of the repository will be built and loaded inside the kind cluster.

Having done that, the operator deployment will be patched with a sidecar
containing the hello-world plugin.
