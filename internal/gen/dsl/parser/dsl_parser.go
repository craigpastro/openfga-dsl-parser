// Code generated from DSL.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser // DSL

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type DSLParser struct {
	*antlr.BaseParser
}

var DSLParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func dslParserInit() {
	staticData := &DSLParserStaticData
	staticData.LiteralNames = []string{
		"", "'type'", "'relations'", "'define'", "'as'", "':'", "'['", "']'",
		"','", "'#'", "':*'", "'self'", "'from'", "'or'", "'and'", "'but not'",
		"'('", "')'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "ID", "WS",
	}
	staticData.RuleNames = []string{
		"dsl", "typeDefinition", "relation", "typeRestriction", "relationReferences",
		"relationReference", "rewrite",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 19, 87, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 1, 0, 4, 0, 16, 8, 0, 11, 0, 12, 0, 17, 1, 0,
		1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 4, 1, 26, 8, 1, 11, 1, 12, 1, 27, 3, 1, 30,
		8, 1, 1, 2, 1, 2, 1, 2, 3, 2, 35, 8, 2, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1,
		3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 5, 4, 48, 8, 4, 10, 4, 12, 4, 51, 9, 4,
		1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 3, 5, 59, 8, 5, 1, 6, 1, 6, 1, 6, 1,
		6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 3, 6, 71, 8, 6, 1, 6, 1, 6, 1, 6,
		1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 5, 6, 82, 8, 6, 10, 6, 12, 6, 85, 9,
		6, 1, 6, 0, 1, 12, 7, 0, 2, 4, 6, 8, 10, 12, 0, 0, 92, 0, 15, 1, 0, 0,
		0, 2, 21, 1, 0, 0, 0, 4, 31, 1, 0, 0, 0, 6, 39, 1, 0, 0, 0, 8, 44, 1, 0,
		0, 0, 10, 58, 1, 0, 0, 0, 12, 70, 1, 0, 0, 0, 14, 16, 3, 2, 1, 0, 15, 14,
		1, 0, 0, 0, 16, 17, 1, 0, 0, 0, 17, 15, 1, 0, 0, 0, 17, 18, 1, 0, 0, 0,
		18, 19, 1, 0, 0, 0, 19, 20, 5, 0, 0, 1, 20, 1, 1, 0, 0, 0, 21, 22, 5, 1,
		0, 0, 22, 29, 5, 18, 0, 0, 23, 25, 5, 2, 0, 0, 24, 26, 3, 4, 2, 0, 25,
		24, 1, 0, 0, 0, 26, 27, 1, 0, 0, 0, 27, 25, 1, 0, 0, 0, 27, 28, 1, 0, 0,
		0, 28, 30, 1, 0, 0, 0, 29, 23, 1, 0, 0, 0, 29, 30, 1, 0, 0, 0, 30, 3, 1,
		0, 0, 0, 31, 32, 5, 3, 0, 0, 32, 34, 5, 18, 0, 0, 33, 35, 3, 6, 3, 0, 34,
		33, 1, 0, 0, 0, 34, 35, 1, 0, 0, 0, 35, 36, 1, 0, 0, 0, 36, 37, 5, 4, 0,
		0, 37, 38, 3, 12, 6, 0, 38, 5, 1, 0, 0, 0, 39, 40, 5, 5, 0, 0, 40, 41,
		5, 6, 0, 0, 41, 42, 3, 8, 4, 0, 42, 43, 5, 7, 0, 0, 43, 7, 1, 0, 0, 0,
		44, 49, 3, 10, 5, 0, 45, 46, 5, 8, 0, 0, 46, 48, 3, 8, 4, 0, 47, 45, 1,
		0, 0, 0, 48, 51, 1, 0, 0, 0, 49, 47, 1, 0, 0, 0, 49, 50, 1, 0, 0, 0, 50,
		9, 1, 0, 0, 0, 51, 49, 1, 0, 0, 0, 52, 59, 5, 18, 0, 0, 53, 54, 5, 18,
		0, 0, 54, 55, 5, 9, 0, 0, 55, 59, 5, 18, 0, 0, 56, 57, 5, 18, 0, 0, 57,
		59, 5, 10, 0, 0, 58, 52, 1, 0, 0, 0, 58, 53, 1, 0, 0, 0, 58, 56, 1, 0,
		0, 0, 59, 11, 1, 0, 0, 0, 60, 61, 6, 6, -1, 0, 61, 71, 5, 11, 0, 0, 62,
		63, 5, 18, 0, 0, 63, 64, 5, 12, 0, 0, 64, 71, 5, 18, 0, 0, 65, 71, 5, 18,
		0, 0, 66, 67, 5, 16, 0, 0, 67, 68, 3, 12, 6, 0, 68, 69, 5, 17, 0, 0, 69,
		71, 1, 0, 0, 0, 70, 60, 1, 0, 0, 0, 70, 62, 1, 0, 0, 0, 70, 65, 1, 0, 0,
		0, 70, 66, 1, 0, 0, 0, 71, 83, 1, 0, 0, 0, 72, 73, 10, 5, 0, 0, 73, 74,
		5, 13, 0, 0, 74, 82, 3, 12, 6, 6, 75, 76, 10, 4, 0, 0, 76, 77, 5, 14, 0,
		0, 77, 82, 3, 12, 6, 5, 78, 79, 10, 3, 0, 0, 79, 80, 5, 15, 0, 0, 80, 82,
		3, 12, 6, 4, 81, 72, 1, 0, 0, 0, 81, 75, 1, 0, 0, 0, 81, 78, 1, 0, 0, 0,
		82, 85, 1, 0, 0, 0, 83, 81, 1, 0, 0, 0, 83, 84, 1, 0, 0, 0, 84, 13, 1,
		0, 0, 0, 85, 83, 1, 0, 0, 0, 9, 17, 27, 29, 34, 49, 58, 70, 81, 83,
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

// DSLParserInit initializes any static state used to implement DSLParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewDSLParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func DSLParserInit() {
	staticData := &DSLParserStaticData
	staticData.once.Do(dslParserInit)
}

// NewDSLParser produces a new parser instance for the optional input antlr.TokenStream.
func NewDSLParser(input antlr.TokenStream) *DSLParser {
	DSLParserInit()
	this := new(DSLParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &DSLParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "DSL.g4"

	return this
}

// DSLParser tokens.
const (
	DSLParserEOF   = antlr.TokenEOF
	DSLParserT__0  = 1
	DSLParserT__1  = 2
	DSLParserT__2  = 3
	DSLParserT__3  = 4
	DSLParserT__4  = 5
	DSLParserT__5  = 6
	DSLParserT__6  = 7
	DSLParserT__7  = 8
	DSLParserT__8  = 9
	DSLParserT__9  = 10
	DSLParserT__10 = 11
	DSLParserT__11 = 12
	DSLParserT__12 = 13
	DSLParserT__13 = 14
	DSLParserT__14 = 15
	DSLParserT__15 = 16
	DSLParserT__16 = 17
	DSLParserID    = 18
	DSLParserWS    = 19
)

// DSLParser rules.
const (
	DSLParserRULE_dsl                = 0
	DSLParserRULE_typeDefinition     = 1
	DSLParserRULE_relation           = 2
	DSLParserRULE_typeRestriction    = 3
	DSLParserRULE_relationReferences = 4
	DSLParserRULE_relationReference  = 5
	DSLParserRULE_rewrite            = 6
)

// IDslContext is an interface to support dynamic dispatch.
type IDslContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EOF() antlr.TerminalNode
	AllTypeDefinition() []ITypeDefinitionContext
	TypeDefinition(i int) ITypeDefinitionContext

	// IsDslContext differentiates from other interfaces.
	IsDslContext()
}

type DslContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDslContext() *DslContext {
	var p = new(DslContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DSLParserRULE_dsl
	return p
}

func InitEmptyDslContext(p *DslContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DSLParserRULE_dsl
}

func (*DslContext) IsDslContext() {}

func NewDslContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DslContext {
	var p = new(DslContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DSLParserRULE_dsl

	return p
}

func (s *DslContext) GetParser() antlr.Parser { return s.parser }

func (s *DslContext) EOF() antlr.TerminalNode {
	return s.GetToken(DSLParserEOF, 0)
}

func (s *DslContext) AllTypeDefinition() []ITypeDefinitionContext {
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

func (s *DslContext) TypeDefinition(i int) ITypeDefinitionContext {
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

func (s *DslContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DslContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DslContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DSLListener); ok {
		listenerT.EnterDsl(s)
	}
}

func (s *DslContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DSLListener); ok {
		listenerT.ExitDsl(s)
	}
}

func (p *DSLParser) Dsl() (localctx IDslContext) {
	localctx = NewDslContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, DSLParserRULE_dsl)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(15)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == DSLParserT__0 {
		{
			p.SetState(14)
			p.TypeDefinition()
		}

		p.SetState(17)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(19)
		p.Match(DSLParserEOF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
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

	// Getter signatures
	ID() antlr.TerminalNode
	AllRelation() []IRelationContext
	Relation(i int) IRelationContext

	// IsTypeDefinitionContext differentiates from other interfaces.
	IsTypeDefinitionContext()
}

type TypeDefinitionContext struct {
	antlr.BaseParserRuleContext
	parser     antlr.Parser
	objectType antlr.Token
}

func NewEmptyTypeDefinitionContext() *TypeDefinitionContext {
	var p = new(TypeDefinitionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DSLParserRULE_typeDefinition
	return p
}

func InitEmptyTypeDefinitionContext(p *TypeDefinitionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DSLParserRULE_typeDefinition
}

func (*TypeDefinitionContext) IsTypeDefinitionContext() {}

func NewTypeDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeDefinitionContext {
	var p = new(TypeDefinitionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DSLParserRULE_typeDefinition

	return p
}

func (s *TypeDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeDefinitionContext) GetObjectType() antlr.Token { return s.objectType }

func (s *TypeDefinitionContext) SetObjectType(v antlr.Token) { s.objectType = v }

func (s *TypeDefinitionContext) ID() antlr.TerminalNode {
	return s.GetToken(DSLParserID, 0)
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
	if listenerT, ok := listener.(DSLListener); ok {
		listenerT.EnterTypeDefinition(s)
	}
}

func (s *TypeDefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DSLListener); ok {
		listenerT.ExitTypeDefinition(s)
	}
}

func (p *DSLParser) TypeDefinition() (localctx ITypeDefinitionContext) {
	localctx = NewTypeDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, DSLParserRULE_typeDefinition)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(21)
		p.Match(DSLParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(22)

		var _m = p.Match(DSLParserID)

		localctx.(*TypeDefinitionContext).objectType = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(29)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == DSLParserT__1 {
		{
			p.SetState(23)
			p.Match(DSLParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(25)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == DSLParserT__2 {
			{
				p.SetState(24)
				p.Relation()
			}

			p.SetState(27)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
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

	// Getter signatures
	Rewrite() IRewriteContext
	ID() antlr.TerminalNode
	TypeRestriction() ITypeRestrictionContext

	// IsRelationContext differentiates from other interfaces.
	IsRelationContext()
}

type RelationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	name   antlr.Token
}

func NewEmptyRelationContext() *RelationContext {
	var p = new(RelationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DSLParserRULE_relation
	return p
}

func InitEmptyRelationContext(p *RelationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DSLParserRULE_relation
}

func (*RelationContext) IsRelationContext() {}

func NewRelationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RelationContext {
	var p = new(RelationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DSLParserRULE_relation

	return p
}

func (s *RelationContext) GetParser() antlr.Parser { return s.parser }

func (s *RelationContext) GetName() antlr.Token { return s.name }

func (s *RelationContext) SetName(v antlr.Token) { s.name = v }

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

func (s *RelationContext) ID() antlr.TerminalNode {
	return s.GetToken(DSLParserID, 0)
}

func (s *RelationContext) TypeRestriction() ITypeRestrictionContext {
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

func (s *RelationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RelationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RelationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DSLListener); ok {
		listenerT.EnterRelation(s)
	}
}

func (s *RelationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DSLListener); ok {
		listenerT.ExitRelation(s)
	}
}

func (p *DSLParser) Relation() (localctx IRelationContext) {
	localctx = NewRelationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, DSLParserRULE_relation)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(31)
		p.Match(DSLParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(32)

		var _m = p.Match(DSLParserID)

		localctx.(*RelationContext).name = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(34)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == DSLParserT__4 {
		{
			p.SetState(33)
			p.TypeRestriction()
		}

	}
	{
		p.SetState(36)
		p.Match(DSLParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(37)
		p.rewrite(0)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITypeRestrictionContext is an interface to support dynamic dispatch.
type ITypeRestrictionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	RelationReferences() IRelationReferencesContext

	// IsTypeRestrictionContext differentiates from other interfaces.
	IsTypeRestrictionContext()
}

type TypeRestrictionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeRestrictionContext() *TypeRestrictionContext {
	var p = new(TypeRestrictionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DSLParserRULE_typeRestriction
	return p
}

func InitEmptyTypeRestrictionContext(p *TypeRestrictionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DSLParserRULE_typeRestriction
}

func (*TypeRestrictionContext) IsTypeRestrictionContext() {}

func NewTypeRestrictionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeRestrictionContext {
	var p = new(TypeRestrictionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DSLParserRULE_typeRestriction

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
	if listenerT, ok := listener.(DSLListener); ok {
		listenerT.EnterTypeRestriction(s)
	}
}

func (s *TypeRestrictionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DSLListener); ok {
		listenerT.ExitTypeRestriction(s)
	}
}

func (p *DSLParser) TypeRestriction() (localctx ITypeRestrictionContext) {
	localctx = NewTypeRestrictionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, DSLParserRULE_typeRestriction)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(39)
		p.Match(DSLParserT__4)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(40)
		p.Match(DSLParserT__5)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(41)
		p.RelationReferences()
	}
	{
		p.SetState(42)
		p.Match(DSLParserT__6)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IRelationReferencesContext is an interface to support dynamic dispatch.
type IRelationReferencesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	RelationReference() IRelationReferenceContext
	AllRelationReferences() []IRelationReferencesContext
	RelationReferences(i int) IRelationReferencesContext

	// IsRelationReferencesContext differentiates from other interfaces.
	IsRelationReferencesContext()
}

type RelationReferencesContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRelationReferencesContext() *RelationReferencesContext {
	var p = new(RelationReferencesContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DSLParserRULE_relationReferences
	return p
}

func InitEmptyRelationReferencesContext(p *RelationReferencesContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DSLParserRULE_relationReferences
}

func (*RelationReferencesContext) IsRelationReferencesContext() {}

func NewRelationReferencesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RelationReferencesContext {
	var p = new(RelationReferencesContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DSLParserRULE_relationReferences

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
	if listenerT, ok := listener.(DSLListener); ok {
		listenerT.EnterRelationReferences(s)
	}
}

func (s *RelationReferencesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DSLListener); ok {
		listenerT.ExitRelationReferences(s)
	}
}

func (p *DSLParser) RelationReferences() (localctx IRelationReferencesContext) {
	localctx = NewRelationReferencesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, DSLParserRULE_relationReferences)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(44)
		p.RelationReference()
	}
	p.SetState(49)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(45)
				p.Match(DSLParserT__7)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(46)
				p.RelationReferences()
			}

		}
		p.SetState(51)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
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
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRelationReferenceContext() *RelationReferenceContext {
	var p = new(RelationReferenceContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DSLParserRULE_relationReference
	return p
}

func InitEmptyRelationReferenceContext(p *RelationReferenceContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DSLParserRULE_relationReference
}

func (*RelationReferenceContext) IsRelationReferenceContext() {}

func NewRelationReferenceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RelationReferenceContext {
	var p = new(RelationReferenceContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DSLParserRULE_relationReference

	return p
}

func (s *RelationReferenceContext) GetParser() antlr.Parser { return s.parser }

func (s *RelationReferenceContext) CopyAll(ctx *RelationReferenceContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *RelationReferenceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RelationReferenceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type RrTypeAndRelationContext struct {
	RelationReferenceContext
	t antlr.Token
	r antlr.Token
}

func NewRrTypeAndRelationContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *RrTypeAndRelationContext {
	var p = new(RrTypeAndRelationContext)

	InitEmptyRelationReferenceContext(&p.RelationReferenceContext)
	p.parser = parser
	p.CopyAll(ctx.(*RelationReferenceContext))

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
	return s.GetTokens(DSLParserID)
}

func (s *RrTypeAndRelationContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(DSLParserID, i)
}

func (s *RrTypeAndRelationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DSLListener); ok {
		listenerT.EnterRrTypeAndRelation(s)
	}
}

func (s *RrTypeAndRelationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DSLListener); ok {
		listenerT.ExitRrTypeAndRelation(s)
	}
}

type RrTypeAndWildcardContext struct {
	RelationReferenceContext
	t antlr.Token
}

func NewRrTypeAndWildcardContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *RrTypeAndWildcardContext {
	var p = new(RrTypeAndWildcardContext)

	InitEmptyRelationReferenceContext(&p.RelationReferenceContext)
	p.parser = parser
	p.CopyAll(ctx.(*RelationReferenceContext))

	return p
}

func (s *RrTypeAndWildcardContext) GetT() antlr.Token { return s.t }

func (s *RrTypeAndWildcardContext) SetT(v antlr.Token) { s.t = v }

func (s *RrTypeAndWildcardContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RrTypeAndWildcardContext) ID() antlr.TerminalNode {
	return s.GetToken(DSLParserID, 0)
}

func (s *RrTypeAndWildcardContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DSLListener); ok {
		listenerT.EnterRrTypeAndWildcard(s)
	}
}

func (s *RrTypeAndWildcardContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DSLListener); ok {
		listenerT.ExitRrTypeAndWildcard(s)
	}
}

type RrTypeContext struct {
	RelationReferenceContext
	t antlr.Token
}

func NewRrTypeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *RrTypeContext {
	var p = new(RrTypeContext)

	InitEmptyRelationReferenceContext(&p.RelationReferenceContext)
	p.parser = parser
	p.CopyAll(ctx.(*RelationReferenceContext))

	return p
}

func (s *RrTypeContext) GetT() antlr.Token { return s.t }

func (s *RrTypeContext) SetT(v antlr.Token) { s.t = v }

func (s *RrTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RrTypeContext) ID() antlr.TerminalNode {
	return s.GetToken(DSLParserID, 0)
}

func (s *RrTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DSLListener); ok {
		listenerT.EnterRrType(s)
	}
}

func (s *RrTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DSLListener); ok {
		listenerT.ExitRrType(s)
	}
}

func (p *DSLParser) RelationReference() (localctx IRelationReferenceContext) {
	localctx = NewRelationReferenceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, DSLParserRULE_relationReference)
	p.SetState(58)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 5, p.GetParserRuleContext()) {
	case 1:
		localctx = NewRrTypeContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(52)

			var _m = p.Match(DSLParserID)

			localctx.(*RrTypeContext).t = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		localctx = NewRrTypeAndRelationContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(53)

			var _m = p.Match(DSLParserID)

			localctx.(*RrTypeAndRelationContext).t = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(54)
			p.Match(DSLParserT__8)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(55)

			var _m = p.Match(DSLParserID)

			localctx.(*RrTypeAndRelationContext).r = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		localctx = NewRrTypeAndWildcardContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(56)

			var _m = p.Match(DSLParserID)

			localctx.(*RrTypeAndWildcardContext).t = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(57)
			p.Match(DSLParserT__9)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
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
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRewriteContext() *RewriteContext {
	var p = new(RewriteContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DSLParserRULE_rewrite
	return p
}

func InitEmptyRewriteContext(p *RewriteContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DSLParserRULE_rewrite
}

func (*RewriteContext) IsRewriteContext() {}

func NewRewriteContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RewriteContext {
	var p = new(RewriteContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DSLParserRULE_rewrite

	return p
}

func (s *RewriteContext) GetParser() antlr.Parser { return s.parser }

func (s *RewriteContext) CopyAll(ctx *RewriteContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *RewriteContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RewriteContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ComputedUsersetContext struct {
	RewriteContext
	computedUserset antlr.Token
}

func NewComputedUsersetContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ComputedUsersetContext {
	var p = new(ComputedUsersetContext)

	InitEmptyRewriteContext(&p.RewriteContext)
	p.parser = parser
	p.CopyAll(ctx.(*RewriteContext))

	return p
}

func (s *ComputedUsersetContext) GetComputedUserset() antlr.Token { return s.computedUserset }

func (s *ComputedUsersetContext) SetComputedUserset(v antlr.Token) { s.computedUserset = v }

func (s *ComputedUsersetContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ComputedUsersetContext) ID() antlr.TerminalNode {
	return s.GetToken(DSLParserID, 0)
}

func (s *ComputedUsersetContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DSLListener); ok {
		listenerT.EnterComputedUserset(s)
	}
}

func (s *ComputedUsersetContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DSLListener); ok {
		listenerT.ExitComputedUserset(s)
	}
}

type IntersectionContext struct {
	RewriteContext
}

func NewIntersectionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IntersectionContext {
	var p = new(IntersectionContext)

	InitEmptyRewriteContext(&p.RewriteContext)
	p.parser = parser
	p.CopyAll(ctx.(*RewriteContext))

	return p
}

func (s *IntersectionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IntersectionContext) AllRewrite() []IRewriteContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IRewriteContext); ok {
			len++
		}
	}

	tst := make([]IRewriteContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IRewriteContext); ok {
			tst[i] = t.(IRewriteContext)
			i++
		}
	}

	return tst
}

func (s *IntersectionContext) Rewrite(i int) IRewriteContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRewriteContext); ok {
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

	return t.(IRewriteContext)
}

func (s *IntersectionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DSLListener); ok {
		listenerT.EnterIntersection(s)
	}
}

func (s *IntersectionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DSLListener); ok {
		listenerT.ExitIntersection(s)
	}
}

