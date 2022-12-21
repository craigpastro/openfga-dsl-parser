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
				TypeDefinitions: MustParseDSL(test.Model),
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
			_, err := ParseDSL(test.model)
			require.Error(t, err)
		})
	}
}

func TestTupleParser(t *testing.T) {
	tests := []struct {
		name string
		in   string
		out  *pb.TupleKey
	}{
		{
			name: "userID",
			in:   "document:1#viewer@aardvark",
			out: &pb.TupleKey{
				Object:   "document:1",
				Relation: "viewer",
				User:     "aardvark",
			},
		},
		{
			name: "userObject",
			in:   "document:1#viewer@folder:x",
			out: &pb.TupleKey{
				Object:   "document:1",
				Relation: "viewer",
				User:     "folder:x",
			},
		},
		{
			name: "userUserset",
			in:   "document:1#viewer@group:k#viewer",
			out: &pb.TupleKey{
				Object:   "document:1",
				Relation: "viewer",
				User:     "group:k#viewer",
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			res, err := ParseTuple(test.in)
			require.NoError(t, err)
			require.Equal(t, test.out.Object, res.Object)
			require.Equal(t, test.out.Relation, res.Relation)
			require.Equal(t, test.out.User, res.User)
		})
	}
}

func TestTupleParserErrors(t *testing.T) {
	tests := []struct {
		name string
		in   string
	}{
		{
			name: "noNamespace",
			in:   ":1#viewer@aardvark",
		},
		{
			name: "noObjectID",
			in:   "document:#viewer@aardvark",
		},
		{
			name: "noRelation",
			in:   "document:1#@aardvark",
		},
		{
			name: "noUserNamespace",
			in:   "document:1#@:x",
		},
		{
			name: "noUserObjectID",
			in:   "document:1#@folder:",
		},
		{
			name: "noUserRelation",
			in:   "document:1#@folder:x#",
		},
		{
			name: "weirdUser",
			in:   "document:1#@id#relation",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			_, err := ParseTuple(test.in)
			require.Error(t, err)
		})
	}
}
