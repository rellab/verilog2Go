package main

import (
	//#include <stdio.h>
	"C"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	parser "github.com/verilog2Go/antlr/verilog"
	"github.com/verilog2Go/src/builder"
)

//export createGo
func createGo(src *C.char) {
	//verilogソースの読み取り
	input := antlr.NewInputStream(C.GoString(src))
	//lexerの作成
	lexer := parser.NewVerilogLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	//parserの作成
	p := parser.NewVerilogParser(stream)

	// リスナーをセットする
	listener := builder.NewVerilogListener()
	p.AddParseListener(listener)

	tree := p.Source_text()
	tree.ToStringTree([]string{}, p)
	// fmt.Println(tree.ToStringTree([]string{}, p))

	builder.Write()
}

func main() {
	createGo(C.CString("module adder(a, b, q);input [3:0] a, b;output [3:0] q;assign q = a + b;endmodule"))
}
