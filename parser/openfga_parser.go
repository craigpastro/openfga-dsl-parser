// Code generated from java-escape by ANTLR 4.11.1. DO NOT EDIT.

package parser // OpenFGA

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type OpenFGAParser struct {
	*antlr.BaseParser
}

var openfgaParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func openfgaParserInit() {
	staticData := &openfgaParserStaticData
	staticData.literalNames = []string{
		"", "'type'", "'relations'", "'define'", "':'", "'['", "']'", "','",
		"'#'", "':*'", "'or'", "'from'", "'and'", "'but not'",
	}
	staticData.symbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "ID", "WS",
	}
	staticData.ruleNames = []string{
		"start", "typeDefinition", "relation", "typeRestrictionOrId", "typeRestriction",
		"relationReferences", "relationReference", "rewrite", "orTTU", "ors",
		"or", "ands", "and", "exclusion",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 15, 111, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 1, 0, 4, 0, 30, 8, 0, 11,
		0, 12, 0, 31, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 4, 1, 40, 8, 1, 11, 1,
		12, 1, 41, 3, 1, 44, 8, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 3, 2, 51, 8, 2,
		1, 3, 1, 3, 3, 3, 55, 8, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 5,
		5, 64, 8, 5, 10, 5, 12, 5, 67, 9, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6,
		3, 6, 75, 8, 6, 1, 7, 1, 7, 1, 7, 1, 7, 3, 7, 81, 8, 7, 1, 8, 1, 8, 1,
		8, 1, 8, 1, 8, 1, 9, 1, 9, 5, 9, 90, 8, 9, 10, 9, 12, 9, 93, 9, 9, 1, 10,
		1, 10, 1, 10, 1, 11, 1, 11, 5, 11, 100, 8, 11, 10, 11, 12, 11, 103, 9,
		11, 1, 12, 1, 12, 1, 12, 1, 13, 1, 13, 1, 13, 1, 13, 0, 0, 14, 0, 2, 4,
		6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 0, 0, 109, 0, 29, 1, 0, 0, 0,
		2, 35, 1, 0, 0, 0, 4, 45, 1, 0, 0, 0, 6, 54, 1, 0, 0, 0, 8, 56, 1, 0, 0,
		0, 10, 60, 1, 0, 0, 0, 12, 74, 1, 0, 0, 0, 14, 80, 1, 0, 0, 0, 16, 82,
		1, 0, 0, 0, 18, 87, 1, 0, 0, 0, 20, 94, 1, 0, 0, 0, 22, 97, 1, 0, 0, 0,
		24, 104, 1, 0, 0, 0, 26, 107, 1, 0, 0, 0, 28, 30, 3, 2, 1, 0, 29, 28, 1,
		0, 0, 0, 30, 31, 1, 0, 0, 0, 31, 29, 1, 0, 0, 0, 31, 32, 1, 0, 0, 0, 32,
		33, 1, 0, 0, 0, 33, 34, 5, 0, 0, 1, 34, 1, 1, 0, 0, 0, 35, 36, 5, 1, 0,
		0, 36, 43, 5, 14, 0, 0, 37, 39, 5, 2, 0, 0, 38, 40, 3, 4, 2, 0, 39, 38,
		1, 0, 0, 0, 40, 41, 1, 0, 0, 0, 41, 39, 1, 0, 0, 0, 41, 42, 1, 0, 0, 0,
		42, 44, 1, 0, 0, 0, 43, 37, 1, 0, 0, 0, 43, 44, 1, 0, 0, 0, 44, 3, 1, 0,
		0, 0, 45, 46, 5, 3, 0, 0, 46, 47, 5, 14, 0, 0, 47, 48, 5, 4, 0, 0, 48,
		50, 3, 6, 3, 0, 49, 51, 3, 14, 7, 0, 50, 49, 1, 0, 0, 0, 50, 51, 1, 0,
		0, 0, 51, 5, 1, 0, 0, 0, 52, 55, 3, 8, 4, 0, 53, 55, 5, 14, 0, 0, 54, 52,
		1, 0, 0, 0, 54, 53, 1, 0, 0, 0, 55, 7, 1, 0, 0, 0, 56, 57, 5, 5, 0, 0,
		57, 58, 3, 10, 5, 0, 58, 59, 5, 6, 0, 0, 59, 9, 1, 0, 0, 0, 60, 65, 3,
		12, 6, 0, 61, 62, 5, 7, 0, 0, 62, 64, 3, 10, 5, 0, 63, 61, 1, 0, 0, 0,
		64, 67, 1, 0, 0, 0, 65, 63, 1, 0, 0, 0, 65, 66, 1, 0, 0, 0, 66, 11, 1,
		0, 0, 0, 67, 65, 1, 0, 0, 0, 68, 75, 5, 14, 0, 0, 69, 70, 5, 14, 0, 0,
		70, 71, 5, 8, 0, 0, 71, 75, 5, 14, 0, 0, 72, 73, 5, 14, 0, 0, 73, 75, 5,
		9, 0, 0, 74, 68, 1, 0, 0, 0, 74, 69, 1, 0, 0, 0, 74, 72, 1, 0, 0, 0, 75,
		13, 1, 0, 0, 0, 76, 81, 3, 16, 8, 0, 77, 81, 3, 18, 9, 0, 78, 81, 3, 22,
		11, 0, 79, 81, 3, 26, 13, 0, 80, 76, 1, 0, 0, 0, 80, 77, 1, 0, 0, 0, 80,
		78, 1, 0, 0, 0, 80, 79, 1, 0, 0, 0, 81, 15, 1, 0, 0, 0, 82, 83, 5, 10,
		0, 0, 83, 84, 5, 14, 0, 0, 84, 85, 5, 11, 0, 0, 85, 86, 5, 14, 0, 0, 86,
		17, 1, 0, 0, 0, 87, 91, 3, 20, 10, 0, 88, 90, 3, 18, 9, 0, 89, 88, 1, 0,
		0, 0, 90, 93, 1, 0, 0, 0, 91, 89, 1, 0, 0, 0, 91, 92, 1, 0, 0, 0, 92, 19,
		1, 0, 0, 0, 93, 91, 1, 0, 0, 0, 94, 95, 5, 10, 0, 0, 95, 96, 5, 14, 0,
		0, 96, 21, 1, 0, 0, 0, 97, 101, 3, 24, 12, 0, 98, 100, 3, 22, 11, 0, 99,
		98, 1, 0, 0, 0, 100, 103, 1, 0, 0, 0, 101, 99, 1, 0, 0, 0, 101, 102, 1,
		0, 0, 0, 102, 23, 1, 0, 0, 0, 103, 101, 1, 0, 0, 0, 104, 105, 5, 12, 0,
		0, 105, 106, 5, 14, 0, 0, 106, 25, 1, 0, 0, 0, 107, 108, 5, 13, 0, 0, 108,
		109, 5, 14, 0, 0, 109, 27, 1, 0, 0, 0, 10, 31, 41, 43, 50, 54, 65, 74,
		80, 91, 101,
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

// OpenFGAParserInit initializes any static state used to implement OpenFGAParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewOpenFGAParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func OpenFGAParserInit() {
	staticData := &openfgaParserStaticData
	staticData.once.Do(openfgaParserInit)
}

