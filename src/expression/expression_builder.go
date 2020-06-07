package expression

import (
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
	expression.Push(id)
}

// OperateUniary は単項演算を行う
func OperateUniary(op string) {
	var value string = expression.Pop().(string)
	var result string = ""
	switch op {
	case "~":
		result = value + "not()"
		break
	case "|":
		result = value + "reductionor()"
		break
	default:

	}
	expression.Push(result)
}

// OperateBinary は２項演算を行う
func OperateBinary(op string) {
	left := expression.Pop().(string)
	right := expression.Pop().(string)
	var result string = ""
	switch op {
	case "+":
		result = left + ".add(" + right + ")"
		break
	case "-":
		result = left + ".sub(" + right + ")"
		break
	case "*":
		result = left + ".mul(" + right + ")"
		break
	case "/":
		result = left + ".div(" + right + ")"
		break
	case "&&":
		result = left + ".and(" + right + ")"
		break
	case "==":
		result = left + ".equal(" + right + ")"
		break
	case "^":
		result = left + ".bitxor(" + right + ")"
		break
	case "|":
		result = left + ".bitor(" + right + ")"
		break
	case "&":
		result = left + ".bitand(" + right + ")"
		break
	default:
		break

	}
	expression.Push(result)
}
