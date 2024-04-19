# Implementation Workflow

## Identity

1. Define a struct inside the `internal/identity` package that implements the `pluginhelper.IdentityServer` interface.
2. Implement the following methods:
    - `GetPluginMetadata`: Return human-readable information about the plugin.
    - `GetPluginCapabilities`: Specify the features supported by the plugin.
      In the hello-world example, the `PluginCapability_Service_TYPE_LIFECYCLE_SERVICE` is defined [here](internal/lifecycle/lifecycle.go).
    - `Probe`: Indicate whether the plugin is ready to serve requests. Since the hello-example plugin is stateless, it
    - will always return ready.

## Implement the supported features

Since the hello-world example supports the lifecycle service capabilities, implement the `OperatorLifecycleServer`
interface inside the `internal/lifecycle` package.

The `OperatorLifecycleServer` interface requires several methods:
- `GetCapabilities`: Describe the resources and operations for which the plugin should be notified.
- `LifecycleHook`: This function will be invoked with the `OperatorLifecycleRequest`.
In this function, the plugin is expected to do pattern matching with the `Kind` and the operation `Type` so it can
proceed to execute the proper logic.

## Register the plugin webhooks on cluster creation and mutation

The operator interface offers a way for the plugin to interact with the Cluster resource webhooks.
Implement the [operator](https://github.com/cloudnative-pg/cnpg-i/blob/main/proto/operator.proto)
interfaces `MutateCluster`, `ValidateClusterCreate`, and `ValidateClusterChange`.

- `MutateCluster`: Set the default parameters for the plugin when needed.
- `ValidateClusterCreate` and `ValidateClusterChange`: Execute the webhook validation logic for the plugin fields.

In the example, this is done in the package `internal/operator`.

## Startup Command

1. Invoke `pluginhelper.CreateMainCmd` with the implemented Identity struct created during the `Identity` step.
`pluginhelper.CreateMainCmd(identity.Implementation{}, func(server *grpc.Server)`
2. Register any implementations for the declared features within the callback function. In the hello-world example, it would be the `RegisterOperatorServer` and the Lifecycle Service:
```
operator.RegisterOperatorServer(server, operatorImpl.Implementation{})
lifecycle.RegisterOperatorLifecycleServer(server, lifecycleImpl.Implementation{})
```

# Local Testing

The repository provides a `Makefile` that contains several helpful commands to execute the local testing of the plugin.
It assumes that you have set up a cluster running CNPG with the instructions contained in the CNPG operator repository.

By executing `make run`, a Docker image containing the executable of the repository will be built and loaded inside
the kind cluster. After that, the operator deployment will be patched with a sidecar containing the hello-world plugin.
