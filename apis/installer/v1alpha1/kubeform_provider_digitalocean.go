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
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	ResourceKindKubeformProviderDigitalocean = "KubeformProviderDigitalocean"
	ResourceKubeformProviderDigitalocean     = "kubeformproviderdigitalocean"
	ResourceKubeformProviderDigitaloceans    = "kubeformproviderdigitaloceans"
)

// KubeformProviderDigitalocean defines the schama for one-click installer of Kubeform Provider Digitalocean.

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// +kubebuilder:object:root=true
// +kubebuilder:resource:path=kubeformproviderdigitaloceans,singular=kubeformproviderdigitalocean,categories={kubeformproviderdigitalocean,appscode}
type KubeformProviderDigitalocean struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              KubeformProviderDigitaloceanSpec `json:"spec,omitempty"`
}

// KubeformProviderDigitaloceanSpec is the schema for kubeform-provider-digitalocean chart values file
type KubeformProviderDigitaloceanSpec struct {
	Global GlobalValues `json:"global"`

	CRDs CRDs `json:"crds"`

	//+optional
	Controller KubeformProviderSpec `json:"kubeform-provider"`
}

type GlobalValues struct {
	License      string `json:"license"`
	Registry     string `json:"registry"`
	RegistryFQDN string `json:"registryFQDN"`
	//+optional
	ImagePullSecrets []core.LocalObjectReference `json:"imagePullSecrets"`
	SkipCleaner      bool                        `json:"skipCleaner"`
}

type CRDs struct {
	App               bool `json:"app"`
	Cdn               bool `json:"cdn"`
	Certificate       bool `json:"certificate"`
	Containerregistry bool `json:"containerregistry"`
	Custom            bool `json:"custom"`
	Database          bool `json:"database"`
	Domain            bool `json:"domain"`
	Droplet           bool `json:"droplet"`
	Firewall          bool `json:"firewall"`
	Floatingip        bool `json:"floatingip"`
	Kubernetes        bool `json:"kubernetes"`
	Loadbalancer      bool `json:"loadbalancer"`
	Project           bool `json:"project"`
	Record            bool `json:"record"`
	Spacesbucket      bool `json:"spacesbucket"`
	Ssh               bool `json:"ssh"`
	Tag               bool `json:"tag"`
	Volume            bool `json:"volume"`
	Vpc               bool `json:"vpc"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// KubeformProviderDigitaloceanList is a list of KubeformProviderDigitaloceans
type KubeformProviderDigitaloceanList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of KubeformProviderDigitalocean CRD objects
	Items []KubeformProviderDigitalocean `json:"items,omitempty"`
}
