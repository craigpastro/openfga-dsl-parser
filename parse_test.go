package parser

import (
	"testing"

	openfgav1 "buf.build/gen/go/openfga/api/protocolbuffers/go/openfga/v1"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

var opts = cmpopts.IgnoreUnexported(
	openfgav1.TypeDefinition{},
	openfgav1.Userset{},
	openfgav1.ObjectRelation{},
	openfgav1.Userset_Union{},
	openfgav1.Usersets{},
	openfgav1.TupleToUserset{},
	openfgav1.Difference{},
)

// Only testing one type definition as don't want to deal with ordering
func TestParser(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		output []*openfgav1.TypeDefinition
	}{
		{
			name: "Test1",
			input: `
type document
	relations
		define parent as self
		define viewer as self or viewer from parent
`,
			output: []*openfgav1.TypeDefinition{
				{
					Type: "document",
					Relations: map[string]*openfgav1.Userset{
						"parent": {
							Userset: &openfgav1.Userset_This{},
						},
						"viewer": {
							Userset: &openfgav1.Userset_Union{
								Union: &openfgav1.Usersets{
									Child: []*openfgav1.Userset{
										{
											Userset: &openfgav1.Userset_This{},
										},
										{
											Userset: &openfgav1.Userset_TupleToUserset{
												TupleToUserset: &openfgav1.TupleToUserset{
													Tupleset:        &openfgav1.ObjectRelation{Relation: "parent"},
													ComputedUserset: &openfgav1.ObjectRelation{Relation: "viewer"},
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
		{
			name: "Test2",
			input: `
type document
	relations
		define parent as (reader and writer) or (viewer and editor)
`,
			output: []*openfgav1.TypeDefinition{
				{
					Type: "document",
					Relations: map[string]*openfgav1.Userset{
						"parent": {
							Userset: &openfgav1.Userset_Union{
								Union: &openfgav1.Usersets{
									Child: []*openfgav1.Userset{
										{
											Userset: &openfgav1.Userset_Intersection{
												Intersection: &openfgav1.Usersets{
													Child: []*openfgav1.Userset{
														{
															Userset: &openfgav1.Userset_ComputedUserset{
																ComputedUserset: &openfgav1.ObjectRelation{Relation: "reader"},
															},
														},
														{
															Userset: &openfgav1.Userset_ComputedUserset{
																ComputedUserset: &openfgav1.ObjectRelation{Relation: "writer"},
															},
														},
													},
												},
											},
										},
										{
											Userset: &openfgav1.Userset_Intersection{
												Intersection: &openfgav1.Usersets{
													Child: []*openfgav1.Userset{
														{
															Userset: &openfgav1.Userset_ComputedUserset{
																ComputedUserset: &openfgav1.ObjectRelation{Relation: "viewer"},
															},
														},
														{
															Userset: &openfgav1.Userset_ComputedUserset{
																ComputedUserset: &openfgav1.ObjectRelation{Relation: "editor"},
															},
														},
													},
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
		{
			name: "Test3",
			input: `
type document
	relations
		define parent as (reader or writer) and (viewer or editor)
`,
			output: []*openfgav1.TypeDefinition{
				{
					Type: "document",
					Relations: map[string]*openfgav1.Userset{
						"parent": {
							Userset: &openfgav1.Userset_Intersection{
								Intersection: &openfgav1.Usersets{
									Child: []*openfgav1.Userset{
										{
											Userset: &openfgav1.Userset_Union{
												Union: &openfgav1.Usersets{
													Child: []*openfgav1.Userset{
														{
															Userset: &openfgav1.Userset_ComputedUserset{
																ComputedUserset: &openfgav1.ObjectRelation{Relation: "reader"},
															},
														},
														{
															Userset: &openfgav1.Userset_ComputedUserset{
																ComputedUserset: &openfgav1.ObjectRelation{Relation: "writer"},
															},
														},
													},
												},
											},
										},
										{
											Userset: &openfgav1.Userset_Union{
												Union: &openfgav1.Usersets{
													Child: []*openfgav1.Userset{
														{
															Userset: &openfgav1.Userset_ComputedUserset{
																ComputedUserset: &openfgav1.ObjectRelation{Relation: "viewer"},
															},
														},
														{
															Userset: &openfgav1.Userset_ComputedUserset{
																ComputedUserset: &openfgav1.ObjectRelation{Relation: "editor"},
															},
														},
													},
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
		{
			name: "Test4",
			input: `
type document
	relations
		define viewer as writer but not banned`,
			output: []*openfgav1.TypeDefinition{
				{
					Type: "document",
					Relations: map[string]*openfgav1.Userset{
						"viewer": {
							Userset: &openfgav1.Userset_Difference{
								Difference: &openfgav1.Difference{
									Base: &openfgav1.Userset{
										Userset: &openfgav1.Userset_ComputedUserset{
											ComputedUserset: &openfgav1.ObjectRelation{Relation: "writer"},
										},
									},
									Subtract: &openfgav1.Userset{
										Userset: &openfgav1.Userset_ComputedUserset{
											ComputedUserset: &openfgav1.ObjectRelation{Relation: "banned"},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if diff := cmp.Diff(test.output, MustParse(test.input), opts); diff != "" {
				t.Error(diff)
			}
		})
	}
}
