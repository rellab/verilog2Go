package builder

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/verilog2Go/src/expression"
)

var nonBlockingStatementCount int
var ifBlock string
var leftBlock string
var rightBlock string

func AddPosedgeObserver(id string) {
	Observer += InputIndent(1) + "p." + id + ".AddPosedgeObserver(p)\n"
}

func AddNegedgeObserver(id string) {
	Observer += InputIndent(1) + "p." + id + ".AddNegedgeObserver(p)\n"
}

func CreateAlways() {
	Always += "func (" + ModuleName + " " + ModuleName + ") Always(){\n"
}

func EndAlways() {
	Always += leftBlock + rightBlock + "}\n"
}

func IfStart() {
	ifBlock += "{\n"
}

func IfStatement(conditionalStatement string) {
	leftBlock += InputIndent(IfDepth) + "if " + conditionalStatement + " {\n"
	rightBlock += InputIndent(IfDepth) + "if " + conditionalStatement + " {\n"
}

func ElseStatement() {
	leftBlock += InputIndent(IfDepth) + "else{\n"
	rightBlock += InputIndent(IfDepth) + "else{\n"
}

func EndIfStatement() {
	leftBlock += InputIndent(IfDepth) + "}\n"
	rightBlock += InputIndent(IfDepth) + "}\n"
}

func DeclarateVariable(exp string) {
	nonBlockingStatementCount++
	//一次格納する変数
	temp := "var" + strconv.Itoa(nonBlockingStatementCount)
	slice := strings.Split(exp, "<=")

	Always += InputIndent(1) + temp + " := variable.InitBitArray(8)\n"
	fmt.Println(slice[0])
	leftBlock += InputIndent(IfDepth+1) + temp + ".Assign(" + expression.CompileExpression(slice[1], "") + ")\n"
	rightBlock += InputIndent(IfDepth+1) + slice[0] + ".Assign(" + temp + ")\n"
	//ノンブロッキング代入の右辺を格納しておく変数を宣言
	//var varable1 variable.BitArray
}
