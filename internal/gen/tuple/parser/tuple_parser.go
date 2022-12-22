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
		"", "'#'", "'@'", "':'",
	}
	staticData.symbolicNames = []string{
		"", "", "", "", "ID", "WS",
	}
	staticData.ruleNames = []string{
		"tuple", "object", "user",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 5, 26, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 1, 0, 1, 0, 1, 0, 1, 0,
		1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2,
		1, 2, 3, 2, 24, 8, 2, 1, 2, 0, 0, 3, 0, 2, 4, 0, 0, 24, 0, 6, 1, 0, 0,
		0, 2, 13, 1, 0, 0, 0, 4, 23, 1, 0, 0, 0, 6, 7, 3, 2, 1, 0, 7, 8, 5, 1,
		0, 0, 8, 9, 5, 4, 0, 0, 9, 10, 5, 2, 0, 0, 10, 11, 3, 4, 2, 0, 11, 12,
		5, 0, 0, 1, 12, 1, 1, 0, 0, 0, 13, 14, 5, 4, 0, 0, 14, 15, 5, 3, 0, 0,
		15, 16, 5, 4, 0, 0, 16, 3, 1, 0, 0, 0, 17, 18, 3, 2, 1, 0, 18, 19, 5, 1,
		0, 0, 19, 20, 5, 4, 0, 0, 20, 24, 1, 0, 0, 0, 21, 24, 3, 2, 1, 0, 22, 24,
		5, 4, 0, 0, 23, 17, 1, 0, 0, 0, 23, 21, 1, 0, 0, 0, 23, 22, 1, 0, 0, 0,
		24, 5, 1, 0, 0, 0, 1, 23,
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
	TupleParserRULE_tuple  = 0
	TupleParserRULE_object = 1
	TupleParserRULE_user   = 2
)

// ITupleContext is an interface to support dynamic dispatch.
type ITupleContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetRelation returns the relation token.
	GetRelation() antlr.Token

	// SetRelation sets the relation token.
	SetRelation(antlr.Token)

	// GetObj returns the obj rule contexts.
	GetObj() IObjectContext

	// SetObj sets the obj rule contexts.
	SetObj(IObjectContext)

	// IsTupleContext differentiates from other interfaces.
	IsTupleContext()
}

type TupleContext struct {
	*antlr.BaseParserRuleContext
	parser   antlr.Parser
	obj      IObjectContext
	relation antlr.Token
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

func (s *TupleContext) GetRelation() antlr.Token { return s.relation }

func (s *TupleContext) SetRelation(v antlr.Token) { s.relation = v }

func (s *TupleContext) GetObj() IObjectContext { return s.obj }

func (s *TupleContext) SetObj(v IObjectContext) { s.obj = v }

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

func (s *TupleContext) Object() IObjectContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IObjectContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IObjectContext)
}

func (s *TupleContext) ID() antlr.TerminalNode {
	return s.GetToken(TupleParserID, 0)
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
		p.SetState(6)

		var _x = p.Object()

		localctx.(*TupleContext).obj = _x
	}
	{
		p.SetState(7)
		p.Match(TupleParserT__0)
	}
	{
		p.SetState(8)

		var _m = p.Match(TupleParserID)

		localctx.(*TupleContext).relation = _m
	}
	{
		p.SetState(9)
		p.Match(TupleParserT__1)
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

// IObjectContext is an interface to support dynamic dispatch.
type IObjectContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetNamespace returns the namespace token.
	GetNamespace() antlr.Token

	// GetObjectID returns the objectID token.
	GetObjectID() antlr.Token

	// SetNamespace sets the namespace token.
	SetNamespace(antlr.Token)

	// SetObjectID sets the objectID token.
	SetObjectID(antlr.Token)

	// IsObjectContext differentiates from other interfaces.
	IsObjectContext()
}

type ObjectContext struct {
	*antlr.BaseParserRuleContext
	parser    antlr.Parser
	namespace antlr.Token
	objectID  antlr.Token
}

