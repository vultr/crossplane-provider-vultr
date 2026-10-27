package kubernetes

import (
	"github.com/crossplane/upjet/v2/pkg/config"

	"github.com/crossplane-contrib/provider-vultr/internal/kubeconfig"
)

func Configure(p *config.Provider) {
	p.AddResourceConfigurator("vultr_kubernetes", func(r *config.Resource) {
		r.ShortGroup = "kubernetes"
		r.Sensitive.AdditionalConnectionDetailsFn = kubeconfig.ConnectionDetails
	})

	p.AddResourceConfigurator("vultr_kubernetes_node_pools", func(r *config.Resource) {
	})
}
