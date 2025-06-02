GO ?= go

APP_NAME=my-trader

.PHONY: clean env build install uninstall run style vendor

help: ## Show help message
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST) | sort

clean: ## Remove build artifacts and temporary files
	@${GO} clean
	@rm -f ./${APP_NAME}-server ./${APP_NAME}-client

env:
	@cp .env.example .env

build: ## Build the binaries
	@${GO} build -o ./${APP_NAME}-server ./server/main.go
	@${GO} build -o ./${APP_NAME}-cli ./client/main.go

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
