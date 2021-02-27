# Common variables
VERSION := 0.0.1
BUILD_INFO := Manual build 
#SRC_DIR := cmd
SRC_DIR := ./cmd

# Most likely want to override these when calling `make image`
IMAGE_REG ?= ghcr.io
IMAGE_REPO ?= benc-uk/k6-prometheus-exporter
IMAGE_TAG ?= latest
IMAGE_PREFIX := $(IMAGE_REG)/$(IMAGE_REPO)

.PHONY: help image push build run lint lint-fix
.DEFAULT_GOAL := help

help:  ## This help message :)
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

lint:  ## Lint & format, will not fix but sets exit code on error
	@which golangci-lint > /dev/null || go get github.com/golangci/golangci-lint/cmd/golangci-lint
	`go env GOPATH`/bin/golangci-lint run $(SRC_DIR)/...

lint-fix:  ## Lint & format, will try to fix errors and modify code
	@which golangci-lint > /dev/null || go get github.com/golangci/golangci-lint/cmd/golangci-lint
	`go env GOPATH`/bin/golangci-lint run $(SRC_DIR)/... --fix 

image:  ## Build container image from Dockerfile
	docker build --file ./build/Dockerfile \
	--build-arg BUILD_INFO="$(BUILD_INFO)" \
	--build-arg VERSION="$(VERSION)" \
	--tag $(IMAGE_PREFIX):$(IMAGE_TAG) . 

push:  ## Push container image to registry
	docker push $(IMAGE_PREFIX):$(IMAGE_TAG)

build:  ## Run a local build without a container
	CGO_ENABLED=0 GOOS=linux go build -o bin/k6-prometheus-exporter $(SRC_DIR)/...

run:  ## Run application, used for local development
	air -c .air.toml