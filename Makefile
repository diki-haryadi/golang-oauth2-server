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

.PHONY: fmt lint golint test test-with-coverage ci
# TODO: When Go 1.9 is released vendor folder should be ignored automatically
PACKAGES=`go list ./... | grep -v vendor | grep -v mocks`

fmt:
	for pkg in ${PACKAGES}; do \
		go fmt $$pkg; \
	done;

lint:
	gometalinter --tests --disable-all --deadline=120s -E vet -E gofmt -E misspell -E ineffassign -E goimports -E deadcode ./...

golint:
	for pkg in ${PACKAGES}; do \
		golint $$pkg; \
	done;

test:
	TEST_FAILED= ; \
	for pkg in ${PACKAGES}; do \
		go test $$pkg || TEST_FAILED=1; \
	done; \
	[ -z "$$TEST_FAILED" ]

test-with-coverage:
	echo "" > coverage.out
	echo "mode: set" > coverage-all.out
	TEST_FAILED= ; \
	for pkg in ${PACKAGES}; do \
		go test -coverprofile=coverage.out -covermode=set $$pkg || TEST_FAILED=1; \
		tail -n +2 coverage.out >> coverage-all.out; \
	done; \
	[ -z "$$TEST_FAILED" ]
	#go tool cover -html=coverage-all.out

ci:
	bash -c 'docker-compose -f docker-compose.test.yml -p go_oauth2_server_ci up --build --abort-on-container-exit --exit-code-from sut'
