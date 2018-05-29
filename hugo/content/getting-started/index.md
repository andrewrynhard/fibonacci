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

#### `dep`

```bash
go get -u github.com/golang/dep
```

> Note: `dep` is not a hard requirement to build and deploy Fibonacci. It is only required for development.

#### Build and Deploy

Once you have ensured that the above dependecies have been installed, you are ready to start.

First, the `Makefile` requires a valid Docker Hub username. Export the username before
invoking `make`:

```bash
export USERNAME=${your docker hub username goes here}
```

The quick start uses the `all` target of the `Makefile`. This target will
perform the following:

- Create a minikube cluster
- Install
  - Prometheus
  - Grafana
- Build and push the Docker image
- Install Fibonacci

Execute the steps outlined above by running:

```bash
make
```

> Note: You can build on the minikube host by executing: `eval $(minikube docker-env)`


You can now use the API! For example:

```bash
curl fibonacci.local/v1/sequence/5
{"sequence":["0","1","1","2","3"]}
```
