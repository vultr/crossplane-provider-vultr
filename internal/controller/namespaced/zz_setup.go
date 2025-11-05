// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	baremetalserver "github.com/crossplane-contrib/provider-vultr/internal/controller/namespaced/baremetal/baremetalserver"
	instance "github.com/crossplane-contrib/provider-vultr/internal/controller/namespaced/compute/instance"
	connectionpool "github.com/crossplane-contrib/provider-vultr/internal/controller/namespaced/database/connectionpool"
	database "github.com/crossplane-contrib/provider-vultr/internal/controller/namespaced/database/database"
	db "github.com/crossplane-contrib/provider-vultr/internal/controller/namespaced/database/db"
	replica "github.com/crossplane-contrib/provider-vultr/internal/controller/namespaced/database/replica"
	user "github.com/crossplane-contrib/provider-vultr/internal/controller/namespaced/database/user"
	kubernetes "github.com/crossplane-contrib/provider-vultr/internal/controller/namespaced/kubernetes/kubernetes"
	nodepools "github.com/crossplane-contrib/provider-vultr/internal/controller/namespaced/kubernetes/nodepools"
	loadbalancer "github.com/crossplane-contrib/provider-vultr/internal/controller/namespaced/loadbalancer/loadbalancer"
	providerconfig "github.com/crossplane-contrib/provider-vultr/internal/controller/namespaced/providerconfig"
	blockstorage "github.com/crossplane-contrib/provider-vultr/internal/controller/namespaced/storage/blockstorage"
	objectstorage "github.com/crossplane-contrib/provider-vultr/internal/controller/namespaced/storage/objectstorage"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		baremetalserver.Setup,
		instance.Setup,
		connectionpool.Setup,
		database.Setup,
		db.Setup,
		replica.Setup,
		user.Setup,
		kubernetes.Setup,
		nodepools.Setup,
		loadbalancer.Setup,
		providerconfig.Setup,
		blockstorage.Setup,
		objectstorage.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		baremetalserver.SetupGated,
		instance.SetupGated,
		connectionpool.SetupGated,
		database.SetupGated,
		db.SetupGated,
		replica.SetupGated,
		user.SetupGated,
		kubernetes.SetupGated,
		nodepools.SetupGated,
		loadbalancer.SetupGated,
		providerconfig.SetupGated,
		blockstorage.SetupGated,
		objectstorage.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
