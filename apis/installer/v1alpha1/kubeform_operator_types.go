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
	ResourceKindKubeformOperator = "KubeformOperator"
	ResourceKubeformOperator     = "kubeformoperator"
	ResourceKubeformOperators    = "kubeformoperators"
)

// KubeformOperator defines the schama for KubeVault Operator Installer.

// +genclient
// +genclient:skipVerbs=updateStatus
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// +kubebuilder:object:root=true
// +kubebuilder:resource:path=kubeformoperators,singular=kubeformoperator,categories={kubeform,appscode}
type KubeformOperator struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Spec              KubeformOperatorSpec `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`
}

// KubeformOperatorSpec is the schema for Kubeform Operator values file
type KubeformOperatorSpec struct {
	//+optional
	NameOverride string `json:"nameOverride" protobuf:"bytes,1,opt,name=nameOverride"`
	//+optional
	FullnameOverride string       `json:"fullnameOverride" protobuf:"bytes,2,opt,name=fullnameOverride"`
	ReplicaCount     int32        `json:"replicaCount" protobuf:"varint,3,opt,name=replicaCount"`
	Operator         ContianerRef `json:"operator" protobuf:"bytes,4,opt,name=operator"`
	ImagePullPolicy  string       `json:"imagePullPolicy" protobuf:"bytes,5,opt,name=imagePullPolicy"`
	//+optional
	ImagePullSecrets []string `json:"imagePullSecrets" protobuf:"bytes,6,rep,name=imagePullSecrets"`
	//+optional
	CriticalAddon bool `json:"criticalAddon" protobuf:"varint,7,opt,name=criticalAddon"`
	//+optional
	LogLevel int32 `json:"logLevel" protobuf:"varint,8,opt,name=logLevel"`
	//+optional
	Annotations map[string]string `json:"annotations" protobuf:"bytes,9,rep,name=annotations"`
	//+optional
	PodAnnotations map[string]string `json:"podAnnotations" protobuf:"bytes,10,rep,name=podAnnotations"`
	//+optional
	NodeSelector map[string]string `json:"nodeSelector" protobuf:"bytes,11,rep,name=nodeSelector"`
	// If specified, the pod's tolerations.
	// +optional
	Tolerations []core.Toleration `json:"tolerations" protobuf:"bytes,12,rep,name=tolerations"`
	// If specified, the pod's scheduling constraints
	// +optional
	Affinity *core.Affinity `json:"affinity" protobuf:"bytes,13,opt,name=affinity"`
	// PodSecurityContext holds pod-level security attributes and common container settings.
	// Optional: Defaults to empty.  See type description for default values of each field.
	// +optional
	PodSecurityContext *core.PodSecurityContext `json:"podSecurityContext" protobuf:"bytes,14,opt,name=podSecurityContext"`
	ServiceAccount     ServiceAccountSpec       `json:"serviceAccount" protobuf:"bytes,15,opt,name=serviceAccount"`
	//+optional
	SecretKey *string `json:"secretKey" protobuf:"bytes,16,opt,name=secretKey"`
	//+optional
	EnableAnalytics bool `json:"enableAnalytics" protobuf:"varint,17,opt,name=enableAnalytics"`
}

type ImageRef struct {
	Registry   string `json:"registry" protobuf:"bytes,1,opt,name=registry"`
	Repository string `json:"repository" protobuf:"bytes,2,opt,name=repository"`
	Tag        string `json:"tag" protobuf:"bytes,3,opt,name=tag"`
}

type ContianerRef struct {
	ImageRef `json:",inline" protobuf:"bytes,1,opt,name=imageRef"`
	// Compute Resources required by the sidecar container.
	// +optional
	Resources core.ResourceRequirements `json:"resources" protobuf:"bytes,2,opt,name=resources"`
	// Security options the pod should run with.
	// +optional
	SecurityContext *core.SecurityContext `json:"securityContext" protobuf:"bytes,3,opt,name=securityContext"`
}

type ServiceAccountSpec struct {
	Create bool `json:"create" protobuf:"varint,1,opt,name=create"`
	//+optional
	Name *string `json:"name" protobuf:"bytes,2,opt,name=name"`
	//+optional
	Annotations map[string]string `json:"annotations" protobuf:"bytes,3,rep,name=annotations"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// KubeformOperatorList is a list of KubeformOperators
type KubeformOperatorList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	// Items is a list of KubeformOperator CRD objects
	Items []KubeformOperator `json:"items,omitempty" protobuf:"bytes,2,rep,name=items"`
}
