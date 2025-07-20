# --- Tooling & Variables ----------------------------------------------------------------
export SHELL := bash

# Exporting bin folder to the path for makefile
export PATH   := $(PWD)/scripts/bin:$(PATH)

export VERSION := $(shell git describe --tags 2>/dev/null || echo "v0.1.0")
export COMMIT :=$(shell git rev-parse HEAD)
export BRANCH :=$(shell git rev-parse --abbrev-ref HEAD)

# Type of OS: Linux or Darwin.
export GOBASE := $(shell pwd)
export GOBUILDBASE := $(shell pwd)/build
export OSTYPE := $(shell uname -s | tr A-Z a-z)
export ARCH := $(shell uname -m)
export PROJECTNAME := $(shell basename "$(PWD)")
export GOFILES := $(wildcard *.go)

export BINARY := superdo

export CMD := $(GOBASE)/cmd
export HTTPSERVER := $(CMD)/httpserver/main.go
export MIGRATION := $(CMD)/migrations/main.go

MYSQL_USER ?= admin
MYSQL_PASSWORD ?= password123
MYSQL_ADDRESS ?= 127.0.0.1:3306
MYSQL_DATABASE ?= go-otp-auth_db
MYSQL_DSN := "mysql://$(MYSQL_USER):$(MYSQL_PASSWORD)@tcp($(MYSQL_ADDRESS))/$(MYSQL_DATABASE)"
MYSQL_MIGRATION_PATH := "./internal/repository/mysql/migrations"

include ./scripts/help.Makefile
include ./scripts/tools.Makefile

# ==================================================================================== #
# Dependencies
# ==================================================================================== #
install-deps:
	@ $(MAKE) scripts/bin/migrate

# ==================================================================================== #
# DEVELOPMENT
# ==================================================================================== #
## development-up: Startup / Build services from docker-compose and air for live reloading
.PHONY: development-up
development-up:
	@echo
	@echo " > Startup / Build services from docker-compose and air for live reloading"
	@echo
	@docker-compose -f deployments/development/docker-compose.yaml up

## go/env: print Go environment information
.PHONY: go/env
go/env:
	@echo "  >  Environment information"
	go env

## go/clean: remove object files and cached files
.PHONY: go/clean
go/clean:
	@echo "  >  Cleaning build cache"
    @GOPATH=$(GOPATH) GOBIN=$(GOBIN) go clean

# ==================================================================================== #
# DATABASE MIGRATIONS
# ==================================================================================== #
.PHONY: migrate-force
migrate-force: $(MIGRATE) ##  Set version V but don't run migration (ignores dirty state).
	migrate -database $(MYSQL_DSN) -path $(MYSQL_MIGRATION_PATH) version
	migrate -database $(MYSQL_DSN) -path $(MYSQL_MIGRATION_PATH) force 1

.PHONY: migrate-up
migrate-up: $(MIGRATE) ## Apply all (or N up) migrations.
	@ read -p "How many migration you wants to perform (default value: [all]): " N; \
	migrate  -database $(MYSQL_DSN) -path $(MYSQL_MIGRATION_PATH) up ${NN}

.PHONY: migrate-down
migrate-down: $(MIGRATE) ## Apply all (or N down) migrations.
	@ read -p "How many migration you wants to perform (default value: [all]): " N; \
	migrate  -database $(MYSQL_DSN) -path $(MYSQL_MIGRATION_PATH) down ${NN}

.PHONY: migrate-drop
migrate-drop: $(MIGRATE) ## Drop everything inside the database.
	migrate  -database $(MYSQL_DSN) -path $(MYSQL_MIGRATION_PATH) drop

.PHONY: migrate-create
migrate-create: $(MIGRATE) ## Create a set of up/down migrations with a specified name.
	@ read -p "Please provide name for the migration: " Name; \
	migrate create -ext sql -dir $(MYSQL_MIGRATION_PATH) $${Name}

# ==================================================================================== #
# QUALITY CONTROL
# ==================================================================================== #
## gosec: The go security checker
.PHONY: gosec
gosec:
	@if ! command -v gosec &> /dev/null; then \
    	echo "gosec not found, installing..."; \
		go install github.com/securego/gosec/v2/cmd/gosec@latest; \
    else \
    	echo "gosec is already installed"; \
	fi
	gosec --version
	gosec ./...

## staticcheck: The advanced Go linter
.PHONY: staticcheck
staticcheck:
	@if ! command -v staticcheck &> /dev/null; then \
    	echo "staticcheck not found, installing..."; \
    	go install honnef.co/go/tools/cmd/staticcheck@latest; \
    else \
    	echo "staticcheck is already installed"; \
	fi
	staticcheck --version
	staticcheck ./...

## govulncheck: looks for vulnerabilities in Go programs using a specific build configuration. For analyzing source code
.PHONY: govulncheck
govulncheck:
	@if ! command -v govulncheck &> /dev/null; then \
    	echo "govulncheck not found, installing..."; \
		go install golang.org/x/vuln/cmd/govulncheck@latest ; \
    else \
    	echo "govulncheck is already installed"; \
	fi
	govulncheck --version
	govulncheck ./...

## golangci-lint: Smart, fast linters runner
.PHONY: golangci-lint
lint:
	@if ! command -v golangci-lint &> /dev/null; then \
		echo "golangci-lint not found, installing..."; \
		go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest; \
	else \
		echo "golangci-lint is already installed"; \
	fi
	golangci-lint --version
	golangci-lint run --config .golangci.yml

## goimports: This tool updates your Go import lines, adding missing ones and removing unreferenced ones
.PHONY: goimports
goimports:
	goimports -w .
	@if ! command -v goimports &> /dev/null; then \
		echo "golangci-lint not found, installing..."; \
		go install golang.org/x/tools/cmd/goimports@latest; \
    else \
		echo "goimports is already installed"; \
    fi
	goimports -w .

## tidy: format code and tidy mod file
.PHONY: tidy
tidy:
	go fmt ./...
	go mod tidy -v

## audit: run quality control checks
.PHONY: audit
audit:
	go mod download
	go mod verify
	go vet ./...
	go fmt ./...
#	go install golang.org/x/tools/cmd/goimports@latest && goimports -w .
	goimports -w .
#	go install honnef.co/go/tools/cmd/staticcheck@latest && staticcheck -checks=all,-ST1000,-U1000 ./...
	staticcheck ./...
#	go install golang.org/x/vuln/cmd/govulncheck@latest && govulncheck ./...
	#govulncheck ./...
#	go install github.com/securego/gosec/v2/cmd/gosec@latest && gosec ./...
	gosec ./...
#	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest && golangci-lint run --config .golangci.yml
	golangci-lint run --config .golangci.yml