// Code generated from java-escape by ANTLR 4.11.1. DO NOT EDIT.

package parser // Tuple

import "github.com/antlr/antlr4/runtime/Go/antlr/v4"

// TupleListener is a complete listener for a parse tree produced by TupleParser.
type TupleListener interface {
	antlr.ParseTreeListener

	// EnterTuple is called when entering the tuple production.
	EnterTuple(c *TupleContext)

	// EnterUserUserset is called when entering the userUserset production.
	EnterUserUserset(c *UserUsersetContext)

	// EnterUserObject is called when entering the userObject production.
	EnterUserObject(c *UserObjectContext)

	// EnterUserId is called when entering the userId production.
	EnterUserId(c *UserIdContext)

	// ExitTuple is called when exiting the tuple production.
	ExitTuple(c *TupleContext)

	// ExitUserUserset is called when exiting the userUserset production.
	ExitUserUserset(c *UserUsersetContext)

	// ExitUserObject is called when exiting the userObject production.
	ExitUserObject(c *UserObjectContext)

	// ExitUserId is called when exiting the userId production.
	ExitUserId(c *UserIdContext)
}
