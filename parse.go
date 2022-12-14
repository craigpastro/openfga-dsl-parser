package parser

import (
	"errors"
	"fmt"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/craigpastro/openfga-dsl-parser/v2/internal/gen/parser"
	pb "go.buf.build/openfga/go/openfga/api/openfga/v1"
	"go.uber.org/multierr"
)

type openFGAListener struct {
	*parser.BaseOpenFGAListener
	*antlr.DefaultErrorListener

	typeDefinitions []*pb.TypeDefinition
	relations       map[string]*pb.Relation

	rewrite  []*pb.Userset
	typeInfo []*pb.RelationReference

	errors []error
}

func newOpenFGAListener() *openFGAListener {
	return &openFGAListener{
		relations: map[string]*pb.Relation{},
	}
}

type syntaxError struct {
	line, column int
	msg          string
}

func (e *syntaxError) Error() string {
	return fmt.Sprintf("error at %d:%d: %s", e.line, e.column, e.msg)
}

func (l *openFGAListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	l.errors = append(l.errors, &syntaxError{
		line:   line,
		column: column,
		msg:    msg,
	})
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

func (l *openFGAListener) ExitRrType(ctx *parser.RrTypeContext) {
	l.typeInfo = append(l.typeInfo, &pb.RelationReference{
		Type: ctx.GetT().GetText(),
	})
}

func (l *openFGAListener) ExitRrTypeAndRelation(ctx *parser.RrTypeAndRelationContext) {
	l.typeInfo = append(l.typeInfo, &pb.RelationReference{
		Type: ctx.GetT().GetText(),
		RelationOrWildcard: &pb.RelationReference_Relation{
			Relation: ctx.GetR().GetText(),
		},
	})
}

func (l *openFGAListener) ExitRrTypeAndWildcard(ctx *parser.RrTypeAndWildcardContext) {
	l.typeInfo = append(l.typeInfo, &pb.RelationReference{
		Type: ctx.GetT().GetText(),
		RelationOrWildcard: &pb.RelationReference_Wildcard{
			Wildcard: &pb.Wildcard{},
		},
	})
}

func (l *openFGAListener) ExitThis(_ *parser.ThisContext) {
	l.push(&pb.Userset{Userset: &pb.Userset_This{}})
}

func (l *openFGAListener) ExitTupleToUserset(ctx *parser.TupleToUsersetContext) {
	l.push(&pb.Userset{
		Userset: &pb.Userset_TupleToUserset{
			TupleToUserset: &pb.TupleToUserset{
				Tupleset:        &pb.ObjectRelation{Relation: ctx.GetTupleset().GetText()},
				ComputedUserset: &pb.ObjectRelation{Relation: ctx.GetComputedUserset().GetText()},
			},
		},
	})
}

func (l *openFGAListener) ExitComputedUserset(ctx *parser.ComputedUsersetContext) {
	l.push(&pb.Userset{
		Userset: &pb.Userset_ComputedUserset{
			ComputedUserset: &pb.ObjectRelation{Relation: ctx.GetComputedUserset().GetText()},
		},
	})
}

func (l *openFGAListener) ExitUnion(_ *parser.UnionContext) {
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

func (l *openFGAListener) ExitIntersection(_ *parser.IntersectionContext) {
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

func (l *openFGAListener) ExitExclusion(_ *parser.ExclusionContext) {
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

func (l *openFGAListener) push(u *pb.Userset) {
	l.rewrite = append(l.rewrite, u)
}

func (l *openFGAListener) pop() *pb.Userset {
	if len(l.rewrite) < 1 {
		l.errors = append(l.errors, errors.New("missing operand"))
		return nil
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

func MustParse(s string) []*pb.TypeDefinition {
	typeDefinitions, err := Parse(s)
	if err != nil {
		panic(err)
	}
	return typeDefinitions
}

func Parse(s string) ([]*pb.TypeDefinition, error) {
	is := antlr.NewInputStream(s)

	listener := newOpenFGAListener()

	lexer := parser.NewOpenFGALexer(is)
	lexer.RemoveErrorListeners()
	lexer.AddErrorListener(listener)

	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	parser := parser.NewOpenFGAParser(stream)
	parser.RemoveErrorListeners()
	parser.AddErrorListener(listener)

	// Finally parse the expression
	antlr.ParseTreeWalkerDefault.Walk(listener, parser.Start())

	if len(listener.errors) > 0 {
		var err error
		for _, e := range listener.errors {
			err = multierr.Append(err, e)
		}
		return nil, err
	}

	return listener.typeDefinitions, nil
}
