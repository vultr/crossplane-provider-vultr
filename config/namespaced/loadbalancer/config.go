package loadbalancer

import "github.com/crossplane/upjet/v2/pkg/config"

func Configure(p *config.Provider) {
	p.AddResourceConfigurator("vultr_load_balancer", func(r *config.Resource) {

		r.ShortGroup = "loadbalancer"
	})
}
