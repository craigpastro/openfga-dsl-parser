.PHONY: gen
gen:
	antlr -Dlanguage=Go -o internal/gen/parser OpenFGA.g4

.PHONY: test
test:
	go test ./...
