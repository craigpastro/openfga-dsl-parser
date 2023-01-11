# OpenFGA DSL Parser

OpenFGA DSL Parser is a parser for the OpenFGA DSL.

> **Note:** This library parses an unofficial version of the OpenFGA DSL. Some differences:
> - No `schema version` field
> - Still expects `self`
> - Add brackets (possibly at the expense of more "deeply" generated JSON)
>
> See the [tests](./testdata/tests.yaml) for some examples.

## Example usage

```go
package main

import (
	"fmt"

	parser "github.com/craigpastro/openfga-dsl-parser/v2"
)

func main() {
	model := `
type user

type document
  relations
    define viewer: [user] as self
    define writer: [user] as self or writer
`

	typeDefinitions, err := parser.Parse(model)
	if err != nil {
		panic(err)
	}

	fmt.Println(typeDefinitions)
}
```

## Contributing

1. Change the grammar in `DSL.g4` or `Tuple.g4`.
2. Regenerate the `parser` package by `make gen`.
3. Make the appropriate changes to the parsing function. 

## Wasm module

Create a directory called `assets` to hold all the code that will be served by a webserver.

Build the Wasm module:
```
GOOS=js GOARCH=wasm go build -o assets/parse.wasm cmd/wasm/main.go
```

In order to use the Wasm module you will need `wasm_exec.js` which you find in your GOROOT (assuming you have Go installed.)
```
cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" assets
```

Finally, create an `index.html` in `assets` that will import the Wasm module. Below is a very simple example:
```html
<html>
    <head>
        <meta charset="utf-8"/>
        <script src="wasm_exec.js"></script>
        <script>
            const go = new Go();
            WebAssembly.instantiateStreaming(fetch("parse.wasm"), go.importObject).then((result) => {
                go.run(result.instance);
            });
        </script>
    </head>
    
    <body>
         <textarea id="input" name="input" cols="80" rows="20"></textarea>
         <input id="button" type="submit" name="button" value="parse" onclick="parse(input.value)"/>
         <textarea id="output" name="output" cols="80" rows="20"></textarea>
    </body>
    
    <script>
        var parse = function(input) {
            output.value = parseDSL(input)
        }
     </script>
</html>
```
