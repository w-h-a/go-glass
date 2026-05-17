.PHONY: tidy style test build

tidy:
	go mod tidy

style:
	goimports -l -w $(shell find . -name '*.go' -not -path './proto/*')

test:
	go test ./...

build:
	go build -o bin/go-glass ./cmd/go-glass
