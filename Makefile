.DEFAULT_GOAL := help

export LC_ALL=en_US.UTF-8

-include .env
export $(shell sed 's/=.*//' .env)


.PHONY: help
help: ## Help command
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n"} /^[$$()% a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-25s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

.PHONY: run
mod-download:
	go mod download

install-dependencies: 
	go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@v1.12.4

generate: install-dependencies mod-download
	go generate ./...

run: generate
	go run cmd/app/main.go

tests: generate
	go test -v ./... -cover

lint: generate
	golangci-lint run --tests=0 ./...

build: generate 
	go build -o bin/app cmd/app/main.go