// Code generated from Tuple.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser // Tuple

import "github.com/antlr4-go/antlr/v4"

// TupleListener is a complete listener for a parse tree produced by TupleParser.
type TupleListener interface {
	antlr.ParseTreeListener

	// EnterTuple is called when entering the tuple production.
	EnterTuple(c *TupleContext)

	// EnterObject is called when entering the object production.
	EnterObject(c *ObjectContext)

	// EnterUserUserset is called when entering the userUserset production.
	EnterUserUserset(c *UserUsersetContext)

	// EnterUserObject is called when entering the userObject production.
	EnterUserObject(c *UserObjectContext)

	// EnterUserID is called when entering the userID production.
	EnterUserID(c *UserIDContext)

	// ExitTuple is called when exiting the tuple production.
	ExitTuple(c *TupleContext)

	// ExitObject is called when exiting the object production.
	ExitObject(c *ObjectContext)

	// ExitUserUserset is called when exiting the userUserset production.
	ExitUserUserset(c *UserUsersetContext)

	// ExitUserObject is called when exiting the userObject production.
	ExitUserObject(c *UserObjectContext)

	// ExitUserID is called when exiting the userID production.
	ExitUserID(c *UserIDContext)
}
