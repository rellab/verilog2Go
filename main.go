package main

import (
	"fmt"

	"github.com/verilog2Go/parser"

	"github.com/verilog2Go/src/expression"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

func main() {
	input := antlr.NewInputStream("a+b*c+ d[2] + e")

	lexer := parser.NewExpressionLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewExpressionParser(stream)

	// リスナーをセットする
	listener := expression.NewExpressionListener()
	p.AddParseListener(listener)

	tree := p.Start()
	fmt.Println(tree.ToStringTree([]string{}, p))
}
