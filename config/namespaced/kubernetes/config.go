package kubernetes

import (
	"github.com/crossplane/upjet/v2/pkg/config"

	"github.com/crossplane-contrib/provider-vultr/internal/kubeconfig"
)

func Configure(p *config.Provider) {
	p.AddResourceConfigurator("vultr_kubernetes", func(r *config.Resource) {
		r.Sensitive.AdditionalConnectionDetailsFn = kubeconfig.ConnectionDetails
	})
}
