# Changelog

## v0.3.0

First release from this repository, published to `docker.io/vultr/provider-vultr`.
Requires Crossplane v2.

### Breaking changes

- Resources are available namespaced (`*.vultr.m.upbound.io`) and
  cluster-scoped (`*.vultr.upbound.io`). Instance, BareMetalServer,
  LoadBalancer, ObjectStorage and block Storage moved to new API groups.
- The `Kubernetes` connection secret has the kubeconfig under `kubeconfig`.
- `LoadBalancer.globalRegions` is now a list of objects, and `vpc2Ids` was
  removed.

### Added

- 36 new resources, including firewalls, VPCs, DNS, SSH keys, snapshots, IAM
  and container registries.
- Cross-resource references, e.g. `vpcIdRef` and `clusterIdRef`.

### Changed

- Vultr Terraform provider 2.33.2.

### Upgrading from v0.2.0

v0.2.0 was published as `xpkg.upbound.io/crossplane-contrib/provider-vultr`,
from the now archived crossplane-contrib repository. To upgrade, set
`spec.package` of the existing `Provider` to
`docker.io/vultr/provider-vultr:v0.3.0`.

Kubernetes and database resources upgrade in place. For resources whose API
group moved, create the resource in its new group with the same
`crossplane.io/external-name`, then remove the finalizer from the old one and
delete it. The cloud resource is kept.

## v0.2.0

Updated the package metadata.

## v0.1.0

Initial release.
