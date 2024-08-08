GO ?= go

APP_NAME=my-trader

.PHONY: clean env build install uninstall run style vendor

clean:
	@${GO} clean
	@rm -f ./${APP_NAME}-server ./${APP_NAME}-client

env:
	@cp .env.example .env

build:
	@${GO} build -o ./${APP_NAME}-server ./server/main.go
	@${GO} build -o ./${APP_NAME}-cli ./client/main.go

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
