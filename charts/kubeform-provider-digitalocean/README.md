# kubeform-provider-digitalocean

[Kubeform Provider Digitalocean Controller by AppsCode](https://github.com/kubeform) - Provision Digitalocean resources using Terraform

## TL;DR;

```console
$ helm repo add appscode https://charts.appscode.com/stable/
$ helm repo update
$ helm install kubeform-provider-digitalocean appscode/kubeform-provider-digitalocean -n kubeform
```

## Introduction

This chart deploys a Kubeform Provider Digitalocean Controller on a [Kubernetes](http://kubernetes.io) cluster using the [Helm](https://helm.sh) package manager.

## Prerequisites

- Kubernetes 1.16+

## Installing the Chart

To install the chart with the release name `kubeform-provider-digitalocean`:

```console
$ helm install kubeform-provider-digitalocean appscode/kubeform-provider-digitalocean -n kubeform
```

The command deploys a Kubeform Provider Digitalocean Controller on the Kubernetes cluster in the default configuration. The [configuration](#configuration) section lists the parameters that can be configured during installation.

> **Tip**: List all releases using `helm list`

## Uninstalling the Chart

To uninstall/delete the `kubeform-provider-digitalocean`:

```console
$ helm delete kubeform-provider-digitalocean -n kubeform
```

The command removes all the Kubernetes components associated with the chart and deletes the release.

## Configuration

The following table lists the configurable parameters of the `kubeform-provider-digitalocean` chart and their default values.

|       Parameter        |                                    Description                                    |                              Default                               |
|------------------------|-----------------------------------------------------------------------------------|--------------------------------------------------------------------|
| kubeform-provider      | Pass values to kubeform-provider chart                                            | `{"operator":{"tag":"v0.3.0"},"provider":{"name":"digitalocean"}}` |
| crds.app               | If true, installs the kubeform-provider-digitalocean-app-crds chart               | `false`                                                            |
| crds.cdn               | If true, installs the kubeform-provider-digitalocean-cdn-crds chart               | `false`                                                            |
| crds.certificate       | If true, installs the kubeform-provider-digitalocean-certificate-crds chart       | `false`                                                            |
| crds.containerregistry | If true, installs the kubeform-provider-digitalocean-containerregistry-crds chart | `false`                                                            |
| crds.custom            | If true, installs the kubeform-provider-digitalocean-custom-crds chart            | `false`                                                            |
| crds.database          | If true, installs the kubeform-provider-digitalocean-database-crds chart          | `false`                                                            |
| crds.domain            | If true, installs the kubeform-provider-digitalocean-domain-crds chart            | `false`                                                            |
| crds.droplet           | If true, installs the kubeform-provider-digitalocean-droplet-crds chart           | `false`                                                            |
| crds.firewall          | If true, installs the kubeform-provider-digitalocean-firewall-crds chart          | `false`                                                            |
| crds.floatingip        | If true, installs the kubeform-provider-digitalocean-floatingip-crds chart        | `false`                                                            |
| crds.kubernetes        | If true, installs the kubeform-provider-digitalocean-kubernetes-crds chart        | `false`                                                            |
| crds.loadbalancer      | If true, installs the kubeform-provider-digitalocean-loadbalancer-crds chart      | `false`                                                            |
| crds.project           | If true, installs the kubeform-provider-digitalocean-project-crds chart           | `false`                                                            |
| crds.record            | If true, installs the kubeform-provider-digitalocean-record-crds chart            | `false`                                                            |
| crds.spacesbucket      | If true, installs the kubeform-provider-digitalocean-spacesbucket-crds chart      | `false`                                                            |
| crds.ssh               | If true, installs the kubeform-provider-digitalocean-ssh-crds chart               | `false`                                                            |
| crds.tag               | If true, installs the kubeform-provider-digitalocean-tag-crds chart               | `false`                                                            |
| crds.volume            | If true, installs the kubeform-provider-digitalocean-volume-crds chart            | `false`                                                            |
| crds.vpc               | If true, installs the kubeform-provider-digitalocean-vpc-crds chart               | `false`                                                            |


Specify each parameter using the `--set key=value[,key=value]` argument to `helm install`. For example:

```console
$ helm install kubeform-provider-digitalocean appscode/kubeform-provider-digitalocean -n kubeform --set kubeform-provider={"operator":{"tag":"v0.3.0"},"provider":{"name":"digitalocean"}}
```

Alternatively, a YAML file that specifies the values for the parameters can be provided while
installing the chart. For example:

```console
$ helm install kubeform-provider-digitalocean appscode/kubeform-provider-digitalocean -n kubeform --values values.yaml
```
