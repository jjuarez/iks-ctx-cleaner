#!/usr/bin/env make

.DEFAULT_GOAL  := help
.DEFAULT_SHELL := /bin/sh

GOCOMMAND := go
GOBUILD   := $(GOCOMMAND) build
GOCLEAN   := $(GOCOMMAND) clean
GOTEST    := $(GOCOMMAND) test
GOLINT    := $(shell command -v golangci-lint 2>/dev/null)

EXECUTABLE   := ikscc
INSTALL_PATH := /usr/local/bin
INSTALL      := $(INSTALL_PATH)/$(EXECUTABLE)

PROJECT_CHANGESET := $(shell git rev-parse --verify HEAD 2>/dev/null)


define assert-command
	@$(if $(shell command -v $1 2>/dev/null),,$(error $(1) command not found))
endef


.PHONY: help
help:
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make <target>\n\nTargets:\n"} /^[a-z0-9\/_-]+:.*?##/ { printf " %-15s %s\n", $$1, $$2 }' $(MAKEFILE_LIST)


.PHONY: lint
lint: ## Lint
	$(call assert-command,$(GOLINT))
	@$(GOLINT) --verbose run ./...


./bin/$(EXECUTABLE):
ifdef PROJECT_VERSION
	@$(GOBUILD) -v -ldflags="-s -w -X github.com/jjuarez/iks-ctx-cleaner/cmd.Version='v$(PROJECT_VERSION)'"          -o ./bin/$(EXECUTABLE) main.go
else
	@$(GOBUILD) -v -ldflags="-s -w -X github.com/jjuarez/iks-ctx-cleaner/cmd.Version='nightly:$(PROJECT_CHANGESET)'" -o ./bin/$(EXECUTABLE) main.go
endif


.PHONY: build
build: ./bin/$(EXECUTABLE) ## Builds the binary


.PHONY: test
test: ## Testing all the things
	@$(GOTEST) -v ./...


.PHONY: clean
clean: ## Clean the generated products
	@$(GOCLEAN)
	@rm -fr ./bin/$(EXECUTABLE) ./dist/*


$(INSTALL):
	@install -m 0755 ./bin/$(EXECUTABLE) $(INSTALL_PATH)

.PHONY: build install
install: $(INSTALL) ## Installs the program in the system


.PHONY: uninstall
uninstall: ## Uninstalls the program from the system
	@rm -f $(INSTALL)


.PHONY: all
all: lint build test install
