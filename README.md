# Provider Vultr

`provider-vultr` is a [Crossplane](https://crossplane.io/) provider for
[Vultr](https://www.vultr.com/). It is generated with
[Upjet](https://github.com/crossplane/upjet) from the
[Vultr Terraform provider](https://github.com/vultr/terraform-provider-vultr)
and exposes Vultr infrastructure as Kubernetes managed resources.

Every managed resource is served in two flavours:

* namespaced, under `*.vultr.m.upbound.io` (Crossplane v2 style), and
* cluster-scoped, under `*.vultr.upbound.io` (legacy, for Crossplane v1
  compositions).

> [!NOTE]
> Upbound publishes a separate provider for Vultr
> ([upbound/provider-upjet-vultr](https://github.com/upbound/provider-upjet-vultr))
> that uses the same API groups. The two providers cannot be installed in the
> same cluster.

## Install

Crossplane v2 is recommended. Install the provider package:

```yaml
apiVersion: pkg.crossplane.io/v1
kind: Provider
metadata:
  name: provider-vultr
spec:
  package: ghcr.io/vultr/provider-vultr:<version>
```

See [releases](https://github.com/vultr/crossplane-provider-vultr/releases) for
available versions.

## Configure credentials

Create a secret holding your
[Vultr API key](https://my.vultr.com/settings/#settingsapi) as JSON.
`rate_limit` (milliseconds between requests) and `retry_limit` are optional.

```yaml
apiVersion: v1
kind: Secret
metadata:
  name: vultr-creds
  namespace: crossplane-system
type: Opaque
stringData:
  credentials: |
    {
      "api_key": "<your API key>"
    }
```

Then point a provider config at it. Namespaced managed resources use a
`ClusterProviderConfig` (or a `ProviderConfig` in their namespace) from
`vultr.m.upbound.io`; cluster-scoped managed resources use a `ProviderConfig`
from `vultr.upbound.io`. Both default to the one named `default`.

```yaml
apiVersion: vultr.m.upbound.io/v1beta1
kind: ClusterProviderConfig
metadata:
  name: default
spec:
  credentials:
    source: Secret
    secretRef:
      name: vultr-creds
      namespace: crossplane-system
      key: credentials
```

More examples are in [examples/](./examples).

## Create resources

Example manifests for every managed resource are generated into
[examples-generated/](./examples-generated), e.g. a VKE cluster:

```console
kubectl apply -f examples-generated/namespaced/kubernetes/v1alpha1/kubernetes.yaml
kubectl get managed
```

Resource schemas are in [package/crds](./package/crds) and follow the
arguments of the corresponding Terraform resource.

## Developing

See [INSTALL.md](./INSTALL.md) for building, code generation and running the
provider locally.

## Report a bug

Please open an
[issue](https://github.com/vultr/crossplane-provider-vultr/issues).

## License

provider-vultr is under the Apache 2.0 license.
