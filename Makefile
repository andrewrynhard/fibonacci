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
UNAME_S := $(shell uname -s)
CONTEXT ?= minikube

all: deploy

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

coverage:
	docker build -f ./hack/docker/Dockerfile.$@ -t $(REPOSITORY):$@ .
	docker run --rm -i --volume $(shell pwd):/out --entrypoint=/bin/cp $(REPOSITORY):$@ coverage.txt /out/

.PHONY: push
push: test
	docker push $(REPOSITORY):$(TAG)

.PHONY: deploy
deploy: push
	helm upgrade \
		--debug \
		--wait \
		--kube-context $(CONTEXT) \
		--install \
		$(SET_ARGS) \
		--namespace $(NAMESPACE) $(CHART) helm/$(CHART)
	kubectl --context $(CONTEXT) rollout status deployment/fibonacci

.PHONY: docs
docs:
	-rm -rf ./docs
	docker build -f ./hack/docker/Dockerfile.$@ -t $(REPOSITORY):$@ .
	docker run --rm -i --volume $(shell pwd):/out --entrypoint=/bin/cp $(REPOSITORY):$@ -R ./docs /out/docs

.PHONY: kubernetes
kubernetes: minikube helm monitoring

.PHONY: minikube
minikube:
	minikube status || minikube start --vm-driver=hyperkit --kubernetes-version=v1.10.3 --cpus 4 --memory 8192 \
	&& kubectl --context minikube rollout status deployment/kube-dns --namespace kube-system
	minikube addons enable ingress

.PHONY: helm
helm:
	-kubectl --context $(CONTEXT) create serviceaccount tiller --namespace kube-system
	-kubectl --context $(CONTEXT) create clusterrolebinding tiller --clusterrole cluster-admin --serviceaccount=kube-system:tiller
	helm init --kube-context $(CONTEXT)
	kubectl --context $(CONTEXT) rollout status deployment/tiller-deploy --namespace kube-system

.PHONY: monitoring
monitoring:
	helm repo update
	helm upgrade \
		--debug \
		--wait \
		--kube-context $(CONTEXT) \
		--install \
		--version 6.7.0 \
		-f ./helm/prometheus/values.yaml \
		--namespace $@ prometheus stable/prometheus
	kubectl --context $(CONTEXT) rollout status deployment/prometheus-server --namespace monitoring
	helm upgrade \
		--debug \
		--wait \
		--kube-context $(CONTEXT) \
		--install \
		--version 1.10.0 \
		-f ./helm/grafana/values.yaml \
		--namespace $@ grafana stable/grafana
	kubectl --context $(CONTEXT) rollout status deployment/grafana --namespace monitoring
	-curl -X POST -H "Host: grafana.local" -H "Content-Type: application/json" --data "@./hack/fibonacci-datasource.json" http://admin:admin@$$(minikube ip)/api/datasources
	-curl -X POST -H "Host: grafana.local" -H "Content-Type: application/json" --data "@./hack/fibonacci-dashboard.json" http://admin:admin@$$(minikube ip)/api/dashboards/db

.PHONY: clean-dns
clean-dns:
    ifeq ($(UNAME_S),Linux)
		sed -i '/grafana.local/d' /private/etc/hosts
		sed -i '/fibonacci.local/d' /private/etc/hosts
    endif
    ifeq ($(UNAME_S),Darwin)
		sed -i '' '/grafana.local/d' /private/etc/hosts
		sed -i '' '/fibonacci.local/d' /private/etc/hosts
    endif

.PHONY: dns
dns: clean-dns
    ifeq ($(UNAME_S),Linux)
		echo "$$(minikube ip) grafana.local" | tee -a /etc/hosts
		echo "$$(minikube ip) fibonacci.local" | tee -a /etc/hosts
    endif
    ifeq ($(UNAME_S),Darwin)
		echo "$$(minikube ip) grafana.local" | tee -a /private/etc/hosts
		echo "$$(minikube ip) fibonacci.local" | tee -a /private/etc/hosts
    endif

.PHONY: clean
clean:
	minikube delete
