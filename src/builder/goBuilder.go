package builder

import (
	"fmt"
	"strings"
)

var ModuleName string
var Ports string
var Constructor string
var Observer string
var Exec string
var Always string
var Source string

// StartModule はモジュールの初期化を行う
func StartModule(moduleName string) {
	ModuleName = moduleName
	Ports = "type " + moduleName + " struct{\n" + InputIndent(1)
}

// EndModule はモジュール内の要素を一つにまとめる
func EndModule() {
	Source = "package generated\n\n"
	Source += "import \"github.com/verilog2Go/src/variable\"\n\n"
	//ポートの宣言
	Ports += "}\n"
	Source += Ports + "\n"
	//コンストラクタ
	Source += Constructor
	//Exec
	Exec = "func (" + ModuleName + " *" + ModuleName + ") Exec() {\n" + Exec + "}\n\n"
	Source += Exec
	//Always
	Source += Always
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
	Constructor = InputIndent(1) + "p := new(" + ModuleName + ")\n"
	for i := 0; i < len(ports)-1; i++ {
		id := ports[i].id //変数名
		ConstructorArgument += id + " *variable.BitArray, "
		Constructor += InputIndent(1) + "p." + id + " = *" + id + "\n"
	}
	id := ports[len(ports)-1].id //変数名
	ConstructorArgument += id + " *variable.BitArray) " + ModuleName + " {\n"
	Constructor += InputIndent(1) + "p." + id + " = *" + id + "\n"
	Constructor = ConstructorArgument + Constructor + Observer + InputIndent(1) + "return *p\n}\n\n"
}

// CreateExec はExecを生成する
func CreateExec(id string, expression string) {
	Exec += InputIndent(1) + ModuleName + "." + id + ".Assign(" + expression + ")\n"
}

// CreateInstance はインスタンス化を生成する
func CreateInstance(instance Instance) {
	fmt.Println(instance.instanceName)
	Exec += InputIndent(1) + instance.instanceName + " := " + strings.Title(instance.moduleName) + "("
	for i, exp := range instance.ports {
		Exec += exp
		if i == len(instance.ports)-1 {
			Exec += ")\n"
		} else {
			Exec += ", "
		}
	}
	Exec += InputIndent(1) + instance.instanceName + ".Exec()\n"
}

func InputIndent(indent int) string {
	var result string
	for i := 0; i < indent*4; i++ {
		result += " "
	}
	return result
}
