package builder

import (
	"fmt"
	"strings"

	parser "github.com/verilog2Go/antlr/verilog"
)

// EnterProcedural_timing_control_statement is called when production procedural_timing_control_statement is entered.
func (s *CustomVerilogListener) EnterProcedural_timing_control_statement(ctx *parser.Procedural_timing_control_statementContext) {
	//ブロック開始
	CreateAlways()
}

// ExitProcedural_timing_control_statement is called when production procedural_timing_control_statement is exited.
func (s *CustomVerilogListener) ExitProcedural_timing_control_statement(ctx *parser.Procedural_timing_control_statementContext) {
	//ブロック終了
	EndAlways()
}

// ExitEvent_expression is called when production event_expression is exited.
func (s *CustomVerilogListener) ExitEvent_expression(ctx *parser.Event_expressionContext) {
	//Alwaysの条件を全て調べる
	for _, v := range ctx.AllEvent_primary() {
		//条件にposedgeを含む
		if strings.Contains(v.GetText(), "posedge") {
			//posedgeという文字列を消してBitArrayを取り出す
			id := strings.Replace(v.GetText(), "posedge", "", 1)
			AddPosedgeObserver(id)
		} else if strings.Contains(v.GetText(), "negedge") {
			//negedge
			id := strings.Replace(v.GetText(), "negedge", "", 1)
			AddNegedgeObserver(id)
		}
	}
}

// EnterConditional_statement is called when production conditional_statement is exited.
func (s *CustomVerilogListener) EnterConditional_statement(ctx *parser.Conditional_statementContext) {
	//if文開始
}

// ExitConditional_statement is called when production conditional_statement is exited.
func (s *CustomVerilogListener) ExitConditional_statement(ctx *parser.Conditional_statementContext) {
	fmt.Println(ctx.Expression().GetText())
}

// ExitStatement_or_null is called when production statement_or_null is exited.
func (s *CustomVerilogListener) ExitStatement_or_null(ctx *parser.Statement_or_nullContext) {
	fmt.Println("else " + ctx.GetText())
}

// ExitNonblocking_assignment is called when production nonblocking_assignment is exited.
func (s *CustomVerilogListener) ExitNonblocking_assignment(ctx *parser.Nonblocking_assignmentContext) {
	fmt.Println("nonblocking " + ctx.GetText())
}
