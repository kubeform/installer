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
	ResourceKindKubeformProviderAws = "KubeformProviderAws"
	ResourceKubeformProviderAws     = "kubeformprovideraws"
	ResourceKubeformProviderAwss    = "kubeformproviderawss"
)

// KubeformProviderAws defines the schama for one-click installer of Kubeform Provider Aws.

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// +kubebuilder:object:root=true
// +kubebuilder:resource:path=kubeformproviderawss,singular=kubeformprovideraws,categories={kubeformprovideraws,appscode}
type KubeformProviderAws struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              KubeformProviderAwsSpec `json:"spec,omitempty"`
}

// KubeformProviderAwsSpec is the schema for kubeform-provider-aws chart values file
type KubeformProviderAwsSpec struct {
	CRDs AwsCRDs `json:"crds"`

	//+optional
	Controller KubeformProviderSpec `json:"kubeform-provider"`
}

type AwsCRDs struct {
	Accessanalyzer                  bool `json:"accessanalyzer"`
	Acm                             bool `json:"acm"`
	Acmpca                          bool `json:"acmpca"`
	Alb                             bool `json:"alb"`
	Ami                             bool `json:"ami"`
	Amplify                         bool `json:"amplify"`
	Apigateway                      bool `json:"apigateway"`
	Apigatewayv2                    bool `json:"apigatewayv2"`
	App                             bool `json:"app"`
	Appautoscaling                  bool `json:"appautoscaling"`
	Appmesh                         bool `json:"appmesh"`
	Apprunner                       bool `json:"apprunner"`
	Appsync                         bool `json:"appsync"`
	Athena                          bool `json:"athena"`
	Autoscaling                     bool `json:"autoscaling"`
	Autoscalingplans                bool `json:"autoscalingplans"`
	Backup                          bool `json:"backup"`
	Batch                           bool `json:"batch"`
	Budgets                         bool `json:"budgets"`
	Cloud9                          bool `json:"cloud9"`
	Cloudformationstack             bool `json:"cloudformationstack"`
	Cloudformationtype              bool `json:"cloudformationtype"`
	Cloudfront                      bool `json:"cloudfront"`
	Cloudhsmv2                      bool `json:"cloudhsmv2"`
	Cloudtrail                      bool `json:"cloudtrail"`
	Cloudwatch                      bool `json:"cloudwatch"`
	Codeartifact                    bool `json:"codeartifact"`
	Codebuild                       bool `json:"codebuild"`
	Codecommit                      bool `json:"codecommit"`
	Codedeploy                      bool `json:"codedeploy"`
	Codepipeline                    bool `json:"codepipeline"`
	Codestarconnections             bool `json:"codestarconnections"`
	Codestarnotifications           bool `json:"codestarnotifications"`
	Cognito                         bool `json:"cognito"`
	Config                          bool `json:"config"`
	Cur                             bool `json:"cur"`
	Customer                        bool `json:"customer"`
	Datapipeline                    bool `json:"datapipeline"`
	Datasync                        bool `json:"datasync"`
	Dax                             bool `json:"dax"`
	Db                              bool `json:"db"`
	Default                         bool `json:"default"`
	Devicefarm                      bool `json:"devicefarm"`
	Directoryservice                bool `json:"directoryservice"`
	Dlm                             bool `json:"dlm"`
	Dms                             bool `json:"dms"`
	Docdb                           bool `json:"docdb"`
	Dx                              bool `json:"dx"`
	Dynamodb                        bool `json:"dynamodb"`
	Ebs                             bool `json:"ebs"`
	Ec2                             bool `json:"ec2"`
	Ecr                             bool `json:"ecr"`
	Ecrpublic                       bool `json:"ecrpublic"`
	Ecs                             bool `json:"ecs"`
	Efs                             bool `json:"efs"`
	Egress                          bool `json:"egress"`
	Eip                             bool `json:"eip"`
	Eks                             bool `json:"eks"`
	Elasticache                     bool `json:"elasticache"`
	Elasticbeanstalk                bool `json:"elasticbeanstalk"`
	Elasticsearchdomain             bool `json:"elasticsearchdomain"`
	Elastictranscoder               bool `json:"elastictranscoder"`
	Elb                             bool `json:"elb"`
	Emr                             bool `json:"emr"`
	Flow                            bool `json:"flow"`
	Fms                             bool `json:"fms"`
	Fsx                             bool `json:"fsx"`
	Gamelift                        bool `json:"gamelift"`
	Glaciervault                    bool `json:"glaciervault"`
	Globalaccelerator               bool `json:"globalaccelerator"`
	Glue                            bool `json:"glue"`
	Guardduty                       bool `json:"guardduty"`
	Iam                             bool `json:"iam"`
	Imagebuilder                    bool `json:"imagebuilder"`
	Inspector                       bool `json:"inspector"`
	Instance                        bool `json:"instance"`
	Internet                        bool `json:"internet"`
	Iot                             bool `json:"iot"`
	Key                             bool `json:"key"`
	Kinesis                         bool `json:"kinesis"`
	Kinesisanalyticsv2              bool `json:"kinesisanalyticsv2"`
	Kms                             bool `json:"kms"`
	Lakeformation                   bool `json:"lakeformation"`
	Lambda                          bool `json:"lambda"`
	Launch                          bool `json:"launch"`
	Lb                              bool `json:"lb"`
	Lex                             bool `json:"lex"`
	Licensemanager                  bool `json:"licensemanager"`
	Lightsail                       bool `json:"lightsail"`
	Loadbalancer                    bool `json:"loadbalancer"`
	Macie                           bool `json:"macie"`
	Macie2                          bool `json:"macie2"`
	Main                            bool `json:"main"`
	Media                           bool `json:"media"`
	Mq                              bool `json:"mq"`
	Msk                             bool `json:"msk"`
	Mwaa                            bool `json:"mwaa"`
	Nat                             bool `json:"nat"`
	Neptune                         bool `json:"neptune"`
	Network                         bool `json:"network"`
	Networkfirewall                 bool `json:"networkfirewall"`
	Opsworks                        bool `json:"opsworks"`
	Organizations                   bool `json:"organizations"`
	Pinpoint                        bool `json:"pinpoint"`
	Placement                       bool `json:"placement"`
	Prometheus                      bool `json:"prometheus"`
	Proxy                           bool `json:"proxy"`
	Qldb                            bool `json:"qldb"`
	Quicksight                      bool `json:"quicksight"`
	Ram                             bool `json:"ram"`
	Rds                             bool `json:"rds"`
	Redshift                        bool `json:"redshift"`
	Resourcegroups                  bool `json:"resourcegroups"`
	Route                           bool `json:"route"`
	Route53                         bool `json:"route53"`
	S3                              bool `json:"s3"`
	S3control                       bool `json:"s3control"`
	S3outposts                      bool `json:"s3outposts"`
	Sagemaker                       bool `json:"sagemaker"`
	Schemas                         bool `json:"schemas"`
	Secretsmanager                  bool `json:"secretsmanager"`
	Security                        bool `json:"security"`
	Securityhub                     bool `json:"securityhub"`
	Serverlessapplicationrepository bool `json:"serverlessapplicationrepository"`
	Servicecatalog                  bool `json:"servicecatalog"`
	Servicediscovery                bool `json:"servicediscovery"`
	Servicequotas                   bool `json:"servicequotas"`
	Ses                             bool `json:"ses"`
	Sfn                             bool `json:"sfn"`
	Shield                          bool `json:"shield"`
	Signer                          bool `json:"signer"`
	Simpledb                        bool `json:"simpledb"`
	Snapshot                        bool `json:"snapshot"`
	Sns                             bool `json:"sns"`
	Spot                            bool `json:"spot"`
	Sqsqueue                        bool `json:"sqsqueue"`
	Ssm                             bool `json:"ssm"`
	Ssoadmin                        bool `json:"ssoadmin"`
	Storagegateway                  bool `json:"storagegateway"`
	Subnet                          bool `json:"subnet"`
	Swf                             bool `json:"swf"`
	Synthetics                      bool `json:"synthetics"`
	Timestreamwrite                 bool `json:"timestreamwrite"`
	Transfer                        bool `json:"transfer"`
	Volume                          bool `json:"volume"`
	Vpc                             bool `json:"vpc"`
	Vpn                             bool `json:"vpn"`
	Waf                             bool `json:"waf"`
	Wafregional                     bool `json:"wafregional"`
	Wafv2                           bool `json:"wafv2"`
	Worklink                        bool `json:"worklink"`
	Workspaces                      bool `json:"workspaces"`
	Xray                            bool `json:"xray"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// KubeformProviderAwsList is a list of KubeformProviderAwss
type KubeformProviderAwsList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of KubeformProviderAws CRD objects
	Items []KubeformProviderAws `json:"items,omitempty"`
}
