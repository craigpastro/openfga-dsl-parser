// Code generated from java-escape by ANTLR 4.11.1. DO NOT EDIT.

package parser

import (
	"fmt"
	"sync"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type TupleLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var tuplelexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	channelNames           []string
	modeNames              []string
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func tuplelexerLexerInit() {
	staticData := &tuplelexerLexerStaticData
	staticData.channelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.modeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.literalNames = []string{
		"", "'#'", "'@'", "':'",
	}
	staticData.symbolicNames = []string{
		"", "", "", "", "ID", "WS",
	}
	staticData.ruleNames = []string{
		"T__0", "T__1", "T__2", "ID", "WS",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 5, 31, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 1, 0, 1, 0, 1, 1, 1, 1, 1, 2, 1, 2, 1, 3, 1, 3, 5, 3, 20, 8, 3,
		10, 3, 12, 3, 23, 9, 3, 1, 4, 4, 4, 26, 8, 4, 11, 4, 12, 4, 27, 1, 4, 1,
		4, 0, 0, 5, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 1, 0, 3, 2, 0, 65, 90, 97, 122,
		5, 0, 45, 45, 48, 57, 65, 90, 95, 95, 97, 122, 3, 0, 9, 10, 12, 13, 32,
		32, 32, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1,
		0, 0, 0, 0, 9, 1, 0, 0, 0, 1, 11, 1, 0, 0, 0, 3, 13, 1, 0, 0, 0, 5, 15,
		1, 0, 0, 0, 7, 17, 1, 0, 0, 0, 9, 25, 1, 0, 0, 0, 11, 12, 5, 35, 0, 0,
		12, 2, 1, 0, 0, 0, 13, 14, 5, 64, 0, 0, 14, 4, 1, 0, 0, 0, 15, 16, 5, 58,
		0, 0, 16, 6, 1, 0, 0, 0, 17, 21, 7, 0, 0, 0, 18, 20, 7, 1, 0, 0, 19, 18,
		1, 0, 0, 0, 20, 23, 1, 0, 0, 0, 21, 19, 1, 0, 0, 0, 21, 22, 1, 0, 0, 0,
		22, 8, 1, 0, 0, 0, 23, 21, 1, 0, 0, 0, 24, 26, 7, 2, 0, 0, 25, 24, 1, 0,
		0, 0, 26, 27, 1, 0, 0, 0, 27, 25, 1, 0, 0, 0, 27, 28, 1, 0, 0, 0, 28, 29,
		1, 0, 0, 0, 29, 30, 6, 4, 0, 0, 30, 10, 1, 0, 0, 0, 3, 0, 21, 27, 1, 6,
		0, 0,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// TupleLexerInit initializes any static state used to implement TupleLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewTupleLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func TupleLexerInit() {
	staticData := &tuplelexerLexerStaticData
	staticData.once.Do(tuplelexerLexerInit)
}

// NewTupleLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewTupleLexer(input antlr.CharStream) *TupleLexer {
	TupleLexerInit()
	l := new(TupleLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &tuplelexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	l.channelNames = staticData.channelNames
	l.modeNames = staticData.modeNames
	l.RuleNames = staticData.ruleNames
	l.LiteralNames = staticData.literalNames
	l.SymbolicNames = staticData.symbolicNames
	l.GrammarFileName = "Tuple.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// TupleLexer tokens.
const (
	TupleLexerT__0 = 1
	TupleLexerT__1 = 2
	TupleLexerT__2 = 3
	TupleLexerID   = 4
	TupleLexerWS   = 5
)
