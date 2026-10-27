// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	baremetalserver "github.com/crossplane-contrib/provider-vultr/internal/controller/cluster/compute/baremetalserver"
	instance "github.com/crossplane-contrib/provider-vultr/internal/controller/cluster/compute/instance"
	instanceipv4 "github.com/crossplane-contrib/provider-vultr/internal/controller/cluster/compute/instanceipv4"
	iso "github.com/crossplane-contrib/provider-vultr/internal/controller/cluster/compute/iso"
	reverseipv4 "github.com/crossplane-contrib/provider-vultr/internal/controller/cluster/compute/reverseipv4"
	reverseipv6 "github.com/crossplane-contrib/provider-vultr/internal/controller/cluster/compute/reverseipv6"
	snapshot "github.com/crossplane-contrib/provider-vultr/internal/controller/cluster/compute/snapshot"
	snapshotfromurl "github.com/crossplane-contrib/provider-vultr/internal/controller/cluster/compute/snapshotfromurl"
	sshkey "github.com/crossplane-contrib/provider-vultr/internal/controller/cluster/compute/sshkey"
	startupscript "github.com/crossplane-contrib/provider-vultr/internal/controller/cluster/compute/startupscript"
	connectionpool "github.com/crossplane-contrib/provider-vultr/internal/controller/cluster/database/connectionpool"
	connector "github.com/crossplane-contrib/provider-vultr/internal/controller/cluster/database/connector"
	database "github.com/crossplane-contrib/provider-vultr/internal/controller/cluster/database/database"
	db "github.com/crossplane-contrib/provider-vultr/internal/controller/cluster/database/db"
	quota "github.com/crossplane-contrib/provider-vultr/internal/controller/cluster/database/quota"
	replica "github.com/crossplane-contrib/provider-vultr/internal/controller/cluster/database/replica"
	topic "github.com/crossplane-contrib/provider-vultr/internal/controller/cluster/database/topic"
	user "github.com/crossplane-contrib/provider-vultr/internal/controller/cluster/database/user"
	domain "github.com/crossplane-contrib/provider-vultr/internal/controller/cluster/dns/domain"
	record "github.com/crossplane-contrib/provider-vultr/internal/controller/cluster/dns/record"
	group "github.com/crossplane-contrib/provider-vultr/internal/controller/cluster/iam/group"
	oidcissuer "github.com/crossplane-contrib/provider-vultr/internal/controller/cluster/iam/oidcissuer"
	oidcprovider "github.com/crossplane-contrib/provider-vultr/internal/controller/cluster/iam/oidcprovider"
	organization "github.com/crossplane-contrib/provider-vultr/internal/controller/cluster/iam/organization"
	policy "github.com/crossplane-contrib/provider-vultr/internal/controller/cluster/iam/policy"
	policygroupattachment "github.com/crossplane-contrib/provider-vultr/internal/controller/cluster/iam/policygroupattachment"
	policyuserattachment "github.com/crossplane-contrib/provider-vultr/internal/controller/cluster/iam/policyuserattachment"
	role "github.com/crossplane-contrib/provider-vultr/internal/controller/cluster/iam/role"
	rolegroupattachment "github.com/crossplane-contrib/provider-vultr/internal/controller/cluster/iam/rolegroupattachment"
	rolepolicyattachment "github.com/crossplane-contrib/provider-vultr/internal/controller/cluster/iam/rolepolicyattachment"
	roletrust "github.com/crossplane-contrib/provider-vultr/internal/controller/cluster/iam/roletrust"
	useriam "github.com/crossplane-contrib/provider-vultr/internal/controller/cluster/iam/user"
	inference "github.com/crossplane-contrib/provider-vultr/internal/controller/cluster/inference/inference"
	kubernetes "github.com/crossplane-contrib/provider-vultr/internal/controller/cluster/kubernetes/kubernetes"
	nodepools "github.com/crossplane-contrib/provider-vultr/internal/controller/cluster/kubernetes/nodepools"
	loadbalancer "github.com/crossplane-contrib/provider-vultr/internal/controller/cluster/loadbalancer/loadbalancer"
	firewallgroup "github.com/crossplane-contrib/provider-vultr/internal/controller/cluster/network/firewallgroup"
	firewallrule "github.com/crossplane-contrib/provider-vultr/internal/controller/cluster/network/firewallrule"
	natgateway "github.com/crossplane-contrib/provider-vultr/internal/controller/cluster/network/natgateway"
	natgatewayfirewallrule "github.com/crossplane-contrib/provider-vultr/internal/controller/cluster/network/natgatewayfirewallrule"
	natgatewayportforwardingrule "github.com/crossplane-contrib/provider-vultr/internal/controller/cluster/network/natgatewayportforwardingrule"
	reservedip "github.com/crossplane-contrib/provider-vultr/internal/controller/cluster/network/reservedip"
	vpc "github.com/crossplane-contrib/provider-vultr/internal/controller/cluster/network/vpc"
	providerconfig "github.com/crossplane-contrib/provider-vultr/internal/controller/cluster/providerconfig"
	containerregistry "github.com/crossplane-contrib/provider-vultr/internal/controller/cluster/registry/containerregistry"
	blockstorage "github.com/crossplane-contrib/provider-vultr/internal/controller/cluster/storage/blockstorage"
	objectstorage "github.com/crossplane-contrib/provider-vultr/internal/controller/cluster/storage/objectstorage"
	objectstoragebucket "github.com/crossplane-contrib/provider-vultr/internal/controller/cluster/storage/objectstoragebucket"
	virtualfilesystemstorage "github.com/crossplane-contrib/provider-vultr/internal/controller/cluster/storage/virtualfilesystemstorage"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		baremetalserver.Setup,
		instance.Setup,
		instanceipv4.Setup,
		iso.Setup,
		reverseipv4.Setup,
		reverseipv6.Setup,
		snapshot.Setup,
		snapshotfromurl.Setup,
		sshkey.Setup,
		startupscript.Setup,
		connectionpool.Setup,
		connector.Setup,
		database.Setup,
		db.Setup,
		quota.Setup,
		replica.Setup,
		topic.Setup,
		user.Setup,
		domain.Setup,
		record.Setup,
		group.Setup,
		oidcissuer.Setup,
		oidcprovider.Setup,
		organization.Setup,
		policy.Setup,
		policygroupattachment.Setup,
		policyuserattachment.Setup,
		role.Setup,
		rolegroupattachment.Setup,
		rolepolicyattachment.Setup,
		roletrust.Setup,
		useriam.Setup,
		inference.Setup,
		kubernetes.Setup,
		nodepools.Setup,
		loadbalancer.Setup,
		firewallgroup.Setup,
		firewallrule.Setup,
		natgateway.Setup,
		natgatewayfirewallrule.Setup,
		natgatewayportforwardingrule.Setup,
		reservedip.Setup,
		vpc.Setup,
		providerconfig.Setup,
		containerregistry.Setup,
		blockstorage.Setup,
		objectstorage.Setup,
		objectstoragebucket.Setup,
		virtualfilesystemstorage.Setup,
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
		instanceipv4.SetupGated,
		iso.SetupGated,
		reverseipv4.SetupGated,
		reverseipv6.SetupGated,
		snapshot.SetupGated,
		snapshotfromurl.SetupGated,
		sshkey.SetupGated,
		startupscript.SetupGated,
		connectionpool.SetupGated,
		connector.SetupGated,
		database.SetupGated,
		db.SetupGated,
		quota.SetupGated,
		replica.SetupGated,
		topic.SetupGated,
		user.SetupGated,
		domain.SetupGated,
		record.SetupGated,
		group.SetupGated,
		oidcissuer.SetupGated,
		oidcprovider.SetupGated,
		organization.SetupGated,
		policy.SetupGated,
		policygroupattachment.SetupGated,
		policyuserattachment.SetupGated,
		role.SetupGated,
		rolegroupattachment.SetupGated,
		rolepolicyattachment.SetupGated,
		roletrust.SetupGated,
		useriam.SetupGated,
		inference.SetupGated,
		kubernetes.SetupGated,
		nodepools.SetupGated,
		loadbalancer.SetupGated,
		firewallgroup.SetupGated,
		firewallrule.SetupGated,
		natgateway.SetupGated,
		natgatewayfirewallrule.SetupGated,
		natgatewayportforwardingrule.SetupGated,
		reservedip.SetupGated,
		vpc.SetupGated,
		providerconfig.SetupGated,
		containerregistry.SetupGated,
		blockstorage.SetupGated,
		objectstorage.SetupGated,
		objectstoragebucket.SetupGated,
		virtualfilesystemstorage.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupWebhookWithManager registers conversion webhooks for all resource kinds in the group.
