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

func AddPosedgeObserver(id string) {
	Observer += InputIndent(1) + id + ".AddPosedgeObserver(p)\n"
}

func AddNegedgeObserver(id string) {
	Observer += InputIndent(1) + id + ".AddNegedgeObserver(p)\n"
}

func CreateAlways() {
	PreAlways += "func (" + ModuleName + " *" + ModuleName + ") PreAlways() []variable.BitArray{\n"
	Always += "func (" + ModuleName + " *" + ModuleName + ") Always(vars []variable.BitArray){\n"
}

func EndAlways() {

	PreAlways += leftBlock + createPreAlwaysReturn() + "}\n"

	Always += rightBlock
	Always += InputIndent(1) + ModuleName + ".Exec()\n}\n"
}

func IfStart() {
	ifBlock += "{\n"
}

func IfStatement(conditionalStatement string) {
	leftBlock += InputIndent(IfDepth) + "if " + conditionalStatement + " {\n"
	rightBlock += InputIndent(IfDepth) + "if " + conditionalStatement + " {\n"
}

func ElseStatement() {
	//直前の改行コードを削除
	leftBlock = leftBlock[:len(leftBlock)-1] + " else{\n"
	rightBlock = rightBlock[:len(rightBlock)-1] + " else{\n"
}

func EndIfStatement() {
	leftBlock += InputIndent(IfDepth) + "}\n"
	rightBlock += InputIndent(IfDepth) + "}\n"
}

func DeclarateVariable(exp string) {
	nonBlockingStatementCount++
	//一次格納する変数
	temp := "var" + strconv.Itoa(nonBlockingStatementCount)
	alwaysTemp := "vars[" + strconv.Itoa(nonBlockingStatementCount-1) + "]"
	slice := strings.Split(exp, "<=")

	PreAlways += InputIndent(1) + temp + " := *variable.CreateBitArray(8, 0)\n"
	right := expression.CompileExpression(slice[1], ModuleName)
	if strings.Contains(right[len(right)-7:len(right)], "Get(") {
		right = "*" + right
	} else if len(right) > 20 {
		if strings.Contains(right[len(right)-20:len(right)], "CreateBitArray(") {
			right = "*" + right
		}
	}
	leftBlock += InputIndent(IfDepth+1) + temp + ".Assign(" + right + ")\n"
	rightBlock += InputIndent(IfDepth+1) + ModuleName + "." + slice[0] + ".Assign(" + alwaysTemp + ")\n"
}

func createPreAlwaysReturn() string {
	result := InputIndent(1) + "return []variable.BitArray{"
	for i := 1; i <= nonBlockingStatementCount; i++ {
		result += "var" + strconv.Itoa(nonBlockingStatementCount)
		if i != nonBlockingStatementCount {
			result += ", "
		} else {
			result += "}\n"
		}
	}
	return result
}
