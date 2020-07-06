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
	expression.Push(ModuleName + "." + id)
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
		result = value + "Not()"
		break
	case "|":
		result = value + "Reductionor()"
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
	case "&&":
		result = left + ".And(" + right + ")"
		break
	case "==":
		result = left + ".Equal(" + right + ")"
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
	default:
		break

	}
	expression.Push(result)
}
