.PHONY: run build prod


GO_BUILD_ENV := CGO_ENABLED=0 GOOS=linux GOARCH=amd64 PORT=80
DOCKER_BUILD=$(shell pwd)/.docker_build
DOCKER_CMD=$(DOCKER_BUILD)/api

$(DOCKER_CMD): clean
	mkdir -p $(DOCKER_BUILD)
	$(GO_BUILD_ENV) go build -v -o $(DOCKER_CMD) .

clean:
	rm -rf $(DOCKER_BUILD)

dev: build
	PORT=8080 ./api

build:
	go build -o api