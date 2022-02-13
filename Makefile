# Makefile for go-packfile

.DEFAULT_GOAL := help

# -----------------------------------------------------------------
#    ENV VARIABLE
# -----------------------------------------------------------------
NAME       ?= go-packfile
BIN_SUFFIX ?= ""
VERSION    := $(shell git describe --tags --abbrev=0 2> /dev/null || echo 0)
REVISION   := $(shell git rev-parse --short HEAD 2> /dev/null || echo 0)

DESTDIR    ?= ./bin
SOURCEDIR  := .
SOURCES    := $(shell find . -type f -name '*.go' | grep -v vendor)
LDFLAGS    := -ldflags="-s -w -X 'main.version=$(VERSION)' -X 'main.revision=$(REVISION)' -extldflags '-static'"
NOVENDOR   := $(shell go list $(SOURCEDIR)/... | grep -v vendor)
BUILD_OPTS := -v -a -tags netgo -installsuffix netgo $(LDFLAGS)

# Tools version
GOLANGCI_LINT_VERSION  := 1.41.1

# -----------------------------------------------------------------
#    Main targets
# -----------------------------------------------------------------

.PHONY: env
env: ## Print useful environment variables to stdout
	@echo NAME = $(NAME)
	@echo VERSION = $(VERSION)
	@echo REVISION = $(REVISION)

.PHONY: clean
clean: ## Remove temporary files
	@rm -rf cover.*
	@rm -rf bin/*
	@go clean --cache --testcache

.PHONY: tidy
tidy: ## Remove unnecessary go module packages
	@go mod tidy
	@go mod verify

.PHONY: fmt
fmt: ## Format all packages
	@goimports -w $(SOURCES)
	@go fmt $(NOVENDOR)

.PHONY: lint
lint: ## Code check
	@golangci-lint run -v ./...

.PHONY: test
test: ## Run all the tests
	@go test -race -cover $(SOURCEDIR)/...

.PHONY: cover
cover: ## Run unit test and out coverage file for local environment
	@go test -race -timeout 10m -coverprofile=cover.out -covermode=atomic $(SOURCEDIR)/...
	@go tool cover -html=cover.out -o cover.html
	@rm cover.out
	@open cover.html

.PHONY: help
help: env
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

# -----------------------------------------------------------------
#    Setup targets
# -----------------------------------------------------------------

.PHONY: setup
setup: ## Setup dev tools
	@go mod download
