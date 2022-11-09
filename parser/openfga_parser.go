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
		"'#'", "'from'", "'or'", "'and'", "'but not'",
	}
	staticData.symbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "ID", "WS",
	}
	staticData.ruleNames = []string{
		"start", "typedef", "relation", "typesOrID", "relationReferences", "relationReference",
		"rewrite",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 14, 78, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 1, 0, 4, 0, 16, 8, 0, 11, 0, 12, 0, 17, 1, 0,
		1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 4, 1, 26, 8, 1, 11, 1, 12, 1, 27, 3, 1, 30,
		8, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 3, 2, 37, 8, 2, 1, 3, 1, 3, 1, 3, 1,
		3, 1, 3, 3, 3, 44, 8, 3, 1, 4, 1, 4, 1, 4, 5, 4, 49, 8, 4, 10, 4, 12, 4,
		52, 9, 4, 1, 5, 1, 5, 1, 5, 1, 5, 3, 5, 58, 8, 5, 1, 6, 1, 6, 1, 6, 1,
		6, 4, 6, 64, 8, 6, 11, 6, 12, 6, 65, 1, 6, 1, 6, 4, 6, 70, 8, 6, 11, 6,
		12, 6, 71, 1, 6, 1, 6, 3, 6, 76, 8, 6, 1, 6, 0, 0, 7, 0, 2, 4, 6, 8, 10,
		12, 0, 0, 82, 0, 15, 1, 0, 0, 0, 2, 21, 1, 0, 0, 0, 4, 31, 1, 0, 0, 0,
		6, 43, 1, 0, 0, 0, 8, 45, 1, 0, 0, 0, 10, 57, 1, 0, 0, 0, 12, 75, 1, 0,
		0, 0, 14, 16, 3, 2, 1, 0, 15, 14, 1, 0, 0, 0, 16, 17, 1, 0, 0, 0, 17, 15,
		1, 0, 0, 0, 17, 18, 1, 0, 0, 0, 18, 19, 1, 0, 0, 0, 19, 20, 5, 0, 0, 1,
		20, 1, 1, 0, 0, 0, 21, 22, 5, 1, 0, 0, 22, 29, 5, 13, 0, 0, 23, 25, 5,
		2, 0, 0, 24, 26, 3, 4, 2, 0, 25, 24, 1, 0, 0, 0, 26, 27, 1, 0, 0, 0, 27,
		25, 1, 0, 0, 0, 27, 28, 1, 0, 0, 0, 28, 30, 1, 0, 0, 0, 29, 23, 1, 0, 0,
		0, 29, 30, 1, 0, 0, 0, 30, 3, 1, 0, 0, 0, 31, 32, 5, 3, 0, 0, 32, 33, 5,
		13, 0, 0, 33, 34, 5, 4, 0, 0, 34, 36, 3, 6, 3, 0, 35, 37, 3, 12, 6, 0,
		36, 35, 1, 0, 0, 0, 36, 37, 1, 0, 0, 0, 37, 5, 1, 0, 0, 0, 38, 39, 5, 5,
		0, 0, 39, 40, 3, 8, 4, 0, 40, 41, 5, 6, 0, 0, 41, 44, 1, 0, 0, 0, 42, 44,
		5, 13, 0, 0, 43, 38, 1, 0, 0, 0, 43, 42, 1, 0, 0, 0, 44, 7, 1, 0, 0, 0,
		45, 50, 3, 10, 5, 0, 46, 47, 5, 7, 0, 0, 47, 49, 3, 8, 4, 0, 48, 46, 1,
		0, 0, 0, 49, 52, 1, 0, 0, 0, 50, 48, 1, 0, 0, 0, 50, 51, 1, 0, 0, 0, 51,
		9, 1, 0, 0, 0, 52, 50, 1, 0, 0, 0, 53, 58, 5, 13, 0, 0, 54, 55, 5, 13,
		0, 0, 55, 56, 5, 8, 0, 0, 56, 58, 5, 13, 0, 0, 57, 53, 1, 0, 0, 0, 57,
		54, 1, 0, 0, 0, 58, 11, 1, 0, 0, 0, 59, 60, 5, 9, 0, 0, 60, 76, 5, 13,
		0, 0, 61, 62, 5, 10, 0, 0, 62, 64, 5, 13, 0, 0, 63, 61, 1, 0, 0, 0, 64,
		65, 1, 0, 0, 0, 65, 63, 1, 0, 0, 0, 65, 66, 1, 0, 0, 0, 66, 76, 1, 0, 0,
		0, 67, 68, 5, 11, 0, 0, 68, 70, 5, 13, 0, 0, 69, 67, 1, 0, 0, 0, 70, 71,
		1, 0, 0, 0, 71, 69, 1, 0, 0, 0, 71, 72, 1, 0, 0, 0, 72, 76, 1, 0, 0, 0,
		73, 74, 5, 12, 0, 0, 74, 76, 5, 13, 0, 0, 75, 59, 1, 0, 0, 0, 75, 63, 1,
		0, 0, 0, 75, 69, 1, 0, 0, 0, 75, 73, 1, 0, 0, 0, 76, 13, 1, 0, 0, 0, 10,
		17, 27, 29, 36, 43, 50, 57, 65, 71, 75,
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
	OpenFGAParserID    = 13
	OpenFGAParserWS    = 14
)

