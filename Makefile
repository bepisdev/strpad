GOX := $(shell which go)

.PHONY: test prepare publish

test: prepare
	go test -v ./...

prepare:
	go mod tidy

publish:
	git push --all
	GOPROXY=proxy.golang.org go list -m github.com/joshburnsxyz/strpad@v0.1.0

