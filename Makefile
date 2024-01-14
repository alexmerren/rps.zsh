GO ?= go
LINTER := golangci-lint

GOFLAGS :=
# Set to 1 to use static linking for all builds (including tests).
STATIC :=

ifeq ($(STATIC),1)
LDFLAGS += -s -w -extldflags "-static"
endif

DIST_DIR := $(CURDIR)/dist
INTERNAL_DIR := $(CURDIR)/internal
CMD_DIR := $(CURDIR)/cmd

.PHONY: help 
help: ## Show this help message.	
	@egrep -h '\s##\s' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m  %-30s\033[0m %s\n", $$1, $$2}'

.PHONY: all 
all: vendor fmt lint test build ## Download dependencies, run unit tests, and build the project.

.PHONY: build 
build: ## Download dependencies and build the project. GOFLAGS can be specified for build flags.
	@mkdir -p $(DIST_DIR)
	$(GO) build $(GOFLAGS) -ldflags '$(LDFLAGS)' -mod=vendor -o $(DIST_DIR) $(CMD_DIR)/...

.PHONY: vendor
vendor: ## Vendor dependencies.
	$(GO) mod tidy
	$(GO) mod vendor

.PHONY: test
test: ## Run unit tests.
	$(GO) test ./...

.PHONY: lint
lint: ## Lint the project
	$(LINTER) run ./...

.PHONY: fmt
fmt: ## Format all the code in the project, must be done prior to building for maximum effectiveness.
	$(GO) fmt ./...

DEST_DIR :=
PREFIX := /usr/local
BIN_DIR := ${PREFIX}/bin
DATA_DIR := ${PREFIX}/share
LIST_NAME := rps_list
UNIQUE_NAME := rps_unique

# TODO DONT USE THESE THEY NEED TO BE FIXED

.PHONY: install 
install: ## Install rps to /usr/local/bin
	install -d ${DEST_DIR}${BIN_DIR}
	install -m755 $(DIST_DIR)/$(LIST_NAME) ${DEST_DIR}${BIN_DIR}/_$(LIST_NAME)
	install -m755 $(DIST_DIR)/$(UNIQUE_NAME) ${DEST_DIR}${BIN_DIR}/_$(UNIQUE_NAME)

.PHONY: uninstall
uninstall: ## Uninstall rps and config files
	rm -f ${DEST_DIR}${BIN_DIR}/_$(LIST_NAME)
	rm -f ${DEST_DIR}${BIN_DIR}/_$(UNIQUE_NAME)
