package builder

import (
	"bytes"
	"strconv"
	"strings"

	parser "github.com/verilog2Go/antlr/verilog"
	"github.com/verilog2Go/src/expression"
)

// CustomVerilogListener declares BaseVerilogListener as a structure
type CustomVerilogListener struct {
	*parser.BaseVerilogListener
	ports       []Port
	dimensions  []string
	params      []Param
	currentPort Port
}

//Port is a structure with the id and number of bits of the port
type Port struct {
	id          string
	length      int
	portType    string
	isDimension bool
}

type Param struct {
	id         string
	initiation string
}

var builder Builder
var isInstance bool

// NewVerilogListener initializes the listener
func NewVerilogListener() *CustomVerilogListener {
	return &CustomVerilogListener{}
}

// ExitSource_text is called when production source_text is exited.
func (s *CustomVerilogListener) EnterSource_text(ctx *parser.Source_textContext) {
	builder = Builder{
		moduleName:     "",
		ports:          bytes.NewBuffer(make([]byte, 0, 100)),
		inputs:         []Port{},
		constructor:    bytes.NewBuffer(make([]byte, 0, 100)),
		subConstructor: bytes.NewBuffer(make([]byte, 0, 100)),
		observer:       bytes.NewBuffer(make([]byte, 0, 100)),
		assigns:        bytes.NewBuffer(make([]byte, 0, 100)),
		instances:      bytes.NewBuffer(make([]byte, 0, 100)),
		runMethod:      bytes.NewBuffer(make([]byte, 0, 100)),
		preAlways:      bytes.NewBuffer(make([]byte, 0, 100)),
		always:         bytes.NewBuffer(make([]byte, 0, 100)),
		function:       bytes.NewBuffer(make([]byte, 0, 100)),
		source:         bytes.NewBuffer(make([]byte, 0, 1000)),
	}
}

// ExitSource_text is called when production source_text is exited.
func (s *CustomVerilogListener) ExitSource_text(ctx *parser.Source_textContext) {
	//End of perse
	builder.WriteFile(moduleName)
}

// EnterModule_declaration is called when production module_declaration is entered.
func (s *CustomVerilogListener) EnterModule_declaration(ctx *parser.Module_declarationContext) {
	//Module initialization
	isInstance = false
}

// ExitModule_identifier is called when production module_identifier is exited.
func (s *CustomVerilogListener) ExitModule_identifier(ctx *parser.Module_identifierContext) {
	if !isInstance {
		builder.StartModule(ctx.GetText())
	}
	isInstance = true
}

// ExitModule_declaration is called when production module_declaration is exited.
func (s *CustomVerilogListener) ExitModule_declaration(ctx *parser.Module_declarationContext) {
	// StartModule(ctx.Module_identifier().GetText())
	builder.DeclarePorts(s.ports)
	builder.CreateConstructor(ctx.Module_identifier().GetText(), s.ports, s.params)
	builder.CreateRunMethod(s.ports)
}

// ExitInput_declaration is called when production Input_declaration is exited.
func (s *CustomVerilogListener) ExitInput_declaration(ctx *parser.Input_declarationContext) {
	strs := strings.Split(ctx.List_of_port_identifiers().GetText(), ",")
	for i := 0; i < len(strs); i++ {
		if strings.Contains(strs[i], "[") {
			s.currentPort.id = strs[i][:strings.Index(strs[i], "[")]
			s.currentPort.length, _ = strconv.Atoi(ctx.Range_().GetText()[strings.Index(ctx.Range_().GetText(), "[")+1 : strings.Index(ctx.Range_().GetText(), "[")+2])
			s.currentPort.length++
			s.currentPort.isDimension = true
			s.currentPort.portType = "input"
			s.dimensions = append(s.dimensions, strs[i][:strings.Index(strs[i], "[")])
		} else {
			s.currentPort.id = strs[i]
			if strings.Contains(ctx.GetText(), "[") {
				s.currentPort.length, _ = strconv.Atoi(ctx.Range_().GetText()[strings.Index(ctx.Range_().GetText(), "[")+1 : strings.Index(ctx.Range_().GetText(), "[")+2])
				s.currentPort.length++
			} else {
				s.currentPort.length = 1
			}
			s.currentPort.isDimension = false
			s.currentPort.portType = "input"
		}
		s.ports = append(s.ports, s.currentPort)
		builder.DeclareInput(s.currentPort)
	}
}

