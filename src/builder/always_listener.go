package builder

import (
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
