#!/usr/bin/env make

.DEFAULT_GOAL  := help
.DEFAULT_SHELL := /bin/sh

GOCOMMAND := go
GOBUILD   := $(GOCOMMAND) build
GOCLEAN   := $(GOCOMMAND) clean
GOTEST    := $(GOCOMMAND) test

BINARY    := ./bin/ikscc

.PHONY: help
help:
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make <target>\n\nTargets:\n"} /^[a-z0-9//_-]+:.*?##/ { printf " %-15s %s\n", $$1, $$2 }' $(MAKEFILE_LIST)

.PHONY: build
build: ## Builds the binary
	@$(GOBUILD) -o $(BINARY) -v main.go

.PHONY: run
run: build ## Runs the latest binary
	@$(BINARY)

.PHONY: test
test: ## Testing all the things
	@$(GOTEST) -v ./...

.PHONY: clean
clean: ## Clean the generated products
	@$(GOCLEAN)
	@rm -f $(BINARY)

.PHONY: all
all: clean build test
