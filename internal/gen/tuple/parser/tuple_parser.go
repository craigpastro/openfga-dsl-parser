// Code generated from java-escape by ANTLR 4.11.1. DO NOT EDIT.

package parser // Tuple

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

type TupleParser struct {
	*antlr.BaseParser
}

var tupleParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func tupleParserInit() {
	staticData := &tupleParserStaticData
	staticData.literalNames = []string{
		"", "':'", "'#'", "'@'",
	}
	staticData.symbolicNames = []string{
		"", "", "", "", "ID", "WS",
	}
	staticData.ruleNames = []string{
		"tuple", "user",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 5, 25, 2, 0, 7, 0, 2, 1, 7, 1, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0,
		1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		3, 1, 23, 8, 1, 1, 1, 0, 0, 2, 0, 2, 0, 0, 24, 0, 4, 1, 0, 0, 0, 2, 22,
		1, 0, 0, 0, 4, 5, 5, 4, 0, 0, 5, 6, 5, 1, 0, 0, 6, 7, 5, 4, 0, 0, 7, 8,
		5, 2, 0, 0, 8, 9, 5, 4, 0, 0, 9, 10, 5, 3, 0, 0, 10, 11, 3, 2, 1, 0, 11,
		12, 5, 0, 0, 1, 12, 1, 1, 0, 0, 0, 13, 14, 5, 4, 0, 0, 14, 15, 5, 1, 0,
		0, 15, 16, 5, 4, 0, 0, 16, 17, 5, 2, 0, 0, 17, 23, 5, 4, 0, 0, 18, 19,
		5, 4, 0, 0, 19, 20, 5, 1, 0, 0, 20, 23, 5, 4, 0, 0, 21, 23, 5, 4, 0, 0,
		22, 13, 1, 0, 0, 0, 22, 18, 1, 0, 0, 0, 22, 21, 1, 0, 0, 0, 23, 3, 1, 0,
		0, 0, 1, 22,
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

// TupleParserInit initializes any static state used to implement TupleParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewTupleParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func TupleParserInit() {
	staticData := &tupleParserStaticData
	staticData.once.Do(tupleParserInit)
}

// NewTupleParser produces a new parser instance for the optional input antlr.TokenStream.
func NewTupleParser(input antlr.TokenStream) *TupleParser {
	TupleParserInit()
	this := new(TupleParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &tupleParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "java-escape"

	return this
}

// TupleParser tokens.
const (
	TupleParserEOF  = antlr.TokenEOF
	TupleParserT__0 = 1
	TupleParserT__1 = 2
	TupleParserT__2 = 3
	TupleParserID   = 4
	TupleParserWS   = 5
)

// TupleParser rules.
const (
	TupleParserRULE_tuple = 0
	TupleParserRULE_user  = 1
)

// ITupleContext is an interface to support dynamic dispatch.
type ITupleContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetNamespace returns the namespace token.
	GetNamespace() antlr.Token

	// GetObjectId returns the objectId token.
	GetObjectId() antlr.Token

	// GetRelation returns the relation token.
	GetRelation() antlr.Token

	// SetNamespace sets the namespace token.
	SetNamespace(antlr.Token)

	// SetObjectId sets the objectId token.
	SetObjectId(antlr.Token)

	// SetRelation sets the relation token.
	SetRelation(antlr.Token)

	// IsTupleContext differentiates from other interfaces.
	IsTupleContext()
}

type TupleContext struct {
	*antlr.BaseParserRuleContext
	parser    antlr.Parser
	namespace antlr.Token
	objectId  antlr.Token
	relation  antlr.Token
}

func NewEmptyTupleContext() *TupleContext {
	var p = new(TupleContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TupleParserRULE_tuple
	return p
}

func (*TupleContext) IsTupleContext() {}

func NewTupleContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TupleContext {
	var p = new(TupleContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TupleParserRULE_tuple

	return p
}

func (s *TupleContext) GetParser() antlr.Parser { return s.parser }

func (s *TupleContext) GetNamespace() antlr.Token { return s.namespace }

func (s *TupleContext) GetObjectId() antlr.Token { return s.objectId }

func (s *TupleContext) GetRelation() antlr.Token { return s.relation }

func (s *TupleContext) SetNamespace(v antlr.Token) { s.namespace = v }

func (s *TupleContext) SetObjectId(v antlr.Token) { s.objectId = v }

func (s *TupleContext) SetRelation(v antlr.Token) { s.relation = v }

func (s *TupleContext) User() IUserContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUserContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUserContext)
}

func (s *TupleContext) EOF() antlr.TerminalNode {
	return s.GetToken(TupleParserEOF, 0)
}

func (s *TupleContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(TupleParserID)
}

func (s *TupleContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(TupleParserID, i)
}

func (s *TupleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TupleContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TupleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TupleListener); ok {
		listenerT.EnterTuple(s)
	}
}

func (s *TupleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TupleListener); ok {
		listenerT.ExitTuple(s)
	}
}

func (p *TupleParser) Tuple() (localctx ITupleContext) {
	this := p
	_ = this

	localctx = NewTupleContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, TupleParserRULE_tuple)

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
		p.SetState(4)

		var _m = p.Match(TupleParserID)

		localctx.(*TupleContext).namespace = _m
	}
	{
		p.SetState(5)
		p.Match(TupleParserT__0)
	}
	{
		p.SetState(6)

		var _m = p.Match(TupleParserID)

		localctx.(*TupleContext).objectId = _m
	}
	{
		p.SetState(7)
		p.Match(TupleParserT__1)
	}
	{
		p.SetState(8)

		var _m = p.Match(TupleParserID)

		localctx.(*TupleContext).relation = _m
	}
	{
		p.SetState(9)
		p.Match(TupleParserT__2)
	}
	{
		p.SetState(10)
		p.User()
	}
	{
		p.SetState(11)
		p.Match(TupleParserEOF)
	}

	return localctx
}

