package parser

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	pb "go.buf.build/openfga/go/openfga/api/openfga/v1"
)

var opts = cmpopts.IgnoreUnexported(
	pb.TypeDefinition{},
	pb.Userset{},
	pb.ObjectRelation{},
	pb.Userset_Union{},
	pb.Usersets{},
	pb.TupleToUserset{},
	pb.Difference{},
)

// Only testing one type definition as don't want to deal with ordering
func TestParser(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		output []*pb.TypeDefinition
	}{
		{
			name: "Test1",
			input: `
type document
	relations
		define parent as self
		define viewer as self or viewer from parent
`,
			output: []*pb.TypeDefinition{
				{
					Type: "document",
					Relations: map[string]*pb.Userset{
						"parent": {
							Userset: &pb.Userset_This{},
						},
						"viewer": {
							Userset: &pb.Userset_Union{
								Union: &pb.Usersets{
									Child: []*pb.Userset{
										{
											Userset: &pb.Userset_This{},
										},
										{
											Userset: &pb.Userset_TupleToUserset{
												TupleToUserset: &pb.TupleToUserset{
													Tupleset:        &pb.ObjectRelation{Relation: "parent"},
													ComputedUserset: &pb.ObjectRelation{Relation: "viewer"},
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
			output: []*pb.TypeDefinition{
				{
					Type: "document",
					Relations: map[string]*pb.Userset{
						"parent": {
							Userset: &pb.Userset_Union{
								Union: &pb.Usersets{
									Child: []*pb.Userset{
										{
											Userset: &pb.Userset_Intersection{
												Intersection: &pb.Usersets{
													Child: []*pb.Userset{
														{
															Userset: &pb.Userset_ComputedUserset{
																ComputedUserset: &pb.ObjectRelation{Relation: "reader"},
															},
														},
														{
															Userset: &pb.Userset_ComputedUserset{
																ComputedUserset: &pb.ObjectRelation{Relation: "writer"},
															},
														},
													},
												},
											},
										},
										{
											Userset: &pb.Userset_Intersection{
												Intersection: &pb.Usersets{
													Child: []*pb.Userset{
														{
															Userset: &pb.Userset_ComputedUserset{
																ComputedUserset: &pb.ObjectRelation{Relation: "viewer"},
															},
														},
														{
															Userset: &pb.Userset_ComputedUserset{
																ComputedUserset: &pb.ObjectRelation{Relation: "editor"},
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
			output: []*pb.TypeDefinition{
				{
					Type: "document",
					Relations: map[string]*pb.Userset{
						"parent": {
							Userset: &pb.Userset_Intersection{
								Intersection: &pb.Usersets{
									Child: []*pb.Userset{
										{
											Userset: &pb.Userset_Union{
												Union: &pb.Usersets{
													Child: []*pb.Userset{
														{
															Userset: &pb.Userset_ComputedUserset{
																ComputedUserset: &pb.ObjectRelation{Relation: "reader"},
															},
														},
														{
															Userset: &pb.Userset_ComputedUserset{
																ComputedUserset: &pb.ObjectRelation{Relation: "writer"},
															},
														},
													},
												},
											},
										},
										{
											Userset: &pb.Userset_Union{
												Union: &pb.Usersets{
													Child: []*pb.Userset{
														{
															Userset: &pb.Userset_ComputedUserset{
																ComputedUserset: &pb.ObjectRelation{Relation: "viewer"},
															},
														},
														{
															Userset: &pb.Userset_ComputedUserset{
																ComputedUserset: &pb.ObjectRelation{Relation: "editor"},
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
			output: []*pb.TypeDefinition{
				{
					Type: "document",
					Relations: map[string]*pb.Userset{
						"viewer": {
							Userset: &pb.Userset_Difference{
								Difference: &pb.Difference{
									Base: &pb.Userset{
										Userset: &pb.Userset_ComputedUserset{
											ComputedUserset: &pb.ObjectRelation{Relation: "writer"},
										},
									},
									Subtract: &pb.Userset{
										Userset: &pb.Userset_ComputedUserset{
											ComputedUserset: &pb.ObjectRelation{Relation: "banned"},
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