// NewOpenFGAParser produces a new parser instance for the optional input antlr.TokenStream.
func NewOpenFGAParser(input antlr.TokenStream) *OpenFGAParser {
	OpenFGAParserInit()
	this := new(OpenFGAParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &openfgaParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "java-escape"

	return this
}

// OpenFGAParser tokens.
const (
	OpenFGAParserEOF   = antlr.TokenEOF
	OpenFGAParserT__0  = 1
	OpenFGAParserT__1  = 2
	OpenFGAParserT__2  = 3
	OpenFGAParserT__3  = 4
	OpenFGAParserT__4  = 5
	OpenFGAParserT__5  = 6
	OpenFGAParserT__6  = 7
	OpenFGAParserT__7  = 8
	OpenFGAParserT__8  = 9
	OpenFGAParserT__9  = 10
	OpenFGAParserT__10 = 11
	OpenFGAParserT__11 = 12
	OpenFGAParserT__12 = 13
	OpenFGAParserID    = 14
	OpenFGAParserWS    = 15
)

// OpenFGAParser rules.
const (
	OpenFGAParserRULE_start               = 0
	OpenFGAParserRULE_typeDefinition      = 1
	OpenFGAParserRULE_relation            = 2
	OpenFGAParserRULE_typeRestrictionOrId = 3
	OpenFGAParserRULE_typeRestriction     = 4
	OpenFGAParserRULE_relationReferences  = 5
	OpenFGAParserRULE_relationReference   = 6
	OpenFGAParserRULE_rewrite             = 7
	OpenFGAParserRULE_orTTU               = 8
	OpenFGAParserRULE_ors                 = 9
	OpenFGAParserRULE_or                  = 10
	OpenFGAParserRULE_ands                = 11
	OpenFGAParserRULE_and                 = 12
	OpenFGAParserRULE_exclusion           = 13
)

// IStartContext is an interface to support dynamic dispatch.
type IStartContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStartContext differentiates from other interfaces.
	IsStartContext()
}

type StartContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStartContext() *StartContext {
	var p = new(StartContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = OpenFGAParserRULE_start
	return p
}

func (*StartContext) IsStartContext() {}

func NewStartContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StartContext {
	var p = new(StartContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = OpenFGAParserRULE_start

	return p
}

func (s *StartContext) GetParser() antlr.Parser { return s.parser }

func (s *StartContext) EOF() antlr.TerminalNode {
	return s.GetToken(OpenFGAParserEOF, 0)
}

func (s *StartContext) AllTypeDefinition() []ITypeDefinitionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITypeDefinitionContext); ok {
			len++
		}
	}

	tst := make([]ITypeDefinitionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITypeDefinitionContext); ok {
			tst[i] = t.(ITypeDefinitionContext)
			i++
		}
	}

	return tst
}

