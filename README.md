# Provider Vultr

`provider-vultr` is a [Crossplane](https://crossplane.io/) provider that
is built using [Upjet](https://github.com/upbound/upjet) code
generation tools and exposes XRM-conformant managed resources for the
Vultr API.

## Overview

This `provider-vultr` repository is the Crossplane infrastructure provider for
[vultr](https://vultr.com/). The provider that is built from
the source code in this repository can be installed into a Crossplane control
plane and adds the following new functionality:

* Custom Resource Definitions (CRDs) that model vultr infrastructure and managed
  services (e.g. [Vultr Kubernetes Engine](https://www.vultr.com/kubernetes/), [Vultr Loadbalancers](https://www.vultr.com/products/load-balancers/),
  etc.)
* Controllers to provision these resources in vultr based on the users desired
  state captured in CRDs they create
* Implementations of Crossplane's portable resource abstractions, enabling vultr
  resources to fulfill a user's general need for cloud services

## Getting Started and Documentation

For getting started guides, installation, deployment, and administration, see
our [Documentation](https://crossplane.io/docs/latest).

## Contributing

provider-vultr is a community driven project and we welcome contributions. See
the Crossplane
[Contributing](https://github.com/crossplane/crossplane/blob/master/CONTRIBUTING.md)
guidelines to get started.

## Report a Bug

For filing bugs, suggesting improvements, or requesting new features, please
open an [issue](https://github.com/vultr/crossplane-provider-vultr/issues).

## Contact

Please use the following to reach members of the community:

- Slack: Join our [slack channel](https://slack.crossplane.io)
- Forums:
  [crossplane-dev](https://groups.google.com/forum/#!forum/crossplane-dev)
- Twitter: [@crossplane_io](https://twitter.com/crossplane_io)
- Email: [info@crossplane.io](mailto:info@crossplane.io)

## Roadmap

provider-vultr goals and milestones are currently tracked in the Crossplane
repository. More information can be found in
[ROADMAP.md](https://github.com/crossplane/crossplane/blob/master/ROADMAP.md).

## Governance and Owners

provider-vultr is run according to the same
[Governance](https://github.com/crossplane/crossplane/blob/master/GOVERNANCE.md)
and [Ownership](https://github.com/crossplane/crossplane/blob/master/OWNERS.md)
structure as the core Crossplane project.

## Code of Conduct

provider-vultr adheres to the same [Code of
Conduct](https://github.com/crossplane/crossplane/blob/master/CODE_OF_CONDUCT.md)
as the core Crossplane project.

## Licensing

provider-vultr is under the Apache 2.0 license.

