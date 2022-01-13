// Code generated from Expression.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // Expression

import "github.com/antlr/antlr4/runtime/Go/antlr"

// ExpressionListener is a complete listener for a parse tree produced by ExpressionParser.
type ExpressionListener interface {
	antlr.ParseTreeListener

	// EnterStart is called when entering the start production.
	EnterStart(c *StartContext)

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// ExitStart is called when exiting the start production.
	ExitStart(c *StartContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)
}
