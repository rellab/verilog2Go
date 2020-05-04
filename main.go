package main

import (
	"fmt"

	parser "github.com/verilog2Go/generated/verilog"
	verilog "github.com/verilog2Go/src/verilog2Go"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

func main() {
	input := antlr.NewInputStream("a+b*c+ d[2] + e")

	lexer := parser.NewVerilogLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewVerilogParser(stream)

	// リスナーをセットする
	listener := verilog.NewVerilogListener()
	p.AddParseListener(listener)

	tree := p.Expression()
	fmt.Println(tree.ToStringTree([]string{}, p))
}
