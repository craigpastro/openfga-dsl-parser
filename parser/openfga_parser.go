// Code generated from OpenFGA.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser // OpenFGA

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

type OpenFGAParser struct {
	*antlr.BaseParser
}

var OpenFGAParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func openfgaParserInit() {
	staticData := &OpenFGAParserStaticData
	staticData.LiteralNames = []string{
		"", "'type'", "'relations'", "'define'", "'as'", "'self'", "'from'",
		"'or'", "'and'", "'but not'", "'('", "')'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "ID", "WS",
	}
	staticData.RuleNames = []string{
		"prog", "typedef", "relations", "rewrite",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 13, 58, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 1, 0, 5,
		0, 10, 8, 0, 10, 0, 12, 0, 13, 9, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1,
		4, 1, 21, 8, 1, 11, 1, 12, 1, 22, 3, 1, 25, 8, 1, 1, 2, 1, 2, 1, 2, 1,
		2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 3,
		3, 42, 8, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 5, 3,
		53, 8, 3, 10, 3, 12, 3, 56, 9, 3, 1, 3, 0, 1, 6, 4, 0, 2, 4, 6, 0, 0, 62,
		0, 11, 1, 0, 0, 0, 2, 16, 1, 0, 0, 0, 4, 26, 1, 0, 0, 0, 6, 41, 1, 0, 0,
		0, 8, 10, 3, 2, 1, 0, 9, 8, 1, 0, 0, 0, 10, 13, 1, 0, 0, 0, 11, 9, 1, 0,
		0, 0, 11, 12, 1, 0, 0, 0, 12, 14, 1, 0, 0, 0, 13, 11, 1, 0, 0, 0, 14, 15,
		5, 0, 0, 1, 15, 1, 1, 0, 0, 0, 16, 17, 5, 1, 0, 0, 17, 24, 5, 12, 0, 0,
		18, 20, 5, 2, 0, 0, 19, 21, 3, 4, 2, 0, 20, 19, 1, 0, 0, 0, 21, 22, 1,
		0, 0, 0, 22, 20, 1, 0, 0, 0, 22, 23, 1, 0, 0, 0, 23, 25, 1, 0, 0, 0, 24,
		18, 1, 0, 0, 0, 24, 25, 1, 0, 0, 0, 25, 3, 1, 0, 0, 0, 26, 27, 5, 3, 0,
		0, 27, 28, 5, 12, 0, 0, 28, 29, 5, 4, 0, 0, 29, 30, 3, 6, 3, 0, 30, 5,
		1, 0, 0, 0, 31, 32, 6, 3, -1, 0, 32, 42, 5, 5, 0, 0, 33, 34, 5, 12, 0,
		0, 34, 35, 5, 6, 0, 0, 35, 42, 5, 12, 0, 0, 36, 42, 5, 12, 0, 0, 37, 38,
		5, 10, 0, 0, 38, 39, 3, 6, 3, 0, 39, 40, 5, 11, 0, 0, 40, 42, 1, 0, 0,
		0, 41, 31, 1, 0, 0, 0, 41, 33, 1, 0, 0, 0, 41, 36, 1, 0, 0, 0, 41, 37,
		1, 0, 0, 0, 42, 54, 1, 0, 0, 0, 43, 44, 10, 5, 0, 0, 44, 45, 5, 7, 0, 0,
		45, 53, 3, 6, 3, 6, 46, 47, 10, 4, 0, 0, 47, 48, 5, 8, 0, 0, 48, 53, 3,
		6, 3, 5, 49, 50, 10, 3, 0, 0, 50, 51, 5, 9, 0, 0, 51, 53, 3, 6, 3, 4, 52,
		43, 1, 0, 0, 0, 52, 46, 1, 0, 0, 0, 52, 49, 1, 0, 0, 0, 53, 56, 1, 0, 0,
		0, 54, 52, 1, 0, 0, 0, 54, 55, 1, 0, 0, 0, 55, 7, 1, 0, 0, 0, 56, 54, 1,
		0, 0, 0, 6, 11, 22, 24, 41, 52, 54,
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
	staticData := &OpenFGAParserStaticData
	staticData.once.Do(openfgaParserInit)
}

// NewOpenFGAParser produces a new parser instance for the optional input antlr.TokenStream.
func NewOpenFGAParser(input antlr.TokenStream) *OpenFGAParser {
	OpenFGAParserInit()
	this := new(OpenFGAParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &OpenFGAParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "OpenFGA.g4"

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
	OpenFGAParserID    = 12
	OpenFGAParserWS    = 13
)

// OpenFGAParser rules.
const (
	OpenFGAParserRULE_prog      = 0
	OpenFGAParserRULE_typedef   = 1
	OpenFGAParserRULE_relations = 2
	OpenFGAParserRULE_rewrite   = 3
)

// IProgContext is an interface to support dynamic dispatch.
type IProgContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EOF() antlr.TerminalNode
	AllTypedef() []ITypedefContext
	Typedef(i int) ITypedefContext

	// IsProgContext differentiates from other interfaces.
	IsProgContext()
}

type ProgContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgContext() *ProgContext {
	var p = new(ProgContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = OpenFGAParserRULE_prog
	return p
}

func InitEmptyProgContext(p *ProgContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = OpenFGAParserRULE_prog
}

func (*ProgContext) IsProgContext() {}

func NewProgContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgContext {
	var p = new(ProgContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = OpenFGAParserRULE_prog

	return p
}

func (s *ProgContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgContext) EOF() antlr.TerminalNode {
	return s.GetToken(OpenFGAParserEOF, 0)
}

func (s *ProgContext) AllTypedef() []ITypedefContext {
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

func (s *ProgContext) Typedef(i int) ITypedefContext {
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

func (s *ProgContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OpenFGAListener); ok {
		listenerT.EnterProg(s)
	}
}

func (s *ProgContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OpenFGAListener); ok {
		listenerT.ExitProg(s)
	}
}

func (p *OpenFGAParser) Prog() (localctx IProgContext) {
	localctx = NewProgContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, OpenFGAParserRULE_prog)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(11)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == OpenFGAParserT__0 {
		{
			p.SetState(8)
			p.Typedef()
		}

		p.SetState(13)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(14)
		p.Match(OpenFGAParserEOF)
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

// ITypedefContext is an interface to support dynamic dispatch.
type ITypedefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetObjectType returns the objectType token.
	GetObjectType() antlr.Token

	// SetObjectType sets the objectType token.
	SetObjectType(antlr.Token)

	// Getter signatures
	ID() antlr.TerminalNode
	AllRelations() []IRelationsContext
	Relations(i int) IRelationsContext

	// IsTypedefContext differentiates from other interfaces.
	IsTypedefContext()
}

type TypedefContext struct {
	antlr.BaseParserRuleContext
	parser     antlr.Parser
	objectType antlr.Token
}

func NewEmptyTypedefContext() *TypedefContext {
	var p = new(TypedefContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = OpenFGAParserRULE_typedef
	return p
}

func InitEmptyTypedefContext(p *TypedefContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = OpenFGAParserRULE_typedef
}

func (*TypedefContext) IsTypedefContext() {}

func NewTypedefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypedefContext {
	var p = new(TypedefContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

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

func (s *TypedefContext) AllRelations() []IRelationsContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IRelationsContext); ok {
			len++
		}
	}

	tst := make([]IRelationsContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IRelationsContext); ok {
			tst[i] = t.(IRelationsContext)
			i++
		}
	}

	return tst
}

func (s *TypedefContext) Relations(i int) IRelationsContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRelationsContext); ok {
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

	return t.(IRelationsContext)
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
	localctx = NewTypedefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, OpenFGAParserRULE_typedef)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(16)
		p.Match(OpenFGAParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(17)

		var _m = p.Match(OpenFGAParserID)

		localctx.(*TypedefContext).objectType = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(24)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == OpenFGAParserT__1 {
		{
			p.SetState(18)
			p.Match(OpenFGAParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(20)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == OpenFGAParserT__2 {
			{
				p.SetState(19)
				p.Relations()
			}

			p.SetState(22)
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

// IRelationsContext is an interface to support dynamic dispatch.
type IRelationsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetRelation returns the relation token.
	GetRelation() antlr.Token

	// SetRelation sets the relation token.
	SetRelation(antlr.Token)

	// Getter signatures
	Rewrite() IRewriteContext
	ID() antlr.TerminalNode

	// IsRelationsContext differentiates from other interfaces.
	IsRelationsContext()
}

type RelationsContext struct {
	antlr.BaseParserRuleContext
	parser   antlr.Parser
	relation antlr.Token
}

func NewEmptyRelationsContext() *RelationsContext {
	var p = new(RelationsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = OpenFGAParserRULE_relations
	return p
}

func InitEmptyRelationsContext(p *RelationsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = OpenFGAParserRULE_relations
}

func (*RelationsContext) IsRelationsContext() {}

func NewRelationsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RelationsContext {
	var p = new(RelationsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = OpenFGAParserRULE_relations

	return p
}

func (s *RelationsContext) GetParser() antlr.Parser { return s.parser }

func (s *RelationsContext) GetRelation() antlr.Token { return s.relation }

func (s *RelationsContext) SetRelation(v antlr.Token) { s.relation = v }

func (s *RelationsContext) Rewrite() IRewriteContext {
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

func (s *RelationsContext) ID() antlr.TerminalNode {
	return s.GetToken(OpenFGAParserID, 0)
}

func (s *RelationsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RelationsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RelationsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OpenFGAListener); ok {
		listenerT.EnterRelations(s)
	}
}

func (s *RelationsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OpenFGAListener); ok {
		listenerT.ExitRelations(s)
	}
}

func (p *OpenFGAParser) Relations() (localctx IRelationsContext) {
	localctx = NewRelationsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, OpenFGAParserRULE_relations)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(26)
		p.Match(OpenFGAParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(27)

		var _m = p.Match(OpenFGAParserID)

		localctx.(*RelationsContext).relation = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(28)
		p.Match(OpenFGAParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(29)
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
	p.RuleIndex = OpenFGAParserRULE_rewrite
	return p
}

func InitEmptyRewriteContext(p *RewriteContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = OpenFGAParserRULE_rewrite
}

func (*RewriteContext) IsRewriteContext() {}

func NewRewriteContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RewriteContext {
	var p = new(RewriteContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = OpenFGAParserRULE_rewrite

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
	if listenerT, ok := listener.(OpenFGAListener); ok {
		listenerT.EnterIntersection(s)
	}
}

func (s *IntersectionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OpenFGAListener); ok {
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
	if listenerT, ok := listener.(OpenFGAListener); ok {
		listenerT.EnterThis(s)
	}
}

func (s *ThisContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OpenFGAListener); ok {
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
	return s.GetTokens(OpenFGAParserID)
}

func (s *TupleToUsersetContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(OpenFGAParserID, i)
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
	if listenerT, ok := listener.(OpenFGAListener); ok {
		listenerT.EnterGrouping(s)
	}
}

func (s *GroupingContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(OpenFGAListener); ok {
		listenerT.ExitGrouping(s)
	}
}

func (p *OpenFGAParser) Rewrite() (localctx IRewriteContext) {
	return p.rewrite(0)
}

func (p *OpenFGAParser) rewrite(_p int) (localctx IRewriteContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewRewriteContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IRewriteContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 6
	p.EnterRecursionRule(localctx, 6, OpenFGAParserRULE_rewrite, _p)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(41)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 3, p.GetParserRuleContext()) {
	case 1:
		localctx = NewThisContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(32)
			p.Match(OpenFGAParserT__4)
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
			p.SetState(33)

			var _m = p.Match(OpenFGAParserID)

			localctx.(*TupleToUsersetContext).computedUserset = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(34)
			p.Match(OpenFGAParserT__5)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(35)

			var _m = p.Match(OpenFGAParserID)

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
			p.SetState(36)

			var _m = p.Match(OpenFGAParserID)

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
			p.SetState(37)
			p.Match(OpenFGAParserT__9)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(38)
			p.rewrite(0)
		}
		{
			p.SetState(39)
			p.Match(OpenFGAParserT__10)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(54)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 5, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(52)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext()) {
			case 1:
				localctx = NewUnionContext(p, NewRewriteContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, OpenFGAParserRULE_rewrite)
				p.SetState(43)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
					goto errorExit
				}
				{
					p.SetState(44)
					p.Match(OpenFGAParserT__6)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(45)
					p.rewrite(6)
				}

			case 2:
				localctx = NewIntersectionContext(p, NewRewriteContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, OpenFGAParserRULE_rewrite)
				p.SetState(46)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
					goto errorExit
				}
				{
					p.SetState(47)
					p.Match(OpenFGAParserT__7)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(48)
					p.rewrite(5)
				}

			case 3:
				localctx = NewExclusionContext(p, NewRewriteContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, OpenFGAParserRULE_rewrite)
				p.SetState(49)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
					goto errorExit
				}
				{
					p.SetState(50)
					p.Match(OpenFGAParserT__8)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(51)
					p.rewrite(4)
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(56)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 5, p.GetParserRuleContext())
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

func (p *OpenFGAParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 3:
		var t *RewriteContext = nil
		if localctx != nil {
			t = localctx.(*RewriteContext)
		}
		return p.Rewrite_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *OpenFGAParser) Rewrite_Sempred(localctx antlr.RuleContext, predIndex int) bool {
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
