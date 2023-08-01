package main

import (
	"log"
	"syscall/js"

	parser "github.com/craigpastro/openfga-dsl-parser"
	openfgav1 "github.com/openfga/api/proto/openfga/v1"
	"google.golang.org/protobuf/encoding/protojson"
)

// Basically copied from https://golangbot.com/webassembly-using-go/.
func parserWrapper() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) any {
		if len(args) != 1 {
			return "invalid no of arguments passed"
		}

		typeDefinitions, err := parser.Parse(args[0].String())
		if err != nil {
			log.Printf("unable to parse: %s\n", err)
			return err.Error()
		}

		model := &openfgav1.AuthorizationModel{
			SchemaVersion:   "1.0",
			TypeDefinitions: typeDefinitions,
		}

		return protojson.MarshalOptions{Multiline: true}.Format(model)
	})
}

func main() {
	js.Global().Set("parseDSL", parserWrapper())
	<-make(chan bool) // so main doesn't exit
}
