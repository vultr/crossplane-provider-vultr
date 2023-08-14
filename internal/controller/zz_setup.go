/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	connectionpool "github.com/crossplane-contrib/provider-vultr/internal/controller/database/connectionpool"
	db "github.com/crossplane-contrib/provider-vultr/internal/controller/database/db"
	replica "github.com/crossplane-contrib/provider-vultr/internal/controller/database/replica"
	user "github.com/crossplane-contrib/provider-vultr/internal/controller/database/user"
	nodepools "github.com/crossplane-contrib/provider-vultr/internal/controller/kubernetes/nodepools"
	storage "github.com/crossplane-contrib/provider-vultr/internal/controller/object/storage"
	providerconfig "github.com/crossplane-contrib/provider-vultr/internal/controller/providerconfig"
	baremetalserver "github.com/crossplane-contrib/provider-vultr/internal/controller/vultr/baremetalserver"
	database "github.com/crossplane-contrib/provider-vultr/internal/controller/vultr/database"
	instance "github.com/crossplane-contrib/provider-vultr/internal/controller/vultr/instance"
	kubernetes "github.com/crossplane-contrib/provider-vultr/internal/controller/vultr/kubernetes"
	loadbalancer "github.com/crossplane-contrib/provider-vultr/internal/controller/vultr/loadbalancer"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		connectionpool.Setup,
		db.Setup,
		replica.Setup,
		user.Setup,
		nodepools.Setup,
		storage.Setup,
		providerconfig.Setup,
		baremetalserver.Setup,
		database.Setup,
		instance.Setup,
		kubernetes.Setup,
		loadbalancer.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
