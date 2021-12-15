package expression

import (
	"strings"

	parser "github.com/verilog2Go/antlr/expression"
)

var expCount int

// CustomExpressionListener はBaseExpressionListenerを構造体として宣言
type CustomExpressionListener struct {
	*parser.BaseExpressionListener
}

// NewExpressionListener はリスナーの初期化を行う
func NewExpressionListener() *CustomExpressionListener {
	return &CustomExpressionListener{}
}

func (s *CustomExpressionListener) EnterStart(ctx *parser.StartContext) {
	InitializeStack()
}

// ExitStart is called when production start is exited.
func (s *CustomExpressionListener) ExitStart(ctx *parser.StartContext) {
	expCount = 0
}

// EnterExpression is called when production expression is entered.
func (s *CustomExpressionListener) EnterExpression(ctx *parser.ExpressionContext) {
	expCount++
}

// ExitExpression is called when production expression is entered.
func (s *CustomExpressionListener) ExitExpression(ctx *parser.ExpressionContext) {
	switch ctx.GetNumber() {
	//単項演算
	case 1:
		OperateUniary(ctx.GetOp().GetText())
	//値
	case 13:
		PushValue(ToInt(ctx.GetId().GetText()))
		//()内の演算
	case 14:
		return
		//メソッド
	case 15:
		CreateFunction()
	case 16:
		// PushValue(ToInt(ctx.GetText()))
		CreateDimension(ctx.GetId().GetText())
	//数字
	case 17:
		AddValue(ToInt(ctx.GetId().GetText()))
	case 18:
		AddValue(ToInt(ctx.GetText()))
	default:
		OperateBinary(ctx.GetOp().GetText())
	}
}

func ToInt(str string) (result string) {
	ports := strings.Split(str, "[")
	values := strings.Split(str, "'")
	//ポート名[ビット]
	result = str
	if len(ports) > 1 {
		for _, v := range Dimensions {
			if v == ports[0] {
				result = str
				return
			}
		}
		result = ports[0] + ".Get(" + strings.TrimRight(ports[1], "]") + ")"
	} else if len(values) > 1 {
		if expCount == 1 {
			result = "*variable.CreateBits(\"" + str + "\")"
		} else {
			result = "variable.CreateBits(\"" + str + "\")"
		}
	}
	return
}
