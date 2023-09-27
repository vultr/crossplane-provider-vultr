package kubernetes

import "github.com/upbound/upjet/pkg/config"

func Configure(p *config.Provider) {
	p.AddResourceConfigurator("vultr_kubernetes", func(r *config.Resource) {
		r.ShortGroup = "kubernetes"
	})

	p.AddResourceConfigurator("vultr_kubernetes_node_pools", func(r *config.Resource) {
	})
}
