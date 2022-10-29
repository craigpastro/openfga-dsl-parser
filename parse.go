package parser

import (
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/craigpastro/openfga-dsl-parser/parser"
	pb "go.buf.build/openfga/go/openfga/api/openfga/v1"
)

type openFGAListener struct {
	*parser.BaseOpenFGAListener

	typeDefinitions []*pb.TypeDefinition
	relations       map[string]*pb.Userset
	stack           []*pb.Userset
}

func newOpenFGAListener() *openFGAListener {
	return &openFGAListener{
		relations: map[string]*pb.Userset{},
	}
}

func (l *openFGAListener) push(u *pb.Userset) {
	l.stack = append(l.stack, u)
}

func (l *openFGAListener) pop() *pb.Userset {
	if len(l.stack) < 1 {
		panic("stack is empty unable to pop")
	}

	// Get the last value from the stack.
	result := l.stack[len(l.stack)-1]

	// Remove the last element from the stack.
	l.stack = l.stack[:len(l.stack)-1]

	return result
}

func (l *openFGAListener) ExitTypedef(ctx *parser.TypedefContext) {
	objectType := ctx.GetObjectType().GetText()

	l.typeDefinitions = append(l.typeDefinitions, &pb.TypeDefinition{
		Type:      objectType,
		Relations: l.relations,
	})

	// Clear the relations map
	l.relations = map[string]*pb.Userset{}
}

func (l *openFGAListener) ExitRelations(ctx *parser.RelationsContext) {
	relation := ctx.GetRelation().GetText()
	l.relations[relation] = l.pop()
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

func MustParse(data string) []*pb.TypeDefinition {
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
