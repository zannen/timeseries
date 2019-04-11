# Makefile

.PHONY: bench coverage deps test

deps:
	@echo "Making $@"
	@go mod download

bench:
	@echo "Making $@"
	@go test -run=X -bench=. -benchmem ./...

coverage:
	@echo "Making $@"
	@go test -coverprofile=coverage.txt -covermode=atomic ./...

codecov_io: coverage
	@echo "Making $@"
	@curl --silent -D response_headers --output codecov.sh https://codecov.io/bash || exit 1
	@head -n1 response_headers | grep -q "^HTTP/.* 200" || exit 2
	@bash codecov.sh

test:
	@echo "Making $@"
	@go test ./...
