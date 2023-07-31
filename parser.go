package parser

import (
	"errors"
	"fmt"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/craigpastro/openfga-dsl-parser/v2/internal/gen/dsl/parser"

	openfgav1 "buf.build/gen/go/openfga/api/protocolbuffers/go/openfga/v1"
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

	typeDefinitions []*openfgav1.TypeDefinition
	relations       map[string]*openfgav1.Relation

	rewrite  []*openfgav1.Userset
	typeInfo []*openfgav1.RelationReference

	errors []error
}

func newDSLListener() *dslListener {
	return &dslListener{
		relations: map[string]*openfgav1.Relation{},
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

	l.typeDefinitions = append(l.typeDefinitions, &openfgav1.TypeDefinition{
		Type:      objectType,
		Relations: l.getRelations(),
		Metadata:  l.getMetadata(),
	})

	// Clear the relations map
	l.relations = map[string]*openfgav1.Relation{}
}

func (l *dslListener) ExitRelation(ctx *parser.RelationContext) {
	name := ctx.GetName().GetText()

	var typeInfo *openfgav1.RelationTypeInfo
	if l.typeInfo != nil {
		typeInfo = &openfgav1.RelationTypeInfo{DirectlyRelatedUserTypes: l.typeInfo}
	}

	l.relations[name] = &openfgav1.Relation{
		Name:     name,
		Rewrite:  l.pop(),
		TypeInfo: typeInfo,
	}

	// Clear the typeInfo array
	l.typeInfo = nil
}

func (l *dslListener) ExitRrType(ctx *parser.RrTypeContext) {
	l.typeInfo = append(l.typeInfo, &openfgav1.RelationReference{
		Type: ctx.GetT().GetText(),
	})
}

func (l *dslListener) ExitRrTypeAndRelation(ctx *parser.RrTypeAndRelationContext) {
	l.typeInfo = append(l.typeInfo, &openfgav1.RelationReference{
		Type: ctx.GetT().GetText(),
		RelationOrWildcard: &openfgav1.RelationReference_Relation{
			Relation: ctx.GetR().GetText(),
		},
	})
}

func (l *dslListener) ExitRrTypeAndWildcard(ctx *parser.RrTypeAndWildcardContext) {
	l.typeInfo = append(l.typeInfo, &openfgav1.RelationReference{
		Type: ctx.GetT().GetText(),
		RelationOrWildcard: &openfgav1.RelationReference_Wildcard{
			Wildcard: &openfgav1.Wildcard{},
		},
	})
}

func (l *dslListener) ExitThis(_ *parser.ThisContext) {
	l.push(&openfgav1.Userset{Userset: &openfgav1.Userset_This{}})
}

func (l *dslListener) ExitTupleToUserset(ctx *parser.TupleToUsersetContext) {
	l.push(&openfgav1.Userset{
		Userset: &openfgav1.Userset_TupleToUserset{
			TupleToUserset: &openfgav1.TupleToUserset{
				Tupleset:        &openfgav1.ObjectRelation{Relation: ctx.GetTupleset().GetText()},
				ComputedUserset: &openfgav1.ObjectRelation{Relation: ctx.GetComputedUserset().GetText()},
			},
		},
	})
}

func (l *dslListener) ExitComputedUserset(ctx *parser.ComputedUsersetContext) {
	l.push(&openfgav1.Userset{
		Userset: &openfgav1.Userset_ComputedUserset{
			ComputedUserset: &openfgav1.ObjectRelation{Relation: ctx.GetComputedUserset().GetText()},
		},
	})
}

func (l *dslListener) ExitUnion(_ *parser.UnionContext) {
	right := l.pop()
	left := l.pop()

	l.push(&openfgav1.Userset{
		Userset: &openfgav1.Userset_Union{
			Union: &openfgav1.Usersets{
				Child: []*openfgav1.Userset{left, right},
			},
		},
	})
}

func (l *dslListener) ExitIntersection(_ *parser.IntersectionContext) {
	right := l.pop()
	left := l.pop()

	l.push(&openfgav1.Userset{
		Userset: &openfgav1.Userset_Intersection{
			Intersection: &openfgav1.Usersets{
				Child: []*openfgav1.Userset{left, right},
			},
		},
	})
}

func (l *dslListener) ExitExclusion(_ *parser.ExclusionContext) {
	subtract := l.pop()
	base := l.pop()

	l.push(&openfgav1.Userset{
		Userset: &openfgav1.Userset_Difference{
			Difference: &openfgav1.Difference{
				Base:     base,
				Subtract: subtract,
			},
		},
	})
}

func (l *dslListener) push(u *openfgav1.Userset) {
	l.rewrite = append(l.rewrite, u)
}

func (l *dslListener) pop() *openfgav1.Userset {
	if len(l.rewrite) < 1 {
		l.errors = append(l.errors, errors.New("missing operand"))
		return nil
	}

	result := l.rewrite[len(l.rewrite)-1]
	l.rewrite = l.rewrite[:len(l.rewrite)-1]

	return result
}

func (l *dslListener) getRelations() map[string]*openfgav1.Userset {
	relations := map[string]*openfgav1.Userset{}
	for name, relation := range l.relations {
		relations[name] = relation.GetRewrite()
	}
	return relations
}

func (l *dslListener) getMetadata() *openfgav1.Metadata {
	relations := map[string]*openfgav1.RelationMetadata{}
	for name, relation := range l.relations {
		directlyRelatedUserTypes := relation.GetTypeInfo().GetDirectlyRelatedUserTypes()
		if directlyRelatedUserTypes != nil {
			relations[name] = &openfgav1.RelationMetadata{DirectlyRelatedUserTypes: directlyRelatedUserTypes}
		}
	}

	return &openfgav1.Metadata{Relations: relations}
}

func MustParse(s string) []*openfgav1.TypeDefinition {
	typeDefinitions, err := Parse(s)
	if err != nil {
		panic(err)
	}

	return typeDefinitions
}

func Parse(s string) ([]*openfgav1.TypeDefinition, error) {
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
