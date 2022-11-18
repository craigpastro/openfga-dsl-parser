package parser

import (
	"errors"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/craigpastro/openfga-dsl-parser/parser"
	pb "go.buf.build/openfga/go/openfga/api/openfga/v1"
)

type openFGAListener struct {
	*parser.BaseOpenFGAListener

	typeDefinitions []*pb.TypeDefinition
	relations       map[string]*pb.Relation

	rewrite          []*pb.Userset
	computedUserset  string
	typeRestrictions []*pb.RelationReference
}

func newOpenFGAListener() *openFGAListener {
	return &openFGAListener{
		relations: map[string]*pb.Relation{},
	}
}

func (l *openFGAListener) ExitTypeDefinition(ctx *parser.TypeDefinitionContext) {
	objectType := ctx.GetObjectType().GetText()

	l.typeDefinitions = append(l.typeDefinitions, &pb.TypeDefinition{
		Type:      objectType,
		Relations: l.getRelations(),
		Metadata:  l.getMetadata(),
	})

	// Clear the relations map
	l.relations = map[string]*pb.Relation{}
}

func (l *openFGAListener) ExitRelation(ctx *parser.RelationContext) {
	name := ctx.GetName().GetText()

	var typeInfo *pb.RelationTypeInfo
	if l.typeRestrictions != nil {
		typeInfo = &pb.RelationTypeInfo{DirectlyRelatedUserTypes: l.typeRestrictions}
	}

	l.relations[name] = &pb.Relation{
		Name:     name,
		Rewrite:  l.pop(),
		TypeInfo: typeInfo,
	}

	// Clear the typeRestrictions array
	l.typeRestrictions = nil
}

func (l *openFGAListener) ExitTypeRestriction(_ *parser.TypeRestrictionContext) {
	l.push(&pb.Userset{Userset: &pb.Userset_This{}})
}

func (l *openFGAListener) ExitTroiId(ctx *parser.TroiIdContext) {
	l.push(&pb.Userset{
		Userset: &pb.Userset_ComputedUserset{
			ComputedUserset: &pb.ObjectRelation{
				Relation: ctx.GetComputedUserset().GetText(),
			},
		},
	})
}

func (l *openFGAListener) ExitRrType(ctx *parser.RrTypeContext) {
	l.typeRestrictions = append(l.typeRestrictions, &pb.RelationReference{
		Type: ctx.GetT().GetText(),
	})
}

func (l *openFGAListener) ExitRrTypeAndRelation(ctx *parser.RrTypeAndRelationContext) {
	l.typeRestrictions = append(l.typeRestrictions, &pb.RelationReference{
		Type: ctx.GetT().GetText(),
		RelationOrWildcard: &pb.RelationReference_Relation{
			Relation: ctx.GetR().GetText(),
		},
	})
}

func (l *openFGAListener) ExitRrTypeAndWildcard(ctx *parser.RrTypeAndWildcardContext) {
	l.typeRestrictions = append(l.typeRestrictions, &pb.RelationReference{
		Type: ctx.GetT().GetText(),
		RelationOrWildcard: &pb.RelationReference_Wildcard{
			Wildcard: &pb.Wildcard{},
		},
	})
}

func (l *openFGAListener) ExitOrTTU(ctx *parser.OrTTUContext) {
	rewrite := l.pop()

	l.push(&pb.Userset{
		Userset: &pb.Userset_Union{
			Union: &pb.Usersets{
				Child: []*pb.Userset{
					rewrite,
					{
						Userset: &pb.Userset_TupleToUserset{
							TupleToUserset: &pb.TupleToUserset{
								Tupleset:        &pb.ObjectRelation{Relation: ctx.GetTupleset().GetText()},
								ComputedUserset: &pb.ObjectRelation{Relation: ctx.GetComputedUserset().GetText()},
							},
						},
					},
				},
			},
		},
	})
}

func (l *openFGAListener) ExitOrs(_ *parser.OrsContext) {
	children := l.rewrite
	l.rewrite = nil

	l.push(&pb.Userset{
		Userset: &pb.Userset_Union{
			Union: &pb.Usersets{
				Child: children,
			},
		},
	})
}

func (l *openFGAListener) ExitOr(ctx *parser.OrContext) {
	l.push(&pb.Userset{
		Userset: &pb.Userset_ComputedUserset{
			ComputedUserset: &pb.ObjectRelation{
				Relation: ctx.GetId().GetText(),
			},
		},
	})
}

func (l *openFGAListener) ExitAnds(_ *parser.AndsContext) {
	children := l.rewrite
	l.rewrite = nil

	l.push(&pb.Userset{
		Userset: &pb.Userset_Intersection{
			Intersection: &pb.Usersets{
				Child: children,
			},
		},
	})
}

func (l *openFGAListener) ExitAnd(ctx *parser.AndContext) {
	l.push(&pb.Userset{
		Userset: &pb.Userset_ComputedUserset{
			ComputedUserset: &pb.ObjectRelation{
				Relation: ctx.GetId().GetText(),
			},
		},
	})
}

func (l *openFGAListener) ExitExclusion(ctx *parser.ExclusionContext) {
	rewrite := l.pop()

	l.push(&pb.Userset{
		Userset: &pb.Userset_Difference{
			Difference: &pb.Difference{
				Base: rewrite,
				Subtract: &pb.Userset{
					Userset: &pb.Userset_ComputedUserset{
						ComputedUserset: &pb.ObjectRelation{
							Relation: ctx.GetId().GetText(),
						},
					},
				},
			},
		},
	})
}

func (l *openFGAListener) push(u *pb.Userset) {
	l.rewrite = append(l.rewrite, u)
}

func (l *openFGAListener) pop() *pb.Userset {
	if len(l.rewrite) < 1 {
		panic("stack is empty unable to pop")
	}

	result := l.rewrite[len(l.rewrite)-1]
	l.rewrite = l.rewrite[:len(l.rewrite)-1]

	return result
}

func (l *openFGAListener) getRelations() map[string]*pb.Userset {
	relations := map[string]*pb.Userset{}
	for name, relation := range l.relations {
		relations[name] = relation.GetRewrite()
	}
	return relations
}

func (l *openFGAListener) getMetadata() *pb.Metadata {
	relations := map[string]*pb.RelationMetadata{}
	for name, relation := range l.relations {
		directlyRelatedUserTypes := relation.GetTypeInfo().GetDirectlyRelatedUserTypes()
		if directlyRelatedUserTypes != nil {
			relations[name] = &pb.RelationMetadata{DirectlyRelatedUserTypes: directlyRelatedUserTypes}
		}
	}

	return &pb.Metadata{Relations: relations}
}

func MustParse(data string) []*pb.TypeDefinition {
	is := antlr.NewInputStream(data)

	// Create the Lexer
	lexer := parser.NewOpenFGALexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Create the Parser
	p := parser.NewOpenFGAParser(stream)

	// Finally parse the expression
	l := newOpenFGAListener()
	antlr.ParseTreeWalkerDefault.Walk(l, p.Start())

	return l.typeDefinitions
}

func Parse(data string) (typeDefinitions []*pb.TypeDefinition, err error) {
	defer func() {
		if r := recover(); r != nil {
			switch x := r.(type) {
			case string:
				err = errors.New(x)
			case error:
				err = x
			}
			typeDefinitions = nil
		}
	}()

	typeDefinitions = MustParse(data)

	return typeDefinitions, err
}
