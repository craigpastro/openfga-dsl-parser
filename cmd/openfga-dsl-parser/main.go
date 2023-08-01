package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	parser "github.com/craigpastro/openfga-dsl-parser/v2"
	openfgav1 "github.com/openfga/api/proto/openfga/v1"
	"google.golang.org/protobuf/encoding/protojson"
)

func main() {
	readFile := flag.String("f", "", "specify a file to read")
	prettyPrint := flag.Bool("p", false, "pretty print the output")
	outFile := flag.String("o", "", "output to a file")
	flag.Parse()

	var data string
	if *readFile != "" {
		bytes, err := os.ReadFile(*readFile)
		if err != nil {
			panic(err)
		}
		data = string(bytes)
	} else {
		data = os.Args[len(os.Args)-1]
	}

	typeDefinitions, err := parser.Parse(data)
	if err != nil {
		log.Fatalf("error parsing: %s", err)
	}

	model := &openfgav1.AuthorizationModel{
		SchemaVersion:   "1.1",
		TypeDefinitions: typeDefinitions,
	}

	var output string
	if *prettyPrint {
		output = protojson.MarshalOptions{Multiline: true}.Format(model)
	} else {
		output = protojson.MarshalOptions{}.Format(model)
	}

	if *outFile != "" {
		if err := os.WriteFile(*outFile, []byte(output), 0644); err != nil {
			log.Fatalln(err)
		}
	} else {
		fmt.Println(output)
	}
}
