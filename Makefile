#!/usr/bin/env bash
GO_BUILD_ENV := CGO_ENABLED=0 GOOS=linux GOARCH=amd64
DOCKER_BUILD=$(shell pwd)/.docker_build
DOCKER_CMD=$(DOCKER_BUILD)/go-getting-started

$(DOCKER_CMD): clean
	mkdir -p $(DOCKER_BUILD)
	$(GO_BUILD_ENV) go build -v -o $(DOCKER_CMD) .

swag:
	# @export PATH="$HOME/go/bin:$PATH"
	@echo "> Generate Swagger Docs"
	# @if ! command -v swag &> /dev/null; then go install github.com/swaggo/swag/cmd/swag ; fi
	@swag init --parseVendor --parseDependency

clean:
	rm -rf $(DOCKER_BUILD)

heroku: $(DOCKER_CMD)
	heroku container:push web
