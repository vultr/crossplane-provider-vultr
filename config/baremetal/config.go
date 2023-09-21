package baremetal

import "github.com/upbound/upjet/pkg/config"

func Configure(p *config.Provider) {
	p.AddResourceConfigurator("vultr_bare_metal_server", func(r *config.Resource) {
	})
}
