package builder

import (
	parser "github.com/verilog2Go/antlr/verilog"
)

// EnterProcedural_timing_control_statement is called when production procedural_timing_control_statement is entered.
func (s *CustomVerilogListener) EnterProcedural_timing_control_statement(ctx *parser.Procedural_timing_control_statementContext) {
	//ブロック開始
	// fmt.Println("ブロック開始")
}

// ExitProcedural_timing_control_statement is called when production procedural_timing_control_statement is exited.
func (s *CustomVerilogListener) ExitProcedural_timing_control_statement(ctx *parser.Procedural_timing_control_statementContext) {
	//ブロック終了
	// fmt.Println("ブロック終了")
	// fmt.Println(ctx.GetText())
}

// EnterEvent_expression is called when production event_expression is entered.
func (s *CustomVerilogListener) EnterEvent_expression(ctx *parser.Event_expressionContext) {
	// fmt.Println("条件文開始")
}

// ExitEvent_expression is called when production event_expression is exited.
func (s *CustomVerilogListener) ExitEvent_expression(ctx *parser.Event_expressionContext) {
	// fmt.Println("条件文終了")
}
