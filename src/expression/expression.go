package expression

import (
	parser "github.com/verilog2Go/antlr/expression"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

var ModuleName string

func CompileExpression(str string, moduleName string) string {
	ModuleName = moduleName
	input := antlr.NewInputStream(str)

	lexer := parser.NewExpressionLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewExpressionParser(stream)

	// リスナーをセットする
	listener := NewExpressionListener()
	p.AddParseListener(listener)

	tree := p.Start()
	tree.ToStringTree([]string{}, p)
	return expression.Pop().(string)
}
