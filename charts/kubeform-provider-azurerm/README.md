# kubeform-provider-azurerm

[Kubeform Provider Azurerm Controller by AppsCode](https://github.com/kubeform) - Provision Azurerm resources using Terraform

## TL;DR;

```console
$ helm repo add appscode https://charts.appscode.com/stable/
$ helm repo update
$ helm install kubeform-provider-azurerm appscode/kubeform-provider-azurerm -n kubeform
```

## Introduction

This chart deploys a Kubeform Provider Azurerm Controller on a [Kubernetes](http://kubernetes.io) cluster using the [Helm](https://helm.sh) package manager.

## Prerequisites

- Kubernetes 1.16+

## Installing the Chart

To install the chart with the release name `kubeform-provider-azurerm`:

```console
$ helm install kubeform-provider-azurerm appscode/kubeform-provider-azurerm -n kubeform
```

The command deploys a Kubeform Provider Azurerm Controller on the Kubernetes cluster in the default configuration. The [configuration](#configuration) section lists the parameters that can be configured during installation.

> **Tip**: List all releases using `helm list`

## Uninstalling the Chart

To uninstall/delete the `kubeform-provider-azurerm`:

```console
$ helm delete kubeform-provider-azurerm -n kubeform
```

The command removes all the Kubernetes components associated with the chart and deletes the release.

## Configuration

The following table lists the configurable parameters of the `kubeform-provider-azurerm` chart and their default values.

|       Parameter        |                                 Description                                  |                            Default                            |
|------------------------|------------------------------------------------------------------------------|---------------------------------------------------------------|
| kubeform-provider      | Pass values to kubeform-provider chart                                       | `{"operator":{"tag":"v0.1.0"},"provider":{"name":"azurerm"}}` |
| crds.advanced          | If true, installs the kubeform-provider-azurerm-advanced-crds chart          | `false`                                                       |
| crds.analysis          | If true, installs the kubeform-provider-azurerm-analysis-crds chart          | `false`                                                       |
| crds.apimanagement     | If true, installs the kubeform-provider-azurerm-apimanagement-crds chart     | `false`                                                       |
| crds.app               | If true, installs the kubeform-provider-azurerm-app-crds chart               | `false`                                                       |
| crds.application       | If true, installs the kubeform-provider-azurerm-application-crds chart       | `false`                                                       |
| crds.attestation       | If true, installs the kubeform-provider-azurerm-attestation-crds chart       | `false`                                                       |
| crds.automation        | If true, installs the kubeform-provider-azurerm-automation-crds chart        | `false`                                                       |
| crds.availability      | If true, installs the kubeform-provider-azurerm-availability-crds chart      | `false`                                                       |
| crds.backup            | If true, installs the kubeform-provider-azurerm-backup-crds chart            | `false`                                                       |
| crds.bastion           | If true, installs the kubeform-provider-azurerm-bastion-crds chart           | `false`                                                       |
| crds.batch             | If true, installs the kubeform-provider-azurerm-batch-crds chart             | `false`                                                       |
| crds.blueprint         | If true, installs the kubeform-provider-azurerm-blueprint-crds chart         | `false`                                                       |
| crds.bot               | If true, installs the kubeform-provider-azurerm-bot-crds chart               | `false`                                                       |
| crds.cdn               | If true, installs the kubeform-provider-azurerm-cdn-crds chart               | `false`                                                       |
| crds.cognitive         | If true, installs the kubeform-provider-azurerm-cognitive-crds chart         | `false`                                                       |
| crds.communication     | If true, installs the kubeform-provider-azurerm-communication-crds chart     | `false`                                                       |
| crds.consumption       | If true, installs the kubeform-provider-azurerm-consumption-crds chart       | `false`                                                       |
| crds.container         | If true, installs the kubeform-provider-azurerm-container-crds chart         | `false`                                                       |
| crds.cosmosdb          | If true, installs the kubeform-provider-azurerm-cosmosdb-crds chart          | `false`                                                       |
| crds.cost              | If true, installs the kubeform-provider-azurerm-cost-crds chart              | `false`                                                       |
| crds.custom            | If true, installs the kubeform-provider-azurerm-custom-crds chart            | `false`                                                       |
| crds.dashboard         | If true, installs the kubeform-provider-azurerm-dashboard-crds chart         | `false`                                                       |
| crds.data              | If true, installs the kubeform-provider-azurerm-data-crds chart              | `false`                                                       |
| crds.database          | If true, installs the kubeform-provider-azurerm-database-crds chart          | `false`                                                       |
| crds.databox           | If true, installs the kubeform-provider-azurerm-databox-crds chart           | `false`                                                       |
| crds.databricks        | If true, installs the kubeform-provider-azurerm-databricks-crds chart        | `false`                                                       |
| crds.dedicatedhardware | If true, installs the kubeform-provider-azurerm-dedicatedhardware-crds chart | `false`                                                       |
| crds.dedicatedhost     | If true, installs the kubeform-provider-azurerm-dedicatedhost-crds chart     | `false`                                                       |
| crds.devspace          | If true, installs the kubeform-provider-azurerm-devspace-crds chart          | `false`                                                       |
| crds.devtest           | If true, installs the kubeform-provider-azurerm-devtest-crds chart           | `false`                                                       |
| crds.digital           | If true, installs the kubeform-provider-azurerm-digital-crds chart           | `false`                                                       |
| crds.disk              | If true, installs the kubeform-provider-azurerm-disk-crds chart              | `false`                                                       |
| crds.dns               | If true, installs the kubeform-provider-azurerm-dns-crds chart               | `false`                                                       |
| crds.eventgrid         | If true, installs the kubeform-provider-azurerm-eventgrid-crds chart         | `false`                                                       |
| crds.eventhub          | If true, installs the kubeform-provider-azurerm-eventhub-crds chart          | `false`                                                       |
| crds.expressroute      | If true, installs the kubeform-provider-azurerm-expressroute-crds chart      | `false`                                                       |
| crds.firewall          | If true, installs the kubeform-provider-azurerm-firewall-crds chart          | `false`                                                       |
| crds.frontdoor         | If true, installs the kubeform-provider-azurerm-frontdoor-crds chart         | `false`                                                       |
| crds.function          | If true, installs the kubeform-provider-azurerm-function-crds chart          | `false`                                                       |
| crds.hdinsight         | If true, installs the kubeform-provider-azurerm-hdinsight-crds chart         | `false`                                                       |
| crds.healthbot         | If true, installs the kubeform-provider-azurerm-healthbot-crds chart         | `false`                                                       |
| crds.healthcare        | If true, installs the kubeform-provider-azurerm-healthcare-crds chart        | `false`                                                       |
| crds.hpc               | If true, installs the kubeform-provider-azurerm-hpc-crds chart               | `false`                                                       |
| crds.image             | If true, installs the kubeform-provider-azurerm-image-crds chart             | `false`                                                       |
| crds.integration       | If true, installs the kubeform-provider-azurerm-integration-crds chart       | `false`                                                       |
| crds.iotcentral        | If true, installs the kubeform-provider-azurerm-iotcentral-crds chart        | `false`                                                       |
| crds.iothub            | If true, installs the kubeform-provider-azurerm-iothub-crds chart            | `false`                                                       |
| crds.iotsecurity       | If true, installs the kubeform-provider-azurerm-iotsecurity-crds chart       | `false`                                                       |
| crds.iottime           | If true, installs the kubeform-provider-azurerm-iottime-crds chart           | `false`                                                       |
| crds.ip                | If true, installs the kubeform-provider-azurerm-ip-crds chart                | `false`                                                       |
| crds.keyvault          | If true, installs the kubeform-provider-azurerm-keyvault-crds chart          | `false`                                                       |
| crds.kubernetescluster | If true, installs the kubeform-provider-azurerm-kubernetescluster-crds chart | `false`                                                       |
| crds.kusto             | If true, installs the kubeform-provider-azurerm-kusto-crds chart             | `false`                                                       |
| crds.lb                | If true, installs the kubeform-provider-azurerm-lb-crds chart                | `false`                                                       |
| crds.lighthouse        | If true, installs the kubeform-provider-azurerm-lighthouse-crds chart        | `false`                                                       |
| crds.linux             | If true, installs the kubeform-provider-azurerm-linux-crds chart             | `false`                                                       |
| crds.local             | If true, installs the kubeform-provider-azurerm-local-crds chart             | `false`                                                       |
| crds.loganalytics      | If true, installs the kubeform-provider-azurerm-loganalytics-crds chart      | `false`                                                       |
| crds.logicapp          | If true, installs the kubeform-provider-azurerm-logicapp-crds chart          | `false`                                                       |
| crds.machine           | If true, installs the kubeform-provider-azurerm-machine-crds chart           | `false`                                                       |
| crds.maintenance       | If true, installs the kubeform-provider-azurerm-maintenance-crds chart       | `false`                                                       |
| crds.managed           | If true, installs the kubeform-provider-azurerm-managed-crds chart           | `false`                                                       |
| crds.management        | If true, installs the kubeform-provider-azurerm-management-crds chart        | `false`                                                       |
| crds.maps              | If true, installs the kubeform-provider-azurerm-maps-crds chart              | `false`                                                       |
| crds.mariadb           | If true, installs the kubeform-provider-azurerm-mariadb-crds chart           | `false`                                                       |
| crds.marketplace       | If true, installs the kubeform-provider-azurerm-marketplace-crds chart       | `false`                                                       |
| crds.media             | If true, installs the kubeform-provider-azurerm-media-crds chart             | `false`                                                       |
| crds.monitor           | If true, installs the kubeform-provider-azurerm-monitor-crds chart           | `false`                                                       |
| crds.mssql             | If true, installs the kubeform-provider-azurerm-mssql-crds chart             | `false`                                                       |
| crds.mysql             | If true, installs the kubeform-provider-azurerm-mysql-crds chart             | `false`                                                       |
| crds.nat               | If true, installs the kubeform-provider-azurerm-nat-crds chart               | `false`                                                       |
| crds.netapp            | If true, installs the kubeform-provider-azurerm-netapp-crds chart            | `false`                                                       |
| crds.network           | If true, installs the kubeform-provider-azurerm-network-crds chart           | `false`                                                       |
| crds.notificationhub   | If true, installs the kubeform-provider-azurerm-notificationhub-crds chart   | `false`                                                       |
| crds.orchestrated      | If true, installs the kubeform-provider-azurerm-orchestrated-crds chart      | `false`                                                       |
| crds.packet            | If true, installs the kubeform-provider-azurerm-packet-crds chart            | `false`                                                       |
| crds.point             | If true, installs the kubeform-provider-azurerm-point-crds chart             | `false`                                                       |
| crds.policy            | If true, installs the kubeform-provider-azurerm-policy-crds chart            | `false`                                                       |
| crds.portal            | If true, installs the kubeform-provider-azurerm-portal-crds chart            | `false`                                                       |
| crds.postgresql        | If true, installs the kubeform-provider-azurerm-postgresql-crds chart        | `false`                                                       |
| crds.powerbi           | If true, installs the kubeform-provider-azurerm-powerbi-crds chart           | `false`                                                       |
| crds.private           | If true, installs the kubeform-provider-azurerm-private-crds chart           | `false`                                                       |
| crds.proximity         | If true, installs the kubeform-provider-azurerm-proximity-crds chart         | `false`                                                       |
| crds.publicip          | If true, installs the kubeform-provider-azurerm-publicip-crds chart          | `false`                                                       |
| crds.purview           | If true, installs the kubeform-provider-azurerm-purview-crds chart           | `false`                                                       |
| crds.recovery          | If true, installs the kubeform-provider-azurerm-recovery-crds chart          | `false`                                                       |
| crds.redis             | If true, installs the kubeform-provider-azurerm-redis-crds chart             | `false`                                                       |
| crds.relay             | If true, installs the kubeform-provider-azurerm-relay-crds chart             | `false`                                                       |
| crds.resource          | If true, installs the kubeform-provider-azurerm-resource-crds chart          | `false`                                                       |
| crds.role              | If true, installs the kubeform-provider-azurerm-role-crds chart              | `false`                                                       |
| crds.route             | If true, installs the kubeform-provider-azurerm-route-crds chart             | `false`                                                       |
| crds.search            | If true, installs the kubeform-provider-azurerm-search-crds chart            | `false`                                                       |
| crds.security          | If true, installs the kubeform-provider-azurerm-security-crds chart          | `false`                                                       |
| crds.sentinel          | If true, installs the kubeform-provider-azurerm-sentinel-crds chart          | `false`                                                       |
| crds.service           | If true, installs the kubeform-provider-azurerm-service-crds chart           | `false`                                                       |
| crds.servicebus        | If true, installs the kubeform-provider-azurerm-servicebus-crds chart        | `false`                                                       |
| crds.sharedimage       | If true, installs the kubeform-provider-azurerm-sharedimage-crds chart       | `false`                                                       |
| crds.signalr           | If true, installs the kubeform-provider-azurerm-signalr-crds chart           | `false`                                                       |
| crds.siterecovery      | If true, installs the kubeform-provider-azurerm-siterecovery-crds chart      | `false`                                                       |
| crds.snapshot          | If true, installs the kubeform-provider-azurerm-snapshot-crds chart          | `false`                                                       |
| crds.spatial           | If true, installs the kubeform-provider-azurerm-spatial-crds chart           | `false`                                                       |
| crds.spring            | If true, installs the kubeform-provider-azurerm-spring-crds chart            | `false`                                                       |
| crds.sql               | If true, installs the kubeform-provider-azurerm-sql-crds chart               | `false`                                                       |
| crds.ssh               | If true, installs the kubeform-provider-azurerm-ssh-crds chart               | `false`                                                       |
| crds.stack             | If true, installs the kubeform-provider-azurerm-stack-crds chart             | `false`                                                       |
| crds.static            | If true, installs the kubeform-provider-azurerm-static-crds chart            | `false`                                                       |
| crds.storage           | If true, installs the kubeform-provider-azurerm-storage-crds chart           | `false`                                                       |
| crds.stream            | If true, installs the kubeform-provider-azurerm-stream-crds chart            | `false`                                                       |
| crds.subnet            | If true, installs the kubeform-provider-azurerm-subnet-crds chart            | `false`                                                       |
| crds.subscription      | If true, installs the kubeform-provider-azurerm-subscription-crds chart      | `false`                                                       |
| crds.synapse           | If true, installs the kubeform-provider-azurerm-synapse-crds chart           | `false`                                                       |
| crds.template          | If true, installs the kubeform-provider-azurerm-template-crds chart          | `false`                                                       |
| crds.tenant            | If true, installs the kubeform-provider-azurerm-tenant-crds chart            | `false`                                                       |
| crds.trafficmanager    | If true, installs the kubeform-provider-azurerm-trafficmanager-crds chart    | `false`                                                       |
| crds.user              | If true, installs the kubeform-provider-azurerm-user-crds chart              | `false`                                                       |
| crds.virtual           | If true, installs the kubeform-provider-azurerm-virtual-crds chart           | `false`                                                       |
| crds.vmware            | If true, installs the kubeform-provider-azurerm-vmware-crds chart            | `false`                                                       |
| crds.vpn               | If true, installs the kubeform-provider-azurerm-vpn-crds chart               | `false`                                                       |
| crds.web               | If true, installs the kubeform-provider-azurerm-web-crds chart               | `false`                                                       |
| crds.windows           | If true, installs the kubeform-provider-azurerm-windows-crds chart           | `false`                                                       |


Specify each parameter using the `--set key=value[,key=value]` argument to `helm install`. For example:

```console
$ helm install kubeform-provider-azurerm appscode/kubeform-provider-azurerm -n kubeform --set kubeform-provider.provider.name=azurerm
```

Alternatively, a YAML file that specifies the values for the parameters can be provided while
installing the chart. For example:

```console
$ helm install kubeform-provider-azurerm appscode/kubeform-provider-azurerm -n kubeform --values values.yaml
```