// OpenFGAParser rules.
const (
	OpenFGAParserRULE_start              = 0
	OpenFGAParserRULE_typedef            = 1
	OpenFGAParserRULE_relation           = 2
	OpenFGAParserRULE_typesOrID          = 3
	OpenFGAParserRULE_relationReferences = 4
	OpenFGAParserRULE_relationReference  = 5
	OpenFGAParserRULE_rewrite            = 6
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

func (s *StartContext) AllTypedef() []ITypedefContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITypedefContext); ok {
			len++
		}
	}

	tst := make([]ITypedefContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITypedefContext); ok {
			tst[i] = t.(ITypedefContext)
			i++
		}
	}

	return tst
}

func (s *StartContext) Typedef(i int) ITypedefContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypedefContext); ok {
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

	return t.(ITypedefContext)
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
	p.SetState(15)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == OpenFGAParserT__0 {
		{
			p.SetState(14)
			p.Typedef()
		}

		p.SetState(17)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(19)
		p.Match(OpenFGAParserEOF)
	}

	return localctx
}

// ITypedefContext is an interface to support dynamic dispatch.
type ITypedefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetObjectType returns the objectType token.
	GetObjectType() antlr.Token

	// SetObjectType sets the objectType token.
	SetObjectType(antlr.Token)

	// IsTypedefContext differentiates from other interfaces.
	IsTypedefContext()
}

type TypedefContext struct {
	*antlr.BaseParserRuleContext
	parser     antlr.Parser
	objectType antlr.Token
}

func NewEmptyTypedefContext() *TypedefContext {
	var p = new(TypedefContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = OpenFGAParserRULE_typedef
	return p
}

func (*TypedefContext) IsTypedefContext() {}

func NewTypedefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypedefContext {
	var p = new(TypedefContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = OpenFGAParserRULE_typedef

	return p
}

func (s *TypedefContext) GetParser() antlr.Parser { return s.parser }

func (s *TypedefContext) GetObjectType() antlr.Token { return s.objectType }

func (s *TypedefContext) SetObjectType(v antlr.Token) { s.objectType = v }

func (s *TypedefContext) ID() antlr.TerminalNode {
	return s.GetToken(OpenFGAParserID, 0)
}

func (s *TypedefContext) AllRelation() []IRelationContext {
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

func (s *TypedefContext) Relation(i int) IRelationContext {
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

func (s *TypedefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypedefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypedefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OpenFGAListener); ok {
		listenerT.EnterTypedef(s)
	}
}

func (s *TypedefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OpenFGAListener); ok {
		listenerT.ExitTypedef(s)
	}
}

func (p *OpenFGAParser) Typedef() (localctx ITypedefContext) {
	this := p
	_ = this

	localctx = NewTypedefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, OpenFGAParserRULE_typedef)
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
		p.SetState(21)
		p.Match(OpenFGAParserT__0)
	}
	{
		p.SetState(22)

		var _m = p.Match(OpenFGAParserID)

		localctx.(*TypedefContext).objectType = _m
	}
	p.SetState(29)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == OpenFGAParserT__1 {
		{
			p.SetState(23)
			p.Match(OpenFGAParserT__1)
		}
		p.SetState(25)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == OpenFGAParserT__2 {
			{
				p.SetState(24)
				p.Relation()
			}

			p.SetState(27)
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

func (s *RelationContext) TypesOrID() ITypesOrIDContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypesOrIDContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypesOrIDContext)
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
		p.SetState(31)
		p.Match(OpenFGAParserT__2)
	}
	{
		p.SetState(32)

		var _m = p.Match(OpenFGAParserID)

		localctx.(*RelationContext).name = _m
	}
	{
		p.SetState(33)
		p.Match(OpenFGAParserT__3)
	}
	{
		p.SetState(34)
		p.TypesOrID()
	}
	p.SetState(36)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&7680) != 0 {
		{
			p.SetState(35)
			p.Rewrite()
		}

	}

	return localctx
}

