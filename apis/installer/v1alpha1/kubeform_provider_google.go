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
	ResourceKindKubeformProviderGoogle = "KubeformProviderGoogle"
	ResourceKubeformProviderGoogle     = "kubeformprovidergoogle"
	ResourceKubeformProviderGoogles    = "kubeformprovidergoogles"
)

// KubeformProviderGoogle defines the schama for one-click installer of Kubeform Provider Google.

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// +kubebuilder:object:root=true
// +kubebuilder:resource:path=kubeformprovidergoogles,singular=kubeformprovidergoogle,categories={kubeformprovidergoogle,appscode}
type KubeformProviderGoogle struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              KubeformProviderGoogleSpec `json:"spec,omitempty"`
}

// KubeformProviderGoogleSpec is the schema for kubeform-provider-google chart values file
type KubeformProviderGoogleSpec struct {
	CRDs GoogleCRDs `json:"crds"`

	//+optional
	Controller KubeformProviderSpec `json:"kubeform-provider"`
}

type GoogleCRDs struct {
	Accesscontext          bool `json:"accesscontext"`
	Active                 bool `json:"active"`
	Apigee                 bool `json:"apigee"`
	Appengine              bool `json:"appengine"`
	Bigquery               bool `json:"bigquery"`
	Bigtable               bool `json:"bigtable"`
	Billingaccount         bool `json:"billingaccount"`
	Billingbudget          bool `json:"billingbudget"`
	Billingsubaccount      bool `json:"billingsubaccount"`
	Binaryauthorization    bool `json:"binaryauthorization"`
	Cloud                  bool `json:"cloud"`
	Cloudbuild             bool `json:"cloudbuild"`
	Cloudfunctionsfunction bool `json:"cloudfunctionsfunction"`
	Cloudiot               bool `json:"cloudiot"`
	Composer               bool `json:"composer"`
	Compute                bool `json:"compute"`
	Container              bool `json:"container"`
	Data                   bool `json:"data"`
	Dataflow               bool `json:"dataflow"`
	Dataproc               bool `json:"dataproc"`
	Datastore              bool `json:"datastore"`
	Deployment             bool `json:"deployment"`
	Dialogflow             bool `json:"dialogflow"`
	Dns                    bool `json:"dns"`
	Endpoints              bool `json:"endpoints"`
	Eventarc               bool `json:"eventarc"`
	Filestore              bool `json:"filestore"`
	Firestore              bool `json:"firestore"`
	Folder                 bool `json:"folder"`
	Game                   bool `json:"game"`
	Healthcare             bool `json:"healthcare"`
	Iap                    bool `json:"iap"`
	Identity               bool `json:"identity"`
	Kms                    bool `json:"kms"`
	Logging                bool `json:"logging"`
	Ml                     bool `json:"ml"`
	Monitoring             bool `json:"monitoring"`
	Network                bool `json:"network"`
	Notebooks              bool `json:"notebooks"`
	Organization           bool `json:"organization"`
	Os                     bool `json:"os"`
	Project                bool `json:"project"`
	Pubsub                 bool `json:"pubsub"`
	Redis                  bool `json:"redis"`
	Resource               bool `json:"resource"`
	Runtimeconfig          bool `json:"runtimeconfig"`
	Scc                    bool `json:"scc"`
	Secret                 bool `json:"secret"`
	Service                bool `json:"service"`
	Sourcereporepository   bool `json:"sourcereporepository"`
	Spanner                bool `json:"spanner"`
	Sql                    bool `json:"sql"`
	Storage                bool `json:"storage"`
	Tags                   bool `json:"tags"`
	Tpu                    bool `json:"tpu"`
	Vpc                    bool `json:"vpc"`
	Workflows              bool `json:"workflows"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// KubeformProviderGoogleList is a list of KubeformProviderGoogles
type KubeformProviderGoogleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of KubeformProviderGoogle CRD objects
	Items []KubeformProviderGoogle `json:"items,omitempty"`
}
