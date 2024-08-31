/*
Copyright 2024.

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

package controller

import (
	"context"
	"errors"

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	webappv1 "github.com/brbarmex/rollout-lab/api/v1"
	"github.com/brbarmex/rollout-lab/internal/helmservice"
	"github.com/go-logr/logr"
)

// DeployTenantReconciler reconciles a DeployTenant object
type DeployTenantReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

// +kubebuilder:rbac:groups=webapp.brbarmex,resources=deploytenants,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=webapp.brbarmex,resources=deploytenants/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=webapp.brbarmex,resources=deploytenants/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the DeployTenant object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.19.0/pkg/reconcile
func (r *DeployTenantReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {

	log := log.FromContext(ctx)

	var webapp webappv1.DeployTenant
	if err := r.Get(ctx, req.NamespacedName, &webapp); err != nil {
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	installApps(ctx, &webapp, log)

	return ctrl.Result{}, nil
}

func installApps(ctx context.Context, webapp *webappv1.DeployTenant, log logr.Logger) {
	for _, value := range webapp.Spec.Tenants {
		err := helmservice.HelmInstall(ctx, &helmservice.ChartParameter{
			Namespace:     value,
			ReleaseName:   value,
			ReleaseVerion: value,
			ChartName:     webapp.Spec.ChartName,
			DryRun:        false,
		})

		if err != nil {
			log.Error(errors.New(err.Error()), "Details")
		}
	}
}

// SetupWithManager sets up the controller with the Manager.
func (r *DeployTenantReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&webappv1.DeployTenant{}).
		//WithOptions(controller.Options{MaxConcurrentReconciles: 1}).
		Complete(r)
}
