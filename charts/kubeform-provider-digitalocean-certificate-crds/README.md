# Kubeform Provider Digitalocean Certificate CRDs

[Kubeform Provider Digitalocean Certificate CRDs](https://github.com/kubeform) - Kubeform Provider Digitalocean Certificate Custom Resource Definitions

## TL;DR;

```console
$ helm repo add appscode https://charts.appscode.com/stable/
$ helm repo update
$ helm install kubeform-provider-digitalocean-certificate-crds appscode/kubeform-provider-digitalocean-certificate-crds -n kubeform
```

## Introduction

This chart deploys Kubeform Provider Digitalocean Certificate crds on a [Kubernetes](http://kubernetes.io) cluster using the [Helm](https://helm.sh) package manager.

## Prerequisites

- Kubernetes 1.16+

## Installing the Chart

To install the chart with the release name `kubeform-provider-digitalocean-certificate-crds`:

```console
$ helm install kubeform-provider-digitalocean-certificate-crds appscode/kubeform-provider-digitalocean-certificate-crds -n kubeform
```

The command deploys Kubeform Provider Digitalocean Certificate crds on the Kubernetes cluster in the default configuration. The [configuration](#configuration) section lists the parameters that can be configured during installation.

> **Tip**: List all releases using `helm list`

## Uninstalling the Chart

To uninstall/delete the `kubeform-provider-digitalocean-certificate-crds`:

```console
$ helm delete kubeform-provider-digitalocean-certificate-crds -n kubeform
```

The command removes all the Kubernetes components associated with the chart and deletes the release.


