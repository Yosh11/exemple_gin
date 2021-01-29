.PHONY: build test
build:
	go build -v ./cmd/main.go

test:
	richgo test -v -cover -timeout 30s ./...

.DEFAULT_GOAL := build
