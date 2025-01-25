OUTPUT_BINARY=markdown-blog-server
API_VERSION=1.8.0

.PHONY: all
all: clean setup lint test build


.PHONY: clean
clean:
	rm -rf build

.PHONY: setup
setup:
	go get github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen
	go get github.com/oapi-codegen/runtime
	rm -rf ./gen/dist
	curl -LO https://github.com/rikotsev/markdown-blog/releases/download/api-v${API_VERSION}/openapi-spec.zip
	unzip openapi-spec.zip -d gen
	unlink openapi-spec.zip
	go generate ./...

.PHONY: lint
lint:
	golangci-lint run

.PHONY: test
test: setup
	go test ./... -cover

.PHONY: build
build:
	go build -o dist/${OUTPUT_BINARY} cmd/server.go

.PHONY: run
run:
	docker compose -f docker/docker-compose-run.yaml down
	docker compose -f docker/docker-compose-run.yaml up -d
	go run cmd/server.go