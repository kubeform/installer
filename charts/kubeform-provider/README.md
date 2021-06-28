# Kubeform Controller

[Kubeform Controller by AppsCode](https://github.com/kubeform) - Provision cloud resources using Kubernetes CRDs & Terraform

## TL;DR;

```console
$ helm repo add appscode https://charts.appscode.com/stable/
$ helm repo update
$ helm install kubeform-provider appscode/kubeform-provider -n kubeform
```

## Introduction

This chart deploys a Kubeform provider controller on a [Kubernetes](http://kubernetes.io) cluster using the [Helm](https://helm.sh) package manager.

## Prerequisites

- Kubernetes 1.16+

## Installing the Chart

To install the chart with the release name `kubeform-provider`:

```console
$ helm install kubeform-provider appscode/kubeform-provider -n kubeform
```

The command deploys a Kubeform provider controller on the Kubernetes cluster in the default configuration. The [configuration](#configuration) section lists the parameters that can be configured during installation.

> **Tip**: List all releases using `helm list`

## Uninstalling the Chart

To uninstall/delete the `kubeform-provider`:

```console
$ helm delete kubeform-provider -n kubeform
```

The command removes all the Kubernetes components associated with the chart and deletes the release.

## Configuration

The following table lists the configurable parameters of the `kubeform-provider` chart and their default values.

|            Parameter            |                                                                                                                                                                                           Description                                                                                                                                                                                           |    Default     |
|---------------------------------|-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|----------------|
| nameOverride                    | Overrides name template                                                                                                                                                                                                                                                                                                                                                                         | `""`           |
| fullnameOverride                | Overrides fullname template                                                                                                                                                                                                                                                                                                                                                                     | `""`           |
| replicaCount                    | Number of Kubeform operator replicas to create (only 1 is supported)                                                                                                                                                                                                                                                                                                                            | `1`            |
| license                         | License for the product. Get a license by following the steps from [here](https://kubeform.com/docs/latest/setup/install/community/#get-a-license). <br> Example: <br> `helm install appscode/kubeform-provider-enterprise \` <br> `--set-file license=/path/to/license/file` <br> `or` <br> `helm install appscode/kubeform-provider-enterprise \` <br> `--set license=<license file content>` | `""`           |
| registryFQDN                    | Docker registry fqdn used to pull Kubeform related images. Set this to use docker registry hosted at ${registryFQDN}/${registry}/${image}                                                                                                                                                                                                                                                       | `""`           |
| operator.registry               | Docker registry used to pull operator image                                                                                                                                                                                                                                                                                                                                                     | `kubeform`     |
| operator.tag                    | Operator container image tag                                                                                                                                                                                                                                                                                                                                                                    | `v0.3.0`       |
| operator.resources              | Compute Resources required by the operator container                                                                                                                                                                                                                                                                                                                                            | `{}`           |
| operator.securityContext        | Security options the operator container should run with                                                                                                                                                                                                                                                                                                                                         | `{}`           |
| cleaner.registry                | Docker registry used to pull Webhook cleaner image                                                                                                                                                                                                                                                                                                                                              | `appscode`     |
| cleaner.repository              | Webhook cleaner container image                                                                                                                                                                                                                                                                                                                                                                 | `kubectl`      |
| cleaner.tag                     | Webhook cleaner container image tag                                                                                                                                                                                                                                                                                                                                                             | `v1.16`        |
| cleaner.skip                    | Skip generating cleaner YAML                                                                                                                                                                                                                                                                                                                                                                    | `false`        |
| imagePullSecrets                | Specify an array of imagePullSecrets. Secrets must be manually created in the namespace. <br> Example: <br> `helm template charts/kubeform-provider \` <br> `--set imagePullSecrets[0].name=sec0 \` <br> `--set imagePullSecrets[1].name=sec1`                                                                                                                                                  | `[]`           |
| imagePullPolicy                 | Container image pull policy                                                                                                                                                                                                                                                                                                                                                                     | `IfNotPresent` |
| criticalAddon                   | If true, installs kubeform-provider operator as critical addon                                                                                                                                                                                                                                                                                                                                  | `false`        |
| logLevel                        | Log level for operator                                                                                                                                                                                                                                                                                                                                                                          | `3`            |
| annotations                     | Annotations applied to operator deployment                                                                                                                                                                                                                                                                                                                                                      | `{}`           |
| podAnnotations                  | Annotations passed to operator pod(s).                                                                                                                                                                                                                                                                                                                                                          | `{}`           |
| nodeSelector                    | Node labels for pod assignment                                                                                                                                                                                                                                                                                                                                                                  | ``             |
| tolerations                     | Tolerations for pod assignment                                                                                                                                                                                                                                                                                                                                                                  | `[]`           |
| affinity                        | Affinity rules for pod assignment                                                                                                                                                                                                                                                                                                                                                               | `{}`           |
| podSecurityContext              | Security options the operator pod should run with.                                                                                                                                                                                                                                                                                                                                              | `{}`           |
| serviceAccount.create           | Specifies whether a service account should be created                                                                                                                                                                                                                                                                                                                                           | `true`         |
| serviceAccount.annotations      | Annotations to add to the service account                                                                                                                                                                                                                                                                                                                                                       | `{}`           |
| serviceAccount.name             | The name of the service account to use. If not set and create is true, a name is generated using the fullname template                                                                                                                                                                                                                                                                          | ``             |
| secretKey                       | Specifies a base64-encoded key, of length 32 bytes when decoded. It is used to encrypt the state file.                                                                                                                                                                                                                                                                                          | ``             |
| enableAnalytics                 | If true, sends usage analytics                                                                                                                                                                                                                                                                                                                                                                  | `true`         |
| proxy.https                     | To configure HTTPS_PROXY environment variable specify <ip_address>:<port>                                                                                                                                                                                                                                                                                                                       | `''`           |
| proxy.http                      | To configure HTTP_PROXY environment variable specify <ip_address>:<port>                                                                                                                                                                                                                                                                                                                        | `''`           |
| proxy.no                        | To configure NO_PROXY environment variable specify <ip_address>:<port> By default exclude Kubernetes apiserver internal IP.                                                                                                                                                                                                                                                                     | `'10.43.0.1'`  |
| provider.name                   |                                                                                                                                                                                                                                                                                                                                                                                                 | `linode`       |
| webhook.enableValidatingWebhook | If true, validating webhook is configured for KubeDB CRDss                                                                                                                                                                                                                                                                                                                                      | `true`         |
| webhook.servingCerts.generate   | If true, generates on install/upgrade the certs that is required for the webhook server.                                                                                                                                                                                                                                                                                                        | `true`         |
| webhook.servingCerts.caCrt      | CA certificate used by serving certificate of webhook server.                                                                                                                                                                                                                                                                                                                                   | `""`           |
| webhook.servingCerts.serverCrt  | Serving certificate used by webhook server.                                                                                                                                                                                                                                                                                                                                                     | `""`           |
| webhook.servingCerts.serverKey  | Private key for the serving certificate used by webhook server.                                                                                                                                                                                                                                                                                                                                 | `""`           |


Specify each parameter using the `--set key=value[,key=value]` argument to `helm install`. For example:

```console
$ helm install kubeform-provider appscode/kubeform-provider -n kubeform --set replicaCount=1
```

Alternatively, a YAML file that specifies the values for the parameters can be provided while
installing the chart. For example:

```console
$ helm install kubeform-provider appscode/kubeform-provider -n kubeform --values values.yaml
```