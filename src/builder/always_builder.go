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
	// b.always.WriteString(switchStatement)
	b.always.WriteString(rightBlock + "}\n")

	leftBlock = ""
	rightBlock = ""
}

func (b *Builder) IfStatement(conditionalStatement string) {
	rightConditionalStatement := b.checkInput(conditionalStatement)
	leftBlock += "if variable.CheckBit(" + conditionalStatement + ") {\n"
	rightBlock += "if variable.CheckBit(" + rightConditionalStatement + ") {\n"
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

func (b *Builder) CreateNonBlocking(lvalue string, exp string, dimensions []string) {
	nonBlockingStatementCount++
	//Variables for primary storage
	temp := "var" + strconv.Itoa(nonBlockingStatementCount)
	alwaysTemp := "vars[" + strconv.Itoa(nonBlockingStatementCount-1) + "]"

	b.preAlways.WriteString(temp + " := *variable.CreateBitArray(8, 0)\n")
	right := expression.CompileExpression(exp, strings.Title(moduleName), dimensions)
	lvalue = expression.CompileExpression(lvalue, strings.Title(moduleName), dimensions)
	if right[0] != '*' {
		if (strings.Contains(right, "Get(") && len(right) < 20) || (strings.Contains(right, "CreateBitArray(") && len(right) < 31) || !(strings.Contains(right[:len(right)-5], "(")) {
			right = "*" + right
		}
	}
	if lvalue[0] == '*' {
		lvalue = lvalue[1:]
	}
	if isCase {
		caseStatement += lvalue + ".Assign(" + right + ")\n"
	} else if isAlways {
		leftBlock += temp + ".Assign(" + right + ")\n"
		rightBlock += lvalue + ".Assign(" + alwaysTemp + ")\n"
	}
}

func (b *Builder) CreateBlocking(lvalue string, exp string, dimensions []string) {
	lvalue = expression.CompileExpression(lvalue, strings.Title(moduleName), dimensions)
	if lvalue[0] == '*' {
		lvalue = lvalue[1:]
	}
	if isLoop {
		loopStatement += lvalue + ".Assign(" + expression.CompileExpression(exp, strings.Title(moduleName), dimensions) + ")\n"
	} else if isAlways {
		rightBlock += lvalue + ".Assign(" + expression.CompileExpression(exp, strings.Title(moduleName), dimensions) + ")\n"
	}
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

func (b *Builder) CreateAlwaysCase(exp string) {
	if caseStatement != "" {
		cases += "case " + exp + ":\n" + caseStatement[:len(caseStatement)] + "\n"
	} else {
		cases += "case " + exp + ":\n"
	}
	caseStatement = ""
}

func (b *Builder) CreateDefault() {
	if caseStatement != "" {
		cases += "default:\n" + caseStatement[:len(caseStatement)] + "\n"
	} else {
		cases += "default:\n"
	}
	caseStatement = ""
}

func (b *Builder) CreateLoop(exp1 string, exp2 string, exp3 string) {
	loopStatement = "for " + exp1 + ";" + exp2 + ";" + exp3 + "{\n" + loopStatement
	loopStatement += "}\n"
}
