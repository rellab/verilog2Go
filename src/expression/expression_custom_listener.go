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
		PushValue(ToInt(ctx.GetId().GetText()))
		break
	//()内の演算
	case 14:
		return
	//メソッド
	case 15:
		CreateFunction()
		break
	case 16:
		PushValue(ToInt(ctx.GetText()))
		break
	//数字
	case 17:
		AddValue(ToInt(ctx.GetId().GetText()))
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
		// var value int64
		// length := values[0]
		// switch values[1][0] {
		// case 'b':
		// 	// strValue = fmt.Sprintf("%b", value)
		// 	value, _ = strconv.ParseInt(values[1][1:], 2, 64)
		// 	break
		// case 'o':
		// 	value, _ = strconv.ParseInt(values[1][1:], 8, 64)
		// 	break
		// case 'd':
		// 	value, _ = strconv.ParseInt(values[1][1:], 10, 64)
		// 	break
		// case 'h':
		// 	value, _ = strconv.ParseInt(values[1][1:], 16, 64)
		// 	break
		// default:
		// 	break
		// }
		// result = "variable.CreateBitArray(" + length + ", " + strconv.FormatInt(value, 10) + ")"
		result = "variable.CreateBits(\"" + str + "\")"
	}
	return
}
