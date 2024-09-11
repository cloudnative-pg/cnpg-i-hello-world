# CNPG-I Hello World Plugin

A [CNPG-I](https://github.com/cloudnative-pg/cnpg-i) plugin to add 
user-defined labels, annotations and a specific `pause` sidecar
to the pods of
[CloudNativePG](https://github.com/cloudnative-pg/cloudnative-pg/) clusters.

This project serves as an introductory guide to bootstrapping
a [CNPG-I](https://github.com/cloudnative-pg/cnpg-i) plugin and leveraging
lifecycle hooks within a development environment. While similar results can be
achieved through simpler methods such as mutating webhooks or CNPG's built-in
features, this project is specifically designed to familiarize developers with
the plugin workflow. By understanding how lifecycle hooks interact with other
interfaces, developers can gain a deeper insight into implementing complex
resource changes in real-world applications.

This plugin uses
the [pluginhelper](https://github.com/cloudnative-pg/cnpg-i-machinery/tree/main/pkg/pluginhelper)
from [`cnpg-i-machinery`](https://github.com/cloudnative-pg/cnpg-i-machinery) to
simplify the plugin's implementation.

## Running the plugin

To see the plugin in execution, you need to have a Kubernetes cluster running
(we'll use [Kind](https://kind.sigs.k8s.io)) and the
[CloudNativePG](https://github.com/cloudnative-pg/cloudnative-pg/) operator
installed. The plugin also requires certificates to communicate with the
operator, hence we are also installing [cert-manager](https://cert-manager.io/)
to manage them.

``` shell
kind create cluster --name cnpg-i-hello-world
# Choose the latest version of CloudNativePG (at least 1.24)
kubectl apply --server-side -f \
  https://github.com/cloudnative-pg/cloudnative-pg/releases/download/vX.Y.Z/cnpg-X.Y.Z.yaml
# Choose the latest version of cert-manager
kubectl apply -f \
  https://github.com/cert-manager/cert-manager/releases/download/vX.Y.Z/cert-manager.yaml
```

Then install the plugin by applying the manifest:

<!-- TODO: reevaluate on release and set release-please to automatically update it-->

``` shell
kubectl apply -f https://github.com/cloudnative-pg/cnpg-i-hello-world/releases/download/v0.1.0/manifest.yaml
```

Finally, create a cluster resource to see the plugin in action. There are three
examples in the `doc/examples` directory:

1. [Cluster with labels and annotations](doc/examples/cluster-example.yaml):
   adds the defined labels and annotations to the pods. Showcases the plugin
   capability of altering the lifecycle of the CloudNativePG resources.
2. [Cluster with no parameters](doc/examples/cluster-example-no-parameters.yaml):
   defaults the plugin settings of the cluster. Showcases the plugin capability
   of altering the defaulting webhook behavior.
3. [Cluster with wrong parameters](doc/examples/cluster-example-with-mistake.yaml):
   includes an error in the configuration. Showcases the plugin capability of
   validating its own configuration.

## Plugin development

For additional details on the plugin implementation refer to
the [development documentation](doc/development.md).
