package builder

import (
	"fmt"
	"strings"
)

var ModuleName string
var Ports string
var Constructor string

func StartModule(moduleName string) {
	ModuleName = moduleName
	Ports = "var "
}

//ポートの宣言
func DeclarePorts(ports []Port) {
	for i := 0; i < len(ports)-1; i++ {
		Ports += ports[i].id + ", "
	}
	Ports += ports[len(ports)-1].id + " variable.BitArray"
	fmt.Println(Ports)
}

func CreateConstructor(funcName string, ports []Port) {
	ConstructorArgument := "func " + strings.Title(funcName) + "("
	for i := 0; i < len(ports)-1; i++ {
		inputID := "input" + strings.Title(ports[i].id)
		ConstructorArgument += inputID + " variable.BitArray, "
		Constructor += InputIndent(1) + ports[i].id + ".Set(" + inputID + ".ToInt())\n"
	}
	inputID := "input" + strings.Title(ports[len(ports)-1].id)
	ConstructorArgument += inputID + " variable.BitArray) {\n"
	Constructor += InputIndent(1) + ports[len(ports)-1].id + ".Set(" + inputID + ".ToInt())\n"
	Constructor = ConstructorArgument + Constructor + "}"
	fmt.Println(Constructor)
}

func InputIndent(indent int) string {
	var result string
	for i := 0; i < indent*4; i++ {
		result += " "
	}
	return result
}
