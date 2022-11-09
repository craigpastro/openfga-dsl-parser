package parser

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/pelletier/go-toml/v2"
	"github.com/stretchr/testify/require"
	pb "go.buf.build/openfga/go/openfga/api/openfga/v1"
	"google.golang.org/protobuf/encoding/protojson"
)

const testdata = "testdata"

type parseTests struct {
	Tests []parseTest `yaml:"tests"`
}

type parseTest struct {
	Name  string
	Model string
	JSON  string
}

func TestParser(t *testing.T) {
	data, err := os.ReadFile(filepath.Join(testdata, "tests.toml"))
	require.NoError(t, err)

	var tests parseTests
	err = toml.Unmarshal(data, &tests)
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
