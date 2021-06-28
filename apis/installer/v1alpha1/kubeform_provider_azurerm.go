/*
Copyright AppsCode Inc. and Contributors

Licensed under the AppsCode Community License 1.0.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    https://github.com/appscode/licenses/raw/1.0.0/AppsCode-Community-1.0.0.md

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	ResourceKindKubeformProviderAzurerm = "KubeformProviderAzurerm"
	ResourceKubeformProviderAzurerm     = "kubeformproviderazurerm"
	ResourceKubeformProviderAzurerms    = "kubeformproviderazurerms"
)

// KubeformProviderAzurerm defines the schama for one-click installer of Kubeform Provider Azurerm.

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// +kubebuilder:object:root=true
// +kubebuilder:resource:path=kubeformproviderazurerms,singular=kubeformproviderazurerm,categories={kubeformproviderazurerm,appscode}
type KubeformProviderAzurerm struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              KubeformProviderAzurermSpec `json:"spec,omitempty"`
}

// KubeformProviderAzurermSpec is the schema for kubeform-provider-azurerm chart values file
type KubeformProviderAzurermSpec struct {
	CRDs AzurermCRDs `json:"crds"`

	//+optional
	Controller KubeformProviderSpec `json:"kubeform-provider"`
}

type AzurermCRDs struct {
	Advanced          bool `json:"advanced"`
	Analysis          bool `json:"analysis"`
	Apimanagement     bool `json:"apimanagement"`
	App               bool `json:"app"`
	Application       bool `json:"application"`
	Attestation       bool `json:"attestation"`
	Automation        bool `json:"automation"`
	Availability      bool `json:"availability"`
	Backup            bool `json:"backup"`
	Bastion           bool `json:"bastion"`
	Batch             bool `json:"batch"`
	Blueprint         bool `json:"blueprint"`
	Bot               bool `json:"bot"`
	Cdn               bool `json:"cdn"`
	Cognitive         bool `json:"cognitive"`
	Communication     bool `json:"communication"`
	Consumption       bool `json:"consumption"`
	Container         bool `json:"container"`
	Cosmosdb          bool `json:"cosmosdb"`
	Cost              bool `json:"cost"`
	Custom            bool `json:"custom"`
	Dashboard         bool `json:"dashboard"`
	Data              bool `json:"data"`
	Database          bool `json:"database"`
	Databox           bool `json:"databox"`
	Databricks        bool `json:"databricks"`
	Dedicatedhardware bool `json:"dedicatedhardware"`
	Dedicatedhost     bool `json:"dedicatedhost"`
	Devspace          bool `json:"devspace"`
	Devtest           bool `json:"devtest"`
	Digital           bool `json:"digital"`
	Disk              bool `json:"disk"`
	Dns               bool `json:"dns"`
	Eventgrid         bool `json:"eventgrid"`
	Eventhub          bool `json:"eventhub"`
	Expressroute      bool `json:"expressroute"`
	Firewall          bool `json:"firewall"`
	Frontdoor         bool `json:"frontdoor"`
	Function          bool `json:"function"`
	Hdinsight         bool `json:"hdinsight"`
	Healthbot         bool `json:"healthbot"`
	Healthcare        bool `json:"healthcare"`
	Hpc               bool `json:"hpc"`
	Image             bool `json:"image"`
	Integration       bool `json:"integration"`
	Iotcentral        bool `json:"iotcentral"`
	Iothub            bool `json:"iothub"`
	Iotsecurity       bool `json:"iotsecurity"`
	Iottime           bool `json:"iottime"`
	Ip                bool `json:"ip"`
	Keyvault          bool `json:"keyvault"`
	Kubernetescluster bool `json:"kubernetescluster"`
	Kusto             bool `json:"kusto"`
	Lb                bool `json:"lb"`
	Lighthouse        bool `json:"lighthouse"`
	Linux             bool `json:"linux"`
	Local             bool `json:"local"`
	Loganalytics      bool `json:"loganalytics"`
	Logicapp          bool `json:"logicapp"`
	Machine           bool `json:"machine"`
	Maintenance       bool `json:"maintenance"`
	Managed           bool `json:"managed"`
	Management        bool `json:"management"`
	Maps              bool `json:"maps"`
	Mariadb           bool `json:"mariadb"`
	Marketplace       bool `json:"marketplace"`
	Media             bool `json:"media"`
	Monitor           bool `json:"monitor"`
	Mssql             bool `json:"mssql"`
	Mysql             bool `json:"mysql"`
	Nat               bool `json:"nat"`
	Netapp            bool `json:"netapp"`
	Network           bool `json:"network"`
	Notificationhub   bool `json:"notificationhub"`
	Orchestrated      bool `json:"orchestrated"`
	Packet            bool `json:"packet"`
	Point             bool `json:"point"`
	Policy            bool `json:"policy"`
	Postgresql        bool `json:"postgresql"`
	Powerbi           bool `json:"powerbi"`
	Private           bool `json:"private"`
	Proximity         bool `json:"proximity"`
	Publicip          bool `json:"publicip"`
	Purview           bool `json:"purview"`
	Recovery          bool `json:"recovery"`
	Redis             bool `json:"redis"`
	Relay             bool `json:"relay"`
	Resource          bool `json:"resource"`
	Role              bool `json:"role"`
	Route             bool `json:"route"`
	Search            bool `json:"search"`
	Security          bool `json:"security"`
	Sentinel          bool `json:"sentinel"`
	Service           bool `json:"service"`
	Servicebus        bool `json:"servicebus"`
	Sharedimage       bool `json:"sharedimage"`
	Signalr           bool `json:"signalr"`
	Siterecovery      bool `json:"siterecovery"`
	Snapshot          bool `json:"snapshot"`
	Spatial           bool `json:"spatial"`
	Spring            bool `json:"spring"`
	Sql               bool `json:"sql"`
	Ssh               bool `json:"ssh"`
	Stack             bool `json:"stack"`
	Static            bool `json:"static"`
	Storage           bool `json:"storage"`
	Stream            bool `json:"stream"`
	Subnet            bool `json:"subnet"`
	Subscription      bool `json:"subscription"`
	Synapse           bool `json:"synapse"`
	Template          bool `json:"template"`
	Tenant            bool `json:"tenant"`
	Trafficmanager    bool `json:"trafficmanager"`
	User              bool `json:"user"`
	Virtual           bool `json:"virtual"`
	Vmware            bool `json:"vmware"`
	Vpn               bool `json:"vpn"`
	Web               bool `json:"web"`
	Windows           bool `json:"windows"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// KubeformProviderAzurermList is a list of KubeformProviderAzurerms
type KubeformProviderAzurermList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of KubeformProviderAzurerm CRD objects
	Items []KubeformProviderAzurerm `json:"items,omitempty"`
}