func (s *StartContext) TypeDefinition(i int) ITypeDefinitionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeDefinitionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeDefinitionContext)
}

func (s *StartContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StartContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StartContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OpenFGAListener); ok {
		listenerT.EnterStart(s)
	}
}

func (s *StartContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OpenFGAListener); ok {
		listenerT.ExitStart(s)
	}
}

func (p *OpenFGAParser) Start() (localctx IStartContext) {
	this := p
	_ = this

	localctx = NewStartContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, OpenFGAParserRULE_start)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(29)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == OpenFGAParserT__0 {
		{
			p.SetState(28)
			p.TypeDefinition()
		}

		p.SetState(31)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(33)
		p.Match(OpenFGAParserEOF)
	}

	return localctx
}

// ITypeDefinitionContext is an interface to support dynamic dispatch.
type ITypeDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetObjectType returns the objectType token.
	GetObjectType() antlr.Token

	// SetObjectType sets the objectType token.
	SetObjectType(antlr.Token)

	// IsTypeDefinitionContext differentiates from other interfaces.
	IsTypeDefinitionContext()
}

type TypeDefinitionContext struct {
	*antlr.BaseParserRuleContext
	parser     antlr.Parser
	objectType antlr.Token
}

func NewEmptyTypeDefinitionContext() *TypeDefinitionContext {
	var p = new(TypeDefinitionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = OpenFGAParserRULE_typeDefinition
	return p
}

func (*TypeDefinitionContext) IsTypeDefinitionContext() {}

func NewTypeDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeDefinitionContext {
	var p = new(TypeDefinitionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = OpenFGAParserRULE_typeDefinition

	return p
}

func (s *TypeDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeDefinitionContext) GetObjectType() antlr.Token { return s.objectType }

func (s *TypeDefinitionContext) SetObjectType(v antlr.Token) { s.objectType = v }

func (s *TypeDefinitionContext) ID() antlr.TerminalNode {
	return s.GetToken(OpenFGAParserID, 0)
}

func (s *TypeDefinitionContext) AllRelation() []IRelationContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IRelationContext); ok {
			len++
		}
	}

	tst := make([]IRelationContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IRelationContext); ok {
			tst[i] = t.(IRelationContext)
			i++
		}
	}

	return tst
}

func (s *TypeDefinitionContext) Relation(i int) IRelationContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRelationContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRelationContext)
}

func (s *TypeDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeDefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OpenFGAListener); ok {
		listenerT.EnterTypeDefinition(s)
	}
}

func (s *TypeDefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OpenFGAListener); ok {
		listenerT.ExitTypeDefinition(s)
	}
}

func (p *OpenFGAParser) TypeDefinition() (localctx ITypeDefinitionContext) {
	this := p
	_ = this

	localctx = NewTypeDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, OpenFGAParserRULE_typeDefinition)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(35)
		p.Match(OpenFGAParserT__0)
	}
	{
		p.SetState(36)

		var _m = p.Match(OpenFGAParserID)

		localctx.(*TypeDefinitionContext).objectType = _m
	}
	p.SetState(43)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == OpenFGAParserT__1 {
		{
			p.SetState(37)
			p.Match(OpenFGAParserT__1)
		}
		p.SetState(39)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == OpenFGAParserT__2 {
			{
				p.SetState(38)
				p.Relation()
			}

			p.SetState(41)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}

	return localctx
}

// IRelationContext is an interface to support dynamic dispatch.
type IRelationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetName returns the name token.
	GetName() antlr.Token

	// SetName sets the name token.
	SetName(antlr.Token)

	// IsRelationContext differentiates from other interfaces.
	IsRelationContext()
}

type RelationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	name   antlr.Token
}

func NewEmptyRelationContext() *RelationContext {
	var p = new(RelationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = OpenFGAParserRULE_relation
	return p
}

func (*RelationContext) IsRelationContext() {}

func NewRelationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RelationContext {
	var p = new(RelationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = OpenFGAParserRULE_relation

	return p
}

func (s *RelationContext) GetParser() antlr.Parser { return s.parser }

func (s *RelationContext) GetName() antlr.Token { return s.name }

func (s *RelationContext) SetName(v antlr.Token) { s.name = v }

func (s *RelationContext) TypeRestrictionOrId() ITypeRestrictionOrIdContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeRestrictionOrIdContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeRestrictionOrIdContext)
}

func (s *RelationContext) ID() antlr.TerminalNode {
	return s.GetToken(OpenFGAParserID, 0)
}

func (s *RelationContext) Rewrite() IRewriteContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRewriteContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRewriteContext)
}

func (s *RelationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RelationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RelationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OpenFGAListener); ok {
		listenerT.EnterRelation(s)
	}
}

func (s *RelationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OpenFGAListener); ok {
		listenerT.ExitRelation(s)
	}
}

