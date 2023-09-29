GO ?= go

APP_NAME=my-trader

.PHONY: clean env build install uninstall run style vendor

clean:
	@${GO} clean

env:
	@cp .env.example .env

build:
	@${GO} build ./...

install:
	@${GO} install ./...

uninstall:
	@rm -f $(shell ${GO} env GOPATH)/bin/${APP_NAME}

run:
	@${APP_NAME}

style:
	@gofmt -s -w .

vendor:
	@${GO} mod tidy
	@${GO} mod vendor
	@${GO} mod verify
