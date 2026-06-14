.PHONY: help build generate test test-verbose test-coverage clean lint format check-format tidy tidy-check check-generate validate-all install-tools

# Variables
GOBASE=$(shell pwd)
GOBIN=$(GOBASE)/bin
GOLANGCI_LINT_VERSION := v2.10.1
GOLANGCI_LINT := $(GOBIN)/golangci-lint
GOIMPORTS := $(GOBIN)/goimports
GINKGO := $(GOBIN)/ginkgo
COVERAGE_DIR := $(CURDIR)/coverage
COVERAGE_PROFILE := $(COVERAGE_DIR)/coverage.out
COVERAGE_HTML := $(COVERAGE_DIR)/coverage.html

.EXPORT_ALL_VARIABLES:

# Default target
all: build

help:
	@echo "🔧 migration-planner-common - Available Targets:"
	@echo ""
	@echo "📦 Build & Generate:"
	@echo "    build:           build all packages"
	@echo "    generate:        regenerate API types from OpenAPI specs"
	@echo "    clean:           clean build artifacts"
	@echo ""
	@echo "🧪 Testing:"
	@echo "    test:            run all tests with coverage"
	@echo "    test-verbose:    run tests with verbose output"
	@echo "    test-coverage:   run tests and open HTML coverage report"
	@echo ""
	@echo "✅ Validation:"
	@echo "    validate-all:    run all validations (lint, format, tidy, generate)"
	@echo "    lint:            run golangci-lint"
	@echo "    format:          format Go code (gofmt + goimports)"
	@echo "    check-format:    verify formatting is up to date"
	@echo "    tidy:            tidy go modules"
	@echo "    tidy-check:      verify go.mod and go.sum are tidy"
	@echo "    check-generate:  verify generated files are up to date"
	@echo ""
	@echo "🛠️  Tools:"
	@echo "    install-tools:   install all required development tools"
	@echo ""

################################################################################
# Build Targets
################################################################################

build:
	@echo "🔨 Building all packages..."
	@go build ./...
	@echo "✅ Build complete."

generate:
	@echo "⚙️  Generating code from OpenAPI specs..."
	@go generate ./...
	@$(MAKE) format
	@echo "✅ Code generation complete."

clean:
	@echo "🗑️  Cleaning build artifacts..."
	@go clean -cache -testcache
	@rm -rf $(COVERAGE_DIR)
	@echo "✅ Clean complete."

################################################################################
# Testing Targets
################################################################################

# Install ginkgo if not already available
$(GINKGO):
	@echo "📦 Installing ginkgo..."
	@mkdir -p $(GOBIN)
	@go install github.com/onsi/ginkgo/v2/ginkgo@v2.27.2
	@echo "✅ 'ginkgo' installed successfully."

test: $(GINKGO)
	@echo "🧪 Running tests..."
	@mkdir -p $(COVERAGE_DIR)
	@$(GINKGO) -v --coverprofile=coverage.out --output-dir=$(COVERAGE_DIR) ./...
	@echo "✅ All tests passed."
	@go tool cover -func=$(COVERAGE_PROFILE) | tail -1
	@echo "📊 Coverage report: $(COVERAGE_HTML)"

test-verbose: $(GINKGO)
	@echo "🧪 Running tests (verbose)..."
	@mkdir -p $(COVERAGE_DIR)
	@$(GINKGO) -v --show-node-events --trace --coverprofile=coverage.out --output-dir=$(COVERAGE_DIR) ./...
	@echo "✅ All tests passed."

test-coverage: test
	@echo "📊 Generating HTML coverage report..."
	@go tool cover -html=$(COVERAGE_PROFILE) -o $(COVERAGE_HTML)
	@echo "📊 Opening coverage report..."
	@xdg-open $(COVERAGE_HTML) 2>/dev/null || open $(COVERAGE_HTML) 2>/dev/null || echo "Coverage report: $(COVERAGE_HTML)"

################################################################################
# Linting Targets
################################################################################

