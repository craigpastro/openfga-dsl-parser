package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	parser "github.com/craigpastro/openfga-dsl-parser"
)

func main() {
	var filename string
	flag.StringVar(&filename, "f", "", "specify a file to read")
	flag.Parse()

	var data string
	if filename != "" {
		bytes, err := os.ReadFile(filename)
		if err != nil {
			panic(err)
		}
		data = string(bytes)
	} else {
		data = os.Args[1]
	}

	typeDefinitions, err := parser.Parse(data)
	if err != nil {
		log.Fatal(err)
	}

	// TODO: output proper JSON. Need to add a TypeDefinitions to the api and then we can protojson it.
	fmt.Println(typeDefinitions)
}
