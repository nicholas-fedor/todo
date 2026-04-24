# Todo Makefile
# Standardized build, test, lint, and release targets

# Variables
BINARY_NAME=todo
BINARY_DIR=bin
GO=go
GOLANGCI_LINT=golangci-lint
GORELEASER=goreleaser
DOCKER=docker

# Build settings
GOFLAGS=-trimpath
LDFLAGS=-s -w
LDFLAGS_VERSION=-X main.version=$(VERSION)
LDFLAGS_COMMIT=-X main.commit=$(COMMIT)
LDFLAGS_DATE=-X main.date=$(DATE)

# Default target
.PHONY: help
help: ## Show this help message
	@echo "Todo Project Makefile"
	@echo ""
	@echo "Available targets:"
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "  %-15s %s\n", $$1, $$2}' $(MAKEFILE_LIST)

# =============================================================================
# Development Targets
# =============================================================================

.PHONY: build test test-verbose bench lint vet fmt run install tidy

build: ## Build the application binary
	$(GO) build $(GOFLAGS) -o $(BINARY_DIR)/$(BINARY_NAME) ./...

test: ## Run all tests with coverage
	$(GO) test -timeout 30s -v -coverprofile=coverage.out -covermode=atomic ./...

test-verbose: ## Run all tests with verbose output
	$(GO) test -v ./...

bench: ## Run benchmarks
	$(GO) test -bench=. -benchmem ./...

lint: ## Run linter (fails on issues)
	$(GOLANGCI_LINT) run --config build/golangci-lint/golangci-lint.yaml ./...

vet: ## Run Go vet
	$(GO) vet ./...

fmt: ## Format code using goimports
	$(GOLANGCI_LINT) fmt --config build/golangci-lint/golangci-lint.yaml ./...

run: ## Run the application
	$(GO) run ./cmd/server

install: ## Install the application
	$(GO) install ./...

tidy: ## Tidy Go modules and verify consistency
	$(GO) mod tidy
	$(GO) mod verify

# =============================================================================
# Dependency Management
# =============================================================================

.PHONY: deps mod-tidy mod-download

deps: mod-download ## Download all dependencies
	$(GO) get -t -d ./...

mod-tidy: tidy ## Tidy and verify Go modules (alias for tidy)

mod-download: ## Download Go module dependencies
	$(GO) mod download

.PHONY: mod-graph mod-why
mod-graph: ## Show module dependency graph
	$(GO) mod graph

mod-why: ## Explain why a module is needed
	@read -p "Enter module name: " module; \
	$(GO) mod why $$module

# =============================================================================
# Release Targets
# =============================================================================

.PHONY: release release-stable release-nightly

release: release-stable ## Create a new stable release using GoReleaser
	@echo "Release completed"

release-stable: ## Create a new stable release
	$(GORELEASER) release --config build/goreleaser/stable.yaml --clean

release-nightly: ## Create a new nightly release
	$(GORELEASER) release --config build/goreleaser/nightly.yaml --clean

.PHONY: check
check: lint vet test ## Run lint, vet, and test

# =============================================================================
# Docker Targets
# =============================================================================

.PHONY: docker-build docker-run docker-push docker-clean

docker-build: ## Build Docker image
	$(DOCKER) build -t $(BINARY_NAME) -f build/docker/Dockerfile .

docker-run: ## Run Docker container
	$(DOCKER) run -p 3000:3000 $(BINARY_NAME)

docker-push: docker-build ## Push Docker images (requires DOCKER_NAMESPACE and GHCR_NAMESPACE env vars)
	$(DOCKER) push $(DOCKER_NAMESPACE)/$(BINARY_NAME)
	$(DOCKER) push ghcr.io/$(GHCR_NAMESPACE)/$(BINARY_NAME)

docker-clean: ## Remove Docker images
	$(DOCKER) rmi $(BINARY_NAME) || true

# =============================================================================
# Utility Targets
# =============================================================================

.PHONY: clean clean-all generate

clean: ## Clean build artifacts
	rm -rf $(BINARY_DIR)/
	rm -f coverage.out

clean-all: clean ## Clean all generated files including vendor
	rm -rf tmp/
	rm -rf vendor/

generate: ## Generate code (templ assets, etc.)
	templ generate
	@echo "Code generation complete"

