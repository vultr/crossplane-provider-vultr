package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	ujconfig "github.com/crossplane/upjet/v2/pkg/config"

	computeCluster "github.com/crossplane-contrib/provider-vultr/config/cluster/compute"
	kubernetesCluster "github.com/crossplane-contrib/provider-vultr/config/cluster/kubernetes"
	baremetalNamespaced "github.com/crossplane-contrib/provider-vultr/config/namespaced/baremetal"
	blockNamespaced "github.com/crossplane-contrib/provider-vultr/config/namespaced/block"
	computeNamespaced "github.com/crossplane-contrib/provider-vultr/config/namespaced/compute"
	databaseNamespaced "github.com/crossplane-contrib/provider-vultr/config/namespaced/database"
	loadbalancerNamespaced "github.com/crossplane-contrib/provider-vultr/config/namespaced/loadbalancer"
	objectNamespaced "github.com/crossplane-contrib/provider-vultr/config/namespaced/object"
)

const (
	resourcePrefix = "vultr"
	modulePath     = "github.com/crossplane-contrib/provider-vultr"
)

//go:embed schema.json
var providerSchema string

//go:embed provider-metadata.yaml
var providerMetadata string

// GetProvider returns provider configuration
func GetProvider() *ujconfig.Provider {
	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithRootGroup("vultr.upbound.io"),
		ujconfig.WithIncludeList(ExternalNameConfigured()),
		ujconfig.WithFeaturesPackage("internal/features"),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
			GroupKindOverrides(),
		))

	for _, configure := range []func(provider *ujconfig.Provider){
		// add custom config functions
		kubernetesCluster.Configure,
		computeCluster.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}

// GetProviderNamespaced returns the namespaced provider configuration
func GetProviderNamespaced() *ujconfig.Provider {
	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithRootGroup("vultr.m.upbound.io"),
		ujconfig.WithIncludeList(ExternalNameConfigured()),
		ujconfig.WithFeaturesPackage("internal/features"),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
			GroupKindOverrides(),
		),
		ujconfig.WithExampleManifestConfiguration(ujconfig.ExampleManifestConfiguration{
			ManagedResourceNamespace: "crossplane-system",
		}))

	for _, configure := range []func(provider *ujconfig.Provider){
		computeNamespaced.Configure,
		baremetalNamespaced.Configure,
		databaseNamespaced.Configure,
		loadbalancerNamespaced.Configure,
		objectNamespaced.Configure,
		blockNamespaced.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}
