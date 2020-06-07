package expression

import (
	"strings"

	parser "github.com/verilog2Go/antlr/expression"
)

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

// ExitExpression is called when production expression is entered.
func (s *CustomExpressionListener) ExitExpression(ctx *parser.ExpressionContext) {
	switch ctx.GetNumber() {
	//単項演算
	case 1:
		OperateUniary(ctx.GetOp().GetText())
		break
	//値
	case 13:
		// fmt.Println(ctx.GetId().GetText())
		PushValue(ctx.GetId().GetText())
		break
	//()内の演算
	case 14:
		return
	//メソッド
	case 15:
		PushValue(ctx.GetId().GetText())
		break
	case 16:
		PushValue(ToInt(ctx.GetText()))
		break
	default:
		OperateBinary(ctx.GetOp().GetText())
	}
}

func ToInt(str string) (result string) {
	ports := strings.Split(str, "[")
	//ポート名[ビット]
	result = str
	if len(ports) > 1 {
		result = ports[0] + ".get(" + strings.TrimRight(ports[1], "]") + ")"
	}
	return
}
