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

// EnterRelation is called when production relation is entered.
func (s *BaseTupleListener) EnterRelation(ctx *RelationContext) {}

// ExitRelation is called when production relation is exited.
func (s *BaseTupleListener) ExitRelation(ctx *RelationContext) {}

// EnterUser_userset is called when production user_userset is entered.
func (s *BaseTupleListener) EnterUser_userset(ctx *User_usersetContext) {}

// ExitUser_userset is called when production user_userset is exited.
func (s *BaseTupleListener) ExitUser_userset(ctx *User_usersetContext) {}

// EnterUser_object is called when production user_object is entered.
func (s *BaseTupleListener) EnterUser_object(ctx *User_objectContext) {}

// ExitUser_object is called when production user_object is exited.
func (s *BaseTupleListener) ExitUser_object(ctx *User_objectContext) {}

// EnterUser_id is called when production user_id is entered.
func (s *BaseTupleListener) EnterUser_id(ctx *User_idContext) {}

// ExitUser_id is called when production user_id is exited.
func (s *BaseTupleListener) ExitUser_id(ctx *User_idContext) {}
