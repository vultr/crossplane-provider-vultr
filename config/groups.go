/*
Copyright 2021 Upbound Inc.
*/

package config

import (
	"strings"

	"github.com/upbound/upjet/pkg/config"
	"github.com/upbound/upjet/pkg/types/name"
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

var GroupMap = map[string]GroupKindCalculator{
	"vultr_load_balancer":     ReplaceGroupWords("vultr", 0),
	"vultr_bare_metal_server": ReplaceGroupWords("vultr", 0),
	"vultr_object_storage":    ReplaceGroupWords("vultr", 0),
}

// KindMap contains kind string overrides.
var KindMap = map[string]string{
	"vultr_load_balancer":     "loadbalancer",
	"vultr_bare_metal_server": "baremetal",
	"vultr_object_storage":    "object",
}
