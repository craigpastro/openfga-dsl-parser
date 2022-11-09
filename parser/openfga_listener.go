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

	// EnterTypeRestriction is called when entering the typeRestriction production.
	EnterTypeRestriction(c *TypeRestrictionContext)

	// EnterRelationReferences is called when entering the relationReferences production.
	EnterRelationReferences(c *RelationReferencesContext)

	// EnterType is called when entering the type production.
	EnterType(c *TypeContext)

	// EnterTypeAndRelation is called when entering the typeAndRelation production.
	EnterTypeAndRelation(c *TypeAndRelationContext)

	// EnterComputedUserset is called when entering the computedUserset production.
	EnterComputedUserset(c *ComputedUsersetContext)

	// EnterIntersection is called when entering the intersection production.
	EnterIntersection(c *IntersectionContext)

	// EnterThis is called when entering the this production.
	EnterThis(c *ThisContext)

	// EnterExclusion is called when entering the exclusion production.
	EnterExclusion(c *ExclusionContext)

	// EnterUnion is called when entering the union production.
	EnterUnion(c *UnionContext)

	// EnterTupleToUserset is called when entering the tupleToUserset production.
	EnterTupleToUserset(c *TupleToUsersetContext)

	// EnterGrouping is called when entering the grouping production.
	EnterGrouping(c *GroupingContext)

	// ExitStart is called when exiting the start production.
	ExitStart(c *StartContext)

	// ExitTypeDefinition is called when exiting the typeDefinition production.
	ExitTypeDefinition(c *TypeDefinitionContext)

	// ExitRelation is called when exiting the relation production.
	ExitRelation(c *RelationContext)

	// ExitTypeRestriction is called when exiting the typeRestriction production.
	ExitTypeRestriction(c *TypeRestrictionContext)

	// ExitRelationReferences is called when exiting the relationReferences production.
	ExitRelationReferences(c *RelationReferencesContext)

	// ExitType is called when exiting the type production.
	ExitType(c *TypeContext)

	// ExitTypeAndRelation is called when exiting the typeAndRelation production.
	ExitTypeAndRelation(c *TypeAndRelationContext)

	// ExitComputedUserset is called when exiting the computedUserset production.
	ExitComputedUserset(c *ComputedUsersetContext)

	// ExitIntersection is called when exiting the intersection production.
	ExitIntersection(c *IntersectionContext)

	// ExitThis is called when exiting the this production.
	ExitThis(c *ThisContext)

	// ExitExclusion is called when exiting the exclusion production.
	ExitExclusion(c *ExclusionContext)

	// ExitUnion is called when exiting the union production.
	ExitUnion(c *UnionContext)

	// ExitTupleToUserset is called when exiting the tupleToUserset production.
	ExitTupleToUserset(c *TupleToUsersetContext)

	// ExitGrouping is called when exiting the grouping production.
	ExitGrouping(c *GroupingContext)
}
