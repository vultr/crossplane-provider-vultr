package compute

import "github.com/upbound/upjet/pkg/config"

func Configure(p *config.Provider) {
	p.AddResourceConfigurator("vultr_instance", func(r *config.Resource) {
	})
}
