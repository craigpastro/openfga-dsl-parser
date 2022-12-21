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
		"tuple", "object", "relation", "user",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 5, 30, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 1, 0, 1, 0,
		1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 3,
		1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 3, 3, 28, 8, 3, 1, 3, 0, 0, 4, 0, 2, 4, 6,
		0, 0, 27, 0, 8, 1, 0, 0, 0, 2, 15, 1, 0, 0, 0, 4, 19, 1, 0, 0, 0, 6, 27,
		1, 0, 0, 0, 8, 9, 3, 2, 1, 0, 9, 10, 5, 1, 0, 0, 10, 11, 3, 4, 2, 0, 11,
		12, 5, 2, 0, 0, 12, 13, 3, 6, 3, 0, 13, 14, 5, 0, 0, 1, 14, 1, 1, 0, 0,
		0, 15, 16, 5, 4, 0, 0, 16, 17, 5, 3, 0, 0, 17, 18, 5, 4, 0, 0, 18, 3, 1,
		0, 0, 0, 19, 20, 5, 4, 0, 0, 20, 5, 1, 0, 0, 0, 21, 22, 3, 2, 1, 0, 22,
		23, 5, 1, 0, 0, 23, 24, 3, 4, 2, 0, 24, 28, 1, 0, 0, 0, 25, 28, 3, 2, 1,
		0, 26, 28, 5, 4, 0, 0, 27, 21, 1, 0, 0, 0, 27, 25, 1, 0, 0, 0, 27, 26,
		1, 0, 0, 0, 28, 7, 1, 0, 0, 0, 1, 27,
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
	TupleParserRULE_tuple    = 0
	TupleParserRULE_object   = 1
	TupleParserRULE_relation = 2
	TupleParserRULE_user     = 3
)

// ITupleContext is an interface to support dynamic dispatch.
type ITupleContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTupleContext differentiates from other interfaces.
	IsTupleContext()
}

type TupleContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
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

func (s *TupleContext) Relation() IRelationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRelationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRelationContext)
}

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
		p.SetState(8)
		p.Object()
	}
	{
		p.SetState(9)
		p.Match(TupleParserT__0)
	}
	{
		p.SetState(10)
		p.Relation()
	}
	{
		p.SetState(11)
		p.Match(TupleParserT__1)
	}
	{
		p.SetState(12)
		p.User()
	}
	{
		p.SetState(13)
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

	// GetObject_id returns the object_id token.
	GetObject_id() antlr.Token

	// SetNamespace sets the namespace token.
	SetNamespace(antlr.Token)

	// SetObject_id sets the object_id token.
	SetObject_id(antlr.Token)

	// IsObjectContext differentiates from other interfaces.
	IsObjectContext()
}

type ObjectContext struct {
	*antlr.BaseParserRuleContext
	parser    antlr.Parser
	namespace antlr.Token
	object_id antlr.Token
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

func (s *ObjectContext) GetObject_id() antlr.Token { return s.object_id }

func (s *ObjectContext) SetNamespace(v antlr.Token) { s.namespace = v }

func (s *ObjectContext) SetObject_id(v antlr.Token) { s.object_id = v }

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
		p.SetState(15)

		var _m = p.Match(TupleParserID)

		localctx.(*ObjectContext).namespace = _m
	}
	{
		p.SetState(16)
		p.Match(TupleParserT__2)
	}
	{
		p.SetState(17)

		var _m = p.Match(TupleParserID)

		localctx.(*ObjectContext).object_id = _m
	}

	return localctx
}

// IRelationContext is an interface to support dynamic dispatch.
type IRelationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRelationContext differentiates from other interfaces.
	IsRelationContext()
}

type RelationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRelationContext() *RelationContext {
	var p = new(RelationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TupleParserRULE_relation
	return p
}

func (*RelationContext) IsRelationContext() {}

func NewRelationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RelationContext {
	var p = new(RelationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TupleParserRULE_relation

	return p
}

func (s *RelationContext) GetParser() antlr.Parser { return s.parser }

func (s *RelationContext) ID() antlr.TerminalNode {
	return s.GetToken(TupleParserID, 0)
}

func (s *RelationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RelationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RelationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TupleListener); ok {
		listenerT.EnterRelation(s)
	}
}

func (s *RelationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TupleListener); ok {
		listenerT.ExitRelation(s)
	}
}

func (p *TupleParser) Relation() (localctx IRelationContext) {
	this := p
	_ = this

	localctx = NewRelationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, TupleParserRULE_relation)

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
		p.SetState(19)
		p.Match(TupleParserID)
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

type User_idContext struct {
	*UserContext
}

func NewUser_idContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *User_idContext {
	var p = new(User_idContext)

	p.UserContext = NewEmptyUserContext()
	p.parser = parser
	p.CopyFrom(ctx.(*UserContext))

	return p
}

func (s *User_idContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *User_idContext) ID() antlr.TerminalNode {
	return s.GetToken(TupleParserID, 0)
}

func (s *User_idContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TupleListener); ok {
		listenerT.EnterUser_id(s)
	}
}

func (s *User_idContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TupleListener); ok {
		listenerT.ExitUser_id(s)
	}
}

type User_usersetContext struct {
	*UserContext
}

func NewUser_usersetContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *User_usersetContext {
	var p = new(User_usersetContext)

	p.UserContext = NewEmptyUserContext()
	p.parser = parser
	p.CopyFrom(ctx.(*UserContext))

	return p
}

func (s *User_usersetContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *User_usersetContext) Object() IObjectContext {
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

func (s *User_usersetContext) Relation() IRelationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRelationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRelationContext)
}

func (s *User_usersetContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TupleListener); ok {
		listenerT.EnterUser_userset(s)
	}
}

func (s *User_usersetContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TupleListener); ok {
		listenerT.ExitUser_userset(s)
	}
}

type User_objectContext struct {
	*UserContext
}

func NewUser_objectContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *User_objectContext {
	var p = new(User_objectContext)

	p.UserContext = NewEmptyUserContext()
	p.parser = parser
	p.CopyFrom(ctx.(*UserContext))

	return p
}

func (s *User_objectContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *User_objectContext) Object() IObjectContext {
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

func (s *User_objectContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TupleListener); ok {
		listenerT.EnterUser_object(s)
	}
}

func (s *User_objectContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TupleListener); ok {
		listenerT.ExitUser_object(s)
	}
}

func (p *TupleParser) User() (localctx IUserContext) {
	this := p
	_ = this

	localctx = NewUserContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, TupleParserRULE_user)

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

	p.SetState(27)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext()) {
	case 1:
		localctx = NewUser_usersetContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(21)
			p.Object()
		}
		{
			p.SetState(22)
			p.Match(TupleParserT__0)
		}
		{
			p.SetState(23)
			p.Relation()
		}

	case 2:
		localctx = NewUser_objectContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(25)
			p.Object()
		}

	case 3:
		localctx = NewUser_idContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(26)
			p.Match(TupleParserID)
		}

	}

	return localctx
}
