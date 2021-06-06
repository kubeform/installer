/*
Copyright The Kubeform Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

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
	ResourceKindKubeform = "Kubeform"
	ResourceKubeform     = "kubeform"
	ResourceKubeforms    = "kubeforms"
)

// Kubeform defines the schama for Kubeform provider controller Installer.

// +genclient
// +genclient:skipVerbs=updateStatus
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// +kubebuilder:object:root=true
// +kubebuilder:resource:path=kubeforms,singular=kubeform,categories={kubeform,appscode}
type Kubeform struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              KubeformSpec `json:"spec,omitempty"`
}

// KubeformSpec is the schema for Kubeform Operator values file
type KubeformSpec struct {
	//+optional
	NameOverride string `json:"nameOverride"`
	//+optional
	FullnameOverride string `json:"fullnameOverride"`
	ReplicaCount     int32  `json:"replicaCount"`
	// +optional
	License         string       `json:"license"`
	Operator        ContianerRef `json:"operator"`
	Cleaner         CleanerRef   `json:"cleaner"`
	ImagePullPolicy string       `json:"imagePullPolicy"`
	//+optional
	ImagePullSecrets []string `json:"imagePullSecrets"`
	//+optional
	CriticalAddon bool `json:"criticalAddon"`
	//+optional
	LogLevel int32 `json:"logLevel"`
	//+optional
	Annotations map[string]string `json:"annotations"`
	//+optional
	PodAnnotations map[string]string `json:"podAnnotations"`
	//+optional
	NodeSelector map[string]string `json:"nodeSelector"`
	// If specified, the pod's tolerations.
	// +optional
	Tolerations []core.Toleration `json:"tolerations"`
	// If specified, the pod's scheduling constraints
	// +optional
	Affinity *core.Affinity `json:"affinity"`
	// PodSecurityContext holds pod-level security attributes and common container settings.
	// Optional: Defaults to empty.  See type description for default values of each field.
	// +optional
	PodSecurityContext *core.PodSecurityContext `json:"podSecurityContext"`
	ServiceAccount     ServiceAccountSpec       `json:"serviceAccount"`
	//+optional
	SecretKey *string `json:"secretKey"`
	//+optional
	EnableAnalytics bool         `json:"enableAnalytics"`
	Proxy           ProxySpec    `json:"proxy"`
	Provider        ProviderSpec `json:"provider"`
}

type ContianerRef struct {
	Registry string `json:"registry"`
	Tag      string `json:"tag"`
	// Compute Resources required by the sidecar container.
	// +optional
	Resources core.ResourceRequirements `json:"resources"`
	// Security options the pod should run with.
	// +optional
	SecurityContext *core.SecurityContext `json:"securityContext"`
}

type CleanerRef struct {
	Registry   string `json:"registry"`
	Repository string `json:"repository"`
	Tag        string `json:"tag"`
	Skip       bool   `json:"skip"`
}

type ServiceAccountSpec struct {
	Create bool `json:"create"`
	//+optional
	Name *string `json:"name"`
	//+optional
	Annotations map[string]string `json:"annotations"`
}

type ProxySpec struct {
	//+optional
	HTTPS string `json:"https"`
	//+optional
	HTTP string `json:"http"`
	//+optional
	No string `json:"no"`
}

type ProviderSpec struct {
	Name string `json:"name"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// KubeformList is a list of Kubeforms
type KubeformList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Kubeform CRD objects
	Items []Kubeform `json:"items,omitempty"`
}
