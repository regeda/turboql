export GOBIN := $(PWD)/bin
export PATH := $(GOBIN):$(PATH)

SHELL := env PATH=$(PATH) /bin/sh

GOLANGLINT_VERSION := 1.55.2

PG_HOST ?= localhost
PG_PORT ?= 5432
PG_USER ?= postgres
PG_DB ?= postgres
PG_DSN ?= $(PG_USER)@$(PG_HOST):$(PG_PORT)/$(PG_DB)?sslmode=disable

export PG_URI := postgresql://$(PG_DSN)

GOFLAGS ?=

$(GOBIN):
	mkdir -p $(GOBIN)

$(GOBIN)/golangci-lint: $(GOBIN)
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(GOBIN) v$(GOLANGLINT_VERSION)

.PHONY: clean
clean:
	rm -rf $(GOBIN)

.PHONY: lint
lint: $(GOBIN)/golangci-lint
	golangci-lint run --timeout 10m ./...

.PHONY: tidy
tidy:
	go mod tidy

.PHONY: docker-up
docker-up:
	@docker-compose up -d --build

.PHONY: docker-down
docker-down:
	@docker-compose down -v

.PHONY: migrate-up
migrate-up:
	bash examples/bookstore/install/sample_db/install.sh

.PHONY: run
run:
	go run examples/bookstore/cmd/bookstore/main.go

.PHONY: generate
generate:
	mkdir -p examples/bookstore/pkg/bookstore
	go run cmd/turboqlgen/main.go --package-name=bookstore > examples/bookstore/pkg/bookstore/graphql_gen.go
	goimports -w .

.PHONY: test
test:
	GOGC=off go test $(GOFLAGS) -race ./... -count 1
