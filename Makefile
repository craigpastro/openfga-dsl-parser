.PHONY: gen
gen:
	antlr -Dlanguage=Go -o internal/gen/parser OpenFGA.g4

.PHONY: lint
lint:
	golangci-lint run

.PHONY: test
test:
	go test ./...

.PHONY: all
all: lint test
