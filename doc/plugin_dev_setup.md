### Plugin development

For setup instructions and development guide, refer to [doc/newplugin\_dev\_setup.md](doc/newplugin_dev_setup.md).

---

# Plugin Development Setup


````markdown
# CNPG-I Plugin Development Setup Guide

This document outlines how to set up a development environment for CNPG-I plugin development.

## Prerequisites
- Git
- Go >= 1.22
- Docker
- Make
- Kind (Kubernetes-in-Docker)
- kubectl
- kustomize (>= v5.0.0)

## 1. Clone the Repository

```bash
git clone https://github.com/cloudnative-pg/cnpg-i-hello-world-plugin.git
cd cnpg-i-hello-world-plugin
````

## 2. Install Go Dependencies

```bash
go mod tidy
```

## 3. Install Kind and Create Cluster

```bash
kind create cluster --name cnpg-i-hello-world
```

## 4. Install CloudNativePG Operator

```bash
kubectl apply --server-side -f \
  https://github.com/cloudnative-pg/cloudnative-pg/releases/download/vX.Y.Z/cnpg-X.Y.Z.yaml
```

## 5. Install cert-manager

```bash
kubectl apply -f \
  https://github.com/cert-manager/cert-manager/releases/download/vX.Y.Z/cert-manager.yaml
```

## 6. Run the Plugin Locally

In one terminal window, run:

```bash
make run
```

This uses the pluginhelper machinery to run the plugin in a local dev mode.

## 7. Deploy and Test Example Clusters

In a separate terminal window:

```bash
kubectl apply -f doc/examples/01_cluster_with_labels_and_annotations.yaml
kubectl get pods -n <namespace>  # Verify sidecars and annotations are applied
```

## 8. Make Your Changes

Edit plugin logic in `internal/` or wherever appropriate. Then rebuild:

```bash
make build
```

## 9. Test Your Plugin

Use example manifests or write new ones to verify plugin behavior.

Happy Hacking!