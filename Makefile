.PHONY: gen
gen:
	antlr -Dlanguage=Go -o internal/gen/dsl/parser DSL.g4
	antlr -Dlanguage=Go -o internal/gen/tuple/parser Tuple.g4

.PHONY: lint
lint:
	golangci-lint run

.PHONY: test
test:
	go test ./...

.PHONY: all
all: lint test