func (p *OpenFGAParser) Relation() (localctx IRelationContext) {
	this := p
	_ = this

	localctx = NewRelationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, OpenFGAParserRULE_relation)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(45)
		p.Match(OpenFGAParserT__2)
	}
	{
		p.SetState(46)

		var _m = p.Match(OpenFGAParserID)

		localctx.(*RelationContext).name = _m
	}
	{
		p.SetState(47)
		p.Match(OpenFGAParserT__3)
	}
	{
		p.SetState(48)
		p.TypeRestrictionOrId()
	}
	p.SetState(50)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&13312) != 0 {
		{
			p.SetState(49)
			p.Rewrite()
		}

	}

	return localctx
}

// ITypeRestrictionOrIdContext is an interface to support dynamic dispatch.
type ITypeRestrictionOrIdContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypeRestrictionOrIdContext differentiates from other interfaces.
	IsTypeRestrictionOrIdContext()
}

type TypeRestrictionOrIdContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeRestrictionOrIdContext() *TypeRestrictionOrIdContext {
	var p = new(TypeRestrictionOrIdContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = OpenFGAParserRULE_typeRestrictionOrId
	return p
}

func (*TypeRestrictionOrIdContext) IsTypeRestrictionOrIdContext() {}

func NewTypeRestrictionOrIdContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeRestrictionOrIdContext {
	var p = new(TypeRestrictionOrIdContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = OpenFGAParserRULE_typeRestrictionOrId

	return p
}

func (s *TypeRestrictionOrIdContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeRestrictionOrIdContext) CopyFrom(ctx *TypeRestrictionOrIdContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *TypeRestrictionOrIdContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeRestrictionOrIdContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type TroiTypeRestrictionContext struct {
	*TypeRestrictionOrIdContext
}

func NewTroiTypeRestrictionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TroiTypeRestrictionContext {
	var p = new(TroiTypeRestrictionContext)

	p.TypeRestrictionOrIdContext = NewEmptyTypeRestrictionOrIdContext()
	p.parser = parser
	p.CopyFrom(ctx.(*TypeRestrictionOrIdContext))

	return p
}

func (s *TroiTypeRestrictionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TroiTypeRestrictionContext) TypeRestriction() ITypeRestrictionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeRestrictionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeRestrictionContext)
}

func (s *TroiTypeRestrictionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OpenFGAListener); ok {
		listenerT.EnterTroiTypeRestriction(s)
	}
}

func (s *TroiTypeRestrictionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OpenFGAListener); ok {
		listenerT.ExitTroiTypeRestriction(s)
	}
}

type TroiIdContext struct {
	*TypeRestrictionOrIdContext
	computedUserset antlr.Token
}

func NewTroiIdContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TroiIdContext {
	var p = new(TroiIdContext)

	p.TypeRestrictionOrIdContext = NewEmptyTypeRestrictionOrIdContext()
	p.parser = parser
	p.CopyFrom(ctx.(*TypeRestrictionOrIdContext))

	return p
}

func (s *TroiIdContext) GetComputedUserset() antlr.Token { return s.computedUserset }

func (s *TroiIdContext) SetComputedUserset(v antlr.Token) { s.computedUserset = v }

func (s *TroiIdContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TroiIdContext) ID() antlr.TerminalNode {
	return s.GetToken(OpenFGAParserID, 0)
}

func (s *TroiIdContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OpenFGAListener); ok {
		listenerT.EnterTroiId(s)
	}
}

func (s *TroiIdContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OpenFGAListener); ok {
		listenerT.ExitTroiId(s)
	}
}

func (p *OpenFGAParser) TypeRestrictionOrId() (localctx ITypeRestrictionOrIdContext) {
	this := p
	_ = this

	localctx = NewTypeRestrictionOrIdContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, OpenFGAParserRULE_typeRestrictionOrId)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(54)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case OpenFGAParserT__4:
		localctx = NewTroiTypeRestrictionContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(52)
			p.TypeRestriction()
		}

	case OpenFGAParserID:
		localctx = NewTroiIdContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(53)

			var _m = p.Match(OpenFGAParserID)

			localctx.(*TroiIdContext).computedUserset = _m
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ITypeRestrictionContext is an interface to support dynamic dispatch.
type ITypeRestrictionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypeRestrictionContext differentiates from other interfaces.
	IsTypeRestrictionContext()
}

type TypeRestrictionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeRestrictionContext() *TypeRestrictionContext {
	var p = new(TypeRestrictionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = OpenFGAParserRULE_typeRestriction
	return p
}

func (*TypeRestrictionContext) IsTypeRestrictionContext() {}

func NewTypeRestrictionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeRestrictionContext {
	var p = new(TypeRestrictionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = OpenFGAParserRULE_typeRestriction

	return p
}

func (s *TypeRestrictionContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeRestrictionContext) RelationReferences() IRelationReferencesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRelationReferencesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRelationReferencesContext)
}

