# Concepts

### Identity

The Identity interface declares the features supported by the plugin, and is the only interface that needs to be always
implemented.

This information is crucial for the operator to discover the plugin's
capabilities during startup.

It exposes:
- A way for the plugins to report readiness probe. The readiness is a requirement to receive events.
It is expected that the plugins always report back the most accurate readiness data available.
- The plugin reported capabilities, that will dictate which subsequent calls the plugin will receive.
- Metadata about the Plugin

[API reference](https://github.com/cloudnative-pg/cnpg-i/blob/main/proto/identity.proto)

## Features

### Operator (ClusterLifecycle)

This feature enables the plugin to receive events about the cluster
creation and mutations, this is defined by the following

```
// ValidateCreate improves the behavior of the validating webhook that
// is called on creation of the Cluster resources
rpc ValidateClusterCreate(OperatorValidateClusterCreateRequest) returns (OperatorValidateClusterCreateResult) {}

// ValidateClusterChange improves the behavior of the validating webhook of
// is called on updates of the Cluster resources
rpc ValidateClusterChange(OperatorValidateClusterChangeRequest) returns (OperatorValidateClusterChangeResult) {}

// MutateCluster fills in the defaults inside a Cluster resource
rpc MutateCluster(OperatorMutateClusterRequest) returns (OperatorMutateClusterResult) {}
```

This interface allows the plugins to implement important features like:

1. validating the cluster manifest during the creation and mutations
   (it is expected that the plugin validate the parameters assigned to
   their configuration).

2. mutating the cluster object before it is submitted to kubernetes API
   server, for example to set default values for the plugin parameters.

[API reference](https://github.com/cloudnative-pg/cnpg-i/blob/main/proto/operator.proto)

### Lifecycle

This feature enables the plugin to receive events and create patches
for Kubernetes resources `before` they are submitted to the API server.

To use this feature, the plugin must specify the resource and operation
it wants to be notified of.

Some examples of what it can be achieved through the lifecycle:
- add volume, volume mounts, sidecar containers, labels, annotations to pods, especially necessary when implementing
custom backup solutions
- modify any resource with some annotations or labels
- add/remove finalizers

[API reference](https://github.com/cloudnative-pg/cnpg-i/blob/main/proto/operator_lifecycle.proto):

### Backup

This feature allows the plugin to implement new physical backup methods.

The backup interfaces is invoked by the `elected instance` for the backup.
The instance is elected based off two criteria:
- the specified `backup.spec.target` value
- the readiness of the instance

It could happen that the custom backup that is being implemented needs to act also on the Backup reconciliation too,
this feature is offered by the `Custom Reconcilers` section.

To add volumes and volume mounts to execute the custom logic please refer to the `Lifecycle` section

[API reference](https://github.com/cloudnative-pg/cnpg-i/blob/main/proto/backup.proto)


### Custom reconcilers

This feature allows the plugin to enhance the operator
reconciliation for resources managed by CNPG, such as Clusters
and Backups.

The custom logic can be injected before the reconciliation and after it is executed.

[API reference](https://github.com/cloudnative-pg/cnpg-i/blob/main/proto/reconciler.proto)
