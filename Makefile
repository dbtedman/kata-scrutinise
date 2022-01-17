.DEFAULT_GOAL := all

all: install lint test build

install:
	@pnpm install && go mod vendor

lint:
	@pnpm run lint && gofmt -l ./main.go ./internal ./cmd

format:
	@pnpm run format && gofmt -w ./main.go ./internal ./cmd

test:
	@go test -cover -coverprofile=coverage.txt ./internal/... ./cmd/...

build:
	@go build -mod vendor -o scrutinise