func (s *TypeRestrictionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeRestrictionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeRestrictionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OpenFGAListener); ok {
		listenerT.EnterTypeRestriction(s)
	}
}

func (s *TypeRestrictionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OpenFGAListener); ok {
		listenerT.ExitTypeRestriction(s)
	}
}

func (p *OpenFGAParser) TypeRestriction() (localctx ITypeRestrictionContext) {
	this := p
	_ = this

	localctx = NewTypeRestrictionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, OpenFGAParserRULE_typeRestriction)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(56)
		p.Match(OpenFGAParserT__4)
	}
	{
		p.SetState(57)
		p.RelationReferences()
	}
	{
		p.SetState(58)
		p.Match(OpenFGAParserT__5)
	}

	return localctx
}

// IRelationReferencesContext is an interface to support dynamic dispatch.
type IRelationReferencesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRelationReferencesContext differentiates from other interfaces.
	IsRelationReferencesContext()
}

type RelationReferencesContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRelationReferencesContext() *RelationReferencesContext {
	var p = new(RelationReferencesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = OpenFGAParserRULE_relationReferences
	return p
}

func (*RelationReferencesContext) IsRelationReferencesContext() {}

func NewRelationReferencesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RelationReferencesContext {
	var p = new(RelationReferencesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = OpenFGAParserRULE_relationReferences

	return p
}

func (s *RelationReferencesContext) GetParser() antlr.Parser { return s.parser }

func (s *RelationReferencesContext) RelationReference() IRelationReferenceContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRelationReferenceContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRelationReferenceContext)
}

func (s *RelationReferencesContext) AllRelationReferences() []IRelationReferencesContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IRelationReferencesContext); ok {
			len++
		}
	}

	tst := make([]IRelationReferencesContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IRelationReferencesContext); ok {
			tst[i] = t.(IRelationReferencesContext)
			i++
		}
	}

	return tst
}

func (s *RelationReferencesContext) RelationReferences(i int) IRelationReferencesContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRelationReferencesContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRelationReferencesContext)
}

func (s *RelationReferencesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RelationReferencesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RelationReferencesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OpenFGAListener); ok {
		listenerT.EnterRelationReferences(s)
	}
}

func (s *RelationReferencesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OpenFGAListener); ok {
		listenerT.ExitRelationReferences(s)
	}
}

func (p *OpenFGAParser) RelationReferences() (localctx IRelationReferencesContext) {
	this := p
	_ = this

	localctx = NewRelationReferencesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, OpenFGAParserRULE_relationReferences)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(60)
		p.RelationReference()
	}
	p.SetState(65)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(61)
				p.Match(OpenFGAParserT__6)
			}
			{
				p.SetState(62)
				p.RelationReferences()
			}

		}
		p.SetState(67)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext())
	}

	return localctx
}

// IRelationReferenceContext is an interface to support dynamic dispatch.
type IRelationReferenceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRelationReferenceContext differentiates from other interfaces.
	IsRelationReferenceContext()
}

type RelationReferenceContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRelationReferenceContext() *RelationReferenceContext {
	var p = new(RelationReferenceContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = OpenFGAParserRULE_relationReference
	return p
}

func (*RelationReferenceContext) IsRelationReferenceContext() {}

func NewRelationReferenceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RelationReferenceContext {
	var p = new(RelationReferenceContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = OpenFGAParserRULE_relationReference

	return p
}

func (s *RelationReferenceContext) GetParser() antlr.Parser { return s.parser }

func (s *RelationReferenceContext) CopyFrom(ctx *RelationReferenceContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *RelationReferenceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RelationReferenceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type RrTypeAndRelationContext struct {
	*RelationReferenceContext
	t antlr.Token
	r antlr.Token
}

func NewRrTypeAndRelationContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *RrTypeAndRelationContext {
	var p = new(RrTypeAndRelationContext)

	p.RelationReferenceContext = NewEmptyRelationReferenceContext()
	p.parser = parser
	p.CopyFrom(ctx.(*RelationReferenceContext))

	return p
}

func (s *RrTypeAndRelationContext) GetT() antlr.Token { return s.t }

func (s *RrTypeAndRelationContext) GetR() antlr.Token { return s.r }

func (s *RrTypeAndRelationContext) SetT(v antlr.Token) { s.t = v }

func (s *RrTypeAndRelationContext) SetR(v antlr.Token) { s.r = v }

func (s *RrTypeAndRelationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RrTypeAndRelationContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(OpenFGAParserID)
}

func (s *RrTypeAndRelationContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(OpenFGAParserID, i)
}

func (s *RrTypeAndRelationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OpenFGAListener); ok {
		listenerT.EnterRrTypeAndRelation(s)
	}
}

func (s *RrTypeAndRelationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OpenFGAListener); ok {
		listenerT.ExitRrTypeAndRelation(s)
	}
}

type RrTypeAndWildcardContext struct {
	*RelationReferenceContext
	t antlr.Token
}

func NewRrTypeAndWildcardContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *RrTypeAndWildcardContext {
	var p = new(RrTypeAndWildcardContext)

	p.RelationReferenceContext = NewEmptyRelationReferenceContext()
	p.parser = parser
	p.CopyFrom(ctx.(*RelationReferenceContext))

	return p
}

func (s *RrTypeAndWildcardContext) GetT() antlr.Token { return s.t }

func (s *RrTypeAndWildcardContext) SetT(v antlr.Token) { s.t = v }

func (s *RrTypeAndWildcardContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RrTypeAndWildcardContext) ID() antlr.TerminalNode {
	return s.GetToken(OpenFGAParserID, 0)
}

func (s *RrTypeAndWildcardContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OpenFGAListener); ok {
		listenerT.EnterRrTypeAndWildcard(s)
	}
}