// ITypesOrIDContext is an interface to support dynamic dispatch.
type ITypesOrIDContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypesOrIDContext differentiates from other interfaces.
	IsTypesOrIDContext()
}

type TypesOrIDContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypesOrIDContext() *TypesOrIDContext {
	var p = new(TypesOrIDContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = OpenFGAParserRULE_typesOrID
	return p
}

func (*TypesOrIDContext) IsTypesOrIDContext() {}

func NewTypesOrIDContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypesOrIDContext {
	var p = new(TypesOrIDContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = OpenFGAParserRULE_typesOrID

	return p
}

func (s *TypesOrIDContext) GetParser() antlr.Parser { return s.parser }

func (s *TypesOrIDContext) CopyFrom(ctx *TypesOrIDContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *TypesOrIDContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypesOrIDContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type TypesContext struct {
	*TypesOrIDContext
}

func NewTypesContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TypesContext {
	var p = new(TypesContext)

	p.TypesOrIDContext = NewEmptyTypesOrIDContext()
	p.parser = parser
	p.CopyFrom(ctx.(*TypesOrIDContext))

	return p
}

func (s *TypesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypesContext) RelationReferences() IRelationReferencesContext {
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

func (s *TypesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OpenFGAListener); ok {
		listenerT.EnterTypes(s)
	}
}

func (s *TypesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OpenFGAListener); ok {
		listenerT.ExitTypes(s)
	}
}

type ComputedUsersetContext struct {
	*TypesOrIDContext
	id antlr.Token
}

func NewComputedUsersetContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ComputedUsersetContext {
	var p = new(ComputedUsersetContext)

	p.TypesOrIDContext = NewEmptyTypesOrIDContext()
	p.parser = parser
	p.CopyFrom(ctx.(*TypesOrIDContext))

	return p
}

func (s *ComputedUsersetContext) GetId() antlr.Token { return s.id }

func (s *ComputedUsersetContext) SetId(v antlr.Token) { s.id = v }

func (s *ComputedUsersetContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ComputedUsersetContext) ID() antlr.TerminalNode {
	return s.GetToken(OpenFGAParserID, 0)
}

func (s *ComputedUsersetContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OpenFGAListener); ok {
		listenerT.EnterComputedUserset(s)
	}
}

func (s *ComputedUsersetContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OpenFGAListener); ok {
		listenerT.ExitComputedUserset(s)
	}
}

func (p *OpenFGAParser) TypesOrID() (localctx ITypesOrIDContext) {
	this := p
	_ = this

	localctx = NewTypesOrIDContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, OpenFGAParserRULE_typesOrID)

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

	p.SetState(43)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case OpenFGAParserT__4:
		localctx = NewTypesContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(38)
			p.Match(OpenFGAParserT__4)
		}
		{
			p.SetState(39)
			p.RelationReferences()
		}
		{
			p.SetState(40)
			p.Match(OpenFGAParserT__5)
		}

	case OpenFGAParserID:
		localctx = NewComputedUsersetContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(42)

			var _m = p.Match(OpenFGAParserID)

			localctx.(*ComputedUsersetContext).id = _m
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IRelationReferencesContext is an interface to support dynamic dispatch.
type IRelationReferencesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetHead returns the head rule contexts.
	GetHead() IRelationReferenceContext

	// GetTail returns the tail rule contexts.
	GetTail() IRelationReferencesContext

	// SetHead sets the head rule contexts.
	SetHead(IRelationReferenceContext)

	// SetTail sets the tail rule contexts.
	SetTail(IRelationReferencesContext)

	// IsRelationReferencesContext differentiates from other interfaces.
	IsRelationReferencesContext()
}

type RelationReferencesContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	head   IRelationReferenceContext
	tail   IRelationReferencesContext
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

func (s *RelationReferencesContext) GetHead() IRelationReferenceContext { return s.head }

func (s *RelationReferencesContext) GetTail() IRelationReferencesContext { return s.tail }

func (s *RelationReferencesContext) SetHead(v IRelationReferenceContext) { s.head = v }

func (s *RelationReferencesContext) SetTail(v IRelationReferencesContext) { s.tail = v }

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
	p.EnterRule(localctx, 8, OpenFGAParserRULE_relationReferences)

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
		p.SetState(45)

		var _x = p.RelationReference()

		localctx.(*RelationReferencesContext).head = _x
	}
	p.SetState(50)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(46)
				p.Match(OpenFGAParserT__6)
			}
			{
				p.SetState(47)

				var _x = p.RelationReferences()

				localctx.(*RelationReferencesContext).tail = _x
			}

		}
		p.SetState(52)
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

type TypeAndRelationContext struct {
	*RelationReferenceContext
	t antlr.Token
	r antlr.Token
}

func NewTypeAndRelationContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TypeAndRelationContext {
	var p = new(TypeAndRelationContext)

	p.RelationReferenceContext = NewEmptyRelationReferenceContext()
	p.parser = parser
	p.CopyFrom(ctx.(*RelationReferenceContext))

	return p
}

func (s *TypeAndRelationContext) GetT() antlr.Token { return s.t }

func (s *TypeAndRelationContext) GetR() antlr.Token { return s.r }

func (s *TypeAndRelationContext) SetT(v antlr.Token) { s.t = v }

func (s *TypeAndRelationContext) SetR(v antlr.Token) { s.r = v }

func (s *TypeAndRelationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeAndRelationContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(OpenFGAParserID)
}

func (s *TypeAndRelationContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(OpenFGAParserID, i)
}

func (s *TypeAndRelationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OpenFGAListener); ok {
		listenerT.EnterTypeAndRelation(s)
	}
}

func (s *TypeAndRelationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OpenFGAListener); ok {
		listenerT.ExitTypeAndRelation(s)
	}
}

type TypeContext struct {
	*RelationReferenceContext
	t antlr.Token
}

func NewTypeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TypeContext {
	var p = new(TypeContext)

	p.RelationReferenceContext = NewEmptyRelationReferenceContext()
	p.parser = parser
	p.CopyFrom(ctx.(*RelationReferenceContext))

	return p
}

func (s *TypeContext) GetT() antlr.Token { return s.t }

func (s *TypeContext) SetT(v antlr.Token) { s.t = v }

func (s *TypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeContext) ID() antlr.TerminalNode {
	return s.GetToken(OpenFGAParserID, 0)
}

func (s *TypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OpenFGAListener); ok {
		listenerT.EnterType(s)
	}
}

func (s *TypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OpenFGAListener); ok {
		listenerT.ExitType(s)
	}
}

func (p *OpenFGAParser) RelationReference() (localctx IRelationReferenceContext) {
	this := p
	_ = this

	localctx = NewRelationReferenceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, OpenFGAParserRULE_relationReference)

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

	p.SetState(57)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext()) {
	case 1:
		localctx = NewTypeContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(53)

			var _m = p.Match(OpenFGAParserID)

			localctx.(*TypeContext).t = _m
		}

	case 2:
		localctx = NewTypeAndRelationContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(54)

			var _m = p.Match(OpenFGAParserID)

			localctx.(*TypeAndRelationContext).t = _m
		}
		{
			p.SetState(55)
			p.Match(OpenFGAParserT__7)
		}
		{
			p.SetState(56)

			var _m = p.Match(OpenFGAParserID)

			localctx.(*TypeAndRelationContext).r = _m
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

func (s *RewriteContext) CopyFrom(ctx *RewriteContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *RewriteContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RewriteContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type IntersectionContext struct {
	*RewriteContext
	id antlr.Token
}

func NewIntersectionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IntersectionContext {
	var p = new(IntersectionContext)

	p.RewriteContext = NewEmptyRewriteContext()
	p.parser = parser
	p.CopyFrom(ctx.(*RewriteContext))

	return p
}

func (s *IntersectionContext) GetId() antlr.Token { return s.id }

func (s *IntersectionContext) SetId(v antlr.Token) { s.id = v }

func (s *IntersectionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IntersectionContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(OpenFGAParserID)
}

func (s *IntersectionContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(OpenFGAParserID, i)
}

func (s *IntersectionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OpenFGAListener); ok {
		listenerT.EnterIntersection(s)
	}
}

func (s *IntersectionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OpenFGAListener); ok {
		listenerT.ExitIntersection(s)
	}
}

