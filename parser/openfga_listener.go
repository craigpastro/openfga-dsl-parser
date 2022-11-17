// Code generated from java-escape by ANTLR 4.11.1. DO NOT EDIT.

package parser // OpenFGA

import "github.com/antlr/antlr4/runtime/Go/antlr/v4"

// OpenFGAListener is a complete listener for a parse tree produced by OpenFGAParser.
type OpenFGAListener interface {
	antlr.ParseTreeListener

	// EnterStart is called when entering the start production.
	EnterStart(c *StartContext)

	// EnterTypeDefinition is called when entering the typeDefinition production.
	EnterTypeDefinition(c *TypeDefinitionContext)

	// EnterRelation is called when entering the relation production.
	EnterRelation(c *RelationContext)

	// EnterTroiTypeRestriction is called when entering the troiTypeRestriction production.
	EnterTroiTypeRestriction(c *TroiTypeRestrictionContext)

	// EnterTroiId is called when entering the troiId production.
	EnterTroiId(c *TroiIdContext)

	// EnterTypeRestriction is called when entering the typeRestriction production.
	EnterTypeRestriction(c *TypeRestrictionContext)

	// EnterRelationReferences is called when entering the relationReferences production.
	EnterRelationReferences(c *RelationReferencesContext)

	// EnterRrType is called when entering the rrType production.
	EnterRrType(c *RrTypeContext)

	// EnterRrTypeAndRelation is called when entering the rrTypeAndRelation production.
	EnterRrTypeAndRelation(c *RrTypeAndRelationContext)

	// EnterRrTypeAndWildcard is called when entering the rrTypeAndWildcard production.
	EnterRrTypeAndWildcard(c *RrTypeAndWildcardContext)

	// EnterRewrite is called when entering the rewrite production.
	EnterRewrite(c *RewriteContext)

	// EnterOrTTU is called when entering the orTTU production.
	EnterOrTTU(c *OrTTUContext)

	// EnterOrs is called when entering the ors production.
	EnterOrs(c *OrsContext)

	// EnterOr is called when entering the or production.
	EnterOr(c *OrContext)

	// EnterAnds is called when entering the ands production.
	EnterAnds(c *AndsContext)

	// EnterAnd is called when entering the and production.
	EnterAnd(c *AndContext)

	// EnterExclusion is called when entering the exclusion production.
	EnterExclusion(c *ExclusionContext)

	// ExitStart is called when exiting the start production.
	ExitStart(c *StartContext)

	// ExitTypeDefinition is called when exiting the typeDefinition production.
	ExitTypeDefinition(c *TypeDefinitionContext)

	// ExitRelation is called when exiting the relation production.
	ExitRelation(c *RelationContext)

	// ExitTroiTypeRestriction is called when exiting the troiTypeRestriction production.
	ExitTroiTypeRestriction(c *TroiTypeRestrictionContext)

	// ExitTroiId is called when exiting the troiId production.
	ExitTroiId(c *TroiIdContext)

	// ExitTypeRestriction is called when exiting the typeRestriction production.
	ExitTypeRestriction(c *TypeRestrictionContext)

	// ExitRelationReferences is called when exiting the relationReferences production.
	ExitRelationReferences(c *RelationReferencesContext)

	// ExitRrType is called when exiting the rrType production.
	ExitRrType(c *RrTypeContext)

	// ExitRrTypeAndRelation is called when exiting the rrTypeAndRelation production.
	ExitRrTypeAndRelation(c *RrTypeAndRelationContext)

	// ExitRrTypeAndWildcard is called when exiting the rrTypeAndWildcard production.
	ExitRrTypeAndWildcard(c *RrTypeAndWildcardContext)

	// ExitRewrite is called when exiting the rewrite production.
	ExitRewrite(c *RewriteContext)

	// ExitOrTTU is called when exiting the orTTU production.
	ExitOrTTU(c *OrTTUContext)

	// ExitOrs is called when exiting the ors production.
	ExitOrs(c *OrsContext)

	// ExitOr is called when exiting the or production.
	ExitOr(c *OrContext)

	// ExitAnds is called when exiting the ands production.
	ExitAnds(c *AndsContext)

	// ExitAnd is called when exiting the and production.
	ExitAnd(c *AndContext)

	// ExitExclusion is called when exiting the exclusion production.
	ExitExclusion(c *ExclusionContext)
}
