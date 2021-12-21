package builder

import (
	"strings"

	parser "github.com/verilog2Go/antlr/verilog"
	"github.com/verilog2Go/src/expression"
)

var isAlways bool
var isCase bool
var isLoop bool
var IfDepth int
var statementDepth int
var hasExpression bool
var statementCount int
var caseStatementCount int
var caseStatement string
var seqStatement string
var loopStatement string

// EnterProcedural_timing_control_statement is called when production procedural_timing_control_statement is entered.
func (s *CustomVerilogListener) EnterProcedural_timing_control_statement(ctx *parser.Procedural_timing_control_statementContext) {
	//Block start
	builder.CreateAlways()
	isAlways = true
}

// ExitProcedural_timing_control_statement is called when production procedural_timing_control_statement is exited.
func (s *CustomVerilogListener) ExitProcedural_timing_control_statement(ctx *parser.Procedural_timing_control_statementContext) {
	//End of block
	statementCount = 0
	builder.EndAlways()
	isAlways = false
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
}

// ExitExpression is called when production expression is exited.
func (s *CustomVerilogListener) ExitExpression(ctx *parser.ExpressionContext) {
	//Conditional expression of if statement
	if IfDepth > 0 && !hasExpression {
		hasExpression = true
		if isCase {
			caseStatement += "if variable.CheckBit(" + expression.CompileExpression(ctx.GetText(), strings.Title(moduleName), s.dimensions) + ") {\n"
		} else if isAlways {
			builder.IfStatement(expression.CompileExpression(ctx.GetText(), strings.Title(moduleName), s.dimensions))
		}
	}
}

// ExitConditional_statement is called when production conditional_statement is exited.
func (s *CustomVerilogListener) ExitConditional_statement(ctx *parser.Conditional_statementContext) {
	//end if statement
	IfDepth--
	builder.EndIfStatement()
	// fmt.Println(ctx.GetId())
}

// EnterElse_statement is called when production else_statement is entered.
func (s *CustomVerilogListener) EnterElse_statement(ctx *parser.Else_statementContext) {
	builder.ElseStatement()
}

// EnterStatement_or_null is called when production statement_or_null is entered.
func (s *CustomVerilogListener) EnterStatement(ctx *parser.StatementContext) {
	if isCase {
		if caseStatementCount%2 == 0 && caseStatementCount != 0 {
			caseStatement = caseStatement[:len(caseStatement)-1] + " else{\n"
		}
		caseStatementCount++
		statementDepth++
	}
}

// ExitStatement_or_null is called when production statement_or_null is exited.
func (s *CustomVerilogListener) ExitStatement(ctx *parser.StatementContext) {
	if isCase {
		if statementDepth != 1 {
			caseStatement += "}\n"
		}
		statementDepth--
	}
}

// ExitNonblocking_assignment is called when production nonblocking_assignment is exited.
func (s *CustomVerilogListener) ExitNonblocking_assignment(ctx *parser.Nonblocking_assignmentContext) {
	//Non-blocking assignment statement
	builder.CreateNonBlocking(ctx.Variable_lvalue().GetText(), ctx.Expression().GetText(), s.dimensions)
}

// ExitBlocking_assignment is called when production blocking_assignment is exited.
func (s *CustomVerilogListener) ExitBlocking_assignment(ctx *parser.Blocking_assignmentContext) {
	builder.CreateBlocking(ctx.Variable_lvalue().GetText(), ctx.Expression().GetText(), s.dimensions)
}

func (s *CustomVerilogListener) ExitCase_statement(ctx *parser.Case_statementContext) {
	// CreateSwitch(expression.CompileExpression(ctx.Expression().GetText(), strings.Title(moduleName), s.dimensions))
	if isAlways {
		rightBlock += "switch " + expression.CompileExpression(ctx.Expression().GetText(), strings.Title(moduleName), s.dimensions)[1:] + ".ToInt()" + "{\n"
		rightBlock += cases
		rightBlock += "}\n"
	}
}

func (s *CustomVerilogListener) EnterCase_item(ctx *parser.Case_itemContext) {
	isCase = true
	caseStatementCount = 0
}

// ExitCase_item is called when production Case_item is exited.
func (s *CustomVerilogListener) ExitCase_item(ctx *parser.Case_itemContext) {
	if len(ctx.AllExpression()) != 0 {
		builder.CreateAlwaysCase(toInt(ctx.Expression(0).GetText()))
	} else {
		builder.CreateDefault()
	}
	isCase = false
}

// EnterLoop_statement is called when production loop_statement is entered.
func (s *CustomVerilogListener) EnterLoop_statement(ctx *parser.Loop_statementContext) {
	isLoop = true
}

// ExitLoop_statement is called when production loop_statement is exited.
func (s *CustomVerilogListener) ExitLoop_statement(ctx *parser.Loop_statementContext) {
	isLoop = false
	builder.CreateLoop(s.insertModuleName(ctx.Variable_assignment(0).GetText()), s.insertModuleName(ctx.Expression().GetText()), s.insertModuleName(ctx.Variable_assignment(1).GetText()))
}

func (s *CustomVerilogListener) insertModuleName(str string) string {
	res := str
	for _, v := range s.variables {
		for strings.Contains(str, v.id) {
			res = res[:strings.LastIndex(str, v.id)] + strings.Title(moduleName) + "." + res[strings.LastIndex(str, v.id):]
			str = str[:strings.LastIndex(str, v.id)] + str[strings.LastIndex(str, v.id)+len(v.id):]
		}
	}
	return res
}
