/*
Copyright 2022 Upbound Inc.
*/

package config

import "github.com/upbound/upjet/pkg/config"

// ExternalNameConfigs contains all external name configurations for this
// provider.
var ExternalNameConfigs = map[string]config.ExternalName{
	// Import requires using a randomly generated ID from provider: nl-2e21sda
	"null_resource":                  config.IdentifierFromProvider,
	"vultr_kubernetes":               config.IdentifierFromProvider,
	"vultr_kubernetes_node_pools":    config.IdentifierFromProvider,
	"vultr_object_storage":           config.IdentifierFromProvider,
	"vultr_instance":                 config.IdentifierFromProvider,
	"vultr_load_balancer":            config.IdentifierFromProvider,
	"vultr_database":                 config.IdentifierFromProvider,
	"vultr_database_connection_pool": config.IdentifierFromProvider,
	"vultr_database_db":              config.IdentifierFromProvider,
	"vultr_database_replica":         config.IdentifierFromProvider,
	"vultr_database_user":            config.IdentifierFromProvider,
	"vultr_bare_metal_server":        config.IdentifierFromProvider,
	"vultr_block_storage":            config.IdentifierFromProvider,
}

// ExternalNameConfigurations applies all external name configs listed in the
// table ExternalNameConfigs and sets the version of those resources to v1beta1
// assuming they will be tested.
func ExternalNameConfigurations() config.ResourceOption {
	return func(r *config.Resource) {
		if e, ok := ExternalNameConfigs[r.Name]; ok {
			r.ExternalName = e
		}
	}
}

// ExternalNameConfigured returns the list of all resources whose external name
// is configured manually.
func ExternalNameConfigured() []string {
	l := make([]string, len(ExternalNameConfigs))
	i := 0
	for name := range ExternalNameConfigs {
		// $ is added to match the exact string since the format is regex.
		l[i] = name + "$"
		i++
	}
	return l
}
