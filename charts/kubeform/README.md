# Kubeform

[Kubeform by AppsCode](https://github.com/kubeform) - Provision cloud resources using Kubernetes CRDs & Terraform

## TL;DR;

```console
$ helm repo add appscode https://charts.appscode.com/stable/
$ helm repo update
$ helm install kfc appscode/kubeform -n kube-system
```

## Introduction

This chart deploys a Kubeform operator on a [Kubernetes](http://kubernetes.io) cluster using the [Helm](https://helm.sh) package manager.

## Prerequisites

- Kubernetes 1.11+

## Installing the Chart

To install the chart with the release name `kfc`:

```console
$ helm install kfc appscode/kubeform -n kube-system
```

The command deploys a Kubeform operator on the Kubernetes cluster in the default configuration. The [configuration](#configuration) section lists the parameters that can be configured during installation.

> **Tip**: List all releases using `helm list`

## Uninstalling the Chart

To uninstall/delete the `kfc`:

```console
$ helm delete kfc -n kube-system
```

The command removes all the Kubernetes components associated with the chart and deletes the release.
## Configuration

The following table lists the configurable parameters of the `kubeform` chart and their default values.

|         Parameter          |                                                                                                                                                                                  Description                                                                                                                                                                                  |    Default     |
|----------------------------|-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|----------------|
| nameOverride               | Overrides name template                                                                                                                                                                                                                                                                                                                                                       | `""`           |
| fullnameOverride           | Overrides fullname template                                                                                                                                                                                                                                                                                                                                                   | `""`           |
| replicaCount               | Number of Kubeform operator replicas to create (only 1 is supported)                                                                                                                                                                                                                                                                                                          | `1`            |
| license                    | License for the product. Get a license by following the steps from [here](https://kubeform.com/docs/latest/setup/install/community/#get-a-license). <br> Example: <br> `helm install appscode/kubeform-enterprise \` <br> `--set-file license=/path/to/license/file` <br> `or` <br> `helm install appscode/kubeform-enterprise \` <br> `--set license=<license file content>` | `""`           |
| operator.registry          | Docker registry used to pull operator image                                                                                                                                                                                                                                                                                                                                   | `kubeform`     |
| operator.repository        | Name of operator container image                                                                                                                                                                                                                                                                                                                                              | `kfc`          |
| operator.tag               | Operator container image tag                                                                                                                                                                                                                                                                                                                                                  | `v0.3.0`       |
| operator.resources         | Compute Resources required by the operator container                                                                                                                                                                                                                                                                                                                          | `{}`           |
| operator.securityContext   | Security options the operator container should run with                                                                                                                                                                                                                                                                                                                       | `{}`           |
| imagePullSecrets           | Specify an array of imagePullSecrets. Secrets must be manually created in the namespace. <br> Example: <br> `helm template charts/kubeform \` <br> `--set imagePullSecrets[0].name=sec0 \` <br> `--set imagePullSecrets[1].name=sec1`                                                                                                                                         | `[]`           |
| imagePullPolicy            | Container image pull policy                                                                                                                                                                                                                                                                                                                                                   | `IfNotPresent` |
| criticalAddon              | If true, installs kubeform operator as critical addon                                                                                                                                                                                                                                                                                                                         | `false`        |
| logLevel                   | Log level for operator                                                                                                                                                                                                                                                                                                                                                        | `3`            |
| annotations                | Annotations applied to operator deployment                                                                                                                                                                                                                                                                                                                                    | `{}`           |
| podAnnotations             | Annotations passed to operator pod(s).                                                                                                                                                                                                                                                                                                                                        | `{}`           |
| nodeSelector               | Node labels for pod assignment                                                                                                                                                                                                                                                                                                                                                | ``             |
| tolerations                | Tolerations for pod assignment                                                                                                                                                                                                                                                                                                                                                | `[]`           |
| affinity                   | Affinity rules for pod assignment                                                                                                                                                                                                                                                                                                                                             | `{}`           |
| podSecurityContext         | Security options the operator pod should run with.                                                                                                                                                                                                                                                                                                                            | `{}`           |
| serviceAccount.create      | Specifies whether a service account should be created                                                                                                                                                                                                                                                                                                                         | `true`         |
| serviceAccount.annotations | Annotations to add to the service account                                                                                                                                                                                                                                                                                                                                     | `{}`           |
| serviceAccount.name        | The name of the service account to use. If not set and create is true, a name is generated using the fullname template                                                                                                                                                                                                                                                        | ``             |
| secretKey                  | Specifies a base64-encoded key, of length 32 bytes when decoded. It is used to encrypt the state file.                                                                                                                                                                                                                                                                        | ``             |
| enableAnalytics            | If true, sends usage analytics                                                                                                                                                                                                                                                                                                                                                | `true`         |
| proxy.https                | To configure HTTPS_PROXY environment variable specify <ip_address>:<port>                                                                                                                                                                                                                                                                                                     | `''`           |
| proxy.http                 | To configure HTTP_PROXY environment variable specify <ip_address>:<port>                                                                                                                                                                                                                                                                                                      | `''`           |
| proxy.no                   | To configure NO_PROXY environment variable specify <ip_address>:<port> By default exclude Kubernetes apiserver internal IP.                                                                                                                                                                                                                                                   | `'10.43.0.1'`  |


Specify each parameter using the `--set key=value[,key=value]` argument to `helm install`. For example:

```console
$ helm install kfc appscode/kubeform -n kube-system --set replicaCount=1
```

Alternatively, a YAML file that specifies the values for the parameters can be provided while
installing the chart. For example:

```console
$ helm install kfc appscode/kubeform -n kube-system --values values.yaml
```
