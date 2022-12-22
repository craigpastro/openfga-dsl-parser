// Code generated from java-escape by ANTLR 4.11.1. DO NOT EDIT.

package parser // Tuple

import "github.com/antlr/antlr4/runtime/Go/antlr/v4"

// BaseTupleListener is a complete listener for a parse tree produced by TupleParser.
type BaseTupleListener struct{}

var _ TupleListener = &BaseTupleListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseTupleListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseTupleListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseTupleListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseTupleListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterTuple is called when production tuple is entered.
func (s *BaseTupleListener) EnterTuple(ctx *TupleContext) {}

// ExitTuple is called when production tuple is exited.
func (s *BaseTupleListener) ExitTuple(ctx *TupleContext) {}

// EnterObject is called when production object is entered.
func (s *BaseTupleListener) EnterObject(ctx *ObjectContext) {}

// ExitObject is called when production object is exited.
func (s *BaseTupleListener) ExitObject(ctx *ObjectContext) {}

// EnterUserUserset is called when production userUserset is entered.
func (s *BaseTupleListener) EnterUserUserset(ctx *UserUsersetContext) {}

// ExitUserUserset is called when production userUserset is exited.
func (s *BaseTupleListener) ExitUserUserset(ctx *UserUsersetContext) {}

// EnterUserObject is called when production userObject is entered.
func (s *BaseTupleListener) EnterUserObject(ctx *UserObjectContext) {}

// ExitUserObject is called when production userObject is exited.
func (s *BaseTupleListener) ExitUserObject(ctx *UserObjectContext) {}

// EnterUserID is called when production userID is entered.
func (s *BaseTupleListener) EnterUserID(ctx *UserIDContext) {}

// ExitUserID is called when production userID is exited.
func (s *BaseTupleListener) ExitUserID(ctx *UserIDContext) {}
