SHELL := /bin/bash
REPOSITORY = arynhard/fibonacci
SRC = github.com/$(REPOSITORY)
TAG = $(shell ./hack/tag.sh)
NAMESPACE = default
CHART = fibonacci
SET_ARGS = --set image.repository=$(REPOSITORY) --set image.tag=$(TAG) --set redis.usePassword="false"
IP = $(shell ip route get 1 | awk '{print $$NF;exit}')

all: test

.PHONY: build clean generate deploy test

start:

generate:
	-find $(shell pwd)/pkg/generated ! -name 'configure_fibonacci.go' -delete
	docker build -f ./hack/docker/Dockerfile.$@ -t $(REPOSITORY):$@ .
	docker run --rm -i --volume $(shell pwd):/out --entrypoint=/bin/cp $(REPOSITORY):$@ -R ./pkg/generated /out/pkg

build: generate
	docker build -f ./hack/docker/Dockerfile.$@ -t $(REPOSITORY):latest .
	docker tag $(REPOSITORY):latest $(REPOSITORY):$(TAG)
	helm dependency build helm/$(CHART)

test: build
	-docker rm -f fibonacci-latest
	docker run -d -p 8080:8080 --name fibonacci-latest  $(REPOSITORY):latest serve api --api-port 8080
	docker build --build-arg TEST_HOST="test.local:8080" --add-host test.local:$(IP) -f ./hack/docker/Dockerfile.$@ -t $(REPOSITORY):$@ .
	-docker rm -f fibonacci-latest
	helm lint $(SET_ARGS) helm/$(CHART)

push: test
	docker push $(REPOSITORY):$(TAG)

monitoring:
	helm upgrade \
		--debug \
		--wait \
		--kube-context minikube \
		--install \
		-f ./helm/prometheus/values.yaml \
		--namespace $@ prometheus stable/prometheus
	helm upgrade \
		--debug \
		--wait \
		--kube-context minikube \
		--install \
		-f ./helm/grafana/values.yaml \
		--namespace $@ grafana stable/grafana

deploy:
	helm upgrade \
		--debug \
		--wait \
		--kube-context minikube \
		--install \
		$(SET_ARGS) \
		--namespace $(NAMESPACE) $(CHART) helm/$(CHART)

clean-dns:
	sed -i '' '/grafana.local/d' /private/etc/hosts
	sed -i '' '/fibonacci.local/d' /private/etc/hosts

dns: clean-dns
	echo "$$(minikube ip) grafana.local" | tee -a /private/etc/hosts
	echo "$$(minikube ip) fibonacci.local" | tee -a /private/etc/hosts

clean: clean-dns
	helm delete --purge $(CHART)
