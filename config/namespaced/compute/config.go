package compute

import "github.com/crossplane/upjet/v2/pkg/config"

func Configure(p *config.Provider) {
	p.AddResourceConfigurator("vultr_instance", func(r *config.Resource) {
	})
}
