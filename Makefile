export GOBIN := $(PWD)/bin
export PATH := $(GOBIN):$(PATH)

SHELL := env PATH=$(PATH) /bin/sh

PG_HOST ?= localhost
PG_PORT ?= 5432
PG_USER ?= postgres
PG_DB ?= postgres
PG_DSN ?= $(PG_USER)@$(PG_HOST):$(PG_PORT)/$(PG_DB)?sslmode=disable

export PG_URI := postgresql://$(PG_DSN)

GOFLAGS ?=

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
	bash examples/gravity/install/sample_db/install.sh

.PHONY: run
run:
	go run examples/gravity/cmd/httpgraphql/main.go

.PHONY: generate
generate:
	mkdir -p examples/gravity/internal/gravity
	go run cmd/turboqlgen/main.go --package-name=gravity > examples/gravity/internal/gravity/graphql.go
	goimports -w .
