package parser

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
	pb "go.buf.build/openfga/go/openfga/api/openfga/v1"
	"google.golang.org/protobuf/encoding/protojson"
	"gopkg.in/yaml.v3"
)

const testdata = "testdata"

type parseTests struct {
	Tests []parseTest
}

type parseTest struct {
	Name  string
	Model string
	JSON  string
}

func TestDSLParser(t *testing.T) {
	data, err := os.ReadFile(filepath.Join(testdata, "tests.yaml"))
	require.NoError(t, err)

	var tests parseTests
	err = yaml.Unmarshal(data, &tests)
	require.NoError(t, err)

	for _, test := range tests.Tests {
		t.Run(test.Name, func(t *testing.T) {
			model := &pb.AuthorizationModel{
				SchemaVersion:   "1.1",
				TypeDefinitions: MustParse(test.Model),
			}
			bytes, err := protojson.Marshal(model)
			require.NoError(t, err)

			require.JSONEq(t, test.JSON, string(bytes))
		})
	}
}

func TestDSLParserErrors(t *testing.T) {
	tests := []struct {
		name  string
		model string
	}{
		{
			name: "TupleToUsersetWithNoComputedUserset",
			model: `
type user

type document
	relations
    	define parent: [folder] as self
        define viewer as parent from
`,
		},
		{
			name: "UnionWithOneProjection",
			model: `
type user
        
type document
	relations
    	define parent: [folder] as self or
`,
		},
		{
			name: "IntersectionWithOneProjection",
			model: `
type user
        
type document
	relations
    	define parent: [folder] as self and
`,
		},
		{
			name: "ExclusionWithNoSubtrahend",
			model: `
type user
        
type document
	relations
    	define parent: [folder] as self but not
`,
		},
		{
			name: "ColonWithNoType",
			model: `
type document
        relations
          define parent: as self
`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			_, err := Parse(test.model)
			require.Error(t, err)
		})
	}
}
