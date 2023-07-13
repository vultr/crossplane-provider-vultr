/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	nodepools "github.com/crossplane-contrib/provider-vultr/internal/controller/kubernetes/nodepools"
	storage "github.com/crossplane-contrib/provider-vultr/internal/controller/object/storage"
	providerconfig "github.com/crossplane-contrib/provider-vultr/internal/controller/providerconfig"
	instance "github.com/crossplane-contrib/provider-vultr/internal/controller/vultr/instance"
	kubernetes "github.com/crossplane-contrib/provider-vultr/internal/controller/vultr/kubernetes"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		nodepools.Setup,
		storage.Setup,
		providerconfig.Setup,
		instance.Setup,
		kubernetes.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
