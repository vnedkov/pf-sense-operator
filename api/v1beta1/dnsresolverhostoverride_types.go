/*
Copyright 2025 Vess Nedkov.

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

package v1beta1

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// DNSResolverHostOverrideSpec defines the desired state of DNSResolverHostOverride.
type DNSResolverHostOverrideSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Foo is an example field of DNSResolverHostOverride. Edit dnsresolverhostoverride_types.go to remove/update
	Host   string `json:"host" yaml:"host"`
	Domain string `json:"domain" yaml:"domain"`

	//+optional
	Description string `json:"descr,omitempty" yaml:"descr,omitempty"`
	//+optional
	Aliases *[]DNSResolverHostOverrideAlias `json:"aliases,omitempty" yaml:"aliases,omitempty"`
	Service string                          `json:"service,omitempty" yaml:"service,omitempty"`
}

// DNSResolverHostOverrideAlias defines an alias (CNAME record) for the host override.
type DNSResolverHostOverrideAlias struct {
	Host   string `json:"host" yaml:"host"`
	Domain string `json:"domain" yaml:"domain"`
	//+optional
	Description string `json:"descr,omitempty" yaml:"descr,omitempty"`
}

// DNSResolverHostOverrideStatus defines the observed state of DNSResolverHostOverride.
type DNSResolverHostOverrideStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// A reference to the service that created the .
	// +optional
	Service corev1.ObjectReference `json:"active,omitempty"`
	// Information when the last time the job was run.
	// +optional
	LastRun *metav1.Time `json:"lastRun,omitempty"`
	// Information when the last time the job successfully finished.
	// +optional
	LastSuccessfulRun *metav1.Time `json:"lastSuccessfulRun,omitempty"`

	// Id of the external resource being managed
	// +optional
	ExternalID string `json:"externalId,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// DNSResolverHostOverride is the Schema for the dnsresolverhostoverrides API.
type DNSResolverHostOverride struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DNSResolverHostOverrideSpec   `json:"spec,omitempty"`
	Status DNSResolverHostOverrideStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DNSResolverHostOverrideList contains a list of DNSResolverHostOverride.
type DNSResolverHostOverrideList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DNSResolverHostOverride `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DNSResolverHostOverride{}, &DNSResolverHostOverrideList{})
}
