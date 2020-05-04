package verilog2Go

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

// ExitInput_declaration is called when production input_declaration is exited.
func (s *CustomVerilogListener) ExitInput_declaration(ctx *parser.Input_declarationContext) {
	fmt.Println(ctx)
}
