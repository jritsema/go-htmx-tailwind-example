PACKAGES := $(shell go list ./...)

all: help

.PHONY: help
help: Makefile
	@echo
	@echo " Choose a make command to run"
	@echo
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'
	@echo

## vet: vet code
.PHONY: vet
vet:
	go vet $(PACKAGES)

## test: run unit tests
.PHONY: test
test:
	go test -race -cover $(PACKAGES)

## build: build a binary
.PHONY: build
build: test
	go build -o ./app -v

## autobuild: auto build when source files change
.PHONY: autobuild
autobuild:
	# curl -sf https://gobinaries.com/cespare/reflex | sh
	reflex -g '*.go' -- sh -c 'echo "\n\n\n\n\n\n" && make build'

## dockerbuild: build project into a docker container image
.PHONY: dockerbuild
dockerbuild: test
	docker-compose build

## start: build and run local project
.PHONY: start
start: build
	clear
	@echo ""
	./app

## deploy: build code into a container and deploy it to the cloud dev environment
.PHONY: deploy
deploy: build
	./deploy.sh

