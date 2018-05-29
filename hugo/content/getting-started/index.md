---
title: Getting Started
date: 2018-05-28T16:27:37-07:00
draft: false
---

## Quick Start

### Install the Dependencies

#### `docker`

The `Dockerfile`s use multi-stage builds. Docker 17.05 or greater is required.
Follow the official installation [guide](https://docs.docker.com/install/) for your platform.

#### `minikube`

```bash
curl -Lo /usr/local/bin/minikube https://storage.googleapis.com/minikube/releases/latest/minikube-darwin-amd64 \
&& chmod +x /usr/local/bin/minikube
curl -LO https://storage.googleapis.com/minikube/releases/latest/docker-machine-driver-hyperkit \
&& chmod +x docker-machine-driver-hyperkit \
&& sudo mv docker-machine-driver-hyperkit /usr/local/bin/ \
&& sudo chown root:wheel /usr/local/bin/docker-machine-driver-hyperkit \
&& sudo chmod u+s /usr/local/bin/docker-machine-driver-hyperkit
```

#### `helm`

```bash
curl -L https://storage.googleapis.com/kubernetes-helm/helm-v2.9.1-darwin-amd64.tar.gz | tar -xz -C /usr/local/bin --strip-components=1 darwin-amd64/helm \
&& chmod +x /usr/local/bin/helm
```

#### `kubectl`

```bash
curl -L https://storage.googleapis.com/kubernetes-release/release/$(curl -s https://storage.googleapis.com/kubernetes-release/release/stable.txt)/bin/darwin/amd64/kubectl -C /usr/local/bin \
&& chmod +x /usr/local/bin/kubectl
```

#### `ip`

```bash
brew install iproute2mac
```

#### `dep`

```bash
go get -u github.com/golang/dep
```

> Note: `dep` is not a hard requirement to build and deploy Fibonacci. It is only required for development.

### Set the Required Environment Variables

The `Makefile` requires a valid Docker Hub username. Export the username before
invoking `make`:

```bash
export USERNAME=${your docker hub username goes here}
```

### Create the Kubernetes Cluster

These steps are optional, and exist to provide a local development environment.
You can skip them if you already have a Kubernetes cluster.
To target an existing cluster be sure to export the `CONTEXT` variable to the context of the cluster.

```bash
make kubernetes
sudo make dns
```

> Note: `sudo` is required since the `dns` target will modify `/etc/hosts`.
> DNS is optional.

### Build and Deploy

Once you have ensured that the above dependecies have been installed, you are ready to start.
The `all` target will perform the following:

- Build and push the Docker image
- Install the Fibonacci Helm chart

Execute the steps outlined above by running:

```bash
make
```

> Note: You can build on the minikube host by executing `eval $(minikube docker-env)` before running `make`

Congratulations! You can now use the API. For example:

```bash
curl http://fibonacci.local/v1/sequence/5
{"sequence":["0","1","1","2","3"]}
```

If you have setup DNS, or:

```bash
curl -H "Host: fibonacci.local" http://$(minikube ip)/v1/sequence/5
{"sequence":["0","1","1","2","3"]}
```

if you have not.

#### Cleaning Up

To tear everything down:

```bash
make clean
sudo make clean-dns
```

This is only required if you have created the minikube cluster.
