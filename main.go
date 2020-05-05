package main

import (
	"fmt"
	"io/ioutil"

	parser "github.com/verilog2Go/generated/verilog"
	verilog "github.com/verilog2Go/src/verilog2Go"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

func main() {
	//ファイル読み取り
	bytes, err := ioutil.ReadFile("./src/examples/adder.v")
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
	listener := verilog.NewVerilogListener()
	p.AddParseListener(listener)

	tree := p.Source_text()
	fmt.Println(tree.ToStringTree([]string{}, p))
}
