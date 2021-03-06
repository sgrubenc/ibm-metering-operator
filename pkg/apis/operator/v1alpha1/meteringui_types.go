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

// MeteringUISpec defines the desired state of MeteringUI
type MeteringUISpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html
	OperatorVersion  string                    `json:"operatorVersion,omitempty"`
	ImageRegistry    string                    `json:"imageRegistry,omitempty"`
	IAMnamespace     string                    `json:"iamNamespace,omitempty"`
	IngressNamespace string                    `json:"ingressNamespace,omitempty"`
	UI               MeteringUISpecUI          `json:"ui,omitempty"`
	DataManager      MeteringUISpecDataManager `json:"dm,omitempty"`
	MongoDB          MeteringUISpecMongoDB     `json:"mongodb,omitempty"`
	External         MeteringUISpecExternal    `json:"external,omitempty"`
}

// MeteringUISpecUI defines the metering-ui configuration in the the MeteringUI spec
type MeteringUISpecUI struct {
	ImageTagPostfix string `json:"imageTagPostfix,omitempty"`
}

// MeteringUISpecDataManager defines the metering-datamanager configuration in the the MeteringUI spec
type MeteringUISpecDataManager struct {
	ImageTagPostfix string `json:"imageTagPostfix,omitempty"`
}

// MeteringUISpecMongoDB defines the MongoDB configuration in the the MeteringUI spec
type MeteringUISpecMongoDB struct {
	Host               string `json:"host,omitempty"`
	Port               string `json:"port,omitempty"`
	UsernameSecret     string `json:"usernameSecret,omitempty"`
	UsernameKey        string `json:"usernameKey,omitempty"`
	PasswordSecret     string `json:"passwordSecret,omitempty"`
	PasswordKey        string `json:"passwordKey,omitempty"`
	ClusterCertsSecret string `json:"clustercertssecret,omitempty"`
	ClientCertsSecret  string `json:"clientcertssecret,omitempty"`
}

// MeteringUISpecExternal defines the external cluster configuration in the the MeteringUI spec
type MeteringUISpecExternal struct {
	ClusterIP   string `json:"clusterIP,omitempty"`
	ClusterPort string `json:"clusterPort,omitempty"`
	ClusterName string `json:"clusterName,omitempty"`
}

// MeteringUIStatus defines the observed state of MeteringUI
type MeteringUIStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html
	// Nodes are the names of the metering-ui pods
	Nodes []string `json:"nodes"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// MeteringUI is the Schema for the meteringuis API
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=meteringuis,scope=Namespaced
type MeteringUI struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   MeteringUISpec   `json:"spec,omitempty"`
	Status MeteringUIStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// MeteringUIList contains a list of MeteringUI
type MeteringUIList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MeteringUI `json:"items"`
}

func init() {
	SchemeBuilder.Register(&MeteringUI{}, &MeteringUIList{})
}
