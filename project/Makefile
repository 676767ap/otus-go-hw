APP_BIN := "./bin/banners-rotator"
DOCKER_IMG="banners-rotator:develop"

GIT_HASH := $(shell git log --format="%h" -n 1)
LDFLAGS := -X main.release="develop" -X main.buildDate=$(shell date -u +%Y-%m-%dT%H:%M:%S) -X main.gitHash=$(GIT_HASH)

up:
	docker-compose -f docker-compose.yaml up -d --build

down:
	docker-compose -f docker-compose.yaml down

# integration-tests:
# set -e; \
# docker-compose -f docker-compose.test.yaml up --build -d; \
# status_code=0; \
# docker-compose -f docker-compose.test.yaml run tests go test -v || status_code=$$?; \
# docker-compose -f docker-compose.test.yaml down; \
# exit $$status_code

build:
	go build -v -o $(APP_BIN) -ldflags "$(LDFLAGS)" ./cmd

run: build
	$(APP_BIN)

build-img:
	docker build \
		--build-arg=LDFLAGS="$(LDFLAGS)" \
		-t $(DOCKER_IMG) \
		-f build/Dockerfile .

run-img: build-img
	docker run $(DOCKER_IMG)

version: build
	$(APP_BIN) version

test:
	go test -race ./internal/...

.PHONY: tools build run build-img run-img version test