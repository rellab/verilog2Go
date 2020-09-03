package builder

import "strconv"

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
	Always += leftBlock + "}\n"
}

func IfStart() {
	ifBlock += "{\n"
}

func IfStatement(conditionalStatement string) {
	// ifBlock = InputIndent(1) + "if " + conditionalStatement + ifBlock
	// Always += ifBlock
	leftBlock += InputIndent(IfDepth) + "if " + conditionalStatement + " {\n"
}

func ElseStatement() {
	leftBlock += InputIndent(IfDepth) + "else{\n"
}

func EndIfStatement() {
	leftBlock += InputIndent(IfDepth) + "}\n"
}

func DeclarateVariable(expression string, count int) {
	nonBlockingStatementCount++
	Always += InputIndent(1) + "var" + strconv.Itoa(count) + " = variable.InitBitArray(8)\n"
	leftBlock += InputIndent(IfDepth+1) + expression + "\n"
	//ノンブロッキング代入の右辺を格納しておく変数を宣言
	//var varable1 variable.BitArray
}
