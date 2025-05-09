OUTPUT_BINARY=markdown-blog-server
API_VERSION=1.10.1
DIST_BUILD=CGO_ENABLED=1 GOOS=linux GOARCH=amd64

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

.PHONY: dist
dist:
	$(DIST_BUILD) go build -ldflags="-s -w" -o dist/${OUTPUT_BINARY} cmd/server.go


.PHONY: run
run:
	docker compose -f docker/docker-compose-run.yaml down
	docker compose -f docker/docker-compose-run.yaml up -d
	export SERVER_ADDRESS="0.0.0.0:8080" && export DATABASE_URL="postgres://mdblog:mdblog@localhost:5432/mdblog?sslmode=disable" && go run cmd/server.go