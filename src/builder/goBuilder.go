package builder

import (
	"strings"
)

var ModuleName string
var Ports string
var Constructor string
var Exec string
var Source string

// StartModule はモジュールの初期化を行う
func StartModule(moduleName string) {
	ModuleName = moduleName
	Ports = "var "
}

// EndModule はモジュール内の要素を一つにまとめる
func EndModule() {
	Source = "package generated\n\n"
	Source += "import \"github.com/verilog2Go/src/builder/variable\"\n\n"
	//ポートの宣言
	Source += Ports + "\n"
	//コンストラクタ
	Source += Constructor + "\n"
	//Exec
	Exec = "func Exec() {\n" + Exec + "}\n"
	Source += Exec
}

//DeclarePorts はポートの宣言を行う
func DeclarePorts(ports []Port) {
	for i := 0; i < len(ports)-1; i++ {
		Ports += ports[i].id + ", "
	}
	Ports += ports[len(ports)-1].id + " variable.BitArray\n"
}

// CreateConstructor はコンストラクタを生成する
func CreateConstructor(funcName string, ports []Port) {
	//コンストラクタの１行目
	ConstructorArgument := "func " + strings.Title(funcName) + "("
	for i := 0; i < len(ports)-1; i++ {
		inputID := "input" + strings.Title(ports[i].id) //引数名
		ConstructorArgument += inputID + " variable.BitArray, "
		Constructor += inputIndent(1) + ports[i].id + ".Set(" + inputID + ".ToInt())\n"
	}
	inputID := "input" + strings.Title(ports[len(ports)-1].id) //引数名
	ConstructorArgument += inputID + " variable.BitArray) {\n"
	Constructor += inputIndent(1) + ports[len(ports)-1].id + ".Set(" + inputID + ".ToInt())\n"
	Constructor = ConstructorArgument + Constructor + "}\n"
}

// CreateExec はExecを生成する
func CreateExec(expression string) {
	Exec += inputIndent(1) + expression + "\n"
}

func inputIndent(indent int) string {
	var result string
	for i := 0; i < indent*4; i++ {
		result += " "
	}
	return result
}
