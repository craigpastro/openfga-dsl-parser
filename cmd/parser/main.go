package main

import (
	"flag"
	"fmt"
	"os"

	parser "github.com/craigpastro/openfga-dsl-parser"
	pb "go.buf.build/openfga/go/openfga/api/openfga/v1"
	"google.golang.org/protobuf/encoding/protojson"
)

func main() {
	var filename string
	flag.StringVar(&filename, "f", "", "specify a file to read")

	var prettyPrint bool
	flag.BoolVar(&prettyPrint, "p", false, "pretty print the output")
	flag.Parse()

	var data string
	if filename != "" {
		bytes, err := os.ReadFile(filename)
		if err != nil {
			panic(err)
		}
		data = string(bytes)
	} else {
		data = os.Args[len(os.Args)-1]
	}

	output := &pb.AuthorizationModel{
		SchemaVersion:   "1.0",
		TypeDefinitions: parser.MustParse(data),
	}

	if prettyPrint {
		fmt.Println(protojson.MarshalOptions{Multiline: true}.Format(output))
	} else {
		fmt.Println(protojson.MarshalOptions{}.Format(output))
	}
}
