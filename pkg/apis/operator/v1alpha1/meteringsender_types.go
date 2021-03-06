//
// Copyright 2020 IBM Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// MeteringSenderSpec defines the desired state of MeteringSender
type MeteringSenderSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html
	OperatorVersion string                    `json:"operatorVersion,omitempty"`
	ImageRegistry   string                    `json:"imageRegistry,omitempty"`
	IAMnamespace    string                    `json:"iamNamespace,omitempty"`
	Sender          MeteringSenderSpecSender  `json:"sender,omitempty"`
	MongoDB         MeteringSenderSpecMongoDB `json:"mongodb,omitempty"`
}

// MeteringSenderSpecSender defines the metering-sender configuration in the the MeteringSender spec
type MeteringSenderSpecSender struct {
	ImageTagPostfix     string `json:"imageTagPostfix,omitempty"`
	ClusterName         string `json:"clusterName,omitempty"`
	ClusterNamespace    string `json:"clusterNamespace,omitempty"`
	HubKubeConfigSecret string `json:"hubKubeConfigSecret,omitempty"`
}

// MeteringSenderSpecMongoDB defines the MongoDB configuration in the the MeteringSender spec
type MeteringSenderSpecMongoDB struct {
	Host               string `json:"host,omitempty"`
	Port               string `json:"port,omitempty"`
	UsernameSecret     string `json:"usernameSecret,omitempty"`
	UsernameKey        string `json:"usernameKey,omitempty"`
	PasswordSecret     string `json:"passwordSecret,omitempty"`
	PasswordKey        string `json:"passwordKey,omitempty"`
	ClusterCertsSecret string `json:"clustercertssecret,omitempty"`
	ClientCertsSecret  string `json:"clientcertssecret,omitempty"`
}

// MeteringSenderStatus defines the observed state of MeteringSender
type MeteringSenderStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html
	// Nodes are the names of the metering-sender pods
	Nodes []string `json:"nodes"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// MeteringSender is the Schema for the meteringsenders API
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=meteringsenders,scope=Namespaced
type MeteringSender struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   MeteringSenderSpec   `json:"spec,omitempty"`
	Status MeteringSenderStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// MeteringSenderList contains a list of MeteringSender
type MeteringSenderList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MeteringSender `json:"items"`
}

func init() {
	SchemeBuilder.Register(&MeteringSender{}, &MeteringSenderList{})
}
