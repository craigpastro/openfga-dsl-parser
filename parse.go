package parser

import (
	"errors"

	openfgav1 "buf.build/gen/go/openfga/api/protocolbuffers/go/openfga/v1"
	"github.com/antlr4-go/antlr/v4"
	"github.com/craigpastro/openfga-dsl-parser/parser"
)

type openFGAListener struct {
	*parser.BaseOpenFGAListener

	typeDefinitions []*openfgav1.TypeDefinition
	relations       map[string]*openfgav1.Userset
	stack           []*openfgav1.Userset
}

func newOpenFGAListener() *openFGAListener {
	return &openFGAListener{
		relations: map[string]*openfgav1.Userset{},
	}
}

func (l *openFGAListener) push(u *openfgav1.Userset) {
	l.stack = append(l.stack, u)
}

func (l *openFGAListener) pop() *openfgav1.Userset {
	if len(l.stack) < 1 {
		panic("stack is empty unable to pop")
	}

	result := l.stack[len(l.stack)-1]
	l.stack = l.stack[:len(l.stack)-1]

	return result
}

func (l *openFGAListener) ExitTypedef(ctx *parser.TypedefContext) {
	objectType := ctx.GetObjectType().GetText()

	l.typeDefinitions = append(l.typeDefinitions, &openfgav1.TypeDefinition{
		Type:      objectType,
		Relations: l.relations,
	})

	// Clear the relations map
	l.relations = map[string]*openfgav1.Userset{}
}

func (l *openFGAListener) ExitRelations(ctx *parser.RelationsContext) {
	relation := ctx.GetRelation().GetText()
	l.relations[relation] = l.pop()
}

func (l *openFGAListener) ExitThis(_ *parser.ThisContext) {
	l.push(&openfgav1.Userset{Userset: &openfgav1.Userset_This{}})
}

func (l *openFGAListener) ExitTupleToUserset(ctx *parser.TupleToUsersetContext) {
	l.push(&openfgav1.Userset{
		Userset: &openfgav1.Userset_TupleToUserset{
			TupleToUserset: &openfgav1.TupleToUserset{
				Tupleset:        &openfgav1.ObjectRelation{Relation: ctx.GetTupleset().GetText()},
				ComputedUserset: &openfgav1.ObjectRelation{Relation: ctx.GetComputedUserset().GetText()},
			},
		},
	})
}

func (l *openFGAListener) ExitComputedUserset(ctx *parser.ComputedUsersetContext) {
	l.push(&openfgav1.Userset{
		Userset: &openfgav1.Userset_ComputedUserset{
			ComputedUserset: &openfgav1.ObjectRelation{Relation: ctx.GetComputedUserset().GetText()},
		},
	})
}

func (l *openFGAListener) ExitUnion(_ *parser.UnionContext) {
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

func (l *openFGAListener) ExitIntersection(_ *parser.IntersectionContext) {
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

func (l *openFGAListener) ExitExclusion(_ *parser.ExclusionContext) {
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

func MustParse(data string) []*openfgav1.TypeDefinition {
	is := antlr.NewInputStream(data)

	// Create the Lexer
	lexer := parser.NewOpenFGALexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Create the Parser
	p := parser.NewOpenFGAParser(stream)

	// Finally parse the expression
	l := newOpenFGAListener()
	antlr.ParseTreeWalkerDefault.Walk(l, p.Prog())

	return l.typeDefinitions
}

func Parse(data string) (typeDefinitions []*openfgav1.TypeDefinition, err error) {
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
