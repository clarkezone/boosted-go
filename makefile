.DEFAULT_GOAL := help

BIN_NAME := "boosted-go"
MAIN_BRANCH := main
HEAD_BRANCH := HEAD
ifeq ($(strip $(VERSION_HASH)),)

# hash of current commit
VERSION_HASH := $(shell git rev-parse --short HEAD)

# tag matching current commit or empty
HEAD_TAG := $(shell git tag --points-at HEAD)

#name of branch
BRANCH_NAME := $(shell git rev-parse --abbrev-ref HEAD)
endif

VERSION_STRING := $(BRANCH_NAME)

#if we are on HEAD and there is a tag pointing at head, use that for version else use branch name as version
ifeq ($(BRANCH_NAME),$(HEAD_BRANCH))

$(info match head)
ifneq ($(strip $(HEAD_TAG)),)
VERSION_STRING := $(HEAD_TAG)
$(info    $(version_string))
endif
endif

BINDIR    := $(CURDIR)/bin
PLATFORMS := linux/amd64/rk-Linux-x86_64 darwin/amd64/rk-Darwin-x86_64 windows/amd64/rk.exe linux/arm64/rk-Linux-arm64 darwin/arm64/rk-Darwin-arm64

UNAME := $(shell uname)
ifeq ($(UNAME), Darwin)
SHACOMMAND := shasum -a 256
else
SHACOMMAND := sha256sum
endif

### -------------------------------------
### Help
### -------------------------------------

##@ Help:

help: ## Display this help.
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[0-9a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-25s\033[0m %s\n    ", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)


### -------------------------------------
### Install
### -------------------------------------

##@ Install:

.PHONY: install-binaries
install-tools:     ## install tooling
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest && \
	go install github.com/uw-labs/strongbox@latesta && \
	go install github.com/mgechev/revive@latest

### -------------------------------------
### Install
### -------------------------------------

##@ Dev utils:
.PHONY: test
test:   ## run tests
	go test -p 4 -coverprofile=coverage.txt -covermode=atomic ./...

.PHONY: dep
dep:    ## ensure dependencies
	go mod tidy

.PHONY: latest
latest: ## write version
	echo ${VERSION_STRING} > bin/latest

.PHONY: lint
lint:   ## run linter
	revive $(shell go list ./...)
	go vet $(shell go list ./...)
	golangci-lint run

### -------------------------------------
### Setup
### -------------------------------------

##@ procommit:
.PHONY: precommit-installhooks
precommit-installhooks:   ## install git hooks
	pre-commit install

.PHONY: precommit
precommit:   ## precommit
	pre-commit run --all-files
