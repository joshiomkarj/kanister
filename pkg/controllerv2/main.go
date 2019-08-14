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

package main

import (
	"context"
	"flag"
	"log"
	"os"

	_ "github.com/kanisterio/kanister/pkg/function"
	"github.com/kanisterio/kanister/pkg/handler"
	"github.com/kanisterio/kanister/pkg/kube"
	"github.com/kanisterio/kanister/pkg/resource"
	setupLog "github.com/sirupsen/logrus"
	"k8s.io/client-go/rest"

	//crv1alpha1 "github.com/kanisterio/kanister/pkg/controllerv2/api/v1alpha1"
	crv1alpha1 "github.com/kanisterio/kanister/pkg/apis/cr/v1alpha1"
	"github.com/kanisterio/kanister/pkg/controllerv2/controllers"
	"k8s.io/apimachinery/pkg/runtime"
	clientgoscheme "k8s.io/client-go/kubernetes/scheme"
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
	ctrl "sigs.k8s.io/controller-runtime"
	//"sigs.k8s.io/controller-runtime/pkg/log/zap"
	// +kubebuilder:scaffold:imports
)

var (
	scheme = runtime.NewScheme()
	//setupLog = ctrl.Log.WithName("setup")
)

func init() {
	_ = clientgoscheme.AddToScheme(scheme)

	_ = crv1alpha1.AddToScheme(scheme)
	// +kubebuilder:scaffold:scheme
}

func main() {
	var metricsAddr string
	var enableLeaderElection bool
	flag.StringVar(&metricsAddr, "metrics-addr", ":8080", "The address the metric endpoint binds to.")
	flag.BoolVar(&enableLeaderElection, "enable-leader-election", false,
		"Enable leader election for controller manager. Enabling this will ensure there is only one active controller manager.")
	flag.Parse()

	//ctrl.SetLogger(zap.Logger(true))

	ctx := context.Background()

	s := handler.NewServer()
	defer func() {
		if err := s.Shutdown(ctx); err != nil {
			setupLog.Error(err, "Failed to shutdown health check server")
		}
	}()
	go func() {
		if err := s.ListenAndServe(); err != nil {
			setupLog.Error(err, "Failed to shutdown health check server")
		}
	}()

	// Initialize the clients.
	setupLog.Info("Getting kubernetes context")
	config, err := rest.InClusterConfig()
	if err != nil {
		setupLog.Fatalf("Failed to get k8s config. %+v", err)
	}

	// Make sure the CRD's exist.
	if err := resource.CreateCustomResources(ctx, config); err != nil {
		log.Fatalf("Failed to create CustomResources. %+v", err)
	}

	ns, err := kube.GetControllerNamespace()
	if err != nil {
		log.Fatalf("Failed to determine this pod's namespace %+v", err)
	}

	// Create controller object
	c, err := controllers.New(config, ns)
	if err != nil {
		log.Fatalf("Failed to start controller. %+v", err)
	}

	mgr, err := ctrl.NewManager(ctrl.GetConfigOrDie(), ctrl.Options{
		Scheme:             scheme,
		MetricsBindAddress: metricsAddr,
		LeaderElection:     enableLeaderElection,
	})
	if err != nil {
		setupLog.Error(err, "unable to start manager")
		os.Exit(1)
	}

	if err = (&controllers.ActionSetReconciler{
		Client:     mgr.GetClient(),
		Log:        ctrl.Log.WithName("controllers").WithName("ActionSet"),
		Controller: c,
	}).SetupWithManager(mgr); err != nil {
		setupLog.Error(err, "unable to create controller", "controller", "ActionSet")
		os.Exit(1)
	}
	//	if err = (&controllers.BlueprintReconciler{
	//		Client:     mgr.GetClient(),
	//		Log:        ctrl.Log.WithName("controllers").WithName("Blueprint"),
	//		Controller: c,
	//	}).SetupWithManager(mgr); err != nil {
	//		setupLog.Error(err, "unable to create controller", "controller", "Blueprint")
	//		os.Exit(1)
	//	}
	// +kubebuilder:scaffold:builder

	setupLog.Info("starting manager")
	if err := mgr.Start(ctrl.SetupSignalHandler()); err != nil {
		setupLog.Error(err, "problem running manager")
		os.Exit(1)
	}
}
