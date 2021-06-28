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
	ResourceKindKubeformProvider{{ .Provider | camelcase }} = "KubeformProvider{{ .Provider | camelcase }}"
	ResourceKubeformProvider{{ .Provider | camelcase }}     = "kubeformprovider{{ .Provider }}"
	ResourceKubeformProvider{{ .Provider | camelcase }}s    = "kubeformprovider{{ .Provider }}s"
)

// KubeformProvider{{ .Provider | camelcase }} defines the schama for one-click installer of Kubeform Provider {{ .Provider | camelcase }}.

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// +kubebuilder:object:root=true
// +kubebuilder:resource:path=kubeformprovider{{ .Provider }}s,singular=kubeformprovider{{ .Provider }},categories={kubeformprovider{{ .Provider }},appscode}
type KubeformProvider{{ .Provider | camelcase }} struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              KubeformProvider{{ .Provider | camelcase }}Spec `json:"spec,omitempty"`
}

// KubeformProvider{{ .Provider | camelcase }}Spec is the schema for kubeform-provider-{{ .Provider }} chart values file
type KubeformProvider{{ .Provider | camelcase }}Spec struct {
	CRDs {{ .Provider | camelcase }}CRDs `json:"crds"`

	//+optional
	Controller KubeformProviderSpec `json:"kubeform-provider"`
}

type {{ .Provider | camelcase }}CRDs struct {
	{{ range .GIDs }}
	{{- . | camelcase }} bool `json:"{{.}}"`
	{{ end }}
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// KubeformProvider{{ .Provider | camelcase }}List is a list of KubeformProvider{{ .Provider | camelcase }}s
type KubeformProvider{{ .Provider | camelcase }}List struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of KubeformProvider{{ .Provider | camelcase }} CRD objects
	Items []KubeformProvider{{ .Provider | camelcase }} `json:"items,omitempty"`
}
