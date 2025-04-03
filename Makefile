########################################################################################################################
# Config
########################################################################################################################
GOBIN = $(shell go env GOPATH)
GO_LINTER = $(GOBIN)/golangci-lint
GO_LINTER_VERSION = 1.55.0


########################################################################################################################
# Examples
########################################################################################################################

.PHONY: example-server
example-server:
ifndef name
	@echo "Error: 'name' variable required (e.g., make example-server name=dashboard)"
	@exit 1
endif
	go run ./go/examples/$(name)


.PHONY: example-client
example-client:
ifndef name
	@echo "Error: 'name' variable required (e.g., make example-client name=dashboard)"
	@exit 1
endif
	cd ts/examples/$(name) && npm install && npm run dev

.PHONY: example
example:
ifndef name
	@echo "Error: 'name' variable required (e.g., make example name=dashboard)"
	@exit 1
endif
	@make -j2 example-server example-client name=$(name)

########################################################################################################################
# Go
########################################################################################################################

# Go - Test
.PHONY: go-test
go-test:
ifeq ($(cover), true)
	go test -p 1 ./go/... -coverprofile cover.out
else
	go test -p 1 ./go/...
endif

# Go - Lint
.PHONY: go-lint
go-lint: go-linter-check
	$(GO_LINTER) run ./go/...

.PHONY: go-lint-fix
go-lint-fix: go-linter-check
	$(GO_LINTER) run --fix ./go/...

.PHONY: go-linter-check
go-linter-check:
	@if ! [ -x "$(GO_LINTER)" ] || [ "$$($(GO_LINTER) --version | grep -o '[0-9]\.[0-9]\+\.[0-9]\+')" != "$(GO_LINTER_VERSION)" ]; then \
		echo "Installing golangci-lint $(GO_LINTER_VERSION)"; \
		curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(GOBIN) v$(GO_LINTER_VERSION); \
	else \
		echo "golangci-lint $(GO_LINTER_VERSION) is already installed"; \
	fi

########################################################################################################################
# TypeScript
########################################################################################################################

# TypeScript - Test
.PHONY: ts-test
ts-test:
	# TODO: Add npm test command for TS tests
	@echo "TypeScript tests not implemented yet"

# TypeScript - Lint
.PHONY: ts-lint
ts-lint:
	# TODO: Add npm lint command for TS
	@echo "TypeScript linting not implemented yet"