func (s *RrTypeAndWildcardContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OpenFGAListener); ok {
		listenerT.ExitRrTypeAndWildcard(s)
	}
}

type RrTypeContext struct {
	*RelationReferenceContext
	t antlr.Token
}

func NewRrTypeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *RrTypeContext {
	var p = new(RrTypeContext)

	p.RelationReferenceContext = NewEmptyRelationReferenceContext()
	p.parser = parser
	p.CopyFrom(ctx.(*RelationReferenceContext))

	return p
}

func (s *RrTypeContext) GetT() antlr.Token { return s.t }

func (s *RrTypeContext) SetT(v antlr.Token) { s.t = v }

func (s *RrTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RrTypeContext) ID() antlr.TerminalNode {
	return s.GetToken(OpenFGAParserID, 0)
}

func (s *RrTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OpenFGAListener); ok {
		listenerT.EnterRrType(s)
	}
}

func (s *RrTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OpenFGAListener); ok {
		listenerT.ExitRrType(s)
	}
}

func (p *OpenFGAParser) RelationReference() (localctx IRelationReferenceContext) {
	this := p
	_ = this

	localctx = NewRelationReferenceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, OpenFGAParserRULE_relationReference)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(74)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext()) {
	case 1:
		localctx = NewRrTypeContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(68)

			var _m = p.Match(OpenFGAParserID)

			localctx.(*RrTypeContext).t = _m
		}

	case 2:
		localctx = NewRrTypeAndRelationContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(69)

			var _m = p.Match(OpenFGAParserID)

			localctx.(*RrTypeAndRelationContext).t = _m
		}
		{
			p.SetState(70)
			p.Match(OpenFGAParserT__7)
		}
		{
			p.SetState(71)

			var _m = p.Match(OpenFGAParserID)

			localctx.(*RrTypeAndRelationContext).r = _m
		}

	case 3:
		localctx = NewRrTypeAndWildcardContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(72)

			var _m = p.Match(OpenFGAParserID)

			localctx.(*RrTypeAndWildcardContext).t = _m
		}
		{
			p.SetState(73)
			p.Match(OpenFGAParserT__8)
		}

	}

	return localctx
}

// IRewriteContext is an interface to support dynamic dispatch.
type IRewriteContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRewriteContext differentiates from other interfaces.
	IsRewriteContext()
}

type RewriteContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRewriteContext() *RewriteContext {
	var p = new(RewriteContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = OpenFGAParserRULE_rewrite
	return p
}

func (*RewriteContext) IsRewriteContext() {}

func NewRewriteContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RewriteContext {
	var p = new(RewriteContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = OpenFGAParserRULE_rewrite

	return p
}

func (s *RewriteContext) GetParser() antlr.Parser { return s.parser }

func (s *RewriteContext) OrTTU() IOrTTUContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOrTTUContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOrTTUContext)
}

func (s *RewriteContext) Ors() IOrsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOrsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOrsContext)
}

func (s *RewriteContext) Ands() IAndsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAndsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAndsContext)
}

func (s *RewriteContext) Exclusion() IExclusionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExclusionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExclusionContext)
}

func (s *RewriteContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RewriteContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RewriteContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OpenFGAListener); ok {
		listenerT.EnterRewrite(s)
	}
}

func (s *RewriteContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OpenFGAListener); ok {
		listenerT.ExitRewrite(s)
	}
}

func (p *OpenFGAParser) Rewrite() (localctx IRewriteContext) {
	this := p
	_ = this

	localctx = NewRewriteContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, OpenFGAParserRULE_rewrite)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(80)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(76)
			p.OrTTU()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(77)
			p.Ors()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(78)
			p.Ands()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(79)
			p.Exclusion()
		}

	}

	return localctx
}

// IOrTTUContext is an interface to support dynamic dispatch.
type IOrTTUContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetComputedUserset returns the computedUserset token.
	GetComputedUserset() antlr.Token

	// GetTupleset returns the tupleset token.
	GetTupleset() antlr.Token

	// SetComputedUserset sets the computedUserset token.
	SetComputedUserset(antlr.Token)

	// SetTupleset sets the tupleset token.
	SetTupleset(antlr.Token)

	// IsOrTTUContext differentiates from other interfaces.
	IsOrTTUContext()
}