type ExclusionContext struct {
	*RewriteContext
	id antlr.Token
}

func NewExclusionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExclusionContext {
	var p = new(ExclusionContext)

	p.RewriteContext = NewEmptyRewriteContext()
	p.parser = parser
	p.CopyFrom(ctx.(*RewriteContext))

	return p
}

func (s *ExclusionContext) GetId() antlr.Token { return s.id }

func (s *ExclusionContext) SetId(v antlr.Token) { s.id = v }

func (s *ExclusionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExclusionContext) ID() antlr.TerminalNode {
	return s.GetToken(OpenFGAParserID, 0)
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

type UnionContext struct {
	*RewriteContext
	id antlr.Token
}

func NewUnionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *UnionContext {
	var p = new(UnionContext)

	p.RewriteContext = NewEmptyRewriteContext()
	p.parser = parser
	p.CopyFrom(ctx.(*RewriteContext))

	return p
}

func (s *UnionContext) GetId() antlr.Token { return s.id }

func (s *UnionContext) SetId(v antlr.Token) { s.id = v }

func (s *UnionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnionContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(OpenFGAParserID)
}

func (s *UnionContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(OpenFGAParserID, i)
}

func (s *UnionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OpenFGAListener); ok {
		listenerT.EnterUnion(s)
	}
}

func (s *UnionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OpenFGAListener); ok {
		listenerT.ExitUnion(s)
	}
}

type TupleToUsersetContext struct {
	*RewriteContext
	id antlr.Token
}

func NewTupleToUsersetContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TupleToUsersetContext {
	var p = new(TupleToUsersetContext)

	p.RewriteContext = NewEmptyRewriteContext()
	p.parser = parser
	p.CopyFrom(ctx.(*RewriteContext))

	return p
}

func (s *TupleToUsersetContext) GetId() antlr.Token { return s.id }

func (s *TupleToUsersetContext) SetId(v antlr.Token) { s.id = v }

func (s *TupleToUsersetContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TupleToUsersetContext) ID() antlr.TerminalNode {
	return s.GetToken(OpenFGAParserID, 0)
}

func (s *TupleToUsersetContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OpenFGAListener); ok {
		listenerT.EnterTupleToUserset(s)
	}
}

func (s *TupleToUsersetContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OpenFGAListener); ok {
		listenerT.ExitTupleToUserset(s)
	}
}

func (p *OpenFGAParser) Rewrite() (localctx IRewriteContext) {
	this := p
	_ = this

	localctx = NewRewriteContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, OpenFGAParserRULE_rewrite)
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

	p.SetState(75)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case OpenFGAParserT__8:
		localctx = NewTupleToUsersetContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(59)
			p.Match(OpenFGAParserT__8)
		}
		{
			p.SetState(60)

			var _m = p.Match(OpenFGAParserID)

			localctx.(*TupleToUsersetContext).id = _m
		}

	case OpenFGAParserT__9:
		localctx = NewUnionContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		p.SetState(63)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == OpenFGAParserT__9 {
			{
				p.SetState(61)
				p.Match(OpenFGAParserT__9)
			}
			{
				p.SetState(62)

				var _m = p.Match(OpenFGAParserID)

				localctx.(*UnionContext).id = _m
			}

			p.SetState(65)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	case OpenFGAParserT__10:
		localctx = NewIntersectionContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		p.SetState(69)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == OpenFGAParserT__10 {
			{
				p.SetState(67)
				p.Match(OpenFGAParserT__10)
			}
			{
				p.SetState(68)

				var _m = p.Match(OpenFGAParserID)

				localctx.(*IntersectionContext).id = _m
			}

			p.SetState(71)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	case OpenFGAParserT__11:
		localctx = NewExclusionContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(73)
			p.Match(OpenFGAParserT__11)
		}
		{
			p.SetState(74)

			var _m = p.Match(OpenFGAParserID)

			localctx.(*ExclusionContext).id = _m
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}
