package builder

import (
	"fmt"
	"strconv"

	parser "github.com/verilog2Go/antlr/verilog"
	"github.com/verilog2Go/src/expression"
)

// CustomVerilogListener はBaseVerilogListenerを構造体として宣言
type CustomVerilogListener struct {
	*parser.BaseVerilogListener
	ports       []Port
	currentPort Port
}

//Port はポートのidとビット数を持つ構造体
type Port struct {
	id     string
	length int
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
	StartModule(ctx.Module_identifier().GetText())
	DeclarePorts(s.ports)
	CreateConstructor(ctx.Module_identifier().GetText(), s.ports)
}

// ExitList_of_port_identifiers is called when production list_of_port_identifiers is exited.
func (s *CustomVerilogListener) ExitList_of_port_identifiers(ctx *parser.List_of_port_identifiersContext) {
	for i := 0; i < len(ctx.AllPort_identifier()); i++ {
		s.currentPort.id = ctx.AllPort_identifier()[i].GetText()
		s.ports = append(s.ports, s.currentPort)
	}
}

// ExitRange_ is called when production range_ is exited.
func (s *CustomVerilogListener) ExitRange_(ctx *parser.Range_Context) {
	//文字列をintに変換
	s.currentPort.length, _ = strconv.Atoi(ctx.Msb_constant_expression().GetText())
}

// ExitNet_assignment is called when production net_assignment is exited.
func (s *CustomVerilogListener) ExitNet_assignment(ctx *parser.Net_assignmentContext) {
	fmt.Println(ctx.Net_lvalue().GetText())
	fmt.Println(expression.CompileExpression(ctx.Expression().GetText()))
}