# Check installed golangci-lint version
.PHONY: check-golangci-lint-version
check-golangci-lint-version:
	@if [ -f '$(GOLANGCI_LINT)' ]; then \
		installed=$$('$(GOLANGCI_LINT)' version 2>/dev/null | sed -n 's/.*version \([0-9.]*\).*/\1/p' | head -1); \
		required=$$(echo '$(GOLANGCI_LINT_VERSION)' | sed 's/^v//'); \
		if [ -n "$$installed" ] && [ "$$installed" != "$$required" ]; then \
			echo "🔍 Installed golangci-lint $$installed != required $(GOLANGCI_LINT_VERSION), re-installing..."; \
			rm -f '$(GOLANGCI_LINT)'; \
		fi; \
	fi

# Download golangci-lint if not present
$(GOLANGCI_LINT):
	@echo "📦 Installing golangci-lint $(GOLANGCI_LINT_VERSION)..."
	@mkdir -p $(GOBIN)
	@curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | \
		sh -s -- -b $(GOBIN) $(GOLANGCI_LINT_VERSION)
	@echo "✅ 'golangci-lint' installed successfully."

lint: check-golangci-lint-version $(GOLANGCI_LINT)
	@echo "🔍 Running golangci-lint..."
	@$(GOLANGCI_LINT) run --timeout=5m
	@echo "✅ Lint passed successfully!"

################################################################################
# Formatting Targets
################################################################################

# Install goimports if not already available
$(GOIMPORTS):
	@echo "📦 Installing goimports..."
	@mkdir -p $(GOBIN)
	@go install golang.org/x/tools/cmd/goimports@latest
	@echo "✅ 'goimports' installed successfully."

format: $(GOIMPORTS)
	@echo "🧹 Formatting Go code..."
	@gofmt -s -w .
	@$(GOIMPORTS) -local github.com/kubev2v/migration-planner-common -w .
	@echo "✅ Format complete."

check-format: format
	@echo "🔍 Checking if formatting is up to date..."
	@git diff --quiet || (echo "❌ Detected uncommitted changes after format. Run 'make format' and commit the result." && git status && exit 1)
	@echo "✅ All formatted files are up to date."

################################################################################
# Go Module Targets
################################################################################

tidy:
	@echo "🧹 Tidying go modules..."
	@git ls-files go.mod '**/*go.mod' -z | xargs -0 -I{} bash -xc 'cd $$(dirname {}) && go mod tidy'
	@echo "✅ Go modules tidied successfully."

tidy-check: tidy
	@echo "🔍 Checking if go.mod and go.sum are tidy..."
	@git diff --quiet go.mod go.sum || (echo "❌ Detected uncommitted changes after tidy. Run 'make tidy' and commit the result." && git diff go.mod go.sum && exit 1)
	@echo "✅ go.mod and go.sum are tidy."

################################################################################
# Code Generation Validation
################################################################################

check-generate: generate
	@echo "🔍 Checking if generated files are up to date..."
	@git diff --quiet || (echo "❌ Detected uncommitted changes after generate. Run 'make generate' and commit the result." && git status && exit 1)
	@echo "✅ All generated files are up to date."

################################################################################
# Validation (CI-friendly)
################################################################################

validate-all: lint check-format tidy-check check-generate
	@echo "✅ All validations passed!"

################################################################################
# Tool Installation
################################################################################

install-tools: $(GOLANGCI_LINT) $(GOIMPORTS) $(GINKGO)
	@echo "✅ All development tools installed."

################################################################################
# Emoji Legend
#
# Action Type        | Emoji | Description
# -------------------|--------|------------------------------------------------
# Install tool        📦     Installing a dependency or binary
# Running task        ⚙️     Executing tasks like generate, build, etc.
# Build              🔨     Building binaries or packages
# Linting/validation  🔍     Checking format, lint, static analysis, etc.
# Formatting          🧹     Formatting source code
# Testing            🧪     Running tests
# Success/complete    ✅     Task completed successfully
# Failure/alert       ❌     An error or failure occurred
# Teardown/cleanup    🗑️     Cleaning up resources
# Documentation      📊     Generating or viewing documentation/reports
# Tools/utilities     🛠️     Installing or managing dev tools
# Configuration      🔧     Configuration-related tasks
################################################################################
