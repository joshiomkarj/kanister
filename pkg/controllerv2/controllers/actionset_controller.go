/*

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

package controllers

import (
	"context"

	"github.com/go-logr/logr"
	log "github.com/sirupsen/logrus"
	"k8s.io/apimachinery/pkg/api/errors"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	//logf "sigs.k8s.io/controller-runtime/pkg/runtime/log"

	//crv1alpha1 "github.com/kanisterio/kanister/pkg/controllerv2/api/v1alpha1"
	crv1alpha1 "github.com/kanisterio/kanister/pkg/apis/cr/v1alpha1"
)

//var log = logf.Log.WithName("actionset_controller")

// ActionSetReconciler reconciles a ActionSet object
type ActionSetReconciler struct {
	client.Client
	Log        logr.Logger
	Controller *Controller
}

// +kubebuilder:rbac:groups=cr.kanister.io,resources=actionsets,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=cr.kanister.io,resources=actionsets/status,verbs=get;update;patch

func (r *ActionSetReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	ctx := context.Background()
	logger := r.Log.WithValues("actionset", req.NamespacedName)

	// Fetch the actionset resource
	actionset := &crv1alpha1.ActionSet{}
	err := r.Client.Get(context.TODO(), req.NamespacedName, actionset)
	if err != nil {
		if errors.IsNotFound(err) {
			// ActionSet resource not found
			// Return and don't requeue
			return ctrl.Result{}, nil
		}
		// Error getting the object - requeue the request.
		return ctrl.Result{}, err
	}

	// Check if delete event
	if actionset.ObjectMeta.DeletionTimestamp != nil {
		log.Info("Handling ActionSet delete request")
		r.Controller.onDelete(actionset)
		r.removeFinalizer(ctx, actionset)
		return ctrl.Result{}, nil
	}

	// Check if it's create event or update event
	// For figuring out this, we are checking status of the resource
	// It status is not set then it is a create event otherwise update event
	if actionset.Status == nil {
		// Create event
		log.Info("Handling ActionSet create request")

		// Add finalizer if does not exist to capture delete events
		err := r.addFinalizer(ctx, actionset)
		if err != nil {
			logger.Error(err, "Add Finalizer to ActionSet failed", "ActionSet.Namespace", actionset.Namespace, "ActionSet.Name", actionset.Name)
			return ctrl.Result{}, err
		}

		err = r.Controller.onAdd(actionset.DeepCopy())
		if err != nil {
			logger.Error(err, "Add ActionSet failed", "ActionSet.Namespace", actionset.Namespace, "ActionSet.Name", actionset.Name)
		}
		return ctrl.Result{}, err
	}

	// Update event
	log.Info("Handling ActionSet update request")
	err = r.Controller.onUpdate(actionset)
	if err != nil {
		logger.Error(err, "Update ActionSet failed", "ActionSet.Namespace", actionset.Namespace, "ActionSet.Name", actionset.Name)
	}
	//actionset.Status.State = crv1alpha1.StateComplete
	//err = r.Client.Status().Update(ctx, actionset)
	return ctrl.Result{}, err
}

func (r *ActionSetReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&crv1alpha1.ActionSet{}).
		Complete(r)
}

func (r *ActionSetReconciler) addFinalizer(ctx context.Context, actionset *crv1alpha1.ActionSet) error {
	if actionset.ObjectMeta.Finalizers == nil {
		log.Info("Adding finalizer")
		actionset.ObjectMeta.Finalizers = []string{"actionsetcleaner.cr.kanister.io"}
		if err := r.Client.Update(ctx, actionset); err != nil {
			return err
		}
	}
	return nil
}

func (r *ActionSetReconciler) removeFinalizer(ctx context.Context, actionset *crv1alpha1.ActionSet) error {
	log.Info("Removing finalizer")
	actionset.ObjectMeta.Finalizers = nil
	err := r.Client.Update(ctx, actionset)
	if err != nil {
		return err
	}
	return nil
}
