/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	storage "github.com/crossplane-contrib/provider-vultr/internal/controller/block/storage"
	connectionpool "github.com/crossplane-contrib/provider-vultr/internal/controller/database/connectionpool"
	database "github.com/crossplane-contrib/provider-vultr/internal/controller/database/database"
	db "github.com/crossplane-contrib/provider-vultr/internal/controller/database/db"
	replica "github.com/crossplane-contrib/provider-vultr/internal/controller/database/replica"
	user "github.com/crossplane-contrib/provider-vultr/internal/controller/database/user"
	kubernetes "github.com/crossplane-contrib/provider-vultr/internal/controller/kubernetes/kubernetes"
	nodepools "github.com/crossplane-contrib/provider-vultr/internal/controller/kubernetes/nodepools"
	providerconfig "github.com/crossplane-contrib/provider-vultr/internal/controller/providerconfig"
	baremetalserver "github.com/crossplane-contrib/provider-vultr/internal/controller/vultr/baremetalserver"
	instance "github.com/crossplane-contrib/provider-vultr/internal/controller/vultr/instance"
	loadbalancer "github.com/crossplane-contrib/provider-vultr/internal/controller/vultr/loadbalancer"
	objectstorage "github.com/crossplane-contrib/provider-vultr/internal/controller/vultr/objectstorage"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		storage.Setup,
		connectionpool.Setup,
		database.Setup,
		db.Setup,
		replica.Setup,
		user.Setup,
		kubernetes.Setup,
		nodepools.Setup,
		providerconfig.Setup,
		baremetalserver.Setup,
		instance.Setup,
		loadbalancer.Setup,
		objectstorage.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
