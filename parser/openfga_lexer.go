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

type OpenFGALexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var openfgalexerLexerStaticData struct {
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

func openfgalexerLexerInit() {
	staticData := &openfgalexerLexerStaticData
	staticData.channelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.modeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.literalNames = []string{
		"", "'type'", "'relations'", "'define'", "':'", "'['", "']'", "','",
		"'#'", "'from'", "'or'", "'and'", "'but not'",
	}
	staticData.symbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "ID", "WS",
	}
	staticData.ruleNames = []string{
		"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8",
		"T__9", "T__10", "T__11", "ID", "WS",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 14, 95, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 1, 0, 1, 0, 1, 0,
		1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 4, 1, 4, 1, 5,
		1, 5, 1, 6, 1, 6, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9,
		1, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1,
		11, 1, 11, 1, 11, 1, 12, 1, 12, 5, 12, 84, 8, 12, 10, 12, 12, 12, 87, 9,
		12, 1, 13, 4, 13, 90, 8, 13, 11, 13, 12, 13, 91, 1, 13, 1, 13, 0, 0, 14,
		1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11,
		23, 12, 25, 13, 27, 14, 1, 0, 3, 3, 0, 65, 90, 95, 95, 97, 122, 4, 0, 48,
		57, 65, 90, 95, 95, 97, 122, 3, 0, 9, 10, 12, 13, 32, 32, 96, 0, 1, 1,
		0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1,
		0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17,
		1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0,
		25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 1, 29, 1, 0, 0, 0, 3, 34, 1, 0, 0, 0,
		5, 44, 1, 0, 0, 0, 7, 51, 1, 0, 0, 0, 9, 53, 1, 0, 0, 0, 11, 55, 1, 0,
		0, 0, 13, 57, 1, 0, 0, 0, 15, 59, 1, 0, 0, 0, 17, 61, 1, 0, 0, 0, 19, 66,
		1, 0, 0, 0, 21, 69, 1, 0, 0, 0, 23, 73, 1, 0, 0, 0, 25, 81, 1, 0, 0, 0,
		27, 89, 1, 0, 0, 0, 29, 30, 5, 116, 0, 0, 30, 31, 5, 121, 0, 0, 31, 32,
		5, 112, 0, 0, 32, 33, 5, 101, 0, 0, 33, 2, 1, 0, 0, 0, 34, 35, 5, 114,
		0, 0, 35, 36, 5, 101, 0, 0, 36, 37, 5, 108, 0, 0, 37, 38, 5, 97, 0, 0,
		38, 39, 5, 116, 0, 0, 39, 40, 5, 105, 0, 0, 40, 41, 5, 111, 0, 0, 41, 42,
		5, 110, 0, 0, 42, 43, 5, 115, 0, 0, 43, 4, 1, 0, 0, 0, 44, 45, 5, 100,
		0, 0, 45, 46, 5, 101, 0, 0, 46, 47, 5, 102, 0, 0, 47, 48, 5, 105, 0, 0,
		48, 49, 5, 110, 0, 0, 49, 50, 5, 101, 0, 0, 50, 6, 1, 0, 0, 0, 51, 52,
		5, 58, 0, 0, 52, 8, 1, 0, 0, 0, 53, 54, 5, 91, 0, 0, 54, 10, 1, 0, 0, 0,
		55, 56, 5, 93, 0, 0, 56, 12, 1, 0, 0, 0, 57, 58, 5, 44, 0, 0, 58, 14, 1,
		0, 0, 0, 59, 60, 5, 35, 0, 0, 60, 16, 1, 0, 0, 0, 61, 62, 5, 102, 0, 0,
		62, 63, 5, 114, 0, 0, 63, 64, 5, 111, 0, 0, 64, 65, 5, 109, 0, 0, 65, 18,
		1, 0, 0, 0, 66, 67, 5, 111, 0, 0, 67, 68, 5, 114, 0, 0, 68, 20, 1, 0, 0,
		0, 69, 70, 5, 97, 0, 0, 70, 71, 5, 110, 0, 0, 71, 72, 5, 100, 0, 0, 72,
		22, 1, 0, 0, 0, 73, 74, 5, 98, 0, 0, 74, 75, 5, 117, 0, 0, 75, 76, 5, 116,
		0, 0, 76, 77, 5, 32, 0, 0, 77, 78, 5, 110, 0, 0, 78, 79, 5, 111, 0, 0,
		79, 80, 5, 116, 0, 0, 80, 24, 1, 0, 0, 0, 81, 85, 7, 0, 0, 0, 82, 84, 7,
		1, 0, 0, 83, 82, 1, 0, 0, 0, 84, 87, 1, 0, 0, 0, 85, 83, 1, 0, 0, 0, 85,
		86, 1, 0, 0, 0, 86, 26, 1, 0, 0, 0, 87, 85, 1, 0, 0, 0, 88, 90, 7, 2, 0,
		0, 89, 88, 1, 0, 0, 0, 90, 91, 1, 0, 0, 0, 91, 89, 1, 0, 0, 0, 91, 92,
		1, 0, 0, 0, 92, 93, 1, 0, 0, 0, 93, 94, 6, 13, 0, 0, 94, 28, 1, 0, 0, 0,
		3, 0, 85, 91, 1, 6, 0, 0,
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

// OpenFGALexerInit initializes any static state used to implement OpenFGALexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewOpenFGALexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func OpenFGALexerInit() {
	staticData := &openfgalexerLexerStaticData
	staticData.once.Do(openfgalexerLexerInit)
}

// NewOpenFGALexer produces a new lexer instance for the optional input antlr.CharStream.
func NewOpenFGALexer(input antlr.CharStream) *OpenFGALexer {
	OpenFGALexerInit()
	l := new(OpenFGALexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &openfgalexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	l.channelNames = staticData.channelNames
	l.modeNames = staticData.modeNames
	l.RuleNames = staticData.ruleNames
	l.LiteralNames = staticData.literalNames
	l.SymbolicNames = staticData.symbolicNames
	l.GrammarFileName = "OpenFGA.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// OpenFGALexer tokens.
const (
	OpenFGALexerT__0  = 1
	OpenFGALexerT__1  = 2
	OpenFGALexerT__2  = 3
	OpenFGALexerT__3  = 4
	OpenFGALexerT__4  = 5
	OpenFGALexerT__5  = 6
	OpenFGALexerT__6  = 7
	OpenFGALexerT__7  = 8
	OpenFGALexerT__8  = 9
	OpenFGALexerT__9  = 10
	OpenFGALexerT__10 = 11
	OpenFGALexerT__11 = 12
	OpenFGALexerID    = 13
	OpenFGALexerWS    = 14
)
