// Code generated from OpenFGA.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser // OpenFGA

import "github.com/antlr4-go/antlr/v4"

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

// EnterProg is called when production prog is entered.
func (s *BaseOpenFGAListener) EnterProg(ctx *ProgContext) {}

// ExitProg is called when production prog is exited.
func (s *BaseOpenFGAListener) ExitProg(ctx *ProgContext) {}

// EnterTypedef is called when production typedef is entered.
func (s *BaseOpenFGAListener) EnterTypedef(ctx *TypedefContext) {}

// ExitTypedef is called when production typedef is exited.
func (s *BaseOpenFGAListener) ExitTypedef(ctx *TypedefContext) {}

// EnterRelations is called when production relations is entered.
func (s *BaseOpenFGAListener) EnterRelations(ctx *RelationsContext) {}

// ExitRelations is called when production relations is exited.
func (s *BaseOpenFGAListener) ExitRelations(ctx *RelationsContext) {}

// EnterComputedUserset is called when production computedUserset is entered.
func (s *BaseOpenFGAListener) EnterComputedUserset(ctx *ComputedUsersetContext) {}

// ExitComputedUserset is called when production computedUserset is exited.
func (s *BaseOpenFGAListener) ExitComputedUserset(ctx *ComputedUsersetContext) {}

// EnterIntersection is called when production intersection is entered.
func (s *BaseOpenFGAListener) EnterIntersection(ctx *IntersectionContext) {}

// ExitIntersection is called when production intersection is exited.
func (s *BaseOpenFGAListener) ExitIntersection(ctx *IntersectionContext) {}

// EnterThis is called when production this is entered.
func (s *BaseOpenFGAListener) EnterThis(ctx *ThisContext) {}

// ExitThis is called when production this is exited.
func (s *BaseOpenFGAListener) ExitThis(ctx *ThisContext) {}

// EnterExclusion is called when production exclusion is entered.
func (s *BaseOpenFGAListener) EnterExclusion(ctx *ExclusionContext) {}

// ExitExclusion is called when production exclusion is exited.
func (s *BaseOpenFGAListener) ExitExclusion(ctx *ExclusionContext) {}

// EnterUnion is called when production union is entered.
func (s *BaseOpenFGAListener) EnterUnion(ctx *UnionContext) {}

// ExitUnion is called when production union is exited.
func (s *BaseOpenFGAListener) ExitUnion(ctx *UnionContext) {}

// EnterTupleToUserset is called when production tupleToUserset is entered.
func (s *BaseOpenFGAListener) EnterTupleToUserset(ctx *TupleToUsersetContext) {}

// ExitTupleToUserset is called when production tupleToUserset is exited.
func (s *BaseOpenFGAListener) ExitTupleToUserset(ctx *TupleToUsersetContext) {}

// EnterGrouping is called when production grouping is entered.
func (s *BaseOpenFGAListener) EnterGrouping(ctx *GroupingContext) {}

// ExitGrouping is called when production grouping is exited.
func (s *BaseOpenFGAListener) ExitGrouping(ctx *GroupingContext) {}
