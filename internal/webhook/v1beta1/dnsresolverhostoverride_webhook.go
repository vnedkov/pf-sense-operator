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
	"context"
	"fmt"

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"

	pfsensev1beta1 "github.com/vnedkov/pf-sense-operator/api/v1beta1"
)

// nolint:unused
// log is for logging in this package.
var dnsresolverhostoverridelog = logf.Log.WithName("dnsresolverhostoverride-resource")

// SetupDNSResolverHostOverrideWebhookWithManager registers the webhook for DNSResolverHostOverride in the manager.
func SetupDNSResolverHostOverrideWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).For(&pfsensev1beta1.DNSResolverHostOverride{}).
		WithValidator(&DNSResolverHostOverrideCustomValidator{}).
		WithDefaulter(&DNSResolverHostOverrideCustomDefaulter{}).
		Complete()
}

// TODO(user): EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!

// +kubebuilder:webhook:path=/mutate-pfsense-lab-nedkoff-com-v1beta1-dnsresolverhostoverride,mutating=true,failurePolicy=fail,sideEffects=None,groups=pfsense.lab.nedkoff.com,resources=dnsresolverhostoverrides,verbs=create;update,versions=v1beta1,name=mdnsresolverhostoverride-v1beta1.kb.io,admissionReviewVersions=v1

// DNSResolverHostOverrideCustomDefaulter struct is responsible for setting default values on the custom resource of the
// Kind DNSResolverHostOverride when those are created or updated.
//
// NOTE: The +kubebuilder:object:generate=false marker prevents controller-gen from generating DeepCopy methods,
// as it is used only for temporary operations and does not need to be deeply copied.
type DNSResolverHostOverrideCustomDefaulter struct {
	// TODO(user): Add more fields as needed for defaulting
}

var _ webhook.CustomDefaulter = &DNSResolverHostOverrideCustomDefaulter{}

// Default implements webhook.CustomDefaulter so a webhook will be registered for the Kind DNSResolverHostOverride.
func (d *DNSResolverHostOverrideCustomDefaulter) Default(ctx context.Context, obj runtime.Object) error {
	dnsresolverhostoverride, ok := obj.(*pfsensev1beta1.DNSResolverHostOverride)

	if !ok {
		return fmt.Errorf("expected an DNSResolverHostOverride object but got %T", obj)
	}
	dnsresolverhostoverridelog.Info("Defaulting for DNSResolverHostOverride", "name", dnsresolverhostoverride.GetName())

	// TODO(user): fill in your defaulting logic.

	return nil
}

// TODO(user): change verbs to "verbs=create;update;delete" if you want to enable deletion validation.
// NOTE: The 'path' attribute must follow a specific pattern and should not be modified directly here.
// Modifying the path for an invalid path can cause API server errors; failing to locate the webhook.
// +kubebuilder:webhook:path=/validate-pfsense-lab-nedkoff-com-v1beta1-dnsresolverhostoverride,mutating=false,failurePolicy=fail,sideEffects=None,groups=pfsense.lab.nedkoff.com,resources=dnsresolverhostoverrides,verbs=create;update,versions=v1beta1,name=vdnsresolverhostoverride-v1beta1.kb.io,admissionReviewVersions=v1

// DNSResolverHostOverrideCustomValidator struct is responsible for validating the DNSResolverHostOverride resource
// when it is created, updated, or deleted.
//
// NOTE: The +kubebuilder:object:generate=false marker prevents controller-gen from generating DeepCopy methods,
// as this struct is used only for temporary operations and does not need to be deeply copied.
type DNSResolverHostOverrideCustomValidator struct {
	// TODO(user): Add more fields as needed for validation
}

var _ webhook.CustomValidator = &DNSResolverHostOverrideCustomValidator{}

// ValidateCreate implements webhook.CustomValidator so a webhook will be registered for the type DNSResolverHostOverride.
func (v *DNSResolverHostOverrideCustomValidator) ValidateCreate(ctx context.Context, obj runtime.Object) (admission.Warnings, error) {
	dnsresolverhostoverride, ok := obj.(*pfsensev1beta1.DNSResolverHostOverride)
	if !ok {
		return nil, fmt.Errorf("expected a DNSResolverHostOverride object but got %T", obj)
	}
	dnsresolverhostoverridelog.Info("Validation for DNSResolverHostOverride upon creation", "name", dnsresolverhostoverride.GetName())

	// TODO(user): fill in your validation logic upon object creation.

	return nil, nil
}

// ValidateUpdate implements webhook.CustomValidator so a webhook will be registered for the type DNSResolverHostOverride.
func (v *DNSResolverHostOverrideCustomValidator) ValidateUpdate(ctx context.Context, oldObj, newObj runtime.Object) (admission.Warnings, error) {
	dnsresolverhostoverride, ok := newObj.(*pfsensev1beta1.DNSResolverHostOverride)
	if !ok {
		return nil, fmt.Errorf("expected a DNSResolverHostOverride object for the newObj but got %T", newObj)
	}
	dnsresolverhostoverridelog.Info("Validation for DNSResolverHostOverride upon update", "name", dnsresolverhostoverride.GetName())

	// TODO(user): fill in your validation logic upon object update.

	return nil, nil
}

// ValidateDelete implements webhook.CustomValidator so a webhook will be registered for the type DNSResolverHostOverride.
func (v *DNSResolverHostOverrideCustomValidator) ValidateDelete(ctx context.Context, obj runtime.Object) (admission.Warnings, error) {
	dnsresolverhostoverride, ok := obj.(*pfsensev1beta1.DNSResolverHostOverride)
	if !ok {
		return nil, fmt.Errorf("expected a DNSResolverHostOverride object but got %T", obj)
	}
	dnsresolverhostoverridelog.Info("Validation for DNSResolverHostOverride upon deletion", "name", dnsresolverhostoverride.GetName())

	// TODO(user): fill in your validation logic upon object deletion.

	return nil, nil
}
