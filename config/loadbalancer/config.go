package loadbalancer

import "github.com/upbound/upjet/pkg/config"

func Configure(p *config.Provider) {
	p.AddResourceConfigurator("vultr_load_balancer", func(r *config.Resource) {
	})
}
