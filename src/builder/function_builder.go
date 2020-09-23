package builder

import (
	"strings"
)

var cases string

func CreateCase(exp string, statement string) {
	cases += InputIndent(1) + "case " + exp + ":\n"
	terms := strings.Split(statement, "=")
	cases += InputIndent(2) + terms[0] + ".Set(" + toInt(terms[1][:len(terms[1])-1]) + ")\n"
	cases += InputIndent(2) + "break\n"
}

func CreateFunction(funcName string, input string) {
	Function += "\nfunc (" + ModuleName + " *" + ModuleName + ") " + funcName + "(" + input + ") variable.BitArray {\n"
	Function += cases
	Function += InputIndent(1) + "return " + funcName + "\n}"
}
