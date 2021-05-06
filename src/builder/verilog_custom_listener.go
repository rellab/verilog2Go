package builder

import (
	"strconv"
	"strings"

	parser "github.com/verilog2Go/antlr/verilog"
	"github.com/verilog2Go/src/expression"
)

// CustomVerilogListener はBaseVerilogListenerを構造体として宣言
type CustomVerilogListener struct {
	*parser.BaseVerilogListener
	ports       []Port
	params      []Param
	currentPort Port
}

//Port はポートのidとビット数を持つ構造体
type Port struct {
	id          string
	length      int
	isDimension bool
}

type Param struct {
	id         string
	initiation string
}

var isInstance bool

// NewVerilogListener はリスナーの初期化を行う
func NewVerilogListener() *CustomVerilogListener {
	return &CustomVerilogListener{}
}

// ExitSource_text is called when production source_text is exited.
func (s *CustomVerilogListener) ExitSource_text(ctx *parser.Source_textContext) {
	//パース終了
	EndModule()
}

// EnterModule_declaration is called when production module_declaration is entered.
func (s *CustomVerilogListener) EnterModule_declaration(ctx *parser.Module_declarationContext) {
	//モジュールの初期化
	isInstance = false
}

// ExitModule_identifier is called when production module_identifier is exited.
func (s *CustomVerilogListener) ExitModule_identifier(ctx *parser.Module_identifierContext) {
	if !isInstance {
		StartModule(ctx.GetText())
	}
	isInstance = true
}

// ExitModule_declaration is called when production module_declaration is exited.
func (s *CustomVerilogListener) ExitModule_declaration(ctx *parser.Module_declarationContext) {
	// StartModule(ctx.Module_identifier().GetText())
	DeclarePorts(s.ports)
	CreateConstructor(ctx.Module_identifier().GetText(), s.ports, s.params)
}

// ExitInput_declaration is called when production Input_declaration is exited.
func (s *CustomVerilogListener) ExitInput_declaration(ctx *parser.Input_declarationContext) {
	strs := strings.Split(ctx.List_of_port_identifiers().GetText(), ",")
	for i := 0; i < len(strs); i++ {
		if strings.Contains(strs[i], "[") {
			s.currentPort.id = strs[i][:strings.Index(strs[i], "[")]
			s.currentPort.isDimension = true
		} else {
			s.currentPort.id = strs[i]
			s.currentPort.isDimension = false
		}
		s.ports = append(s.ports, s.currentPort)
		DeclareInput(s.currentPort)
	}
}

// ExitOutput_declaration is called when production Output_declaration is exited.
func (s *CustomVerilogListener) ExitOutput_declaration(ctx *parser.Output_declarationContext) {
	strs := strings.Split(ctx.List_of_port_identifiers().GetText(), ",")
	for i := 0; i < len(strs); i++ {
		if strings.Contains(strs[i], "[") {
			s.currentPort.id = strs[i][:strings.Index(strs[i], "[")]
			s.currentPort.isDimension = true
		} else {
			s.currentPort.id = strs[i]
			s.currentPort.isDimension = false
		}
		s.ports = append(s.ports, s.currentPort)
	}
}

// ExitReg_declaration is called when production Reg_declaration is exited.
func (s *CustomVerilogListener) ExitReg_declaration(ctx *parser.Reg_declarationContext) {
	strs := strings.Split(ctx.List_of_variable_identifiers().GetText(), ",")
	for i := 0; i < len(strs); i++ {
		if strings.Contains(strs[i], "[") {
			s.currentPort.id = strs[i][:strings.Index(strs[i], "[")]
			s.currentPort.isDimension = true
		} else {
			s.currentPort.id = strs[i]
			s.currentPort.isDimension = false
		}
		s.ports = append(s.ports, s.currentPort)
	}
}

// ExitNet_declaration is called when production Reg_declaration is exited.
func (s *CustomVerilogListener) ExitNet_declaration(ctx *parser.Net_declarationContext) {
	strs := strings.Split(ctx.List_of_net_identifiers().GetText(), ",")
	for i := 0; i < len(strs); i++ {
		if strings.Contains(strs[i], "[") {
			s.currentPort.id = strs[i][:strings.Index(strs[i], "[")]
			s.currentPort.isDimension = true
		} else {
			s.currentPort.id = strs[i]
			s.currentPort.isDimension = false
		}
		s.ports = append(s.ports, s.currentPort)
	}
}

// ExitParam_assignment is called when production param_assignment is exited.
func (s *CustomVerilogListener) ExitParam_assignment(ctx *parser.Param_assignmentContext) {
	s.currentPort.id = ctx.Parameter_identifier().GetText()
	s.ports = append(s.ports, s.currentPort)
	s.params = append(s.params, Param{
		id:         s.currentPort.id,
		initiation: expression.CompileExpression(ctx.Constant_expression().GetText(), ModuleName),
	})
}

// ExitRange_ is called when production range_ is exited.
func (s *CustomVerilogListener) ExitRange_(ctx *parser.Range_Context) {
	//文字列をintに変換
	s.currentPort.length, _ = strconv.Atoi(ctx.Msb_constant_expression().GetText())
}

// ExitNet_assignment is called when production net_assignment is exited.
func (s *CustomVerilogListener) ExitNet_assignment(ctx *parser.Net_assignmentContext) {
	CreateExec(ctx.Net_lvalue().GetText(), expression.CompileExpression(ctx.Expression().GetText(), ModuleName))
}
