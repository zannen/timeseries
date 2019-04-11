# Makefile

.PHONY: bench coverage deps test

deps:
	go mod download

bench:
	go test -run=X -bench=. -benchmem ./...

coverage:
	go test -coverprofile=coverage.txt -covermode=atomic ./...

codecov_io: coverage
	bash <(curl -s https://codecov.io/bash)

test:
	go test ./...
