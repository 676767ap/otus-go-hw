APP_BIN := "./bin/banners-rotator"
DOCKER_IMG="banners-rotator:develop"

GIT_HASH := $(shell git log --format="%h" -n 1)
LDFLAGS := -X main.release="develop" -X main.buildDate=$(shell date -u +%Y-%m-%dT%H:%M:%S) -X main.gitHash=$(GIT_HASH)

down:
	docker-compose -f docker-compose.yaml down

# integration-tests:
# set -e; \
# docker-compose -f docker-compose.test.yaml up --build -d; \
# status_code=0; \
# docker-compose -f docker-compose.test.yaml run tests go test -v || status_code=$$?; \
# docker-compose -f docker-compose.test.yaml down; \
# exit $$status_code

build-local:
	go build -v -o $(APP_BIN) -ldflags "$(LDFLAGS)" ./cmd

run:
	docker-compose -f docker-compose.yaml up -d --build

version: build
	$(APP_BIN) version

tidy:
	go mod tidy

fmt:
	go fmt ./...

prepare_lint:
	go get github.com/golangci/golangci-lint/cmd/golangci-lint@v1.27.0

lint: prepare_lint fmt tidy
	golangci-lint run ./...

test:
	go test -race ./internal/...

.PHONY: tools build run build-img run-img version test