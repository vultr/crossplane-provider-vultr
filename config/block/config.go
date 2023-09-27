package block

import "github.com/upbound/upjet/pkg/config"

func Configure(p *config.Provider) {
	p.AddResourceConfigurator("vultr_block_storage", func(r *config.Resource) {
	})
}