func NewEmptyObjectContext() *ObjectContext {
	var p = new(ObjectContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TupleParserRULE_object
	return p
}

func (*ObjectContext) IsObjectContext() {}

func NewObjectContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ObjectContext {
	var p = new(ObjectContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TupleParserRULE_object

	return p
}

func (s *ObjectContext) GetParser() antlr.Parser { return s.parser }

func (s *ObjectContext) GetNamespace() antlr.Token { return s.namespace }

func (s *ObjectContext) GetObjectID() antlr.Token { return s.objectID }

func (s *ObjectContext) SetNamespace(v antlr.Token) { s.namespace = v }

func (s *ObjectContext) SetObjectID(v antlr.Token) { s.objectID = v }

func (s *ObjectContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(TupleParserID)
}

func (s *ObjectContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(TupleParserID, i)
}

func (s *ObjectContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ObjectContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ObjectContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TupleListener); ok {
		listenerT.EnterObject(s)
	}
}

func (s *ObjectContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TupleListener); ok {
		listenerT.ExitObject(s)
	}
}

func (p *TupleParser) Object() (localctx IObjectContext) {
	this := p
	_ = this

	localctx = NewObjectContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, TupleParserRULE_object)

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
		p.SetState(13)

		var _m = p.Match(TupleParserID)

		localctx.(*ObjectContext).namespace = _m
	}
	{
		p.SetState(14)
		p.Match(TupleParserT__2)
	}
	{
		p.SetState(15)

		var _m = p.Match(TupleParserID)

		localctx.(*ObjectContext).objectID = _m
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
	obj      IObjectContext
	relation antlr.Token
}

func NewUserUsersetContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *UserUsersetContext {
	var p = new(UserUsersetContext)

	p.UserContext = NewEmptyUserContext()
	p.parser = parser
	p.CopyFrom(ctx.(*UserContext))

	return p
}

func (s *UserUsersetContext) GetRelation() antlr.Token { return s.relation }

func (s *UserUsersetContext) SetRelation(v antlr.Token) { s.relation = v }

func (s *UserUsersetContext) GetObj() IObjectContext { return s.obj }

func (s *UserUsersetContext) SetObj(v IObjectContext) { s.obj = v }

func (s *UserUsersetContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UserUsersetContext) Object() IObjectContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IObjectContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IObjectContext)
}

func (s *UserUsersetContext) ID() antlr.TerminalNode {
	return s.GetToken(TupleParserID, 0)
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
	obj IObjectContext
}

func NewUserObjectContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *UserObjectContext {
	var p = new(UserObjectContext)

	p.UserContext = NewEmptyUserContext()
	p.parser = parser
	p.CopyFrom(ctx.(*UserContext))

	return p
}

func (s *UserObjectContext) GetObj() IObjectContext { return s.obj }

func (s *UserObjectContext) SetObj(v IObjectContext) { s.obj = v }

func (s *UserObjectContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UserObjectContext) Object() IObjectContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IObjectContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IObjectContext)
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

type UserIDContext struct {
	*UserContext
}

func NewUserIDContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *UserIDContext {
	var p = new(UserIDContext)

	p.UserContext = NewEmptyUserContext()
	p.parser = parser
	p.CopyFrom(ctx.(*UserContext))

	return p
}

func (s *UserIDContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UserIDContext) ID() antlr.TerminalNode {
	return s.GetToken(TupleParserID, 0)
}

func (s *UserIDContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TupleListener); ok {
		listenerT.EnterUserID(s)
	}
}

func (s *UserIDContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TupleListener); ok {
		listenerT.ExitUserID(s)
	}
}

func (p *TupleParser) User() (localctx IUserContext) {
	this := p
	_ = this

	localctx = NewUserContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, TupleParserRULE_user)

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

	p.SetState(23)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext()) {
	case 1:
		localctx = NewUserUsersetContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(17)

			var _x = p.Object()

			localctx.(*UserUsersetContext).obj = _x
		}
		{
			p.SetState(18)
			p.Match(TupleParserT__0)
		}
		{
			p.SetState(19)

			var _m = p.Match(TupleParserID)

			localctx.(*UserUsersetContext).relation = _m
		}

	case 2:
		localctx = NewUserObjectContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(21)

			var _x = p.Object()

			localctx.(*UserObjectContext).obj = _x
		}

	case 3:
		localctx = NewUserIDContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(22)
			p.Match(TupleParserID)
		}

	}

	return localctx
}
