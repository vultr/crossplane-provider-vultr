package config

import (
	"github.com/crossplane/upjet/v2/pkg/config"
)

// referenceConfigs maps a Terraform resource to its cross-resource
// references, keyed by the Terraform argument holding the referenced ID.
// Groups must not reference each other in both directions or the generated
// resolvers form an import cycle; e.g. compute references network, so network
// resources (vultr_reserved_ip) do not reference compute.
var referenceConfigs = map[string]map[string]string{
	"vultr_bare_metal_server": {
		"script_id":   "vultr_startup_script",
		"snapshot_id": "vultr_snapshot",
		"ssh_key_ids": "vultr_ssh_key",
		"vpc_id":      "vultr_vpc",
	},
	"vultr_block_storage": {
		"snapshot_id": "vultr_snapshot",
	},
	"vultr_database": {
		"vpc_id": "vultr_vpc",
	},
	"vultr_database_connection_pool": {"database_id": "vultr_database"},
	"vultr_database_connector":       {"database_id": "vultr_database"},
	"vultr_database_db":              {"database_id": "vultr_database"},
	"vultr_database_quota":           {"database_id": "vultr_database"},
	"vultr_database_replica":         {"database_id": "vultr_database"},
	"vultr_database_topic":           {"database_id": "vultr_database"},
	"vultr_database_user":            {"database_id": "vultr_database"},
	"vultr_dns_record":               {"domain": "vultr_dns_domain"},
	"vultr_firewall_rule":            {"firewall_group_id": "vultr_firewall_group"},
	"vultr_instance": {
		"firewall_group_id": "vultr_firewall_group",
		"iso_id":            "vultr_iso_private",
		"reserved_ip_id":    "vultr_reserved_ip",
		"script_id":         "vultr_startup_script",
		"snapshot_id":       "vultr_snapshot",
		"ssh_key_ids":       "vultr_ssh_key",
		"vpc_ids":           "vultr_vpc",
	},
	"vultr_instance_ipv4": {"instance_id": "vultr_instance"},
	"vultr_kubernetes": {
		"vpc_id": "vultr_vpc",
	},
	"vultr_kubernetes_node_pools": {"cluster_id": "vultr_kubernetes"},
	"vultr_nat_gateway":           {"vpc_id": "vultr_vpc"},
	"vultr_nat_gateway_firewall_rule": {
		"nat_gateway_id": "vultr_nat_gateway",
		"vpc_id":         "vultr_vpc",
	},
	"vultr_nat_gateway_port_forwarding_rule": {
		"nat_gateway_id": "vultr_nat_gateway",
		"vpc_id":         "vultr_vpc",
	},
	"vultr_object_storage_bucket": {"object_storage_id": "vultr_object_storage"},
	"vultr_organization_policy_group_attachment": {
		"group_id":  "vultr_organization_group",
		"policy_id": "vultr_organization_policy",
	},
	"vultr_organization_policy_user_attachment": {
		"policy_id": "vultr_organization_policy",
		"user_id":   "vultr_user",
	},
	"vultr_organization_role_group_attachment": {
		"group_id": "vultr_organization_group",
		"role_id":  "vultr_organization_role",
	},
	"vultr_organization_role_policy_attachment": {
		"policy_id": "vultr_organization_policy",
		"role_id":   "vultr_organization_role",
	},
	"vultr_organization_role_trust": {
		"group": "vultr_organization_group",
		"role":  "vultr_organization_role",
		"user":  "vultr_user",
	},
	"vultr_reverse_ipv4": {"instance_id": "vultr_instance"},
	"vultr_reverse_ipv6": {"instance_id": "vultr_instance"},
	"vultr_snapshot":     {"instance_id": "vultr_instance"},
}

// ReferenceConfigurations adds the cross-resource references listed in
// referenceConfigs.
func ReferenceConfigurations() config.ResourceOption {
	return func(r *config.Resource) {
		for field, target := range referenceConfigs[r.Name] {
			r.References[field] = config.Reference{
				TerraformName: target,
			}
		}
	}
}
