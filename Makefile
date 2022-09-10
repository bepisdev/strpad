GOX := $(shell which go)

.PHONY: test prepare publish

test: prepare
	go test -v ./...

prepare:
	go mod tidy
