/*
Copyright 2021 Upbound Inc.
*/

package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	ujconfig "github.com/upbound/upjet/pkg/config"

	"github.com/crossplane-contrib/provider-vultr/config/baremetal"
	"github.com/crossplane-contrib/provider-vultr/config/block"
	"github.com/crossplane-contrib/provider-vultr/config/compute"
	"github.com/crossplane-contrib/provider-vultr/config/database"
	"github.com/crossplane-contrib/provider-vultr/config/kubernetes"
	"github.com/crossplane-contrib/provider-vultr/config/loadbalancer"
	"github.com/crossplane-contrib/provider-vultr/config/object"
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
		ujconfig.WithIncludeList(ExternalNameConfigured()),
		ujconfig.WithFeaturesPackage("internal/features"),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
			GroupKindOverrides(),
		))

	for _, configure := range []func(provider *ujconfig.Provider){
		kubernetes.Configure,
		object.Configure,
		compute.Configure,
		loadbalancer.Configure,
		database.Configure,
		baremetal.Configure,
		block.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}
