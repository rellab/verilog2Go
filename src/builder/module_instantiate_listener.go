package builder

import (
	"strings"

	parser "github.com/verilog2Go/antlr/verilog"
	"github.com/verilog2Go/src/expression"
)

type Instance struct {
	moduleName   string
	instanceName string
	//Use ports for ordered_port and portMap for named_port
	ports   []string
	portMap map[string]string
}

var instance Instance

// EnterModule_instantiation is called when production module_instantiation is exited.
func (s *CustomVerilogListener) EnterModule_instantiation(ctx *parser.Module_instantiationContext) {
	//Instance initialization
	instance = Instance{}
	instance.portMap = map[string]string{}
}

// ExitModule_instantiation is called when production module_instantiation is exited.
func (s *CustomVerilogListener) ExitModule_instantiation(ctx *parser.Module_instantiationContext) {
	//Pass instance information to Builder
	instance.moduleName = ctx.Module_identifier().GetText()
	builder.CreateInstance(instance)
}

// ExitModule_instance is called when production module_instance is exited.
func (s *CustomVerilogListener) ExitModule_instance(ctx *parser.Module_instanceContext) {
	instance.instanceName = ctx.Name_of_instance().GetText()
}

// ExitOrdered_port_connection is called when production ordered_port_connection is exited.
func (s *CustomVerilogListener) ExitOrdered_port_connection(ctx *parser.Ordered_port_connectionContext) {
	expression := expression.CompileExpression(ctx.Expression().GetText(), strings.Title(moduleName), s.dimensions)
	// fmt.Println(expression)
	if !strings.Contains(expression, "Get(") && !strings.Contains(expression, "CreateBitArray(") {
		expression = "&" + expression
	}
	instance.ports = append(instance.ports, expression)
}

// ExitNamed_port_connection is called when production named_port_connection is entered.
func (s *CustomVerilogListener) ExitNamed_port_connection(ctx *parser.Named_port_connectionContext) {
	expression := expression.CompileExpression(ctx.Expression().GetText(), strings.Title(moduleName), s.dimensions)
	instance.portMap[ctx.Port_identifier().GetText()] = expression
}