# Note: Mockery mocks generation is not currently configured
# To enable: 1) add mockery to go.mod, 2) create build/mockery/mockery.yaml
# .PHONY: mocks
# mocks: ## Generate mocks (requires mockery config)
# 	$(MOCKERY) --config build/mockery/mockery.yaml

.PHONY: air
air: ## Run development server with Air hot reload
	air

.PHONY: dev
dev: ## Start full development stack (Templ + Tailwind + Go)
	@echo "Starting development: Tailwind (watch) + Air (templ + Go) in parallel"
	@echo "Press Ctrl+C to stop all processes."
	@echo ""
	@# Start both processes in parallel and wait
	@( \
		TEMPLUI_PATH="$$(go list -mod=mod -m -f '{'{{.Dir}'}}' github.com/templui/templui)" && \
		printf '%s\n' '@source "./**/*.templ";' "@source \"$$TEMPLUI_PATH/components/**/*.templ\";" > ./internal/web/assets/css/sources.generated.css && \
		bunx @tailwindcss/cli -i ./internal/web/assets/css/input.css -o ./internal/web/assets/css/output.css --watch \
	) & \
	TAILWIND_PID=$$!; \
	( \
		sleep 0.5 && \
		echo "[Tailwind] started (PID $$TAILWIND_PID)" \
	) & \
	( \
		air \
	) & \
	WAIT_PIDS="$$TAILWIND_PID $$!"; \
	echo "[Air] started (PID $$!)"; \
	echo ""; \
	echo "Both processes running. Press Ctrl+C to stop."; \
	trap "echo ''; echo 'Stopping...'; kill $$TAILWIND_PID 2>/dev/null; wait $$TAILWIND_PID 2>/dev/null; exit 0" INT TERM; \
	wait $$WAIT_PIDS
install-tools: ## Install development tools (Go-based)
	@echo "Installing development tools..."
	@command -v golangci-lint >/dev/null 2>&1 || (echo "Installing golangci-lint..." && go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest)
	@command -v goreleaser >/dev/null 2>&1 || (echo "Installing goreleaser..." && go install github.com/goreleaser/goreleaser/v2@latest)
	@command -v air >/dev/null 2>&1 || (echo "Installing air..." && go install github.com/air-verse/air@latest)
	@command -v templ >/dev/null 2>&1 || (echo "Installing templ..." && go install github.com/a-h/templ/cmd/templ@latest)
	@command -v templui >/dev/null 2>&1 || (echo "Installing templui CLI..." && go install github.com/templui/templui/cmd/templui@latest)
	@echo ""
	@echo "NOTE: Install Tailwind CSS separately:"
	@echo "  Using bun (recommended for this project):"
	@echo "    bun add -D tailwindcss"
	@echo "  Or download standalone from: https://github.com/tailwindlabs/tailwindcss/releases"
	@echo ""
	@echo "All Go tools installed!"

# =============================================================================
# CI/CD Targets
# =============================================================================

.PHONY: ci-check
ci-check: ## Run all checks as in CI
	@echo "Running CI checks..."
	@$(MAKE) tidy
	@$(MAKE) generate
	@$(MAKE) lint
	@$(MAKE) vet
	@$(MAKE) test
	@echo "CI checks passed!"

# =============================================================================
# Security Targets
# =============================================================================

.PHONY: gosec audit
gosec: ## Run security linter
	$(GOLANGCI_LINT) run --enable=gosec ./...

audit: ## Check for outdated dependencies
	@echo "Checking for dependency updates..."
	@$(GO) list -u -m all

.PHONY: help-extra
help-extra: ## Show extended help with all targets
	@echo ""
	@echo "Additional targets:"
	@echo "  tidy          - Tidy and verify Go modules"
	@echo "  deps          - Download all dependencies"
	@echo "  mod-graph     - Show module dependency graph"
	@echo "  mod-why       - Explain why a module is needed"
	@echo "  ci-check      - Run all CI checks (tidy, lint, vet, test)"
	@echo "  gosec         - Run security linter"
	@echo "  audit         - Check for outdated dependencies"
	@echo "  generate      - Generate templ files"
	@echo "  air           - Run development server with hot reload"
	@echo "  install-tools - Install development tools"
	@echo "  clean-all     - Clean all generated files"
