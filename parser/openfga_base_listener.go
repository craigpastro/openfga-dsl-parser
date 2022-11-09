// Code generated from java-escape by ANTLR 4.11.1. DO NOT EDIT.

package parser // OpenFGA

import "github.com/antlr/antlr4/runtime/Go/antlr/v4"

// BaseOpenFGAListener is a complete listener for a parse tree produced by OpenFGAParser.
type BaseOpenFGAListener struct{}

var _ OpenFGAListener = &BaseOpenFGAListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseOpenFGAListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseOpenFGAListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseOpenFGAListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseOpenFGAListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterStart is called when production start is entered.
func (s *BaseOpenFGAListener) EnterStart(ctx *StartContext) {}

// ExitStart is called when production start is exited.
func (s *BaseOpenFGAListener) ExitStart(ctx *StartContext) {}

// EnterTypedef is called when production typedef is entered.
func (s *BaseOpenFGAListener) EnterTypedef(ctx *TypedefContext) {}

// ExitTypedef is called when production typedef is exited.
func (s *BaseOpenFGAListener) ExitTypedef(ctx *TypedefContext) {}

// EnterRelation is called when production relation is entered.
func (s *BaseOpenFGAListener) EnterRelation(ctx *RelationContext) {}

// ExitRelation is called when production relation is exited.
func (s *BaseOpenFGAListener) ExitRelation(ctx *RelationContext) {}

// EnterTypes is called when production types is entered.
func (s *BaseOpenFGAListener) EnterTypes(ctx *TypesContext) {}

// ExitTypes is called when production types is exited.
func (s *BaseOpenFGAListener) ExitTypes(ctx *TypesContext) {}

// EnterComputedUserset is called when production computedUserset is entered.
func (s *BaseOpenFGAListener) EnterComputedUserset(ctx *ComputedUsersetContext) {}

// ExitComputedUserset is called when production computedUserset is exited.
func (s *BaseOpenFGAListener) ExitComputedUserset(ctx *ComputedUsersetContext) {}

// EnterRelationReferences is called when production relationReferences is entered.
func (s *BaseOpenFGAListener) EnterRelationReferences(ctx *RelationReferencesContext) {}

// ExitRelationReferences is called when production relationReferences is exited.
func (s *BaseOpenFGAListener) ExitRelationReferences(ctx *RelationReferencesContext) {}

// EnterType is called when production type is entered.
func (s *BaseOpenFGAListener) EnterType(ctx *TypeContext) {}

// ExitType is called when production type is exited.
func (s *BaseOpenFGAListener) ExitType(ctx *TypeContext) {}

// EnterTypeAndRelation is called when production typeAndRelation is entered.
func (s *BaseOpenFGAListener) EnterTypeAndRelation(ctx *TypeAndRelationContext) {}

// ExitTypeAndRelation is called when production typeAndRelation is exited.
func (s *BaseOpenFGAListener) ExitTypeAndRelation(ctx *TypeAndRelationContext) {}

// EnterTupleToUserset is called when production tupleToUserset is entered.
func (s *BaseOpenFGAListener) EnterTupleToUserset(ctx *TupleToUsersetContext) {}

// ExitTupleToUserset is called when production tupleToUserset is exited.
func (s *BaseOpenFGAListener) ExitTupleToUserset(ctx *TupleToUsersetContext) {}

// EnterUnion is called when production union is entered.
func (s *BaseOpenFGAListener) EnterUnion(ctx *UnionContext) {}

// ExitUnion is called when production union is exited.
func (s *BaseOpenFGAListener) ExitUnion(ctx *UnionContext) {}

// EnterIntersection is called when production intersection is entered.
func (s *BaseOpenFGAListener) EnterIntersection(ctx *IntersectionContext) {}

// ExitIntersection is called when production intersection is exited.
func (s *BaseOpenFGAListener) ExitIntersection(ctx *IntersectionContext) {}

// EnterExclusion is called when production exclusion is entered.
func (s *BaseOpenFGAListener) EnterExclusion(ctx *ExclusionContext) {}

// ExitExclusion is called when production exclusion is exited.
func (s *BaseOpenFGAListener) ExitExclusion(ctx *ExclusionContext) {}
