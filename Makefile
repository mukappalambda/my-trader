GO ?= go
GIT_VERSION := $(shell git --version | awk '{print $$3}')
GIT_COMMIT  := $(shell git rev-parse HEAD)
BUILD_DATE  := $(shell date -u +%Y-%m-%dT%H:%M:%SZ)
USERNAME := mukappalambda
REPO_NAME := my-trader
REPO_URL := https://github.com/${USERNAME}/${REPO_NAME}
LDFLAGS := -X 'github.com/mukappalambda/my-trader/version.GitVersion=$(GIT_VERSION)' \
	-X 'github.com/mukappalambda/my-trader/version.GitCommit=$(GIT_COMMIT)' \
	-X 'github.com/mukappalambda/my-trader/version.BuildDate=$(BUILD_DATE)' \
	-X 'github.com/mukappalambda/my-trader/version.RepoUrl=$(REPO_URL)'

APP_NAME=my-trader

.PHONY: clean env build install uninstall run style vendor

help: ## Show help message
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST) | sort

clean: ## Remove build artifacts and temporary files
	@echo "$(WHALE) $@"
	@rm -rf bin

env:
	@cp .env.example .env

build: clean ## Build the binaries
	@echo "$(WHALE) $@"
	@${GO} build -ldflags "$(LDFLAGS) -s -w" -o bin/${APP_NAME}-server cmd/my-trader-server/main.go
	@${GO} build -ldflags "$(LDFLAGS) -s -w" -o bin/${APP_NAME}-cli cmd/my-trader-cli/main.go

lint: ## Run golangci-lint
	@echo "$(WHALE) $@"
	@golangci-lint run ./...

install: ## Install the binaries
	@${GO} install ./...

uninstall: ## Uinstall the binaries
	@rm -f $(shell ${GO} env GOPATH)/bin/${APP_NAME}

run:
	@${APP_NAME}

style: ## Format the source
	@gofmt -s -w .

vendor:
	@${GO} mod tidy
	@${GO} mod vendor
	@${GO} mod verify
