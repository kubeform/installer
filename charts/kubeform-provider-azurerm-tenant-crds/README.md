# Kubeform Provider Azurerm Tenant CRDs

[Kubeform Provider Azurerm Tenant CRDs](https://github.com/kubeform) - Kubeform Provider Azurerm Tenant Custom Resource Definitions

## TL;DR;

```console
$ helm repo add appscode https://charts.appscode.com/stable/
$ helm repo update
$ helm install kubeform-provider-azurerm-tenant-crds appscode/kubeform-provider-azurerm-tenant-crds -n kubeform
```

## Introduction

This chart deploys Kubeform Provider Azurerm Tenant crds on a [Kubernetes](http://kubernetes.io) cluster using the [Helm](https://helm.sh) package manager.

## Prerequisites

- Kubernetes 1.16+

## Installing the Chart

To install the chart with the release name `kubeform-provider-azurerm-tenant-crds`:

```console
$ helm install kubeform-provider-azurerm-tenant-crds appscode/kubeform-provider-azurerm-tenant-crds -n kubeform
```

The command deploys Kubeform Provider Azurerm Tenant crds on the Kubernetes cluster in the default configuration. The [configuration](#configuration) section lists the parameters that can be configured during installation.

> **Tip**: List all releases using `helm list`

## Uninstalling the Chart

To uninstall/delete the `kubeform-provider-azurerm-tenant-crds`:

```console
$ helm delete kubeform-provider-azurerm-tenant-crds -n kubeform
```

The command removes all the Kubernetes components associated with the chart and deletes the release.


