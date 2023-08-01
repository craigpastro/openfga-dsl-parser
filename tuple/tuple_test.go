package tuple

import (
	"testing"

	openfgav1 "github.com/openfga/api/proto/openfga/v1"
	"github.com/stretchr/testify/require"
)

func TestTupleParser(t *testing.T) {
	tests := []struct {
		name string
		in   string
		out  *openfgav1.TupleKey
	}{
		{
			name: "userID",
			in:   "document:1#viewer@aardvark",
			out: &openfgav1.TupleKey{
				Object:   "document:1",
				Relation: "viewer",
				User:     "aardvark",
			},
		},
		{
			name: "userObject",
			in:   "document:1#viewer@folder:x",
			out: &openfgav1.TupleKey{
				Object:   "document:1",
				Relation: "viewer",
				User:     "folder:x",
			},
		},
		{
			name: "userUserset",
			in:   "document:1#viewer@group:k#viewer",
			out: &openfgav1.TupleKey{
				Object:   "document:1",
				Relation: "viewer",
				User:     "group:k#viewer",
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			res := MustParseTuple(test.in)
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
			_, err := Parse(test.in)
			require.Error(t, err)
		})
	}
}
