package baremetal

import "github.com/crossplane/upjet/v2/pkg/config"

func Configure(p *config.Provider) {
	p.AddResourceConfigurator("vultr_bare_metal_server", func(r *config.Resource) {

		r.ShortGroup = "baremetal"
	})
}
