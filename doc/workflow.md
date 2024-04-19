# Implementation Workflow

## Identity

1. Inside the `internal/identity` package, define a struct that implements the `pluginhelper.IdentityServer` interface.
2. Implement the following methods:
  - `GetPluginMetadata`: Return human-readable information about the plugin.
  - `GetPluginCapabilities`: Specify the features supported by the plugin. 
In the hello-world example we [define](internal/lifecycle/lifecycle.go) the `PluginCapability_Service_TYPE_LIFECYCLE_SERVICE`
  - `Probe`: Indicate whether the plugin is ready to serve requests. The hello-example plugin is stateless so it will return
always ready.

## Implement the supported features

Given that the hello-world example supports the lifecycle service capabilities we will implement the `OperatorLifecycleServer` interface, this is done
inside the `internal/lifecycle` pkg.

The `OperatorLifecycleServer` interface requires several methods:
- `GetCapabilities`, it should describe the resources and operations for which the plugin should be notified
- `LifecycleHook`, is the function that will be invoked with the `OperatorLifecycleRequest`. In this function the plugin is expected
to do a pattern matching with the `Kind` and the operation `Type` so it can proceed to execute the proper logic.

## Register the plugin webhooks on cluster creation and mutation

The operator interface offers a way for the plugin to interact with the Cluster resource webhooks.
This is done by implementing the [operator](https://github.com/cloudnative-pg/cnpg-i/blob/main/proto/operator.proto)
interfaces `MutateCluster`, `ValidateClusterCreate` and `ValidateClusterChange`.

- `MutateCluster` it should set the defaults parameters for the plugin when needed
- `ValidateClusterCreate` and `ValidateClusterChange` are used to execute the webhook validation logic for the plugin
fields

In the example this is done in the pkg `internal/operator`

## Startup Command

1. Invoke `pluginhelper.CreateMainCmd` with the implemented Identity struct created during the `Identity` step.
```
pluginhelper.CreateMainCmd(identity.Implementation{}, func(server *grpc.Server)
```
2. Register any implementations for the declared features within the callback function. In the hello-world example would 
be the RegisterOperatorServer and the Lifecycle Service:
```
operator.RegisterOperatorServer(server, operatorImpl.Implementation{})
lifecycle.RegisterOperatorLifecycleServer(server, lifecycleImpl.Implementation{})
```

# Local Testing

The repository provides a `Makefile` that contains several helpful commands to execute the local testing of the plugin,
it assumes that you did setup a cluster running CNPG with the instructions contained on the CNPG operator repository.

By executing `make run` a docker image containing the executable of the repository will be built and will be loaded inside
the kind cluster, after that the operator deployment will be patched with a sidecar containing the hello-world
plugin.
