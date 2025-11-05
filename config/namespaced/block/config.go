package block

import "github.com/crossplane/upjet/v2/pkg/config"

func Configure(p *config.Provider) {
	p.AddResourceConfigurator("vultr_block_storage", func(r *config.Resource) {
	})
}
