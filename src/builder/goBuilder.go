package builder

import (
	"fmt"
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
	Ports = "type " + moduleName + " struct{\n" + inputIndent(1)
}

// EndModule はモジュール内の要素を一つにまとめる
func EndModule() {
	Source = "package generated\n\n"
	Source += "import \"github.com/verilog2Go/src/builder/variable\"\n\n"
	//ポートの宣言
	Ports += "}\n"
	Source += Ports + "\n"
	//コンストラクタ
	Source += Constructor + "\n"
	//Exec
	Exec = "func (" + ModuleName + "*" + ModuleName + ") Exec() {\n" + Exec + "}\n"
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
	Constructor = inputIndent(1) + "p := new(" + ModuleName + ")\n"
	for i := 0; i < len(ports)-1; i++ {
		id := ports[i].id //変数名
		ConstructorArgument += id + " variable.BitArray, "
		Constructor += inputIndent(1) + "p." + id + " = " + id + "\n"
	}
	id := ports[len(ports)-1].id //変数名
	ConstructorArgument += id + " variable.BitArray) *" + ModuleName + " {\n"
	Constructor += inputIndent(1) + "p." + id + " = " + id + "\n"
	Constructor = ConstructorArgument + Constructor + inputIndent(1) + "return p\n}\n"
}

// CreateExec はExecを生成する
func CreateExec(id string, expression string) {
	Exec += inputIndent(1) + ModuleName + "." + id + ".Assign(" + expression + ")\n"
}

//
func CreateInstance(instance Instance) {
	fmt.Println(instance.instanceName)
}

func inputIndent(indent int) string {
	var result string
	for i := 0; i < indent*4; i++ {
		result += " "
	}
	return result
}
