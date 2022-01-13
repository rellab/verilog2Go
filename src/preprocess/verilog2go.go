package preprocess

import (
	//#include <stdio.h>
	// "C"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	parser "github.com/verilog2Go/antlr/verilog"
	"github.com/verilog2Go/src/builder"
)

func Run(sourceFile string) {
	str := Preprocess(sourceFile)
	// fmt.Println(str)
	input := antlr.NewInputStream(str)
	//create lexer
	lexer := parser.NewVerilogLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	//create parser
	p := parser.NewVerilogParser(stream)

	// set the listener
	listener := builder.NewVerilogListener()
	p.AddParseListener(listener)

	tree := p.Source_text()
	tree.ToStringTree([]string{}, p)
}
