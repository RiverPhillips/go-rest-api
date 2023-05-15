SHELL := /bin/bash
GO = go
export GOBIN ?= $(shell pwd)/bin

ALL_SRC := $(shell find . -name '*.go' \
								-not -path './gen/*' \
							-type f | sort)


GOLANGCI = $(GOBIN)/golangci-lint

$(GOLANGCI): tools/go.sum
	cd tools && go install github.com/golangci/golangci-lint/cmd/golangci-lint

.PHONY: fmt
fmt:
	@echo Running gofmt on ALL_SRC ...
	@$(GOFMT) -e -s -l -w $(ALL_SRC)

.PHONY: vet
vet:
	@echo Running go vet ...
	@$(GO) vet ./...

.PHONY: lint
lint: fmt vet $(GOLANGCI)
	@echo "Running golangci-lint"
	@$(GOLANGCI) run ./...
