all: help
	@echo "running make help"

build: lint
	@echo "Building and installing to $(GOPATH)/bin/gauntlt..."
	@godep go install ./cmd/gauntlt

test:
	@echo "no tests yet"

lint:
	@echo "Linting with golint..."
	@golint ./...
	@echo "Static Analysis using go vet..."
	@godep go vet ./... | grep -v Godeps |tee /tmp/gauntlt-govet.txt
	@test ! -s /tmp/gauntlt-govet.txt
	@echo "Code formatting with gofmt..."
	@gofmt -l -s . | grep -v Godeps |tee /tmp/gauntl-gofmt.txt
	@test ! -s /tmp/gauntl-gofmt.txt

help:
	@echo "try these:"
	@echo "  make build"
	@echo "  make lint"
	@echo "  make test"
	@echo "  make start"


.phony: build
