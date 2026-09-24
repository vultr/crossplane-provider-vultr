# Developing provider-vultr

## Requirements

* Linux or macOS
* Go (the version in [go.mod](./go.mod); `GOTOOLCHAIN=auto` fetches it)
* Docker
* `make`, `git`, `curl`, `unzip`

Other tools (Terraform, kind, kubectl, helm, the Crossplane CLI, uptest,
golangci-lint) are downloaded into `.cache/tools` by the Makefile. The first
`make` also initialises the `build` submodule.

## Code generation

The Terraform provider version is set by `TERRAFORM_PROVIDER_VERSION` in the
[Makefile](./Makefile). Resource configuration lives in [config/](./config):

* `external_name.go` - resources to generate and their external-name config
* `groups.go` - API group and kind of each resource
* `cluster/`, `namespaced/` - per-resource overrides for each scope

Regenerate APIs, controllers, CRDs and examples after any change:

```console
make generate
```

Run `make reviewable` (generate, lint, test) before opening a pull request.

## Run locally

Against the cluster in your current kubeconfig, with CRDs applied from
`package/crds`:

```console
kubectl apply -f package/crds
make run
```

Or build the package and deploy it with Crossplane to a local kind cluster:

```console
make local-deploy
```

## End-to-end tests

[uptest](https://github.com/crossplane/uptest) creates, updates, imports and
deletes the given examples against a real Vultr account:

```console
export UPTEST_CLOUD_CREDENTIALS='{"api_key":"<your API key>"}'
export UPTEST_EXAMPLE_LIST=examples-generated/namespaced/storage/v1alpha1/blockstorage.yaml
make e2e
```

On pull requests, a maintainer can comment `/test-examples="<path>"` to run the
same test in CI.
