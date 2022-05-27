PACKAGE = stmt-builder
TMPDIR ?= $(shell dirname $$(mktemp -u))
COVER_FILE ?= $(TMPDIR)/$(PACKAGE)-coverage.out

.PHONY: tools
tools: ## Install all needed tools, e.g. for static checks
	@echo Installing tools from tools.go
	@grep '_ "' tools.go | grep -o '"[^"]*"' | xargs -tI % go install %

.PHONY: lint
lint: tools ## Check the project with lint
	golint -set_exit_status ./...

.PHONY: test
test: ## Run unit (short) tests
	go test -short ./... -coverprofile=$(COVER_FILE)
	go tool cover -func=$(COVER_FILE) | grep ^total