package parser

import (
	"errors"
	"fmt"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	dslparser "github.com/craigpastro/openfga-dsl-parser/v2/internal/gen/dsl/parser"
	tupleparser "github.com/craigpastro/openfga-dsl-parser/v2/internal/gen/tuple/parser"
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
	*dslparser.BaseDSLListener
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

func (l *dslListener) ExitTypeDefinition(ctx *dslparser.TypeDefinitionContext) {
	objectType := ctx.GetObjectType().GetText()

	l.typeDefinitions = append(l.typeDefinitions, &pb.TypeDefinition{
		Type:      objectType,
		Relations: l.getRelations(),
		Metadata:  l.getMetadata(),
	})

	// Clear the relations map
	l.relations = map[string]*pb.Relation{}
}

func (l *dslListener) ExitRelation(ctx *dslparser.RelationContext) {
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

func (l *dslListener) ExitRrType(ctx *dslparser.RrTypeContext) {
	l.typeInfo = append(l.typeInfo, &pb.RelationReference{
		Type: ctx.GetT().GetText(),
	})
}

func (l *dslListener) ExitRrTypeAndRelation(ctx *dslparser.RrTypeAndRelationContext) {
	l.typeInfo = append(l.typeInfo, &pb.RelationReference{
		Type: ctx.GetT().GetText(),
		RelationOrWildcard: &pb.RelationReference_Relation{
			Relation: ctx.GetR().GetText(),
		},
	})
}

func (l *dslListener) ExitRrTypeAndWildcard(ctx *dslparser.RrTypeAndWildcardContext) {
	l.typeInfo = append(l.typeInfo, &pb.RelationReference{
		Type: ctx.GetT().GetText(),
		RelationOrWildcard: &pb.RelationReference_Wildcard{
			Wildcard: &pb.Wildcard{},
		},
	})
}

func (l *dslListener) ExitThis(_ *dslparser.ThisContext) {
	l.push(&pb.Userset{Userset: &pb.Userset_This{}})
}

func (l *dslListener) ExitTupleToUserset(ctx *dslparser.TupleToUsersetContext) {
	l.push(&pb.Userset{
		Userset: &pb.Userset_TupleToUserset{
			TupleToUserset: &pb.TupleToUserset{
				Tupleset:        &pb.ObjectRelation{Relation: ctx.GetTupleset().GetText()},
				ComputedUserset: &pb.ObjectRelation{Relation: ctx.GetComputedUserset().GetText()},
			},
		},
	})
}

func (l *dslListener) ExitComputedUserset(ctx *dslparser.ComputedUsersetContext) {
	l.push(&pb.Userset{
		Userset: &pb.Userset_ComputedUserset{
			ComputedUserset: &pb.ObjectRelation{Relation: ctx.GetComputedUserset().GetText()},
		},
	})
}

func (l *dslListener) ExitUnion(_ *dslparser.UnionContext) {
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

func (l *dslListener) ExitIntersection(_ *dslparser.IntersectionContext) {
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

func (l *dslListener) ExitExclusion(_ *dslparser.ExclusionContext) {
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

func MustParseDSL(s string) []*pb.TypeDefinition {
	typeDefinitions, err := ParseDSL(s)
	if err != nil {
		panic(err)
	}

	return typeDefinitions
}

func ParseDSL(s string) ([]*pb.TypeDefinition, error) {
	is := antlr.NewInputStream(s)

	listener := newDSLListener()

	lexer := dslparser.NewDSLLexer(is)
	lexer.RemoveErrorListeners()
	lexer.AddErrorListener(listener)

	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	parser := dslparser.NewDSLParser(stream)
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

// Deprecated: please use ParseDSL instead.
func Parse(s string) ([]*pb.TypeDefinition, error) {
	return ParseDSL(s)
}

// Deprecated: please use MustParseDSL instead.
func MustParse(s string) []*pb.TypeDefinition {
	return MustParseDSL(s)
}

type tupleListener struct {
	*tupleparser.BaseTupleListener
	*antlr.DefaultErrorListener

	tuple *pb.TupleKey

	errors []error
}

func newTupleListener() *tupleListener {
	return &tupleListener{
		tuple: &pb.TupleKey{},
	}
}

func (l *tupleListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	l.errors = append(l.errors, &syntaxError{
		line:   line,
		column: column,
		msg:    msg,
	})
}

func (l *tupleListener) ExitTuple(ctx *tupleparser.TupleContext) {
	if len(l.errors) > 0 {
		return
	}

	l.tuple.Object = fmt.Sprintf("%s:%s",
		ctx.GetNamespace().GetText(),
		ctx.GetObjectId().GetText(),
	)

	l.tuple.Relation = ctx.GetRelation().GetText()
}

func (l *tupleListener) ExitUserId(ctx *tupleparser.UserIdContext) {
	l.tuple.User = fmt.Sprintf("%s", ctx.GetUserId().GetText())
}

func (l *tupleListener) ExitUserObject(ctx *tupleparser.UserObjectContext) {
	l.tuple.User = fmt.Sprintf("%s:%s",
		ctx.GetNamespace().GetText(),
		ctx.GetObjectId().GetText(),
	)
}

func (l *tupleListener) ExitUserUserset(ctx *tupleparser.UserUsersetContext) {
	l.tuple.User = fmt.Sprintf("%s:%s#%s",
		ctx.GetNamespace().GetText(),
		ctx.GetObjectId().GetText(),
		ctx.GetRelation().GetText(),
	)
}

func ParseTuple(s string) (*pb.TupleKey, error) {
	is := antlr.NewInputStream(s)

	listener := newTupleListener()

	lexer := tupleparser.NewTupleLexer(is)
	lexer.RemoveErrorListeners()
	lexer.AddErrorListener(listener)

	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	parser := tupleparser.NewTupleParser(stream)
	parser.RemoveErrorListeners()
	parser.AddErrorListener(listener)

	antlr.ParseTreeWalkerDefault.Walk(listener, parser.Tuple())

	if len(listener.errors) > 0 {
		var err error
		for _, e := range listener.errors {
			err = multierr.Append(err, e)
		}
		return nil, err
	}

	return listener.tuple, nil
}

func MustParseTuple(s string) *pb.TupleKey {
	tuple, err := ParseTuple(s)
	if err != nil {
		panic(err)
	}
	return tuple
}