type ThisContext struct {
	RewriteContext
}

func NewThisContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ThisContext {
	var p = new(ThisContext)

	InitEmptyRewriteContext(&p.RewriteContext)
	p.parser = parser
	p.CopyAll(ctx.(*RewriteContext))

	return p
}

func (s *ThisContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ThisContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DSLListener); ok {
		listenerT.EnterThis(s)
	}
}

func (s *ThisContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DSLListener); ok {
		listenerT.ExitThis(s)
	}
}

type ExclusionContext struct {
	RewriteContext
}

func NewExclusionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExclusionContext {
	var p = new(ExclusionContext)

	InitEmptyRewriteContext(&p.RewriteContext)
	p.parser = parser
	p.CopyAll(ctx.(*RewriteContext))

	return p
}

func (s *ExclusionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExclusionContext) AllRewrite() []IRewriteContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IRewriteContext); ok {
			len++
		}
	}

	tst := make([]IRewriteContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IRewriteContext); ok {
			tst[i] = t.(IRewriteContext)
			i++
		}
	}

	return tst
}

func (s *ExclusionContext) Rewrite(i int) IRewriteContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRewriteContext); ok {
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

	return t.(IRewriteContext)
}

func (s *ExclusionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DSLListener); ok {
		listenerT.EnterExclusion(s)
	}
}

