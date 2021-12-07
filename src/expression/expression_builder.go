package expression

import (
	"strings"

	"github.com/golang-collections/collections/stack"
)

// Builder はスタックを定義する
type Builder struct {
	stack *stack.Stack
}

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
		expression.Push(ModuleName + "." + id)
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
		break
	case "|":
		result = value + ".Reductionor()"
		break
	case "!":
		result = value + ".Not()"
		break
	case "&":
		result = value + ".Baot()"
		break
	case "-":
		result = value + ".Neg()"
		break
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
		right = "*" + right
	}
	var result string = ""
	switch op {
	case "+":
		result = left + ".Add(" + right + ")"
		break
	case "-":
		result = left + ".Sub(" + right + ")"
		break
	case "*":
		result = left + ".Mul(" + right + ")"
		break
	case "/":
		result = left + ".Div(" + right + ")"
		break
	case "<<":
		result = left + ".SHL(" + right + ")"
		break
	case ">>":
		result = left + ".SHR(" + right + ")"
		break
	case "&&":
		result = left + ".And(" + right + ")"
		break
	case "||":
		result = left + ".Or(" + right + ")"
		break
	case "==":
		result = left + ".Equal(" + right + ")"
		break
	case "!=":
		result = left + ".NE(" + right + ")"
		break
	case ">=":
		result = left + ".GE(" + right + ")"
		break
	case "<=":
		result = left + ".LE(" + right + ")"
		break
	case ">":
		result = left + ".GT(" + right + ")"
		break
	case "<":
		result = left + ".LT(" + right + ")"
		break
	case "^":
		result = left + ".Bitxor(" + right + ")"
		break
	case "|":
		result = left + ".Bitor(" + right + ")"
		break
	case "&":
		result = left + ".Bitand(" + right + ")"
		break
	case "=":
		result = left + ".Set(" + right + ")"
		break
	default:
		break

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
	expression.Push(funcName + "(" + input + ")")
}