type OrTTUContext struct {
	*antlr.BaseParserRuleContext
	parser          antlr.Parser
	computedUserset antlr.Token
	tupleset        antlr.Token
}

func NewEmptyOrTTUContext() *OrTTUContext {
	var p = new(OrTTUContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = OpenFGAParserRULE_orTTU
	return p
}

func (*OrTTUContext) IsOrTTUContext() {}

func NewOrTTUContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OrTTUContext {
	var p = new(OrTTUContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = OpenFGAParserRULE_orTTU

	return p
}

func (s *OrTTUContext) GetParser() antlr.Parser { return s.parser }

func (s *OrTTUContext) GetComputedUserset() antlr.Token { return s.computedUserset }

func (s *OrTTUContext) GetTupleset() antlr.Token { return s.tupleset }

func (s *OrTTUContext) SetComputedUserset(v antlr.Token) { s.computedUserset = v }

func (s *OrTTUContext) SetTupleset(v antlr.Token) { s.tupleset = v }

func (s *OrTTUContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(OpenFGAParserID)
}

func (s *OrTTUContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(OpenFGAParserID, i)
}

func (s *OrTTUContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OrTTUContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OrTTUContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OpenFGAListener); ok {
		listenerT.EnterOrTTU(s)
	}
}

func (s *OrTTUContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OpenFGAListener); ok {
		listenerT.ExitOrTTU(s)
	}
}

func (p *OpenFGAParser) OrTTU() (localctx IOrTTUContext) {
	this := p
	_ = this

	localctx = NewOrTTUContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, OpenFGAParserRULE_orTTU)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(82)
		p.Match(OpenFGAParserT__9)
	}
	{
		p.SetState(83)

		var _m = p.Match(OpenFGAParserID)

		localctx.(*OrTTUContext).computedUserset = _m
	}
	{
		p.SetState(84)
		p.Match(OpenFGAParserT__10)
	}
	{
		p.SetState(85)

		var _m = p.Match(OpenFGAParserID)

		localctx.(*OrTTUContext).tupleset = _m
	}

	return localctx
}

// IOrsContext is an interface to support dynamic dispatch.
type IOrsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOrsContext differentiates from other interfaces.
	IsOrsContext()
}

type OrsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOrsContext() *OrsContext {
	var p = new(OrsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = OpenFGAParserRULE_ors
	return p
}

func (*OrsContext) IsOrsContext() {}

func NewOrsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OrsContext {
	var p = new(OrsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = OpenFGAParserRULE_ors

	return p
}

func (s *OrsContext) GetParser() antlr.Parser { return s.parser }

func (s *OrsContext) Or() IOrContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOrContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOrContext)
}

func (s *OrsContext) AllOrs() []IOrsContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IOrsContext); ok {
			len++
		}
	}

	tst := make([]IOrsContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IOrsContext); ok {
			tst[i] = t.(IOrsContext)
			i++
		}
	}

	return tst
}

func (s *OrsContext) Ors(i int) IOrsContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOrsContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOrsContext)
}

func (s *OrsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OrsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OrsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OpenFGAListener); ok {
		listenerT.EnterOrs(s)
	}
}

func (s *OrsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OpenFGAListener); ok {
		listenerT.ExitOrs(s)
	}
}

func (p *OpenFGAParser) Ors() (localctx IOrsContext) {
	this := p
	_ = this

	localctx = NewOrsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, OpenFGAParserRULE_ors)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(87)
		p.Or()
	}
	p.SetState(91)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(88)
				p.Ors()
			}

		}
		p.SetState(93)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext())
	}

	return localctx
}

// IOrContext is an interface to support dynamic dispatch.
type IOrContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetId returns the id token.
	GetId() antlr.Token

	// SetId sets the id token.
	SetId(antlr.Token)

	// IsOrContext differentiates from other interfaces.
	IsOrContext()
}

type OrContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	id     antlr.Token
}

func NewEmptyOrContext() *OrContext {
	var p = new(OrContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = OpenFGAParserRULE_or
	return p
}

func (*OrContext) IsOrContext() {}

func NewOrContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OrContext {
	var p = new(OrContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = OpenFGAParserRULE_or

	return p
}

func (s *OrContext) GetParser() antlr.Parser { return s.parser }

func (s *OrContext) GetId() antlr.Token { return s.id }

func (s *OrContext) SetId(v antlr.Token) { s.id = v }

func (s *OrContext) ID() antlr.TerminalNode {
	return s.GetToken(OpenFGAParserID, 0)
}

func (s *OrContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OrContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OrContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OpenFGAListener); ok {
		listenerT.EnterOr(s)
	}
}

func (s *OrContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OpenFGAListener); ok {
		listenerT.ExitOr(s)
	}
}