func (s *ExclusionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DSLListener); ok {
		listenerT.ExitExclusion(s)
	}
}

type UnionContext struct {
	RewriteContext
}

func NewUnionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *UnionContext {
	var p = new(UnionContext)

	InitEmptyRewriteContext(&p.RewriteContext)
	p.parser = parser
	p.CopyAll(ctx.(*RewriteContext))

	return p
}

func (s *UnionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnionContext) AllRewrite() []IRewriteContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IRewriteContext); ok {
			len++
		}
	}

	tst := make([]IRewriteContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IRewriteContext); ok {
			tst[i] = t.(IRewriteContext)
			i++
		}
	}

	return tst
}

func (s *UnionContext) Rewrite(i int) IRewriteContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRewriteContext); ok {
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

	return t.(IRewriteContext)
}

func (s *UnionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DSLListener); ok {
		listenerT.EnterUnion(s)
	}
}

func (s *UnionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DSLListener); ok {
		listenerT.ExitUnion(s)
	}
}

type TupleToUsersetContext struct {
	RewriteContext
	computedUserset antlr.Token
	tupleset        antlr.Token
}

func NewTupleToUsersetContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TupleToUsersetContext {
	var p = new(TupleToUsersetContext)

	InitEmptyRewriteContext(&p.RewriteContext)
	p.parser = parser
	p.CopyAll(ctx.(*RewriteContext))

	return p
}

