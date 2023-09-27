# Building and Installing the Crossplane Vultr Provider

`provider-vultr` is composed of a golang project and can be built directly with standard `golang` tools. We currently support two different platforms for building:

* Linux: most modern distros should work although most testing has been done on Ubuntu
* Mac: macOS 10.6+ is supported

## Build Requirements

An Intel-based machine (recommend 2+ cores, 2+ GB of memory and 128GB of SSD). Inside your build environment (Docker for Mac or a VM), 6+ GB memory is also recommended.

The following tools are need on the host:

* curl
* docker (1.12+) or Docker for Mac (17+)
* git
* make
* golang
* rsync (if you're using the build container on mac)
* helm (v2.8.2+)
* kubebuilder (v1.0.4+)

## Build
You can build the Crossplane Vultr Provider for the host platform by simply running the command below.
Building in parallel with the `-j` option is recommended.

```console
make -j4
```

The first time `make` is run, the build submodule will be synced and
updated. After initial setup, it can be updated by running `make submodules`.

Run `make help` for more options.

## Building inside the cross container

Official Crossplane builds are done inside a build container. This ensures that we get a consistent build, test and release environment. To run the build inside the cross container run:

```console
build/run make -j4
```

The first run of `build/run` will build the container itself and could take a few minutes to complete, but subsequent builds should go much faster.

## Install Crossplane in Your Cluster
Once your Cluster is up and running, you'll need to install Crossplane. 

We recommend using Helm to install Crossplane. You can find the [official documentation here](https://crossplane.io/docs/v1.5/getting-started/install-configure.html#install-crossplane). These are the commands: 

```console
kubectl create namespace crossplane-system

helm repo add crossplane-stable https://charts.crossplane.io/stable
helm repo update

helm install crossplane --namespace crossplane-system crossplane-stable/crossplane
```

## Getting started

Next install this Vultr Crossplane provider:
```bash
# clone this repo
git clone git@github.com:vultr/crossplane-provider-vultr.git
cd provider-vultr

# install the crds
kubectl apply -f package/crds
```

Run the provider:
```bash
make run
```

Configure the provider:
```bash
# change the examples/secret.yaml to use your VULTR_API_KEY
kubectl apply -f examples/providerconfig/secret.yaml
kubectl apply -f examples/providerconfig/providerconfig.yaml
```


## Provision Vultr Resources 

Go to the [examples-generated](./examples-generated) directory in this repo, find the Vultr product you'd like to deploy, make any needed changes to the yaml file, and then create the resource. 

For example, if you'd like to spin up a VKE Cluster, run the command

```console 
kubectl apply -f examples-generated/vultr/kubernetes.yaml
# wait a bit
kubectl get managed
```
and check your Vultr account to see if the resource has been created.