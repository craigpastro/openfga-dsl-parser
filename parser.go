package parser

import (
	"errors"
	"fmt"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/craigpastro/openfga-dsl-parser/v2/internal/gen/dsl/parser"
	pb "go.buf.build/openfga/go/openfga/api/openfga/v1"
	"go.uber.org/multierr"
)

type syntaxError struct {
	line, column int
	msg          string
}

func (e *syntaxError) Error() string {
	return fmt.Sprintf("error at %d:%d: %s", e.line, e.column, e.msg)
}

type dslListener struct {
	*parser.BaseDSLListener
	*antlr.DefaultErrorListener

	typeDefinitions []*pb.TypeDefinition
	relations       map[string]*pb.Relation

	rewrite  []*pb.Userset
	typeInfo []*pb.RelationReference

	errors []error
}

func newDSLListener() *dslListener {
	return &dslListener{
		relations: map[string]*pb.Relation{},
	}
}

func (l *dslListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	l.errors = append(l.errors, &syntaxError{
		line:   line,
		column: column,
		msg:    msg,
	})
}

func (l *dslListener) ExitTypeDefinition(ctx *parser.TypeDefinitionContext) {
	objectType := ctx.GetObjectType().GetText()

	l.typeDefinitions = append(l.typeDefinitions, &pb.TypeDefinition{
		Type:      objectType,
		Relations: l.getRelations(),
		Metadata:  l.getMetadata(),
	})

	// Clear the relations map
	l.relations = map[string]*pb.Relation{}
}

func (l *dslListener) ExitRelation(ctx *parser.RelationContext) {
	name := ctx.GetName().GetText()

	var typeInfo *pb.RelationTypeInfo
	if l.typeInfo != nil {
		typeInfo = &pb.RelationTypeInfo{DirectlyRelatedUserTypes: l.typeInfo}
	}

	l.relations[name] = &pb.Relation{
		Name:     name,
		Rewrite:  l.pop(),
		TypeInfo: typeInfo,
	}

	// Clear the typeInfo array
	l.typeInfo = nil
}

func (l *dslListener) ExitRrType(ctx *parser.RrTypeContext) {
	l.typeInfo = append(l.typeInfo, &pb.RelationReference{
		Type: ctx.GetT().GetText(),
	})
}

func (l *dslListener) ExitRrTypeAndRelation(ctx *parser.RrTypeAndRelationContext) {
	l.typeInfo = append(l.typeInfo, &pb.RelationReference{
		Type: ctx.GetT().GetText(),
		RelationOrWildcard: &pb.RelationReference_Relation{
			Relation: ctx.GetR().GetText(),
		},
	})
}

func (l *dslListener) ExitRrTypeAndWildcard(ctx *parser.RrTypeAndWildcardContext) {
	l.typeInfo = append(l.typeInfo, &pb.RelationReference{
		Type: ctx.GetT().GetText(),
		RelationOrWildcard: &pb.RelationReference_Wildcard{
			Wildcard: &pb.Wildcard{},
		},
	})
}

func (l *dslListener) ExitThis(_ *parser.ThisContext) {
	l.push(&pb.Userset{Userset: &pb.Userset_This{}})
}

func (l *dslListener) ExitTupleToUserset(ctx *parser.TupleToUsersetContext) {
	l.push(&pb.Userset{
		Userset: &pb.Userset_TupleToUserset{
			TupleToUserset: &pb.TupleToUserset{
				Tupleset:        &pb.ObjectRelation{Relation: ctx.GetTupleset().GetText()},
				ComputedUserset: &pb.ObjectRelation{Relation: ctx.GetComputedUserset().GetText()},
			},
		},
	})
}

func (l *dslListener) ExitComputedUserset(ctx *parser.ComputedUsersetContext) {
	l.push(&pb.Userset{
		Userset: &pb.Userset_ComputedUserset{
			ComputedUserset: &pb.ObjectRelation{Relation: ctx.GetComputedUserset().GetText()},
		},
	})
}

func (l *dslListener) ExitUnion(_ *parser.UnionContext) {
	right := l.pop()
	left := l.pop()

	l.push(&pb.Userset{
		Userset: &pb.Userset_Union{
			Union: &pb.Usersets{
				Child: []*pb.Userset{left, right},
			},
		},
	})
}

func (l *dslListener) ExitIntersection(_ *parser.IntersectionContext) {
	right := l.pop()
	left := l.pop()

	l.push(&pb.Userset{
		Userset: &pb.Userset_Intersection{
			Intersection: &pb.Usersets{
				Child: []*pb.Userset{left, right},
			},
		},
	})
}

func (l *dslListener) ExitExclusion(_ *parser.ExclusionContext) {
	subtract := l.pop()
	base := l.pop()

	l.push(&pb.Userset{
		Userset: &pb.Userset_Difference{
			Difference: &pb.Difference{
				Base:     base,
				Subtract: subtract,
			},
		},
	})
}

func (l *dslListener) push(u *pb.Userset) {
	l.rewrite = append(l.rewrite, u)
}

func (l *dslListener) pop() *pb.Userset {
	if len(l.rewrite) < 1 {
		l.errors = append(l.errors, errors.New("missing operand"))
		return nil
	}

	result := l.rewrite[len(l.rewrite)-1]
	l.rewrite = l.rewrite[:len(l.rewrite)-1]

	return result
}

func (l *dslListener) getRelations() map[string]*pb.Userset {
	relations := map[string]*pb.Userset{}
	for name, relation := range l.relations {
		relations[name] = relation.GetRewrite()
	}
	return relations
}

func (l *dslListener) getMetadata() *pb.Metadata {
	relations := map[string]*pb.RelationMetadata{}
	for name, relation := range l.relations {
		directlyRelatedUserTypes := relation.GetTypeInfo().GetDirectlyRelatedUserTypes()
		if directlyRelatedUserTypes != nil {
			relations[name] = &pb.RelationMetadata{DirectlyRelatedUserTypes: directlyRelatedUserTypes}
		}
	}

	return &pb.Metadata{Relations: relations}
}

func MustParse(s string) []*pb.TypeDefinition {
	typeDefinitions, err := Parse(s)
	if err != nil {
		panic(err)
	}

	return typeDefinitions
}

func Parse(s string) ([]*pb.TypeDefinition, error) {
	is := antlr.NewInputStream(s)

	listener := newDSLListener()

	lexer := parser.NewDSLLexer(is)
	lexer.RemoveErrorListeners()
	lexer.AddErrorListener(listener)

	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	parser := parser.NewDSLParser(stream)
	parser.RemoveErrorListeners()
	parser.AddErrorListener(listener)

	antlr.ParseTreeWalkerDefault.Walk(listener, parser.Dsl())

	if len(listener.errors) > 0 {
		var err error
		for _, e := range listener.errors {
			err = multierr.Append(err, e)
		}
		return nil, err
	}

	return listener.typeDefinitions, nil
}