// ExitOutput_declaration is called when production Output_declaration is exited.
func (s *CustomVerilogListener) ExitOutput_declaration(ctx *parser.Output_declarationContext) {
	strs := strings.Split(ctx.List_of_port_identifiers().GetText(), ",")
	for i := 0; i < len(strs); i++ {
		if strings.Contains(strs[i], "[") {
			s.currentPort.id = strs[i][:strings.Index(strs[i], "[")]
			s.currentPort.length, _ = strconv.Atoi(ctx.Range_().GetText()[strings.Index(ctx.Range_().GetText(), "[")+1 : strings.Index(ctx.Range_().GetText(), "[")+2])
			s.currentPort.length++
			s.currentPort.isDimension = true
			s.currentPort.portType = "output"
			s.dimensions = append(s.dimensions, strs[i][:strings.Index(strs[i], "[")])
		} else {
			s.currentPort.id = strs[i]
			if strings.Contains(ctx.GetText(), "[") {
				s.currentPort.length, _ = strconv.Atoi(ctx.Range_().GetText()[strings.Index(ctx.Range_().GetText(), "[")+1 : strings.Index(ctx.Range_().GetText(), "[")+2])
				s.currentPort.length++
			} else {
				s.currentPort.length = 1
			}
			s.currentPort.isDimension = false
			s.currentPort.portType = "output"
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
			s.currentPort.length, _ = strconv.Atoi(ctx.Range_().GetText()[strings.Index(ctx.Range_().GetText(), "[")+1 : strings.Index(ctx.Range_().GetText(), "[")+2])
			s.currentPort.length++
			s.currentPort.isDimension = true
			s.currentPort.portType = "reg"
			s.dimensions = append(s.dimensions, strs[i][:strings.Index(strs[i], "[")])
		} else {
			s.currentPort.id = strs[i]
			if strings.Contains(ctx.GetText(), "[") {
				s.currentPort.length, _ = strconv.Atoi(ctx.Range_().GetText()[strings.Index(ctx.Range_().GetText(), "[")+1 : strings.Index(ctx.Range_().GetText(), "[")+2])
				s.currentPort.length++
			} else {
				s.currentPort.length = 1
			}
			s.currentPort.isDimension = false
			s.currentPort.portType = "reg"
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
			s.currentPort.length, _ = strconv.Atoi(ctx.Range_().GetText()[strings.Index(ctx.Range_().GetText(), "[")+1 : strings.Index(ctx.Range_().GetText(), "[")+2])
			s.currentPort.length++
			s.currentPort.isDimension = true
			s.currentPort.portType = "net"
			s.dimensions = append(s.dimensions, strs[i][:strings.Index(strs[i], "[")])
		} else {
			s.currentPort.id = strs[i]
			if strings.Contains(ctx.GetText(), "[") {
				s.currentPort.length, _ = strconv.Atoi(ctx.Range_().GetText()[strings.Index(ctx.Range_().GetText(), "[")+1 : strings.Index(ctx.Range_().GetText(), "[")+2])
				s.currentPort.length++
			} else {
				s.currentPort.length = 1
			}
			s.currentPort.isDimension = false
			s.currentPort.portType = "net"
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
		initiation: expression.CompileExpression(ctx.Constant_expression().GetText(), moduleName, s.dimensions),
	})
}

// ExitRange_ is called when production range_ is exited.
func (s *CustomVerilogListener) ExitRange_(ctx *parser.Range_Context) {
	//Convert string to int
	s.currentPort.length, _ = strconv.Atoi(ctx.Msb_constant_expression().GetText())
}

// ExitNet_assignment is called when production net_assignment is exited.
func (s *CustomVerilogListener) ExitNet_assignment(ctx *parser.Net_assignmentContext) {
	builder.CreateAssign(ctx.Net_lvalue().GetText(), expression.CompileExpression(ctx.Expression().GetText(), moduleName, s.dimensions))
}
