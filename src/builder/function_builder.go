package builder

import (
	"strings"
)

func CreateCase(exp string, statement string) {
	Function += InputIndent(1) + "case " + exp + ":\n"
	terms := strings.Split(statement, "=")
	Function += InputIndent(2) + terms[0] + ".Set(" + toInt(terms[1][:len(terms[1])-1]) + ")\n"
	Function += InputIndent(2) + "break\n"
}
