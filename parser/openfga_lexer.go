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
		"", "'type'", "'relations'", "'define'", "'as'", "':'", "'['", "']'",
		"','", "'#'", "'self'", "'from'", "'or'", "'and'", "'but not'", "'('",
		"')'",
	}
	staticData.symbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"ID", "WS",
	}
	staticData.ruleNames = []string{
		"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8",
		"T__9", "T__10", "T__11", "T__12", "T__13", "T__14", "T__15", "ID",
		"WS",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 18, 115, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2,
		1, 2, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 5, 1, 5, 1, 6,
		1, 6, 1, 7, 1, 7, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 10, 1, 10,
		1, 10, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 1, 12, 1, 12, 1, 12, 1, 12, 1,
		13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 14, 1, 14, 1, 15,
		1, 15, 1, 16, 1, 16, 5, 16, 104, 8, 16, 10, 16, 12, 16, 107, 9, 16, 1,
		17, 4, 17, 110, 8, 17, 11, 17, 12, 17, 111, 1, 17, 1, 17, 0, 0, 18, 1,
		1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11,
		23, 12, 25, 13, 27, 14, 29, 15, 31, 16, 33, 17, 35, 18, 1, 0, 3, 3, 0,
		65, 90, 95, 95, 97, 122, 4, 0, 48, 57, 65, 90, 95, 95, 97, 122, 3, 0, 9,
		10, 12, 13, 32, 32, 116, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0,
		0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1,
		0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21,
		1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0,
		29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0, 33, 1, 0, 0, 0, 0, 35, 1, 0, 0, 0,
		1, 37, 1, 0, 0, 0, 3, 42, 1, 0, 0, 0, 5, 52, 1, 0, 0, 0, 7, 59, 1, 0, 0,
		0, 9, 62, 1, 0, 0, 0, 11, 64, 1, 0, 0, 0, 13, 66, 1, 0, 0, 0, 15, 68, 1,
		0, 0, 0, 17, 70, 1, 0, 0, 0, 19, 72, 1, 0, 0, 0, 21, 77, 1, 0, 0, 0, 23,
		82, 1, 0, 0, 0, 25, 85, 1, 0, 0, 0, 27, 89, 1, 0, 0, 0, 29, 97, 1, 0, 0,
		0, 31, 99, 1, 0, 0, 0, 33, 101, 1, 0, 0, 0, 35, 109, 1, 0, 0, 0, 37, 38,
		5, 116, 0, 0, 38, 39, 5, 121, 0, 0, 39, 40, 5, 112, 0, 0, 40, 41, 5, 101,
		0, 0, 41, 2, 1, 0, 0, 0, 42, 43, 5, 114, 0, 0, 43, 44, 5, 101, 0, 0, 44,
		45, 5, 108, 0, 0, 45, 46, 5, 97, 0, 0, 46, 47, 5, 116, 0, 0, 47, 48, 5,
		105, 0, 0, 48, 49, 5, 111, 0, 0, 49, 50, 5, 110, 0, 0, 50, 51, 5, 115,
		0, 0, 51, 4, 1, 0, 0, 0, 52, 53, 5, 100, 0, 0, 53, 54, 5, 101, 0, 0, 54,
		55, 5, 102, 0, 0, 55, 56, 5, 105, 0, 0, 56, 57, 5, 110, 0, 0, 57, 58, 5,
		101, 0, 0, 58, 6, 1, 0, 0, 0, 59, 60, 5, 97, 0, 0, 60, 61, 5, 115, 0, 0,
		61, 8, 1, 0, 0, 0, 62, 63, 5, 58, 0, 0, 63, 10, 1, 0, 0, 0, 64, 65, 5,
		91, 0, 0, 65, 12, 1, 0, 0, 0, 66, 67, 5, 93, 0, 0, 67, 14, 1, 0, 0, 0,
		68, 69, 5, 44, 0, 0, 69, 16, 1, 0, 0, 0, 70, 71, 5, 35, 0, 0, 71, 18, 1,
		0, 0, 0, 72, 73, 5, 115, 0, 0, 73, 74, 5, 101, 0, 0, 74, 75, 5, 108, 0,
		0, 75, 76, 5, 102, 0, 0, 76, 20, 1, 0, 0, 0, 77, 78, 5, 102, 0, 0, 78,
		79, 5, 114, 0, 0, 79, 80, 5, 111, 0, 0, 80, 81, 5, 109, 0, 0, 81, 22, 1,
		0, 0, 0, 82, 83, 5, 111, 0, 0, 83, 84, 5, 114, 0, 0, 84, 24, 1, 0, 0, 0,
		85, 86, 5, 97, 0, 0, 86, 87, 5, 110, 0, 0, 87, 88, 5, 100, 0, 0, 88, 26,
		1, 0, 0, 0, 89, 90, 5, 98, 0, 0, 90, 91, 5, 117, 0, 0, 91, 92, 5, 116,
		0, 0, 92, 93, 5, 32, 0, 0, 93, 94, 5, 110, 0, 0, 94, 95, 5, 111, 0, 0,
		95, 96, 5, 116, 0, 0, 96, 28, 1, 0, 0, 0, 97, 98, 5, 40, 0, 0, 98, 30,
		1, 0, 0, 0, 99, 100, 5, 41, 0, 0, 100, 32, 1, 0, 0, 0, 101, 105, 7, 0,
		0, 0, 102, 104, 7, 1, 0, 0, 103, 102, 1, 0, 0, 0, 104, 107, 1, 0, 0, 0,
		105, 103, 1, 0, 0, 0, 105, 106, 1, 0, 0, 0, 106, 34, 1, 0, 0, 0, 107, 105,
		1, 0, 0, 0, 108, 110, 7, 2, 0, 0, 109, 108, 1, 0, 0, 0, 110, 111, 1, 0,
		0, 0, 111, 109, 1, 0, 0, 0, 111, 112, 1, 0, 0, 0, 112, 113, 1, 0, 0, 0,
		113, 114, 6, 17, 0, 0, 114, 36, 1, 0, 0, 0, 3, 0, 105, 111, 1, 6, 0, 0,
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
	OpenFGALexerT__12 = 13
	OpenFGALexerT__13 = 14
	OpenFGALexerT__14 = 15
	OpenFGALexerT__15 = 16
	OpenFGALexerID    = 17
	OpenFGALexerWS    = 18
)
