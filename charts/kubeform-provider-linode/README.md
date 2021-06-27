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

|        Parameter        |                                                                                                                                                                                                 Description                                                                                                                                                                                                 |                           Default                            |
|-------------------------|-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|--------------------------------------------------------------|
| global.license          | License for the product. Get a license by following the steps from [here](https://kubeform.com/docs/latest/setup/install/enterprise#get-a-trial-license). <br> Example: <br> `helm install appscode/kubeform-provider-linode \` <br> `--set-file global.license=/path/to/license/file` <br> `or` <br> `helm install appscode/kubeform-provider-linode \` <br> `--set global.license=<license file content>` | `""`                                                         |
| global.registry         | Docker registry used to pull Kubeform operator image                                                                                                                                                                                                                                                                                                                                                        | `""`                                                         |
| global.registryFQDN     | Docker registry fqdn used to pull Kubeform related images. Set this to use docker registry hosted at ${registryFQDN}/${registry}/${image}                                                                                                                                                                                                                                                                   | `""`                                                         |
| global.imagePullSecrets | Specify an array of imagePullSecrets. Secrets must be manually created in the namespace. <br> Example: <br> `helm template charts/kubeform-provider-linode \` <br> `--set global.imagePullSecrets[0].name=sec0 \` <br> `--set global.imagePullSecrets[1].name=sec1`                                                                                                                                         | `[]`                                                         |
| global.skipCleaner      | Skip generating cleaner job YAML                                                                                                                                                                                                                                                                                                                                                                            | `false`                                                      |
| kubeform-provider       | Pass values to kubeform-provider chart                                                                                                                                                                                                                                                                                                                                                                      | `{"operator":{"tag":"v0.3.0"},"provider":{"name":"linode"}}` |
| crds.object             | If true, installs the kubeform-provider-linode-object-crds chart                                                                                                                                                                                                                                                                                                                                            | `false`                                                      |


Specify each parameter using the `--set key=value[,key=value]` argument to `helm install`. For example:

```console
$ helm install kubeform-provider-linode appscode/kubeform-provider-linode -n kubeform --set kubeform-provider={"operator":{"tag":"v0.3.0"},"provider":{"name":"linode"}}
```

Alternatively, a YAML file that specifies the values for the parameters can be provided while
installing the chart. For example:

```console
$ helm install kubeform-provider-linode appscode/kubeform-provider-linode -n kubeform --values values.yaml
```
