package builder

import (
	"strings"
)

var ModuleName string
var Ports string
var Inputs []Port
var Constructor string
var Observer string
var Exec string
var PreAlways string
var Always string
var Function string
var Source string

// StartModule はモジュールの初期化を行う
func StartModule(moduleName string) {
	ModuleName = moduleName
}

// EndModule はモジュール内の要素を一つにまとめる
func EndModule() {
	Source = "package generated\n\n"
	Source += "import \"github.com/verilog2Go/src/variable\"\n\n"
	//ポートの宣言
	Source += Ports + "\n"
	//コンストラクタ
	Source += Constructor
	//Exec
	Exec = "func (" + ModuleName + " *" + ModuleName + ") Exec() {\n" + Exec + "}\n\n"
	Source += Exec
	//PreAlways
	Source += PreAlways
	//Always
	Source += Always
	//Function
	Source += Function
}

//DeclarePorts はポートの宣言を行う
func DeclarePorts(ports []Port) {
	if len(ports) < 1 {
		return
	}
	Ports = "type " + ModuleName + " struct{\n" + InputIndent(1)
	for i := 0; i < len(ports)-1; i++ {
		Ports += ports[i].id + ", "
	}
	Ports += ports[len(ports)-1].id + " *variable.BitArray\n"
	Ports += "}\n"
}

// DeclareInput はinput信号の配列を作成する
func DeclareInput(input Port) {
	Inputs = append(Inputs, input)
}

// CreateConstructor はコンストラクタを生成する
func CreateConstructor(funcName string, ports []Port, params []Param) {
	if len(ports) < 1 {
		return
	}
	//コンストラクタの１行目
	ConstructorArgument := "func " + strings.Title(funcName) + "(args *" + ModuleName + ") " + ModuleName + "{\n"

	Constructor = ConstructorArgument + Observer

	if len(params) > 0 {
		for _, v := range params {
			Constructor += InputIndent(1) + "args." + v.id + " = " + v.initiation + "\n"
		}
	}

	Constructor += InputIndent(1) + "return *args\n}\n\n"
}

// CreateExec はExecを生成する
func CreateExec(id string, expression string) {
	Exec += InputIndent(1) + ModuleName + "." + id + ".Assign(" + expression + ")\n"
}

// CreateInstance はインスタンス化を生成する
func CreateInstance(instance Instance) {
	// fmt.Println(instance.instanceName)
	Exec += InputIndent(1) + instance.instanceName + " := " + strings.Title(instance.moduleName) + "(&" + instance.moduleName + "{"
	for _, exp := range instance.ports {
		Exec += exp + ", "
	}

	for key, exp := range instance.portMap {
		Exec += key + " : " + exp + ", "
	}
	Exec = Exec[:len(Exec)-2]
	Exec += "})\n" + InputIndent(1) + instance.instanceName + ".Exec()\n"
}

func InputIndent(indent int) string {
	var result string
	for i := 0; i < indent*4; i++ {
		result += " "
	}
	return result
}
