.DEFAULT_GOAL := build

fmt:
	go fmt ./...
.PHONY: fmt

staticcheck: fmt
	staticcheck ./...
.PHONY: staticcheck

vet: fmt
	go vet ./...
.PHONY: vet

build: vet
	mkdir -p build
	go build
	mv main ./build
.PHONY: build