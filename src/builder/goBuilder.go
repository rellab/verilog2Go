package builder

import (
	"strconv"
	"strings"
)

var ModuleName string
var Ports string
var Inputs []Port
var Constructor string
var SubConstructor string
var Observer string
var Exec string
var PreAlways string
var Always string
var Function string
var Source string

// StartModule はモジュールの初期化を行う
func StartModule(moduleName string) {
	ModuleName = strings.Title(moduleName)
	// ModuleName = moduleName
	Exec = ""
}

// EndModule はモジュール内の要素を一つにまとめる
func EndModule() {
	Source = "package generated\n\n"
	Source += "import \"github.com/verilog2Go/src/variable\"\n\n"
	//ポートの宣言
	Source += Ports + "\n"
	//コンストラクタ
	Source += Constructor
	// 並列処理のコンストラクタ
	Source += SubConstructor
	//Exec
	Exec = "func (" + strings.ToLower(ModuleName) + " *" + ModuleName + ") Exec() {\n" + Exec + "}\n\n"
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
	dimensions := ""
	for i := 0; i < len(ports); i++ {
		if !ports[i].isDimension {
			Ports += ports[i].id + ", "
		} else {
			dimensions += ports[i].id + ", "
		}
	}
	Ports = Ports[:len(Ports)-2]
	Ports += " *variable.BitArray\n"
	if dimensions != "" {
		dimensions = dimensions[:len(dimensions)-2]
		Ports += InputIndent(1) + dimensions + " []*variable.BitArray\n"
	}
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
	ConstructorArgument := "func New" + strings.Title(funcName) + "(args *" + ModuleName + ") " + ModuleName + "{\n"

	Constructor = ConstructorArgument + Observer

	if len(params) > 0 {
		for _, v := range params {
			Constructor += InputIndent(1) + "args." + v.id + " = " + v.initiation + "\n"
		}
	}

	Constructor += InputIndent(1) + "return *args\n}\n\n"

	SubConstructor = "func NewGoroutine" + strings.Title(funcName) + " (in []chan variable.BitArray, out []chan variable.BitArray) *" + ModuleName + "{\n"
	SubConstructor += InputIndent(1) + ModuleName + " := &" + strings.Title(ModuleName) + "{"
	if len(ports) > 0 {
		for _, v := range ports {
			SubConstructor += "variable.NewBitArray(" + strconv.Itoa(v.length) + "), "
		}
		SubConstructor = SubConstructor[:len(SubConstructor)-2] + "}\n"
	}
	SubConstructor += InputIndent(1) + ModuleName + ".run(in, out)\n"
	SubConstructor += InputIndent(1) + "return " + ModuleName + "\n}\n\n"
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
