package builder

import (
	"fmt"

	parser "github.com/verilog2Go/generated/verilog"
)

// CustomVerilogListener はBaseVerilogListenerを構造体として宣言
type CustomVerilogListener struct {
	*parser.BaseVerilogListener
}

// NewVerilogListener はリスナーの初期化を行う
func NewVerilogListener() *CustomVerilogListener {
	return &CustomVerilogListener{}
}

// EnterModule_declaration is called when production module_declaration is entered.
func (s *CustomVerilogListener) EnterModule_declaration(ctx *parser.Module_declarationContext) {
	//モジュールの初期化
}

// ExitModule_declaration is called when production module_declaration is exited.
func (s *CustomVerilogListener) ExitModule_declaration(ctx *parser.Module_declarationContext) {
	fmt.Println(ctx.Module_identifier().GetText())
}

// ExitInput_declaration is called when production input_declaration is exited.
func (s *CustomVerilogListener) ExitInput_declaration(ctx *parser.Input_declarationContext) {
	fmt.Println(ctx.GetText())
}

// ExitOutput_declaration is called when production output_declaration is exited.
func (s *CustomVerilogListener) ExitOutput_declaration(ctx *parser.Output_declarationContext) {
	fmt.Println(ctx.GetText())
}
