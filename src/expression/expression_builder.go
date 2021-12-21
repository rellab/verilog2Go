package expression

import (
	"strconv"
	"strings"

	"github.com/golang-collections/collections/stack"
)

var expression *stack.Stack

// InitializeStack はスタックを初期化する
func InitializeStack() {
	expression = stack.New()
}

// GetExpression はスタックの先頭の文字列を返す
func GetExpression() string {
	return expression.Peek().(string)
}

// PushValue は値をスタックにプッシュする
func PushValue(id string) {
	if ModuleName == "" {
		expression.Push(id)
	} else {
		if expCount == 1 {
			expression.Push("*" + ModuleName + "." + id)
		} else {
			expression.Push(ModuleName + "." + id)
		}
	}
}

func AddValue(id string) {
	expression.Push(id)
}

// OperateUniary は単項演算を行う
func OperateUniary(op string) {
	var value string = expression.Pop().(string)
	var result string = ""
	switch op {
	case "~":
		result = value + ".Bnot()"
	case "|":
		result = value + ".Reductionor()"
	case "!":
		result = value + ".Not()"
	case "&":
		result = value + ".Baot()"
	case "-":
		result = value + ".Neg()"
	default:

	}
	expression.Push(result)
}

// OperateBinary は２項演算を行う
func OperateBinary(op string) {
	left := expression.Pop().(string)
	right := expression.Pop().(string)
	if !(strings.Contains(right, "Add(") ||
		strings.Contains(right, "Sub(") ||
		strings.Contains(right, "Mul(") ||
		strings.Contains(right, "Div(") ||
		strings.Contains(right, "And(") ||
		strings.Contains(right, "Equal(") ||
		strings.Contains(right, "BItxor(") ||
		strings.Contains(right, "Bitor(") ||
		strings.Contains(right, "Bitand(") ||
		strings.Contains(right, "Not(") ||
		strings.Contains(right, "Reductionor(") ||
		strings.Contains(right, "Set(")) {
		// right = "*" + right
	}
	var result string = ""
	switch op {
	case "+":
		result = left + ".Add(" + right + ")"
	case "-":
		result = left + ".Sub(" + right + ")"
	case "*":
		result = left + ".Mul(" + right + ")"
	case "/":
		result = left + ".Div(" + right + ")"
	case "<<":
		result = left + ".SHL(" + right + ")"
	case ">>":
		result = left + ".SHR(" + right + ")"
	case "&&":
		result = left + ".And(" + right + ")"
	case "||":
		result = left + ".Or(" + right + ")"
	case "==":
		result = left + ".Equal(" + right + ")"
	case "!=":
		result = left + ".NE(" + right + ")"
	case ">=":
		result = left + ".GE(" + right + ")"
	case "<=":
		result = left + ".LE(" + right + ")"
	case ">":
		result = left + ".GT(" + right + ")"
	case "<":
		result = left + ".LT(" + right + ")"
	case "^":
		result = left + ".Bitxor(" + right + ")"
	case "|":
		result = left + ".Bitor(" + right + ")"
	case "&":
		result = left + ".Bitand(" + right + ")"
	case "=":
		result = left + ".Set(" + right + ")"
	default:
	}
	expression.Push(result)
}

func CreateFunction() {
	input := expression.Pop().(string)
	if !(strings.Contains(input, "Add(") ||
		strings.Contains(input, "Sub(") ||
		strings.Contains(input, "Mul(") ||
		strings.Contains(input, "Div(") ||
		strings.Contains(input, "And(") ||
		strings.Contains(input, "Equal(") ||
		strings.Contains(input, "BItxor(") ||
		strings.Contains(input, "Bitor(") ||
		strings.Contains(input, "Bitand(") ||
		strings.Contains(input, "Not(") ||
		strings.Contains(input, "Reductionor(") ||
		strings.Contains(input, "Set(")) {
		input = "*" + input
	}
	funcName := expression.Pop().(string)
	if funcName[0] == '*' {
		funcName = funcName[1:]
	}
	expression.Push(funcName + "(" + input + ")")
}

func CreateDimension(id string) {
	index := expression.Pop().(string)
	_, ok := strconv.Atoi(index)
	if ok == nil {
		if ModuleName == "" {
			expression.Push(id + "[" + index + "]")
		} else {
			expression.Push("*" + ModuleName + "." + id + "[" + index + "]")
		}
	} else {
		if ModuleName == "" {
			expression.Push(id + "[" + index + ".ToInt()]")
		} else {
			if expCount == 1 {
				expression.Push("*" + ModuleName + "." + id + "[" + index + ".ToInt()]")
			} else {
				expression.Push(ModuleName + "." + id + "[" + index + ".ToInt()]")
			}
		}
	}
}
