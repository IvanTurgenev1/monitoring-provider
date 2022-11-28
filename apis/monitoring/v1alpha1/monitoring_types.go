/*
Copyright 2022 The Crossplane Authors.

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
	"reflect"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

// MonitoringParameters are the configurable fields of a Monitoring.
type MonitoringParameters struct {
	ConfigurableField string `json:"configurableField"`
}

// MonitoringObservation are the observable fields of a Monitoring.
type MonitoringObservation struct {
	ObservableField string `json:"observableField,omitempty"`
}

// A MonitoringSpec defines the desired state of a Monitoring.
type MonitoringSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       MonitoringParameters `json:"forProvider"`
}

// A MonitoringStatus represents the observed state of a Monitoring.
type MonitoringStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          MonitoringObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// A Monitoring is an example API type.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,paas}
type Monitoring struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   MonitoringSpec   `json:"spec"`
	Status MonitoringStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// MonitoringList contains a list of Monitoring
type MonitoringList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Monitoring `json:"items"`
}

// Monitoring type metadata.
var (
	MonitoringKind             = reflect.TypeOf(Monitoring{}).Name()
	MonitoringGroupKind        = schema.GroupKind{Group: Group, Kind: MonitoringKind}.String()
	MonitoringKindAPIVersion   = MonitoringKind + "." + SchemeGroupVersion.String()
	MonitoringGroupVersionKind = SchemeGroupVersion.WithKind(MonitoringKind)
)

func init() {
	SchemeBuilder.Register(&Monitoring{}, &MonitoringList{})
}