func (s *TupleToUsersetContext) GetComputedUserset() antlr.Token { return s.computedUserset }

func (s *TupleToUsersetContext) GetTupleset() antlr.Token { return s.tupleset }

func (s *TupleToUsersetContext) SetComputedUserset(v antlr.Token) { s.computedUserset = v }

func (s *TupleToUsersetContext) SetTupleset(v antlr.Token) { s.tupleset = v }

func (s *TupleToUsersetContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TupleToUsersetContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(DSLParserID)
}

func (s *TupleToUsersetContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(DSLParserID, i)
}

func (s *TupleToUsersetContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DSLListener); ok {
		listenerT.EnterTupleToUserset(s)
	}
}

func (s *TupleToUsersetContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DSLListener); ok {
		listenerT.ExitTupleToUserset(s)
	}
}

type GroupingContext struct {
	RewriteContext
}

func NewGroupingContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *GroupingContext {
	var p = new(GroupingContext)

	InitEmptyRewriteContext(&p.RewriteContext)
	p.parser = parser
	p.CopyAll(ctx.(*RewriteContext))

	return p
}

func (s *GroupingContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GroupingContext) Rewrite() IRewriteContext {
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

func (s *GroupingContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DSLListener); ok {
		listenerT.EnterGrouping(s)
	}
}

