GO ?= go

GOFLAGS :=
# Set to 1 to use static linking for all builds (including tests).
STATIC :=

ifeq ($(STATIC),1)
LDFLAGS += -s -w -extldflags "-static"
endif

CMD_DIR := $(CURDIR)/cmd
DIST_DIR := $(CURDIR)/dist


.PHONY: help 
help: ## Show this help message.	
	@egrep -h '\s##\s' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m  %-30s\033[0m %s\n", $$1, $$2}'


.PHONY: all 
all: build vendor test ## Download dependencies, run unit tests, and build the project.
	@echo "Not implemented yet"


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
	@echo "Not implemented yet"