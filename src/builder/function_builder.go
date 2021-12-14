package builder

import (
	"strings"
)

var switchStatement, cases string

func CreateCase(exp string, statement string) {
	cases += "case " + exp + ":\n"
	terms := strings.Split(statement, "=")
	cases += terms[0] + ".Set(" + toInt(terms[1][:len(terms[1])-1]) + ")\n"
	cases += "break\n"
}

func CreateSwitch(input string) {
	switchStatement += "switch " + input + ".ToInt()" + "{\n"
	switchStatement += cases
	switchStatement += "}\n"
}

func (b *Builder) CreateFunction(funcName string, input string, length string) {
	b.function.WriteString("\nfunc (" + strings.Title(moduleName) + " *" + strings.Title(moduleName) + ") " + funcName + "(" + input + " variable.BitArray) variable.BitArray {\n")
	b.function.WriteString(funcName + " := *variable.CreateBitArray(" + length + ", 0)\n")
	b.function.WriteString(switchStatement)
	b.function.WriteString("return " + funcName + "\n}\n")
	cases = ""
	switchStatement = ""
}

func (b *Builder) CreateInitial() {

}

func (b *Builder) EndInitial() {

}
