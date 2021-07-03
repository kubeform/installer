# kubeform-provider-google

[Kubeform Provider Google Controller by AppsCode](https://github.com/kubeform) - Provision Google resources using Terraform

## TL;DR;

```console
$ helm repo add appscode https://charts.appscode.com/stable/
$ helm repo update
$ helm install kubeform-provider-google appscode/kubeform-provider-google -n kubeform
```

## Introduction

This chart deploys a Kubeform Provider Google Controller on a [Kubernetes](http://kubernetes.io) cluster using the [Helm](https://helm.sh) package manager.

## Prerequisites

- Kubernetes 1.16+

## Installing the Chart

To install the chart with the release name `kubeform-provider-google`:

```console
$ helm install kubeform-provider-google appscode/kubeform-provider-google -n kubeform
```

The command deploys a Kubeform Provider Google Controller on the Kubernetes cluster in the default configuration. The [configuration](#configuration) section lists the parameters that can be configured during installation.

> **Tip**: List all releases using `helm list`

## Uninstalling the Chart

To uninstall/delete the `kubeform-provider-google`:

```console
$ helm delete kubeform-provider-google -n kubeform
```

The command removes all the Kubernetes components associated with the chart and deletes the release.

## Configuration

The following table lists the configurable parameters of the `kubeform-provider-google` chart and their default values.

|          Parameter          |                                   Description                                    |                           Default                            |
|-----------------------------|----------------------------------------------------------------------------------|--------------------------------------------------------------|
| kubeform-provider           | Pass values to kubeform-provider chart                                           | `{"operator":{"tag":"v0.3.0"},"provider":{"name":"google"}}` |
| crds.accesscontext          | If true, installs the kubeform-provider-google-accesscontext-crds chart          | `false`                                                      |
| crds.active                 | If true, installs the kubeform-provider-google-active-crds chart                 | `false`                                                      |
| crds.apigee                 | If true, installs the kubeform-provider-google-apigee-crds chart                 | `false`                                                      |
| crds.appengine              | If true, installs the kubeform-provider-google-appengine-crds chart              | `false`                                                      |
| crds.bigquery               | If true, installs the kubeform-provider-google-bigquery-crds chart               | `false`                                                      |
| crds.bigtable               | If true, installs the kubeform-provider-google-bigtable-crds chart               | `false`                                                      |
| crds.billingaccount         | If true, installs the kubeform-provider-google-billingaccount-crds chart         | `false`                                                      |
| crds.billingbudget          | If true, installs the kubeform-provider-google-billingbudget-crds chart          | `false`                                                      |
| crds.billingsubaccount      | If true, installs the kubeform-provider-google-billingsubaccount-crds chart      | `false`                                                      |
| crds.binaryauthorization    | If true, installs the kubeform-provider-google-binaryauthorization-crds chart    | `false`                                                      |
| crds.cloud                  | If true, installs the kubeform-provider-google-cloud-crds chart                  | `false`                                                      |
| crds.cloudbuild             | If true, installs the kubeform-provider-google-cloudbuild-crds chart             | `false`                                                      |
| crds.cloudfunctionsfunction | If true, installs the kubeform-provider-google-cloudfunctionsfunction-crds chart | `false`                                                      |
| crds.cloudiot               | If true, installs the kubeform-provider-google-cloudiot-crds chart               | `false`                                                      |
| crds.composer               | If true, installs the kubeform-provider-google-composer-crds chart               | `false`                                                      |
| crds.compute                | If true, installs the kubeform-provider-google-compute-crds chart                | `false`                                                      |
| crds.container              | If true, installs the kubeform-provider-google-container-crds chart              | `false`                                                      |
| crds.data                   | If true, installs the kubeform-provider-google-data-crds chart                   | `false`                                                      |
| crds.dataflow               | If true, installs the kubeform-provider-google-dataflow-crds chart               | `false`                                                      |
| crds.dataproc               | If true, installs the kubeform-provider-google-dataproc-crds chart               | `false`                                                      |
| crds.datastore              | If true, installs the kubeform-provider-google-datastore-crds chart              | `false`                                                      |
| crds.deployment             | If true, installs the kubeform-provider-google-deployment-crds chart             | `false`                                                      |
| crds.dialogflow             | If true, installs the kubeform-provider-google-dialogflow-crds chart             | `false`                                                      |
| crds.dns                    | If true, installs the kubeform-provider-google-dns-crds chart                    | `false`                                                      |
| crds.endpoints              | If true, installs the kubeform-provider-google-endpoints-crds chart              | `false`                                                      |
| crds.eventarc               | If true, installs the kubeform-provider-google-eventarc-crds chart               | `false`                                                      |
| crds.filestore              | If true, installs the kubeform-provider-google-filestore-crds chart              | `false`                                                      |
| crds.firestore              | If true, installs the kubeform-provider-google-firestore-crds chart              | `false`                                                      |
| crds.folder                 | If true, installs the kubeform-provider-google-folder-crds chart                 | `false`                                                      |
| crds.game                   | If true, installs the kubeform-provider-google-game-crds chart                   | `false`                                                      |
| crds.healthcare             | If true, installs the kubeform-provider-google-healthcare-crds chart             | `false`                                                      |
| crds.iap                    | If true, installs the kubeform-provider-google-iap-crds chart                    | `false`                                                      |
| crds.identity               | If true, installs the kubeform-provider-google-identity-crds chart               | `false`                                                      |
| crds.kms                    | If true, installs the kubeform-provider-google-kms-crds chart                    | `false`                                                      |
| crds.logging                | If true, installs the kubeform-provider-google-logging-crds chart                | `false`                                                      |
| crds.memcache               | If true, installs the kubeform-provider-google-memcache-crds chart               | `false`                                                      |
| crds.ml                     | If true, installs the kubeform-provider-google-ml-crds chart                     | `false`                                                      |
| crds.monitoring             | If true, installs the kubeform-provider-google-monitoring-crds chart             | `false`                                                      |
| crds.network                | If true, installs the kubeform-provider-google-network-crds chart                | `false`                                                      |
| crds.notebooks              | If true, installs the kubeform-provider-google-notebooks-crds chart              | `false`                                                      |
| crds.organization           | If true, installs the kubeform-provider-google-organization-crds chart           | `false`                                                      |
| crds.os                     | If true, installs the kubeform-provider-google-os-crds chart                     | `false`                                                      |
| crds.project                | If true, installs the kubeform-provider-google-project-crds chart                | `false`                                                      |
| crds.pubsub                 | If true, installs the kubeform-provider-google-pubsub-crds chart                 | `false`                                                      |
| crds.redis                  | If true, installs the kubeform-provider-google-redis-crds chart                  | `false`                                                      |
| crds.resource               | If true, installs the kubeform-provider-google-resource-crds chart               | `false`                                                      |
| crds.runtimeconfig          | If true, installs the kubeform-provider-google-runtimeconfig-crds chart          | `false`                                                      |
| crds.scc                    | If true, installs the kubeform-provider-google-scc-crds chart                    | `false`                                                      |
| crds.secret                 | If true, installs the kubeform-provider-google-secret-crds chart                 | `false`                                                      |
| crds.service                | If true, installs the kubeform-provider-google-service-crds chart                | `false`                                                      |
| crds.sourcereporepository   | If true, installs the kubeform-provider-google-sourcereporepository-crds chart   | `false`                                                      |
| crds.spanner                | If true, installs the kubeform-provider-google-spanner-crds chart                | `false`                                                      |
| crds.sql                    | If true, installs the kubeform-provider-google-sql-crds chart                    | `false`                                                      |
| crds.storage                | If true, installs the kubeform-provider-google-storage-crds chart                | `false`                                                      |
| crds.tags                   | If true, installs the kubeform-provider-google-tags-crds chart                   | `false`                                                      |
| crds.tpu                    | If true, installs the kubeform-provider-google-tpu-crds chart                    | `false`                                                      |
| crds.vertex                 | If true, installs the kubeform-provider-google-vertex-crds chart                 | `false`                                                      |
| crds.vpc                    | If true, installs the kubeform-provider-google-vpc-crds chart                    | `false`                                                      |
| crds.workflows              | If true, installs the kubeform-provider-google-workflows-crds chart              | `false`                                                      |


Specify each parameter using the `--set key=value[,key=value]` argument to `helm install`. For example:

```console
$ helm install kubeform-provider-google appscode/kubeform-provider-google -n kubeform --set kubeform-provider.provider.name=google
```

Alternatively, a YAML file that specifies the values for the parameters can be provided while
installing the chart. For example:

```console
$ helm install kubeform-provider-google appscode/kubeform-provider-google -n kubeform --values values.yaml
```
