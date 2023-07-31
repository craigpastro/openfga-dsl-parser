// Code generated from DSL.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser // DSL

import "github.com/antlr4-go/antlr/v4"

// BaseDSLListener is a complete listener for a parse tree produced by DSLParser.
type BaseDSLListener struct{}

var _ DSLListener = &BaseDSLListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseDSLListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseDSLListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseDSLListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseDSLListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterDsl is called when production dsl is entered.
func (s *BaseDSLListener) EnterDsl(ctx *DslContext) {}

// ExitDsl is called when production dsl is exited.
func (s *BaseDSLListener) ExitDsl(ctx *DslContext) {}

// EnterTypeDefinition is called when production typeDefinition is entered.
func (s *BaseDSLListener) EnterTypeDefinition(ctx *TypeDefinitionContext) {}

// ExitTypeDefinition is called when production typeDefinition is exited.
func (s *BaseDSLListener) ExitTypeDefinition(ctx *TypeDefinitionContext) {}

// EnterRelation is called when production relation is entered.
func (s *BaseDSLListener) EnterRelation(ctx *RelationContext) {}

// ExitRelation is called when production relation is exited.
func (s *BaseDSLListener) ExitRelation(ctx *RelationContext) {}

// EnterTypeRestriction is called when production typeRestriction is entered.
func (s *BaseDSLListener) EnterTypeRestriction(ctx *TypeRestrictionContext) {}

// ExitTypeRestriction is called when production typeRestriction is exited.
func (s *BaseDSLListener) ExitTypeRestriction(ctx *TypeRestrictionContext) {}

// EnterRelationReferences is called when production relationReferences is entered.
func (s *BaseDSLListener) EnterRelationReferences(ctx *RelationReferencesContext) {}

// ExitRelationReferences is called when production relationReferences is exited.
func (s *BaseDSLListener) ExitRelationReferences(ctx *RelationReferencesContext) {}

// EnterRrType is called when production rrType is entered.
func (s *BaseDSLListener) EnterRrType(ctx *RrTypeContext) {}

// ExitRrType is called when production rrType is exited.
func (s *BaseDSLListener) ExitRrType(ctx *RrTypeContext) {}

// EnterRrTypeAndRelation is called when production rrTypeAndRelation is entered.
func (s *BaseDSLListener) EnterRrTypeAndRelation(ctx *RrTypeAndRelationContext) {}

// ExitRrTypeAndRelation is called when production rrTypeAndRelation is exited.
func (s *BaseDSLListener) ExitRrTypeAndRelation(ctx *RrTypeAndRelationContext) {}

// EnterRrTypeAndWildcard is called when production rrTypeAndWildcard is entered.
func (s *BaseDSLListener) EnterRrTypeAndWildcard(ctx *RrTypeAndWildcardContext) {}

// ExitRrTypeAndWildcard is called when production rrTypeAndWildcard is exited.
func (s *BaseDSLListener) ExitRrTypeAndWildcard(ctx *RrTypeAndWildcardContext) {}

// EnterComputedUserset is called when production computedUserset is entered.
func (s *BaseDSLListener) EnterComputedUserset(ctx *ComputedUsersetContext) {}

// ExitComputedUserset is called when production computedUserset is exited.
func (s *BaseDSLListener) ExitComputedUserset(ctx *ComputedUsersetContext) {}

// EnterIntersection is called when production intersection is entered.
func (s *BaseDSLListener) EnterIntersection(ctx *IntersectionContext) {}

// ExitIntersection is called when production intersection is exited.
func (s *BaseDSLListener) ExitIntersection(ctx *IntersectionContext) {}

// EnterThis is called when production this is entered.
func (s *BaseDSLListener) EnterThis(ctx *ThisContext) {}

// ExitThis is called when production this is exited.
func (s *BaseDSLListener) ExitThis(ctx *ThisContext) {}

// EnterExclusion is called when production exclusion is entered.
func (s *BaseDSLListener) EnterExclusion(ctx *ExclusionContext) {}

// ExitExclusion is called when production exclusion is exited.
func (s *BaseDSLListener) ExitExclusion(ctx *ExclusionContext) {}

// EnterUnion is called when production union is entered.
func (s *BaseDSLListener) EnterUnion(ctx *UnionContext) {}

// ExitUnion is called when production union is exited.
func (s *BaseDSLListener) ExitUnion(ctx *UnionContext) {}

// EnterTupleToUserset is called when production tupleToUserset is entered.
func (s *BaseDSLListener) EnterTupleToUserset(ctx *TupleToUsersetContext) {}

// ExitTupleToUserset is called when production tupleToUserset is exited.
func (s *BaseDSLListener) ExitTupleToUserset(ctx *TupleToUsersetContext) {}

// EnterGrouping is called when production grouping is entered.
func (s *BaseDSLListener) EnterGrouping(ctx *GroupingContext) {}

// ExitGrouping is called when production grouping is exited.
func (s *BaseDSLListener) ExitGrouping(ctx *GroupingContext) {}
