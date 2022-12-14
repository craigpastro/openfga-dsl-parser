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

func TestParser(t *testing.T) {
	data, err := os.ReadFile(filepath.Join(testdata, "tests.yaml"))
	require.NoError(t, err)

	var tests parseTests
	err = yaml.Unmarshal(data, &tests)
	require.NoError(t, err)

	for _, test := range tests.Tests {
		t.Run(test.Name, func(t *testing.T) {
			model := &pb.AuthorizationModel{
				SchemaVersion:   "1.1",
				TypeDefinitions: Must(Parse(test.Model)),
			}
			bytes, err := protojson.Marshal(model)
			require.NoError(t, err)

			require.JSONEq(t, test.JSON, string(bytes))
		})
	}
}

type errorTests struct {
	Tests []errorTest
}

type errorTest struct {
	Name  string
	Model string
	Error string
}

func TestParserErrors(t *testing.T) {
	data, err := os.ReadFile(filepath.Join(testdata, "errors.yaml"))
	require.NoError(t, err)

	var tests errorTests
	err = yaml.Unmarshal(data, &tests)
	require.NoError(t, err)

	for _, test := range tests.Tests {
		t.Run(test.Name, func(t *testing.T) {
			_, err := Parse(test.Model)
			require.Error(t, err)
		})
	}
}
