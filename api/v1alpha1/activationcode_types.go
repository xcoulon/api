/*
Copyright 2021 Your name.

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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ActivationCodeSpec defines the desired state of ActivationCode
type ActivationCodeSpec struct {
	StartDate        metav1.Time `json:"startDate"`
	EndDate          metav1.Time `json:"endDate"`
	TierName         string      `json:"tierName"`
	MaxNumberOfUsers int         `json:"maxNumberOfUsers"`
}

// ActivationCodeStatus defines the observed state of ActivationCode
type ActivationCodeStatus struct {
	NumberOfUsers int `json:"numberOfUsers"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// ActivationCode is the Schema for the activationcodes API
type ActivationCode struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ActivationCodeSpec   `json:"spec,omitempty"`
	Status ActivationCodeStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// ActivationCodeList contains a list of ActivationCode
type ActivationCodeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ActivationCode `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ActivationCode{}, &ActivationCodeList{})
}
