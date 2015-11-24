all: help
	@echo "running make help"

build: lint
	@echo "Building to pkg/gauntlt..."
	@godep go build -o pkg/gauntlt-go ./main.go

test:
	@echo "no tests yet"

lint:
	@echo "Linting with golint..."
	@golint ./...
	@echo "Static Analysis using go vet..."
	@godep go vet ./... | grep -v Godeps |tee /tmp/gauntlt-go-govet.txt
	@test ! -s /tmp/gauntlt-go-govet.txt
	@echo "Code formatting with gofmt..."
	@gofmt -l -s . | grep -v Godeps |tee /tmp/gauntl-go-gofmt.txt
	@test ! -s /tmp/gauntl-go-gofmt.txt

help:
	@echo "try these:"
	@echo "  make build"
	@echo "  make lint"
	@echo "  make test"
	@echo "  make start"


.phony: build
