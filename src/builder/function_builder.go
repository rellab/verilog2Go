package builder

import (
	"strings"
)

var switchStatement, cases string

func CreateCase(exp string, statement string) {
	cases += InputIndent(1) + "case " + exp + ":\n"
	terms := strings.Split(statement, "=")
	cases += InputIndent(2) + terms[0] + ".Set(" + toInt(terms[1][:len(terms[1])-1]) + ")\n"
	cases += InputIndent(2) + "break\n"
}

func CreateSwitch(input string) {
	switchStatement += InputIndent(1) + "switch " + input + ".ToInt()" + "{\n"
	switchStatement += cases
	switchStatement += InputIndent(1) + "}\n"
}

func CreateFunction(funcName string, input string) {
	Function += "\nfunc (" + ModuleName + " *" + ModuleName + ") " + funcName + "(" + input + " variable.BitArray) variable.BitArray {\n"
	Function += switchStatement
	Function += InputIndent(1) + "return " + funcName + "\n}"
}
