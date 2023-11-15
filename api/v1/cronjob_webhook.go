/*
Copyright 2023.

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

package v1

import (
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

// log is for logging in this package.
//var cronjoblog = logf.Log.WithName("cronjob-resource")

func (r *CronJob) SetupWebhookWithManager(mgr ctrl.Manager) error {

	create := admission.WithCustomDefaulter(mgr.GetScheme(), r, &customDef{name: "createDefaulter"})
	update := admission.WithCustomDefaulter(mgr.GetScheme(), r, &customDef{name: "updateDefaulter"})
	mgr.GetWebhookServer().Register("/create", create)
	mgr.GetWebhookServer().Register("/update", update)

	return nil

	//return ctrl.NewWebhookManagedBy(mgr).
	//	For(r).
	//	Complete()
}

// TODO(user): EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!

//+kubebuilder:webhook:path=/create,mutating=true,failurePolicy=fail,sideEffects=None,groups=batch.tutorial.kubebuilder.io,resources=cronjobs,verbs=create,versions=v1,name=mcronjob.kb.io.creator,admissionReviewVersions=v1
//+kubebuilder:webhook:path=/update,mutating=true,failurePolicy=fail,sideEffects=None,groups=batch.tutorial.kubebuilder.io,resources=cronjobs,verbs=update,versions=v1,name=mcronjob.kb.io.updator,admissionReviewVersions=v1

//var _ webhook.Defaulter = &CronJob{}
//
//// Default implements webhook.Defaulter so a webhook will be registered for the type
//func (r *CronJob) Default() {
//	cronjoblog.Info("default", "name", r.Name)
//
//	// TODO(user): fill in your defaulting logic.
//}