func (p *OpenFGAParser) Or() (localctx IOrContext) {
	this := p
	_ = this

	localctx = NewOrContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, OpenFGAParserRULE_or)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(94)
		p.Match(OpenFGAParserT__9)
	}
	{
		p.SetState(95)

		var _m = p.Match(OpenFGAParserID)

		localctx.(*OrContext).id = _m
	}

	return localctx
}

// IAndsContext is an interface to support dynamic dispatch.
type IAndsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAndsContext differentiates from other interfaces.
	IsAndsContext()
}

type AndsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAndsContext() *AndsContext {
	var p = new(AndsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = OpenFGAParserRULE_ands
	return p
}

func (*AndsContext) IsAndsContext() {}

func NewAndsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AndsContext {
	var p = new(AndsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = OpenFGAParserRULE_ands

	return p
}

func (s *AndsContext) GetParser() antlr.Parser { return s.parser }

func (s *AndsContext) And() IAndContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAndContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAndContext)
}

func (s *AndsContext) AllAnds() []IAndsContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IAndsContext); ok {
			len++
		}
	}

	tst := make([]IAndsContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IAndsContext); ok {
			tst[i] = t.(IAndsContext)
			i++
		}
	}

	return tst
}

func (s *AndsContext) Ands(i int) IAndsContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAndsContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAndsContext)
}

func (s *AndsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AndsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AndsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OpenFGAListener); ok {
		listenerT.EnterAnds(s)
	}
}

func (s *AndsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OpenFGAListener); ok {
		listenerT.ExitAnds(s)
	}
}

func (p *OpenFGAParser) Ands() (localctx IAndsContext) {
	this := p
	_ = this

	localctx = NewAndsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, OpenFGAParserRULE_ands)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(97)
		p.And()
	}
	p.SetState(101)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(98)
				p.Ands()
			}

		}
		p.SetState(103)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext())
	}

	return localctx
}

// IAndContext is an interface to support dynamic dispatch.
type IAndContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetId returns the id token.
	GetId() antlr.Token

	// SetId sets the id token.
	SetId(antlr.Token)

	// IsAndContext differentiates from other interfaces.
	IsAndContext()
}

type AndContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	id     antlr.Token
}

func NewEmptyAndContext() *AndContext {
	var p = new(AndContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = OpenFGAParserRULE_and
	return p
}

func (*AndContext) IsAndContext() {}

func NewAndContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AndContext {
	var p = new(AndContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = OpenFGAParserRULE_and

	return p
}

func (s *AndContext) GetParser() antlr.Parser { return s.parser }

func (s *AndContext) GetId() antlr.Token { return s.id }

func (s *AndContext) SetId(v antlr.Token) { s.id = v }

func (s *AndContext) ID() antlr.TerminalNode {
	return s.GetToken(OpenFGAParserID, 0)
}

func (s *AndContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AndContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AndContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OpenFGAListener); ok {
		listenerT.EnterAnd(s)
	}
}

func (s *AndContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OpenFGAListener); ok {
		listenerT.ExitAnd(s)
	}
}

func (p *OpenFGAParser) And() (localctx IAndContext) {
	this := p
	_ = this

	localctx = NewAndContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, OpenFGAParserRULE_and)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(104)
		p.Match(OpenFGAParserT__11)
	}
	{
		p.SetState(105)

		var _m = p.Match(OpenFGAParserID)

		localctx.(*AndContext).id = _m
	}

	return localctx
}

// IExclusionContext is an interface to support dynamic dispatch.
type IExclusionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetId returns the id token.
	GetId() antlr.Token

	// SetId sets the id token.
	SetId(antlr.Token)

	// IsExclusionContext differentiates from other interfaces.
	IsExclusionContext()
}

type ExclusionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	id     antlr.Token
}

func NewEmptyExclusionContext() *ExclusionContext {
	var p = new(ExclusionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = OpenFGAParserRULE_exclusion
	return p
}

func (*ExclusionContext) IsExclusionContext() {}

func NewExclusionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExclusionContext {
	var p = new(ExclusionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = OpenFGAParserRULE_exclusion

	return p
}

func (s *ExclusionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExclusionContext) GetId() antlr.Token { return s.id }

func (s *ExclusionContext) SetId(v antlr.Token) { s.id = v }

func (s *ExclusionContext) ID() antlr.TerminalNode {
	return s.GetToken(OpenFGAParserID, 0)
}

func (s *ExclusionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExclusionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExclusionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OpenFGAListener); ok {
		listenerT.EnterExclusion(s)
	}
}

func (s *ExclusionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OpenFGAListener); ok {
		listenerT.ExitExclusion(s)
	}
}

func (p *OpenFGAParser) Exclusion() (localctx IExclusionContext) {
	this := p
	_ = this

	localctx = NewExclusionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, OpenFGAParserRULE_exclusion)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(107)
		p.Match(OpenFGAParserT__12)
	}
	{
		p.SetState(108)

		var _m = p.Match(OpenFGAParserID)

		localctx.(*ExclusionContext).id = _m
	}

	return localctx
}
