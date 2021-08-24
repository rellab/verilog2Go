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

func AddPosedgeObserver(id string) {
	Observer += InputIndent(1) + "args." + id + ".AddPosedgeObserver(args.PreAlways" + strconv.Itoa(AlwaysCounter) + ", args.Always" + strconv.Itoa(AlwaysCounter) + ", args.Exec)\n"
}

func AddNegedgeObserver(id string) {
	Observer += InputIndent(1) + "args." + id + ".AddNegedgeObserver(args.PreAlways" + strconv.Itoa(AlwaysCounter) + ", args.Always" + strconv.Itoa(AlwaysCounter) + ", args.Exec)\n"
}

func CreateAlways() {
	AlwaysCounter++
	PreAlways += "func (" + ModuleName + " *" + ModuleName + ") PreAlways" + strconv.Itoa(AlwaysCounter) + "() []variable.BitArray{\n"
	Always += "func (" + ModuleName + " *" + ModuleName + ") Always" + strconv.Itoa(AlwaysCounter) + "(vars []variable.BitArray){\n"
	nonBlockingStatementCount = 0
	declarateInput()
}

func EndAlways() {

	PreAlways += leftBlock + createPreAlwaysReturn() + "}\n"

	Always += rightBlock + "}\n"

	leftBlock = ""
	rightBlock = ""
}

func IfStart() {
	ifBlock += "{\n"
}

func IfStatement(conditionalStatement string) {
	leftConditionalStatement := checkInputSignal(conditionalStatement)
	rightConditionalStatement := checkInput(conditionalStatement)
	leftBlock += InputIndent(IfDepth) + "if variable.CheckBit(" + leftConditionalStatement + ") {\n"
	rightBlock += InputIndent(IfDepth) + "if variable.CheckBit(" + rightConditionalStatement + ") {\n"
}

// ifの条件内にinput信号があれば変数に変換する
func checkInputSignal(conditionalStatement string) string {
	for i, v := range Inputs {
		if strings.Contains(conditionalStatement, ModuleName+"."+v.id) {
			conditionalStatement = conditionalStatement[:strings.Index(conditionalStatement, ModuleName+"."+v.id)-1] + "var" + strconv.Itoa(i+1) + ")"
		}
	}
	return conditionalStatement
}

func checkInput(conditionalStatement string) string {
	for i, v := range Inputs {
		if strings.Contains(conditionalStatement, ModuleName+"."+v.id) {
			conditionalStatement = conditionalStatement[:strings.Index(conditionalStatement, ModuleName+"."+v.id)-1] + "vars[" + strconv.Itoa(i) + "])"
		}
	}
	return conditionalStatement
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

func declarateInput() {
	for _, v := range Inputs {
		nonBlockingStatementCount++
		temp := "var" + strconv.Itoa(nonBlockingStatementCount)
		PreAlways += InputIndent(1) + temp + " := *variable.CreateBitArray(" + strconv.Itoa(v.length) + ", " + ModuleName + "." + v.id + ".ToInt())\n"
	}
}

func DeclarateVariable(exp string, dimensions []string) {
	nonBlockingStatementCount++
	//一次格納する変数
	temp := "var" + strconv.Itoa(nonBlockingStatementCount)
	alwaysTemp := "vars[" + strconv.Itoa(nonBlockingStatementCount-1) + "]"
	slice := strings.Split(exp, "<=")

	PreAlways += InputIndent(1) + temp + " := *variable.CreateBitArray(8, 0)\n"
	right := expression.CompileExpression(slice[1], ModuleName, dimensions)
	if (strings.Contains(right, "Get(") && len(right) < 20) || (strings.Contains(right, "CreateBitArray(") && len(right) < 31) || !(strings.Contains(right, "(")) {
		right = "*" + right
	}
	leftBlock += InputIndent(IfDepth+1) + temp + ".Assign(" + right + ")\n"
	rightBlock += InputIndent(IfDepth+1) + ModuleName + "." + slice[0] + ".Assign(" + alwaysTemp + ")\n"
}

func createPreAlwaysReturn() string {
	result := InputIndent(1) + "return []variable.BitArray{"
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
