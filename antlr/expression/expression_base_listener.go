// Code generated from Expression.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // Expression

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseExpressionListener is a complete listener for a parse tree produced by ExpressionParser.
type BaseExpressionListener struct{}

var _ ExpressionListener = &BaseExpressionListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseExpressionListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseExpressionListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseExpressionListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseExpressionListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterStart is called when production start is entered.
func (s *BaseExpressionListener) EnterStart(ctx *StartContext) {}

// ExitStart is called when production start is exited.
func (s *BaseExpressionListener) ExitStart(ctx *StartContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseExpressionListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseExpressionListener) ExitExpression(ctx *ExpressionContext) {}
