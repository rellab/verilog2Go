package builder

import (
	"strconv"
	"strings"

	parser "github.com/verilog2Go/antlr/verilog"
)

// EnterFunction_declaration is called when production function_declaration is entered.
func (s *CustomVerilogListener) EnterFunction_declaration(ctx *parser.Function_declarationContext) {

}

// ExitFunction_declaration is called when production function_declaration is exited.
func (s *CustomVerilogListener) ExitFunction_declaration(ctx *parser.Function_declarationContext) {
	methodName := ctx.Function_identifier().GetText()
	terms := strings.Split(ctx.Function_item_declaration(0).GetText(), "]")
	input := terms[1][:len(terms[1])-1]
	terms = strings.Split(ctx.Range_or_type().GetText(), ":")
	length := terms[0][1:len(terms[0])]
	builder.CreateFunction(methodName, input, length)
}

// EnterFunction_case_statement is called when production function_case_statement is entered.
func (s *CustomVerilogListener) EnterFunction_case_statement(ctx *parser.Function_case_statementContext) {

}

// ExitFunction_case_statement is called when production function_case_statement is exited.
func (s *CustomVerilogListener) ExitFunction_case_statement(ctx *parser.Function_case_statementContext) {
	CreateSwitch(ctx.Expression().GetText())
}

// ExitFunction_case_item is called when production function_case_item is exited.
func (s *CustomVerilogListener) ExitFunction_case_item(ctx *parser.Function_case_itemContext) {
	if len(ctx.AllExpression()) != 0 {
		CreateCase(toInt(ctx.Expression(0).GetText()), ctx.Function_statement_or_null().GetText())
	}
}

// Convert conditional expression of case to int
func toInt(str string) (result string) {
	ports := strings.Split(str, "[")
	values := strings.Split(str, "'")
	//Portname [bit]s
	result = str
	if len(ports) > 1 {
		result = ports[0] + ".Get(" + strings.TrimRight(ports[1], "]") + ")"
	} else if len(values) > 1 {
		var value int64
		switch values[1][0] {
		case 'b':
			value, _ = strconv.ParseInt(values[1][1:], 2, 64)
			break
		case 'o':
			value, _ = strconv.ParseInt(values[1][1:], 8, 64)
			break
		case 'd':
			value, _ = strconv.ParseInt(values[1][1:], 10, 64)
			break
		case 'h':
			value, _ = strconv.ParseInt(values[1][1:], 16, 64)
			break
		default:
			break
		}
		result = strconv.FormatInt(value, 10)
	}
	return
}
