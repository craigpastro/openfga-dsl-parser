// Code generated from OpenFGA.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"sync"
	"unicode"
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

var OpenFGALexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	ChannelNames           []string
	ModeNames              []string
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func openfgalexerLexerInit() {
	staticData := &OpenFGALexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.LiteralNames = []string{
		"", "'type'", "'relations'", "'define'", "'as'", "'self'", "'from'",
		"'or'", "'and'", "'but not'", "'('", "')'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "ID", "WS",
	}
	staticData.RuleNames = []string{
		"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8",
		"T__9", "T__10", "ID", "WS",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 13, 95, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1,
		2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1,
		4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1,
		7, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 10, 1,
		10, 1, 11, 1, 11, 5, 11, 84, 8, 11, 10, 11, 12, 11, 87, 9, 11, 1, 12, 4,
		12, 90, 8, 12, 11, 12, 12, 12, 91, 1, 12, 1, 12, 0, 0, 13, 1, 1, 3, 2,
		5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11, 23, 12, 25,
		13, 1, 0, 3, 3, 0, 65, 90, 95, 95, 97, 122, 4, 0, 48, 57, 65, 90, 95, 95,
		97, 122, 3, 0, 9, 10, 12, 13, 32, 32, 96, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0,
		0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0,
		0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1,
		0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 1, 27,
		1, 0, 0, 0, 3, 32, 1, 0, 0, 0, 5, 42, 1, 0, 0, 0, 7, 49, 1, 0, 0, 0, 9,
		52, 1, 0, 0, 0, 11, 57, 1, 0, 0, 0, 13, 62, 1, 0, 0, 0, 15, 65, 1, 0, 0,
		0, 17, 69, 1, 0, 0, 0, 19, 77, 1, 0, 0, 0, 21, 79, 1, 0, 0, 0, 23, 81,
		1, 0, 0, 0, 25, 89, 1, 0, 0, 0, 27, 28, 5, 116, 0, 0, 28, 29, 5, 121, 0,
		0, 29, 30, 5, 112, 0, 0, 30, 31, 5, 101, 0, 0, 31, 2, 1, 0, 0, 0, 32, 33,
		5, 114, 0, 0, 33, 34, 5, 101, 0, 0, 34, 35, 5, 108, 0, 0, 35, 36, 5, 97,
		0, 0, 36, 37, 5, 116, 0, 0, 37, 38, 5, 105, 0, 0, 38, 39, 5, 111, 0, 0,
		39, 40, 5, 110, 0, 0, 40, 41, 5, 115, 0, 0, 41, 4, 1, 0, 0, 0, 42, 43,
		5, 100, 0, 0, 43, 44, 5, 101, 0, 0, 44, 45, 5, 102, 0, 0, 45, 46, 5, 105,
		0, 0, 46, 47, 5, 110, 0, 0, 47, 48, 5, 101, 0, 0, 48, 6, 1, 0, 0, 0, 49,
		50, 5, 97, 0, 0, 50, 51, 5, 115, 0, 0, 51, 8, 1, 0, 0, 0, 52, 53, 5, 115,
		0, 0, 53, 54, 5, 101, 0, 0, 54, 55, 5, 108, 0, 0, 55, 56, 5, 102, 0, 0,
		56, 10, 1, 0, 0, 0, 57, 58, 5, 102, 0, 0, 58, 59, 5, 114, 0, 0, 59, 60,
		5, 111, 0, 0, 60, 61, 5, 109, 0, 0, 61, 12, 1, 0, 0, 0, 62, 63, 5, 111,
		0, 0, 63, 64, 5, 114, 0, 0, 64, 14, 1, 0, 0, 0, 65, 66, 5, 97, 0, 0, 66,
		67, 5, 110, 0, 0, 67, 68, 5, 100, 0, 0, 68, 16, 1, 0, 0, 0, 69, 70, 5,
		98, 0, 0, 70, 71, 5, 117, 0, 0, 71, 72, 5, 116, 0, 0, 72, 73, 5, 32, 0,
		0, 73, 74, 5, 110, 0, 0, 74, 75, 5, 111, 0, 0, 75, 76, 5, 116, 0, 0, 76,
		18, 1, 0, 0, 0, 77, 78, 5, 40, 0, 0, 78, 20, 1, 0, 0, 0, 79, 80, 5, 41,
		0, 0, 80, 22, 1, 0, 0, 0, 81, 85, 7, 0, 0, 0, 82, 84, 7, 1, 0, 0, 83, 82,
		1, 0, 0, 0, 84, 87, 1, 0, 0, 0, 85, 83, 1, 0, 0, 0, 85, 86, 1, 0, 0, 0,
		86, 24, 1, 0, 0, 0, 87, 85, 1, 0, 0, 0, 88, 90, 7, 2, 0, 0, 89, 88, 1,
		0, 0, 0, 90, 91, 1, 0, 0, 0, 91, 89, 1, 0, 0, 0, 91, 92, 1, 0, 0, 0, 92,
		93, 1, 0, 0, 0, 93, 94, 6, 12, 0, 0, 94, 26, 1, 0, 0, 0, 3, 0, 85, 91,
		1, 6, 0, 0,
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
	staticData := &OpenFGALexerLexerStaticData
	staticData.once.Do(openfgalexerLexerInit)
}

// NewOpenFGALexer produces a new lexer instance for the optional input antlr.CharStream.
func NewOpenFGALexer(input antlr.CharStream) *OpenFGALexer {
	OpenFGALexerInit()
	l := new(OpenFGALexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &OpenFGALexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
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
	OpenFGALexerID    = 12
	OpenFGALexerWS    = 13
)
