/*
Copyright 2021 Upbound Inc.
*/

package config

import (
	"strings"

	"github.com/crossplane/upjet/v2/pkg/config"
	"github.com/crossplane/upjet/v2/pkg/types/name"
)

// GroupKindOverrides overrides the group and kind of the resource if it matches
// any entry in the GroupMap.
func GroupKindOverrides() config.ResourceOption {
	return func(r *config.Resource) {
		if f, ok := GroupMap[r.Name]; ok {
			r.ShortGroup, r.Kind = f(r.Name)
		}
	}
}

// GroupKindCalculator returns the correct group and kind name for given TF
// resource.
type GroupKindCalculator func(resource string) (string, string)

// ReplaceGroupWords uses given group as the group of the resource and removes
// a number of words in resource name before calculating the kind of the resource.
func ReplaceGroupWords(group string, count int) GroupKindCalculator {
	return func(resource string) (string, string) {
		words := strings.Split(strings.TrimPrefix(resource, "vultr_"), "_")
		snakeKind := strings.Join(words[count:], "_")
		return group, name.NewFromSnake(snakeKind).Camel
	}
}

// GroupKind uses the given group and kind as-is, for resources whose kind
// can't be derived from the Terraform name (e.g. acronyms like VPC or IP).
func GroupKind(group, kind string) GroupKindCalculator {
	return func(string) (string, string) {
		return group, kind
	}
}

var GroupMap = map[string]GroupKindCalculator{
	"vultr_load_balancer":     ReplaceGroupWords("loadbalancer", 0),
	"vultr_bare_metal_server": ReplaceGroupWords("compute", 0),
	"vultr_instance":          ReplaceGroupWords("compute", 0),
	"vultr_kubernetes":        ReplaceGroupWords("kubernetes", 0),
	"vultr_database":          ReplaceGroupWords("database", 0),
	"vultr_object_storage":    ReplaceGroupWords("storage", 0),
	"vultr_block_storage":     ReplaceGroupWords("storage", 0),

	"vultr_instance_ipv4":     GroupKind("compute", "InstanceIPv4"),
	"vultr_iso_private":       GroupKind("compute", "ISO"),
	"vultr_reverse_ipv4":      GroupKind("compute", "ReverseIPv4"),
	"vultr_reverse_ipv6":      GroupKind("compute", "ReverseIPv6"),
	"vultr_snapshot":          GroupKind("compute", "Snapshot"),
	"vultr_snapshot_from_url": GroupKind("compute", "SnapshotFromURL"),
	"vultr_ssh_key":           GroupKind("compute", "SSHKey"),
	"vultr_startup_script":    GroupKind("compute", "StartupScript"),

	"vultr_firewall_group":                   GroupKind("network", "FirewallGroup"),
	"vultr_firewall_rule":                    GroupKind("network", "FirewallRule"),
	"vultr_nat_gateway":                      GroupKind("network", "NATGateway"),
	"vultr_nat_gateway_firewall_rule":        GroupKind("network", "NATGatewayFirewallRule"),
	"vultr_nat_gateway_port_forwarding_rule": GroupKind("network", "NATGatewayPortForwardingRule"),
	"vultr_reserved_ip":                      GroupKind("network", "ReservedIP"),
	"vultr_vpc":                              GroupKind("network", "VPC"),

	"vultr_dns_domain": GroupKind("dns", "Domain"),
	"vultr_dns_record": GroupKind("dns", "Record"),

	"vultr_object_storage_bucket":       GroupKind("storage", "ObjectStorageBucket"),
	"vultr_virtual_file_system_storage": GroupKind("storage", "VirtualFileSystemStorage"),

	"vultr_container_registry": GroupKind("registry", "ContainerRegistry"),
	"vultr_inference":          GroupKind("inference", "Inference"),

	"vultr_oidc_issuer":                          GroupKind("iam", "OIDCIssuer"),
	"vultr_oidc_provider":                        GroupKind("iam", "OIDCProvider"),
	"vultr_organization":                         GroupKind("iam", "Organization"),
	"vultr_organization_group":                   GroupKind("iam", "Group"),
	"vultr_organization_policy":                  GroupKind("iam", "Policy"),
	"vultr_organization_policy_group_attachment": GroupKind("iam", "PolicyGroupAttachment"),
	"vultr_organization_policy_user_attachment":  GroupKind("iam", "PolicyUserAttachment"),
	"vultr_organization_role":                    GroupKind("iam", "Role"),
	"vultr_organization_role_group_attachment":   GroupKind("iam", "RoleGroupAttachment"),
	"vultr_organization_role_policy_attachment":  GroupKind("iam", "RolePolicyAttachment"),
	"vultr_organization_role_trust":              GroupKind("iam", "RoleTrust"),
	"vultr_user":                                 GroupKind("iam", "User"),
}

// KindMap contains kind string overrides.
var KindMap = map[string]string{
	"vultr_load_balancer":     "LoadBalancer",
	"vultr_bare_metal_server": "BareMetalServer",
	"vultr_object_storage":    "ObjectStorage",
	"vultr_block_storage":     "BlockStorage",
}
