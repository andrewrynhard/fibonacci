ifeq ($(USERNAME),)
  $(error USERNAME is not set)
endif

SHELL := /bin/bash
REPOSITORY = $(USERNAME)/fibonacci
SRC = github.com/$(REPOSITORY)
TAG = $(shell ./hack/tag.sh)
NAMESPACE = default
CHART = fibonacci
SET_ARGS = --set image.repository=$(REPOSITORY) --set image.tag=$(TAG) --set redis.usePassword="false"
IP = $(shell ip route get 1 | awk '{print $$NF;exit}')

all: kubernetes deploy

.PHONY: generate
generate:
	-find $(shell pwd)/pkg/generated ! -name 'configure_fibonacci.go' -delete
	docker build -f ./hack/docker/Dockerfile.$@ -t $(REPOSITORY):$@ .
	docker run --rm -i --volume $(shell pwd):/out --entrypoint=/bin/cp $(REPOSITORY):$@ -R ./pkg/generated /out/pkg

.PHONY: build
build: generate
	docker build -f ./hack/docker/Dockerfile.$@ -t $(REPOSITORY):latest .
	docker tag $(REPOSITORY):latest $(REPOSITORY):$(TAG)
	helm dependency build helm/$(CHART)

.PHONY: test
test: build
	-docker rm -f fibonacci-latest
	docker run -d -p 8080:8080 --name fibonacci-latest  $(REPOSITORY):latest serve api --api-port 8080
	docker build --build-arg TEST_HOST="test.local:8080" --add-host test.local:$(IP) -f ./hack/docker/Dockerfile.$@ -t $(REPOSITORY):$@ .
	-docker rm -f fibonacci-latest
	helm lint $(SET_ARGS) helm/$(CHART)

.PHONY: push
push: test
	docker push $(REPOSITORY):$(TAG)

.PHONY: deploy
deploy:
	helm upgrade \
		--debug \
		--wait \
		--kube-context minikube \
		--install \
		$(SET_ARGS) \
		--namespace $(NAMESPACE) $(CHART) helm/$(CHART)

.PHONY: docs
docs:
	-rm -rf ./docs
	docker build -f ./hack/docker/Dockerfile.$@ -t $(REPOSITORY):$@ .
	docker run --rm -i --volume $(shell pwd):/out --entrypoint=/bin/cp $(REPOSITORY):$@ -R ./docs /out/docs

.PHONY: kubernetes
kubernetes: minikube dns helm monitoring

.PHONY: minikube
minikube:
	minikube status || minikube start --vm-driver=hyperkit --kubernetes-version=v1.10.3 --cpus 4 --memory 8192
	minikube addons enable ingress

.PHONY: helm
helm:
	-kubectl --context minikube create serviceaccount tiller --namespace kube-system
	-kubectl --context minikube create clusterrolebinding tiller --clusterrole cluster-admin --serviceaccount=kube-system:tiller
	helm init --kube-context minikube
	kubectl --context minikube rollout status deployment/tiller-deploy --namespace kube-system

.PHONY: monitoring
monitoring:
	helm upgrade \
		--debug \
		--wait \
		--kube-context minikube \
		--install \
		-f ./helm/prometheus/values.yaml \
		--namespace $@ prometheus stable/prometheus
	kubectl --context minikube rollout status deployment/prometheus-server --namespace monitoring
	helm upgrade \
		--debug \
		--wait \
		--kube-context minikube \
		--install \
		-f ./helm/grafana/values.yaml \
		--namespace $@ grafana stable/grafana
	kubectl --context minikube rollout status deployment/grafana --namespace monitoring

.PHONY: clean-dns
clean-dns:
	sudo sed -i '' '/grafana.local/d' /private/etc/hosts
	sudo sed -i '' '/fibonacci.local/d' /private/etc/hosts

.PHONY: dns
dns: clean-dns
	echo "$$(minikube ip) grafana.local" | sudo tee -a /private/etc/hosts
	echo "$$(minikube ip) fibonacci.local" | sudo tee -a /private/etc/hosts

.PHONY: clean
clean: clean-dns
	minikube delete
