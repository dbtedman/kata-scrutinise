.DEFAULT_GOAL := all

all: install lint test build

install:
	@pnpm install && go mod vendor

lint:
	@pnpm run lint && gofmt -l ./internal ./cmd

format:
	@pnpm run format && gofmt -w ./internal ./cmd

test:
	@go test -cover -coverprofile=coverage.txt ./internal/... ./cmd/...

build:
	@go build -mod vendor -o scrutinise ./cmd/scrutinise
