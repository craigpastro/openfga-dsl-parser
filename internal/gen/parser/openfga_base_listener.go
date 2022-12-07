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

// EnterTypeDefinition is called when production typeDefinition is entered.
func (s *BaseOpenFGAListener) EnterTypeDefinition(ctx *TypeDefinitionContext) {}

// ExitTypeDefinition is called when production typeDefinition is exited.
func (s *BaseOpenFGAListener) ExitTypeDefinition(ctx *TypeDefinitionContext) {}

// EnterRelation is called when production relation is entered.
func (s *BaseOpenFGAListener) EnterRelation(ctx *RelationContext) {}

// ExitRelation is called when production relation is exited.
func (s *BaseOpenFGAListener) ExitRelation(ctx *RelationContext) {}

// EnterTypeRestriction is called when production typeRestriction is entered.
func (s *BaseOpenFGAListener) EnterTypeRestriction(ctx *TypeRestrictionContext) {}

// ExitTypeRestriction is called when production typeRestriction is exited.
func (s *BaseOpenFGAListener) ExitTypeRestriction(ctx *TypeRestrictionContext) {}

// EnterRelationReferences is called when production relationReferences is entered.
func (s *BaseOpenFGAListener) EnterRelationReferences(ctx *RelationReferencesContext) {}

// ExitRelationReferences is called when production relationReferences is exited.
func (s *BaseOpenFGAListener) ExitRelationReferences(ctx *RelationReferencesContext) {}

// EnterRrType is called when production rrType is entered.
func (s *BaseOpenFGAListener) EnterRrType(ctx *RrTypeContext) {}

// ExitRrType is called when production rrType is exited.
func (s *BaseOpenFGAListener) ExitRrType(ctx *RrTypeContext) {}

// EnterRrTypeAndRelation is called when production rrTypeAndRelation is entered.
func (s *BaseOpenFGAListener) EnterRrTypeAndRelation(ctx *RrTypeAndRelationContext) {}

// ExitRrTypeAndRelation is called when production rrTypeAndRelation is exited.
func (s *BaseOpenFGAListener) ExitRrTypeAndRelation(ctx *RrTypeAndRelationContext) {}

// EnterRrTypeAndWildcard is called when production rrTypeAndWildcard is entered.
func (s *BaseOpenFGAListener) EnterRrTypeAndWildcard(ctx *RrTypeAndWildcardContext) {}

// ExitRrTypeAndWildcard is called when production rrTypeAndWildcard is exited.
func (s *BaseOpenFGAListener) ExitRrTypeAndWildcard(ctx *RrTypeAndWildcardContext) {}

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
