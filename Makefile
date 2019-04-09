# Makefile

.PHONY: bench coverage test

bench:
	go test -run=X -bench=. -benchmem ./...

coverage:
	go test -cover ./...

test:
	go test ./...
