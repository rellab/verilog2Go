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
	// createGo(C.CString("module cnt_unit(CK, RES, EN, Q, CA);input CK, RES, EN;output Q, CA;always @(posedge CK or posedge RES)begin	if(RES == 1'b1)		Q <= 1'b0;	else if(EN == 1'b1)		Q <= ~Q;end assign CA = EN & Q;endmodule"))
	createGo(C.CString("module saikoro(ck, reset, enable, lamp);input ck, reset, enable;output [6:0] lamp;reg [2:0] cnt;always@(posedge ck or posedge reset) begin	if(reset == 1'b1)		cnt <= 3'h1;	else if(enable == 1'b1)		if(cnt == 3'h6)			cnt <= 3'h1;		else			cnt <= cnt + 3'h1;end function [6:0] dec;input [2:0] din;	case(din)		3'h1: dec = 7'b0001000;		3'h2: dec = 7'b1000001;		3'h3: dec = 7'b0011100;		3'h4: dec = 7'b1010101;		3'h5: dec = 7'b1011101;		3'h6: dec = 7'b1110111;		default: dec = 7'bxxxxxxx;	endcase endfunction assign lamp = dec(cnt);endmodule"))
}