func (s *GroupingContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DSLListener); ok {
		listenerT.ExitGrouping(s)
	}
}

func (p *DSLParser) Rewrite() (localctx IRewriteContext) {
	return p.rewrite(0)
}

func (p *DSLParser) rewrite(_p int) (localctx IRewriteContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewRewriteContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IRewriteContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 12
	p.EnterRecursionRule(localctx, 12, DSLParserRULE_rewrite, _p)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(70)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 6, p.GetParserRuleContext()) {
	case 1:
		localctx = NewThisContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(61)
			p.Match(DSLParserT__10)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		localctx = NewTupleToUsersetContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(62)

			var _m = p.Match(DSLParserID)

			localctx.(*TupleToUsersetContext).computedUserset = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(63)
			p.Match(DSLParserT__11)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(64)

			var _m = p.Match(DSLParserID)

			localctx.(*TupleToUsersetContext).tupleset = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		localctx = NewComputedUsersetContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(65)

			var _m = p.Match(DSLParserID)

			localctx.(*ComputedUsersetContext).computedUserset = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 4:
		localctx = NewGroupingContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(66)
			p.Match(DSLParserT__15)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(67)
			p.rewrite(0)
		}
		{
			p.SetState(68)
			p.Match(DSLParserT__16)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(83)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 8, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(81)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 7, p.GetParserRuleContext()) {
			case 1:
				localctx = NewUnionContext(p, NewRewriteContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, DSLParserRULE_rewrite)
				p.SetState(72)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
					goto errorExit
				}
				{
					p.SetState(73)
					p.Match(DSLParserT__12)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(74)
					p.rewrite(6)
				}

			case 2:
				localctx = NewIntersectionContext(p, NewRewriteContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, DSLParserRULE_rewrite)
				p.SetState(75)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
					goto errorExit
				}
				{
					p.SetState(76)
					p.Match(DSLParserT__13)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(77)
					p.rewrite(5)
				}

			case 3:
				localctx = NewExclusionContext(p, NewRewriteContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, DSLParserRULE_rewrite)
				p.SetState(78)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
					goto errorExit
				}
				{
					p.SetState(79)
					p.Match(DSLParserT__14)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(80)
					p.rewrite(4)
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(85)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 8, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

func (p *DSLParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 6:
		var t *RewriteContext = nil
		if localctx != nil {
			t = localctx.(*RewriteContext)
		}
		return p.Rewrite_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *DSLParser) Rewrite_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 5)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 4)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 3)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
