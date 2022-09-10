.PHONY: test prepare

test: prepare
	go test -v ./...

prepare:
	go mod tidy
