package builder

var nonBlockingStatementCount int
var ifBlock string

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
	Always += "}\n"
}

func IfStart() {
	ifBlock += "{\n"
}

func IfStatement(conditionalStatement string) {
	// ifBlock = InputIndent(1) + "if " + conditionalStatement + ifBlock
	// Always += ifBlock
	Always += InputIndent(IfDepth) + "if " + conditionalStatement + " {\n"
}

func ElseStatement() {
	Always += InputIndent(IfDepth) + "else{\n"
}

func EndIfStatement() {
	Always += InputIndent(IfDepth) + "}\n"
}

func DeclarateVariable(expression string) {
	nonBlockingStatementCount++
	Always += InputIndent(IfDepth+1) + expression + "\n"
	//ノンブロッキング代入の右辺を格納しておく変数を宣言
	//var varable1 variable.BitArray
}
