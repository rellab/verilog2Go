package main

import (
	"io/ioutil"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	parser "github.com/verilog2Go/antlr/verilog"
	"github.com/verilog2Go/src/builder"
)

func main() {
	//ファイル読み取り
	bytes, err := ioutil.ReadFile("./src/examples/saikoro.v")
	if err != nil {
		panic(err)
	}
	input := antlr.NewInputStream(string(bytes))
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

	builder.CreateNewFile()
	builder.Write()
}
