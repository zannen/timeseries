# Makefile

.PHONY: all bench codecov_io coverage default deps example test

default: all

all: test bench coverage example

deps:
	@echo "Making $@"
	@go mod download

bench: deps
	@echo "Making $@"
	@go test -run=X -bench=. -benchmem ./...

coverage: deps
	@echo "Making $@"
	@go test -coverprofile=coverage.txt -covermode=atomic ./...

codecov_io: coverage
	@echo "Making $@"
	@curl --silent -D response_headers --output codecov.sh https://codecov.io/bash || exit 1
	@head -n1 response_headers | grep -q "^HTTP/.* 200" || exit 2
	@bash codecov.sh

example: deps
	@echo "Making $@"
	@go run ./examples/timeseries_example.go

test: deps
	@echo "Making $@"
	@go test ./...
