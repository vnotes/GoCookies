.PHONY: build

all: tool lint test

tool:
	@ gofmt -w .

lint:
	@ golangci-lint run ./...

test:
	@ go test ./...
