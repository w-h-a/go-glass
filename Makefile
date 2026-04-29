.PHONY: build test

build:
	go build -o bin/go-glass ./cmd/go-glass

test:
	go test ./...
