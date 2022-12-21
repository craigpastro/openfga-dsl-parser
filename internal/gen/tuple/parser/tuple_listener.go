// Code generated from java-escape by ANTLR 4.11.1. DO NOT EDIT.

package parser // Tuple

import "github.com/antlr/antlr4/runtime/Go/antlr/v4"

// TupleListener is a complete listener for a parse tree produced by TupleParser.
type TupleListener interface {
	antlr.ParseTreeListener

	// EnterTuple is called when entering the tuple production.
	EnterTuple(c *TupleContext)

	// EnterObject is called when entering the object production.
	EnterObject(c *ObjectContext)

	// EnterRelation is called when entering the relation production.
	EnterRelation(c *RelationContext)

	// EnterUser_userset is called when entering the user_userset production.
	EnterUser_userset(c *User_usersetContext)

	// EnterUser_object is called when entering the user_object production.
	EnterUser_object(c *User_objectContext)

	// EnterUser_id is called when entering the user_id production.
	EnterUser_id(c *User_idContext)

	// ExitTuple is called when exiting the tuple production.
	ExitTuple(c *TupleContext)

	// ExitObject is called when exiting the object production.
	ExitObject(c *ObjectContext)

	// ExitRelation is called when exiting the relation production.
	ExitRelation(c *RelationContext)

	// ExitUser_userset is called when exiting the user_userset production.
	ExitUser_userset(c *User_usersetContext)

	// ExitUser_object is called when exiting the user_object production.
	ExitUser_object(c *User_objectContext)

	// ExitUser_id is called when exiting the user_id production.
	ExitUser_id(c *User_idContext)
}
