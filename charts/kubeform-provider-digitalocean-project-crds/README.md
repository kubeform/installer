# Kubeform Provider Digitalocean Project CRDs

[Kubeform Provider Digitalocean Project CRDs](https://github.com/kubeform) - Kubeform Provider Digitalocean Project Custom Resource Definitions

## TL;DR;

```console
$ helm repo add appscode https://charts.appscode.com/stable/
$ helm repo update
$ helm install kubeform-provider-digitalocean-project-crds appscode/kubeform-provider-digitalocean-project-crds -n kubeform
```

## Introduction

This chart deploys Kubeform Provider Digitalocean Project crds on a [Kubernetes](http://kubernetes.io) cluster using the [Helm](https://helm.sh) package manager.

## Prerequisites

- Kubernetes 1.16+

## Installing the Chart

To install the chart with the release name `kubeform-provider-digitalocean-project-crds`:

```console
$ helm install kubeform-provider-digitalocean-project-crds appscode/kubeform-provider-digitalocean-project-crds -n kubeform
```

The command deploys Kubeform Provider Digitalocean Project crds on the Kubernetes cluster in the default configuration. The [configuration](#configuration) section lists the parameters that can be configured during installation.

> **Tip**: List all releases using `helm list`

## Uninstalling the Chart

To uninstall/delete the `kubeform-provider-digitalocean-project-crds`:

```console
$ helm delete kubeform-provider-digitalocean-project-crds -n kubeform
```

The command removes all the Kubernetes components associated with the chart and deletes the release.


