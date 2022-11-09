// Code generated from java-escape by ANTLR 4.11.1. DO NOT EDIT.

package parser // OpenFGA

import "github.com/antlr/antlr4/runtime/Go/antlr/v4"

// OpenFGAListener is a complete listener for a parse tree produced by OpenFGAParser.
type OpenFGAListener interface {
	antlr.ParseTreeListener

	// EnterStart is called when entering the start production.
	EnterStart(c *StartContext)

	// EnterTypedef is called when entering the typedef production.
	EnterTypedef(c *TypedefContext)

	// EnterRelation is called when entering the relation production.
	EnterRelation(c *RelationContext)

	// EnterTypes is called when entering the types production.
	EnterTypes(c *TypesContext)

	// EnterComputedUserset is called when entering the computedUserset production.
	EnterComputedUserset(c *ComputedUsersetContext)

	// EnterRelationReferences is called when entering the relationReferences production.
	EnterRelationReferences(c *RelationReferencesContext)

	// EnterType is called when entering the type production.
	EnterType(c *TypeContext)

	// EnterTypeAndRelation is called when entering the typeAndRelation production.
	EnterTypeAndRelation(c *TypeAndRelationContext)

	// EnterTupleToUserset is called when entering the tupleToUserset production.
	EnterTupleToUserset(c *TupleToUsersetContext)

	// EnterUnion is called when entering the union production.
	EnterUnion(c *UnionContext)

	// EnterIntersection is called when entering the intersection production.
	EnterIntersection(c *IntersectionContext)

	// EnterExclusion is called when entering the exclusion production.
	EnterExclusion(c *ExclusionContext)

	// ExitStart is called when exiting the start production.
	ExitStart(c *StartContext)

	// ExitTypedef is called when exiting the typedef production.
	ExitTypedef(c *TypedefContext)

	// ExitRelation is called when exiting the relation production.
	ExitRelation(c *RelationContext)

	// ExitTypes is called when exiting the types production.
	ExitTypes(c *TypesContext)

	// ExitComputedUserset is called when exiting the computedUserset production.
	ExitComputedUserset(c *ComputedUsersetContext)

	// ExitRelationReferences is called when exiting the relationReferences production.
	ExitRelationReferences(c *RelationReferencesContext)

	// ExitType is called when exiting the type production.
	ExitType(c *TypeContext)

	// ExitTypeAndRelation is called when exiting the typeAndRelation production.
	ExitTypeAndRelation(c *TypeAndRelationContext)

	// ExitTupleToUserset is called when exiting the tupleToUserset production.
	ExitTupleToUserset(c *TupleToUsersetContext)

	// ExitUnion is called when exiting the union production.
	ExitUnion(c *UnionContext)

	// ExitIntersection is called when exiting the intersection production.
	ExitIntersection(c *IntersectionContext)

	// ExitExclusion is called when exiting the exclusion production.
	ExitExclusion(c *ExclusionContext)
}
