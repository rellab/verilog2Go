package builder

import (
	"strings"

	parser "github.com/verilog2Go/antlr/verilog"
	"github.com/verilog2Go/src/expression"
)

var IfDepth int
var hasExpression bool
var statementCount int

// EnterProcedural_timing_control_statement is called when production procedural_timing_control_statement is entered.
func (s *CustomVerilogListener) EnterProcedural_timing_control_statement(ctx *parser.Procedural_timing_control_statementContext) {
	//Block start
	builder.CreateAlways()
}

// ExitProcedural_timing_control_statement is called when production procedural_timing_control_statement is exited.
func (s *CustomVerilogListener) ExitProcedural_timing_control_statement(ctx *parser.Procedural_timing_control_statementContext) {
	//End of block
	statementCount = 0
	builder.EndAlways()
}

// ExitEvent_expression is called when production event_expression is exited.
func (s *CustomVerilogListener) ExitEvent_expression(ctx *parser.Event_expressionContext) {
	//Examine all Always conditions
	for _, v := range ctx.AllEvent_primary() {
		//Includes posedge in the condition
		if strings.Contains(v.GetText(), "posedge") {
			//Eliminate the string posedge and retrieve the BitArray
			id := strings.Replace(v.GetText(), "posedge", "", 1)
			builder.AddPosedgeObserver(id)
		} else if strings.Contains(v.GetText(), "negedge") {
			//negedge
			id := strings.Replace(v.GetText(), "negedge", "", 1)
			builder.AddNegedgeObserver(id)
		}
	}
}

// EnterConditional_statement is called when production conditional_statement is exited.
func (s *CustomVerilogListener) EnterConditional_statement(ctx *parser.Conditional_statementContext) {
	//if statement start
	IfDepth++
	hasExpression = false
	statementCount = 0
	builder.IfStart()
}

// ExitExpression is called when production expression is exited.
func (s *CustomVerilogListener) ExitExpression(ctx *parser.ExpressionContext) {
	//Conditional expression of if statement
	if IfDepth > 0 && !hasExpression {
		hasExpression = true
		builder.IfStatement(expression.CompileExpression(ctx.GetText(), strings.Title(moduleName), s.dimensions))
	}
}

// ExitConditional_statement is called when production conditional_statement is exited.
func (s *CustomVerilogListener) ExitConditional_statement(ctx *parser.Conditional_statementContext) {
	//end if statement
	IfDepth--
}

// EnterStatement_or_null is called when production statement_or_null is entered.
func (s *CustomVerilogListener) EnterStatement_or_null(ctx *parser.Statement_or_nullContext) {
	//Judgment whether it is an else statement
	statementCount++
	if statementCount == 2 {
		builder.ElseStatement()
	}
}

// ExitStatement_or_null is called when production statement_or_null is exited.
func (s *CustomVerilogListener) ExitStatement_or_null(ctx *parser.Statement_or_nullContext) {
	//Start accepting processing in if statement
	if IfDepth > 0 {
		builder.EndIfStatement()
	}
}

// ExitNonblocking_assignment is called when production nonblocking_assignment is exited.
func (s *CustomVerilogListener) ExitNonblocking_assignment(ctx *parser.Nonblocking_assignmentContext) {
	//Non-blocking assignment statement
	builder.DeclarateVariable(ctx.GetText(), s.dimensions)
}
