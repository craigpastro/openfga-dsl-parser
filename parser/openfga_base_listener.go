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

// EnterTroiTypeRestriction is called when production troiTypeRestriction is entered.
func (s *BaseOpenFGAListener) EnterTroiTypeRestriction(ctx *TroiTypeRestrictionContext) {}

// ExitTroiTypeRestriction is called when production troiTypeRestriction is exited.
func (s *BaseOpenFGAListener) ExitTroiTypeRestriction(ctx *TroiTypeRestrictionContext) {}

// EnterTroiId is called when production troiId is entered.
func (s *BaseOpenFGAListener) EnterTroiId(ctx *TroiIdContext) {}

// ExitTroiId is called when production troiId is exited.
func (s *BaseOpenFGAListener) ExitTroiId(ctx *TroiIdContext) {}

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

// EnterRewrite is called when production rewrite is entered.
func (s *BaseOpenFGAListener) EnterRewrite(ctx *RewriteContext) {}

// ExitRewrite is called when production rewrite is exited.
func (s *BaseOpenFGAListener) ExitRewrite(ctx *RewriteContext) {}

// EnterOrTTU is called when production orTTU is entered.
func (s *BaseOpenFGAListener) EnterOrTTU(ctx *OrTTUContext) {}

// ExitOrTTU is called when production orTTU is exited.
func (s *BaseOpenFGAListener) ExitOrTTU(ctx *OrTTUContext) {}

// EnterOrs is called when production ors is entered.
func (s *BaseOpenFGAListener) EnterOrs(ctx *OrsContext) {}

// ExitOrs is called when production ors is exited.
func (s *BaseOpenFGAListener) ExitOrs(ctx *OrsContext) {}

// EnterOr is called when production or is entered.
func (s *BaseOpenFGAListener) EnterOr(ctx *OrContext) {}

// ExitOr is called when production or is exited.
func (s *BaseOpenFGAListener) ExitOr(ctx *OrContext) {}

// EnterAnds is called when production ands is entered.
func (s *BaseOpenFGAListener) EnterAnds(ctx *AndsContext) {}

// ExitAnds is called when production ands is exited.
func (s *BaseOpenFGAListener) ExitAnds(ctx *AndsContext) {}

// EnterAnd is called when production and is entered.
func (s *BaseOpenFGAListener) EnterAnd(ctx *AndContext) {}

// ExitAnd is called when production and is exited.
func (s *BaseOpenFGAListener) ExitAnd(ctx *AndContext) {}

// EnterExclusion is called when production exclusion is entered.
func (s *BaseOpenFGAListener) EnterExclusion(ctx *ExclusionContext) {}

// ExitExclusion is called when production exclusion is exited.
func (s *BaseOpenFGAListener) ExitExclusion(ctx *ExclusionContext) {}
