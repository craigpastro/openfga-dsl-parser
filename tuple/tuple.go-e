package tuple

import (
	"fmt"

	openfgav1 "buf.build/gen/go/openfga/api/protocolbuffers/go/openfga/v1"
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/craigpastro/openfga-dsl-parser/v2/internal/gen/tuple/parser"
	"go.uber.org/multierr"
)

type syntaxError struct {
	line, column int
	msg          string
}

func (e *syntaxError) Error() string {
	return fmt.Sprintf("error at %d:%d: %s", e.line, e.column, e.msg)
}

type tupleListener struct {
	*parser.BaseTupleListener
	*antlr.DefaultErrorListener

	tuple *openfgav1.TupleKey

	errors []error
}

func newTupleListener() *tupleListener {
	return &tupleListener{
		tuple: &openfgav1.TupleKey{},
	}
}

func (l *tupleListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	l.errors = append(l.errors, &syntaxError{
		line:   line,
		column: column,
		msg:    msg,
	})
}

func (l *tupleListener) ExitTuple(ctx *parser.TupleContext) {
	if len(l.errors) > 0 {
		return
	}

	l.tuple.Object = fmt.Sprintf("%s:%s",
		ctx.GetObj().GetNamespace().GetText(),
		ctx.GetObj().GetObjectID().GetText(),
	)

	l.tuple.Relation = ctx.GetRelation().GetText()
}

func (l *tupleListener) ExitUserID(ctx *parser.UserIDContext) {
	l.tuple.User = ctx.GetText()
}

func (l *tupleListener) ExitUserObject(ctx *parser.UserObjectContext) {
	l.tuple.User = fmt.Sprintf("%s:%s",
		ctx.GetObj().GetNamespace().GetText(),
		ctx.GetObj().GetObjectID().GetText(),
	)
}

func (l *tupleListener) ExitUserUserset(ctx *parser.UserUsersetContext) {
	l.tuple.User = fmt.Sprintf("%s:%s#%s",
		ctx.GetObj().GetNamespace().GetText(),
		ctx.GetObj().GetObjectID().GetText(),
		ctx.GetRelation().GetText(),
	)
}

func Parse(s string) (*openfgav1.TupleKey, error) {
	is := antlr.NewInputStream(s)

	listener := newTupleListener()

	lexer := parser.NewTupleLexer(is)
	lexer.RemoveErrorListeners()
	lexer.AddErrorListener(listener)

	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	parser := parser.NewTupleParser(stream)
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

func MustParseTuple(s string) *openfgav1.TupleKey {
	tuple, err := Parse(s)
	if err != nil {
		panic(err)
	}
	return tuple
}
