.DEFAULT_GOAL := all

all: install lint test build

install:
	@pnpm install && go mod vendor

lint:
	@pnpm run lint && gofmt -l ./main.go ./lib

format:
	@pnpm run format && gofmt -w ./main.go ./lib

test:
	@go test -cover -coverprofile=coverage.txt ./lib/...

build:
	@go build -mod vendor -o scrutinise