// IUserContext is an interface to support dynamic dispatch.
type IUserContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsUserContext differentiates from other interfaces.
	IsUserContext()
}

type UserContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUserContext() *UserContext {
	var p = new(UserContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TupleParserRULE_user
	return p
}

func (*UserContext) IsUserContext() {}

func NewUserContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UserContext {
	var p = new(UserContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TupleParserRULE_user

	return p
}

func (s *UserContext) GetParser() antlr.Parser { return s.parser }

func (s *UserContext) CopyFrom(ctx *UserContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *UserContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UserContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type UserUsersetContext struct {
	*UserContext
	namespace antlr.Token
	objectId  antlr.Token
	relation  antlr.Token
}

func NewUserUsersetContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *UserUsersetContext {
	var p = new(UserUsersetContext)

	p.UserContext = NewEmptyUserContext()
	p.parser = parser
	p.CopyFrom(ctx.(*UserContext))

	return p
}

func (s *UserUsersetContext) GetNamespace() antlr.Token { return s.namespace }

func (s *UserUsersetContext) GetObjectId() antlr.Token { return s.objectId }

func (s *UserUsersetContext) GetRelation() antlr.Token { return s.relation }

func (s *UserUsersetContext) SetNamespace(v antlr.Token) { s.namespace = v }

func (s *UserUsersetContext) SetObjectId(v antlr.Token) { s.objectId = v }

func (s *UserUsersetContext) SetRelation(v antlr.Token) { s.relation = v }

func (s *UserUsersetContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UserUsersetContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(TupleParserID)
}

func (s *UserUsersetContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(TupleParserID, i)
}

func (s *UserUsersetContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TupleListener); ok {
		listenerT.EnterUserUserset(s)
	}
}

func (s *UserUsersetContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TupleListener); ok {
		listenerT.ExitUserUserset(s)
	}
}

type UserObjectContext struct {
	*UserContext
	namespace antlr.Token
	objectId  antlr.Token
}

func NewUserObjectContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *UserObjectContext {
	var p = new(UserObjectContext)

	p.UserContext = NewEmptyUserContext()
	p.parser = parser
	p.CopyFrom(ctx.(*UserContext))

	return p
}

func (s *UserObjectContext) GetNamespace() antlr.Token { return s.namespace }

func (s *UserObjectContext) GetObjectId() antlr.Token { return s.objectId }

func (s *UserObjectContext) SetNamespace(v antlr.Token) { s.namespace = v }

func (s *UserObjectContext) SetObjectId(v antlr.Token) { s.objectId = v }

func (s *UserObjectContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UserObjectContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(TupleParserID)
}

func (s *UserObjectContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(TupleParserID, i)
}

func (s *UserObjectContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TupleListener); ok {
		listenerT.EnterUserObject(s)
	}
}

func (s *UserObjectContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TupleListener); ok {
		listenerT.ExitUserObject(s)
	}
}

type UserIdContext struct {
	*UserContext
	userId antlr.Token
}

func NewUserIdContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *UserIdContext {
	var p = new(UserIdContext)

	p.UserContext = NewEmptyUserContext()
	p.parser = parser
	p.CopyFrom(ctx.(*UserContext))

	return p
}

func (s *UserIdContext) GetUserId() antlr.Token { return s.userId }

func (s *UserIdContext) SetUserId(v antlr.Token) { s.userId = v }

func (s *UserIdContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UserIdContext) ID() antlr.TerminalNode {
	return s.GetToken(TupleParserID, 0)
}

func (s *UserIdContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TupleListener); ok {
		listenerT.EnterUserId(s)
	}
}

func (s *UserIdContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TupleListener); ok {
		listenerT.ExitUserId(s)
	}
}

func (p *TupleParser) User() (localctx IUserContext) {
	this := p
	_ = this

	localctx = NewUserContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, TupleParserRULE_user)

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

	p.SetState(22)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext()) {
	case 1:
		localctx = NewUserUsersetContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(13)

			var _m = p.Match(TupleParserID)

			localctx.(*UserUsersetContext).namespace = _m
		}
		{
			p.SetState(14)
			p.Match(TupleParserT__0)
		}
		{
			p.SetState(15)

			var _m = p.Match(TupleParserID)

			localctx.(*UserUsersetContext).objectId = _m
		}
		{
			p.SetState(16)
			p.Match(TupleParserT__1)
		}
		{
			p.SetState(17)

			var _m = p.Match(TupleParserID)

			localctx.(*UserUsersetContext).relation = _m
		}

	case 2:
		localctx = NewUserObjectContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(18)

			var _m = p.Match(TupleParserID)

			localctx.(*UserObjectContext).namespace = _m
		}
		{
			p.SetState(19)
			p.Match(TupleParserT__0)
		}
		{
			p.SetState(20)

			var _m = p.Match(TupleParserID)

			localctx.(*UserObjectContext).objectId = _m
		}

	case 3:
		localctx = NewUserIdContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(21)

			var _m = p.Match(TupleParserID)

			localctx.(*UserIdContext).userId = _m
		}

	}

	return localctx
}
