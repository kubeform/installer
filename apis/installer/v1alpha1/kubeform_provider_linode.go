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
	ResourceKindKubeformProviderLinode = "KubeformProviderLinode"
	ResourceKubeformProviderLinode     = "kubeformproviderlinode"
	ResourceKubeformProviderLinodes    = "kubeformproviderlinodes"
)

// KubeformProviderLinode defines the schama for one-click installer of Kubeform Provider Linode.

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// +kubebuilder:object:root=true
// +kubebuilder:resource:path=kubeformproviderlinodes,singular=kubeformproviderlinode,categories={kubeformproviderlinode,appscode}
type KubeformProviderLinode struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              KubeformProviderLinodeSpec `json:"spec,omitempty"`
}

// KubeformProviderLinodeSpec is the schema for kubeform-provider-linode chart values file
type KubeformProviderLinodeSpec struct {
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
	Object bool `json:"object"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// KubeformProviderLinodeList is a list of KubeformProviderLinodes
type KubeformProviderLinodeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of KubeformProviderLinode CRD objects
	Items []KubeformProviderLinode `json:"items,omitempty"`
}
