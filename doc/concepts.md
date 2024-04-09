# Concepts

### Identity

The Identity interface declares the features supported by the plugin. This information is crucial for the operator to discover the plugin's capabilities during startup.

[API REFERENCE](https://github.com/cloudnative-pg/cnpg-i/blob/main/proto/identity.proto)

## Features

### Operator (ClusterLifecycle)

This feature enables the plugin to receive events about the cluster creation and mutations, this is defined by the following signatures:
```
  // ValidateCreate improves the behaviour of the validating webhook that
  // is called on creation of the Cluster resources
  rpc ValidateClusterCreate(OperatorValidateClusterCreateRequest) returns (OperatorValidateClusterCreateResult) {}

  // ValidateClusterChange improves the behavior of the validating webhook of
  // is called on updates of the Cluster resources
  rpc ValidateClusterChange(OperatorValidateClusterChangeRequest) returns (OperatorValidateClusterChangeResult) {}

  // MutateCluster fills in the defaults inside a Cluster resource
  rpc MutateCluster(OperatorMutateClusterRequest) returns (OperatorMutateClusterResult) {}
```

This is a powerful interface that allows the plugins to achieve several things:
1. To validate the cluster manifest during the creation and mutations, it is expected that the plugin validate the parameters
assigned to their section.
2. To mutate the cluster object before it is submitted to kubernetes api sever, for example to set default values for the plugin parameters.


### Lifecycle

This feature enables the plugin to receive events and create patches for Kubernetes resources before they are submitted to the API server.
To utilize this feature, the plugin must specify the resource and operation it wants to be notified of.

[API REFERENCE](https://github.com/cloudnative-pg/cnpg-i/blob/main/proto/operator_lifecycle.proto):

### Backup

This feature allows the plugin to...

[API REFERENCE](https://github.com/cloudnative-pg/cnpg-i/blob/main/proto/backup.proto)

### Custom Reconcilers

This feature allows the plugin to register a custom reconciler for a supported resource.

[API REFERENCE](https://github.com/cloudnative-pg/cnpg-i/blob/main/proto/reconciler.proto)
