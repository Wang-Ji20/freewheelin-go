.DEFAULT_GOAL := build

fmt:
	go fmt ./...
.PHONY: fmt

staticcheck: fmt
	staticcheck ./...
.PHONY: staticcheck

vet: staticcheck
	go vet ./...
.PHONY: vet

build: vet
	mkdir -p build
	go build
	mv main ./build
.PHONY: build