func SetupWebhookWithManager(mgr ctrl.Manager) error {
	for _, setup := range []func(ctrl.Manager) error{
		baremetalserver.SetupWebhookWithManager,
		instance.SetupWebhookWithManager,
		instanceipv4.SetupWebhookWithManager,
		iso.SetupWebhookWithManager,
		reverseipv4.SetupWebhookWithManager,
		reverseipv6.SetupWebhookWithManager,
		snapshot.SetupWebhookWithManager,
		snapshotfromurl.SetupWebhookWithManager,
		sshkey.SetupWebhookWithManager,
		startupscript.SetupWebhookWithManager,
		connectionpool.SetupWebhookWithManager,
		connector.SetupWebhookWithManager,
		database.SetupWebhookWithManager,
		db.SetupWebhookWithManager,
		quota.SetupWebhookWithManager,
		replica.SetupWebhookWithManager,
		topic.SetupWebhookWithManager,
		user.SetupWebhookWithManager,
		domain.SetupWebhookWithManager,
		record.SetupWebhookWithManager,
		group.SetupWebhookWithManager,
		oidcissuer.SetupWebhookWithManager,
		oidcprovider.SetupWebhookWithManager,
		organization.SetupWebhookWithManager,
		policy.SetupWebhookWithManager,
		policygroupattachment.SetupWebhookWithManager,
		policyuserattachment.SetupWebhookWithManager,
		role.SetupWebhookWithManager,
		rolegroupattachment.SetupWebhookWithManager,
		rolepolicyattachment.SetupWebhookWithManager,
		roletrust.SetupWebhookWithManager,
		useriam.SetupWebhookWithManager,
		inference.SetupWebhookWithManager,
		kubernetes.SetupWebhookWithManager,
		nodepools.SetupWebhookWithManager,
		loadbalancer.SetupWebhookWithManager,
		firewallgroup.SetupWebhookWithManager,
		firewallrule.SetupWebhookWithManager,
		natgateway.SetupWebhookWithManager,
		natgatewayfirewallrule.SetupWebhookWithManager,
		natgatewayportforwardingrule.SetupWebhookWithManager,
		reservedip.SetupWebhookWithManager,
		vpc.SetupWebhookWithManager,
		providerconfig.SetupWebhookWithManager,
		containerregistry.SetupWebhookWithManager,
		blockstorage.SetupWebhookWithManager,
		objectstorage.SetupWebhookWithManager,
		objectstoragebucket.SetupWebhookWithManager,
		virtualfilesystemstorage.SetupWebhookWithManager,
	} {
		if err := setup(mgr); err != nil {
			return err
		}
	}
	return nil
}
