package database

import "github.com/upbound/upjet/pkg/config"

func Configure(p *config.Provider) {
	p.AddResourceConfigurator("vultr_database", func(r *config.Resource) {
		r.ShortGroup = "database"
	})

	p.AddResourceConfigurator("vultr_database_connection_pool", func(r *config.Resource) {
	})
	p.AddResourceConfigurator("vultr_database_db", func(r *config.Resource) {
	})
	p.AddResourceConfigurator("vultr_database_replica", func(r *config.Resource) {
	})
	p.AddResourceConfigurator("vultr_database_user", func(r *config.Resource) {
	})
}
