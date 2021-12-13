package builder

import (
	"strconv"
	"strings"

	"github.com/verilog2Go/src/expression"
)

var nonBlockingStatementCount int
var ifBlock string
var leftBlock string
var rightBlock string
var AlwaysCounter = 0

func (b *Builder) AddPosedgeObserver(id string) {
	b.observer.WriteString("args." + id + ".AddPosedgeObserver(args.PreAlways" + strconv.Itoa(AlwaysCounter) + ", args.Always" + strconv.Itoa(AlwaysCounter) + ", args.Exec)\n")
}

func (b *Builder) AddNegedgeObserver(id string) {
	b.observer.WriteString("args." + id + ".AddNegedgeObserver(args.PreAlways" + strconv.Itoa(AlwaysCounter) + ", args.Always" + strconv.Itoa(AlwaysCounter) + ", args.Exec)\n")
}

func (b *Builder) CreateAlways() {
	AlwaysCounter++
	b.preAlways.WriteString("func (" + strings.Title(moduleName) + " *" + strings.Title(moduleName) + ") PreAlways" + strconv.Itoa(AlwaysCounter) + "() []variable.BitArray{\n")
	b.always.WriteString("func (" + strings.Title(moduleName) + " *" + strings.Title(moduleName) + ") Always" + strconv.Itoa(AlwaysCounter) + "(vars []variable.BitArray){\n")
	nonBlockingStatementCount = 0
	b.declarateInput()
}

func (b *Builder) EndAlways() {

	b.preAlways.WriteString(leftBlock + b.createPreAlwaysReturn() + "}\n")
	b.always.WriteString(switchStatement)
	b.always.WriteString(rightBlock + "}\n")

	leftBlock = ""
	rightBlock = ""
}

func (b *Builder) IfStart() {
	ifBlock += "{\n"
}

func (b *Builder) IfStatement(conditionalStatement string) {
	leftConditionalStatement := b.checkInputSignal(conditionalStatement)
	rightConditionalStatement := b.checkInput(conditionalStatement)
	leftBlock += "if variable.CheckBit(" + leftConditionalStatement + ") {\n"
	rightBlock += "if variable.CheckBit(" + rightConditionalStatement + ") {\n"
}

// If there is an input signal in the if condition, convert it to a variable
func (b *Builder) checkInputSignal(conditionalStatement string) string {
	for _, v := range b.inputs {
		if strings.Contains(conditionalStatement, strings.Title(moduleName)+"."+v.id) {
			// conditionalStatement = conditionalStatement[:strings.Index(conditionalStatement, strings.Title(moduleName)+"."+v.id)-1] + "var" + strconv.Itoa(i+1) + ")"
		}
	}
	return conditionalStatement
}

func (b *Builder) checkInput(conditionalStatement string) string {
	for _, v := range b.inputs {
		if strings.Contains(conditionalStatement, strings.Title(moduleName)+"."+v.id) {
			// conditionalStatement = conditionalStatement[:strings.Index(conditionalStatement, strings.Title(moduleName)+"."+v.id)-1] + "vars[" + strconv.Itoa(i) + "])"
		}
	}
	return conditionalStatement
}

func (b *Builder) ElseStatement() {
	//Delete the previous line feed code
	leftBlock = leftBlock[:len(leftBlock)-1] + " else{\n"
	rightBlock = rightBlock[:len(rightBlock)-1] + " else{\n"
}

func (b *Builder) EndIfStatement() {
	leftBlock += "}\n"
	rightBlock += "}\n"
}

func (b *Builder) declarateInput() {
	for _, v := range b.inputs {
		nonBlockingStatementCount++
		temp := "var" + strconv.Itoa(nonBlockingStatementCount)
		b.preAlways.WriteString(temp + " := *variable.CreateBitArray(" + strconv.Itoa(v.length) + ", " + strings.Title(moduleName) + "." + v.id + ".ToInt())\n")
	}
}

func (b *Builder) DeclarateVariable(exp string, dimensions []string) {
	nonBlockingStatementCount++
	//Variables for primary storage
	temp := "var" + strconv.Itoa(nonBlockingStatementCount)
	alwaysTemp := "vars[" + strconv.Itoa(nonBlockingStatementCount-1) + "]"
	slice := strings.Split(exp, "<=")

	b.preAlways.WriteString(temp + " := *variable.CreateBitArray(8, 0)\n")
	right := expression.CompileExpression(slice[1], strings.Title(moduleName), dimensions)
	if (strings.Contains(right, "Get(") && len(right) < 20) || (strings.Contains(right, "CreateBitArray(") && len(right) < 31) || !(strings.Contains(right, "(")) {
		right = "*" + right
	}
	leftBlock += temp + ".Assign(" + right + ")\n"
	rightBlock += strings.Title(moduleName) + "." + slice[0] + ".Assign(" + alwaysTemp + ")\n"
}

func (b *Builder) createPreAlwaysReturn() string {
	result := "return []variable.BitArray{"
	for i := 1; i <= nonBlockingStatementCount; i++ {
		result += "var" + strconv.Itoa(i)
		if i != nonBlockingStatementCount {
			result += ", "
		} else {
			result += "}\n"
		}
	}
	return result
}

func (s *CustomVerilogListener) CreateAlwaysCase(exp string, statement string) {
	cases += "case " + exp + ":\n"
	index := strings.Index(statement, "=")
	cases += expression.CompileExpression(statement[:index], strings.Title(moduleName), s.dimensions) + ".Substitute(" + expression.CompileExpression(statement[index+1:len(statement)-1], strings.Title(moduleName), s.dimensions) + ")\n"
	cases += "break\n"
}

func (s *CustomVerilogListener) CreateDefault(statement string) {
	cases += "default:\n"
	index := strings.Index(statement, "=")
	cases += expression.CompileExpression(statement[:index], strings.Title(moduleName), s.dimensions) + ".Substitute(*" + expression.CompileExpression(statement[index+1:len(statement)-1], strings.Title(moduleName), s.dimensions) + ")\n"
	cases += "break\n"
}
