# kubeform-provider-linode

[Kubeform Provider Linode Controller by AppsCode](https://github.com/kubeform) - Provision Linode resources using Terraform

## TL;DR;

```console
$ helm repo add appscode https://charts.appscode.com/stable/
$ helm repo update
$ helm install kubeform-provider-linode appscode/kubeform-provider-linode -n kubeform
```

## Introduction

This chart deploys a Kubeform Provider Linode Controller on a [Kubernetes](http://kubernetes.io) cluster using the [Helm](https://helm.sh) package manager.

## Prerequisites

- Kubernetes 1.16+

## Installing the Chart

To install the chart with the release name `kubeform-provider-linode`:

```console
$ helm install kubeform-provider-linode appscode/kubeform-provider-linode -n kubeform
```

The command deploys a Kubeform Provider Linode Controller on the Kubernetes cluster in the default configuration. The [configuration](#configuration) section lists the parameters that can be configured during installation.

> **Tip**: List all releases using `helm list`

## Uninstalling the Chart

To uninstall/delete the `kubeform-provider-linode`:

```console
$ helm delete kubeform-provider-linode -n kubeform
```

The command removes all the Kubernetes components associated with the chart and deletes the release.

## Configuration

The following table lists the configurable parameters of the `kubeform-provider-linode` chart and their default values.

|     Parameter     |                              Description                               |                           Default                            |
|-------------------|------------------------------------------------------------------------|--------------------------------------------------------------|
| kubeform-provider | Pass values to kubeform-provider chart                                 | `{"operator":{"tag":"v0.3.0"},"provider":{"name":"linode"}}` |
| crds.domain       | If true, installs the kubeform-provider-linode-domain-crds chart       | `false`                                                      |
| crds.firewall     | If true, installs the kubeform-provider-linode-firewall-crds chart     | `false`                                                      |
| crds.image        | If true, installs the kubeform-provider-linode-image-crds chart        | `false`                                                      |
| crds.instance     | If true, installs the kubeform-provider-linode-instance-crds chart     | `false`                                                      |
| crds.lke          | If true, installs the kubeform-provider-linode-lke-crds chart          | `false`                                                      |
| crds.nodebalancer | If true, installs the kubeform-provider-linode-nodebalancer-crds chart | `false`                                                      |
| crds.object       | If true, installs the kubeform-provider-linode-object-crds chart       | `false`                                                      |
| crds.rdns         | If true, installs the kubeform-provider-linode-rdns-crds chart         | `false`                                                      |
| crds.sshkey       | If true, installs the kubeform-provider-linode-sshkey-crds chart       | `false`                                                      |
| crds.stackscript  | If true, installs the kubeform-provider-linode-stackscript-crds chart  | `false`                                                      |
| crds.token        | If true, installs the kubeform-provider-linode-token-crds chart        | `false`                                                      |
| crds.user         | If true, installs the kubeform-provider-linode-user-crds chart         | `false`                                                      |
| crds.volume       | If true, installs the kubeform-provider-linode-volume-crds chart       | `false`                                                      |


Specify each parameter using the `--set key=value[,key=value]` argument to `helm install`. For example:

```console
$ helm install kubeform-provider-linode appscode/kubeform-provider-linode -n kubeform --set global.registry=kubeform
```

Alternatively, a YAML file that specifies the values for the parameters can be provided while
installing the chart. For example:

```console
$ helm install kubeform-provider-linode appscode/kubeform-provider-linode -n kubeform --values values.yaml
```
