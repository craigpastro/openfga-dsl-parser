// Code generated from OpenFGA.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser // OpenFGA

import "github.com/antlr4-go/antlr/v4"

// OpenFGAListener is a complete listener for a parse tree produced by OpenFGAParser.
type OpenFGAListener interface {
	antlr.ParseTreeListener

	// EnterProg is called when entering the prog production.
	EnterProg(c *ProgContext)

	// EnterTypedef is called when entering the typedef production.
	EnterTypedef(c *TypedefContext)

	// EnterRelations is called when entering the relations production.
	EnterRelations(c *RelationsContext)

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

	// ExitProg is called when exiting the prog production.
	ExitProg(c *ProgContext)

	// ExitTypedef is called when exiting the typedef production.
	ExitTypedef(c *TypedefContext)

	// ExitRelations is called when exiting the relations production.
	ExitRelations(c *RelationsContext)

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
