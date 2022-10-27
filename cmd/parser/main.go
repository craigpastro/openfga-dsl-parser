package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/craigpastro/openfga-dsl-parser"
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

	fmt.Println(parser.Parse(data))
}
