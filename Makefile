SHELL := /bin/bash
REPOSITORY = andrewrynhard/fibonacci
SRC = github.com/$(REPOSITORY)
TAG = $(shell ./hack/tag.sh)
NAMESPACE = default
CHART = fibonacci
SET_ARGS = --set "app.name=$(NAME)" --set image.repository=$(REPOSITORY) --set image.tag=$(TAG)
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

test: build
	-docker rm -f fibonacci-latest
	docker run -d -p 80:80 --name fibonacci-latest  $(REPOSITORY):latest serve api
	docker build --add-host test.local:$(IP) -f ./hack/docker/Dockerfile.$@ -t $(REPOSITORY):$@ .
	-docker rm -f fibonacci-latest
	helm lint $(SET_ARGS) helm/$(CHART)

deploy: build test
	helm upgrade \
		--debug \
		--wait \
		--kube-context minikube \
		--install \
		$(SET_ARGS) \
		--namespace $(NAMESPACE) $(CHART) helm/$(CHART)

clean:
	helm delete --purge $(CHART)
