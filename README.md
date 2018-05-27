# Fibonacci

## Developing Fibonacci

### Install Minikube and the HyperKit Driver

```bash
curl -Lo /usr/local/bin/minikube https://storage.googleapis.com/minikube/releases/latest/minikube-darwin-amd64 \
&& chmod +x /usr/local/bin/minikube
curl -LO https://storage.googleapis.com/minikube/releases/latest/docker-machine-driver-hyperkit \
&& chmod +x docker-machine-driver-hyperkit \
&& sudo mv docker-machine-driver-hyperkit /usr/local/bin/ \
&& sudo chown root:wheel /usr/local/bin/docker-machine-driver-hyperkit \
&& sudo chmod u+s /usr/local/bin/docker-machine-driver-hyperkit
```

### Start the Minkube Cluster

```bash
minikube start --vm-driver=hyperkit --kubernetes-version=v1.10.3
```

### Initialize Helm

```bash
curl -L https://storage.googleapis.com/kubernetes-helm/helm-v2.9.1-darwin-amd64.tar.gz | tar -xz -C /usr/local/bin --strip-components=1 darwin-amd64/helm \
&& chmod +x /usr/local/bin/helm
kubectl --context minikube create serviceaccount tiller --namespace kube-system
kubectl --context minikube create clusterrolebinding tiller --clusterrole cluster-admin --serviceaccount=kube-system:tiller
helm init --kube-context minikube
```

## Deploy to Minkube

```bash
make minikube
```
