package main

import (
	//#include <stdio.h>
	// "C"

	"os"

	"github.com/verilog2Go/src/preprocess"
)

//export createGo
// func createGo(src *C.char) {
// 	//Read verilog source
// 	input := antlr.NewInputStream(C.GoString(src))
// 	//create lexer
// 	lexer := parser.NewVerilogLexer(input)
// 	stream := antlr.NewCommonTokenStream(lexer, 0)
// 	//create parser
// 	p := parser.NewVerilogParser(stream)

// 	// set the listener
// 	listener := builder.NewVerilogListener()
// 	p.AddParseListener(listener)

// 	tree := p.Source_text()
// 	tree.ToStringTree([]string{}, p)
// }

func main() {
	// createGo(C.CString("module adder(a, b, q); input [3:0] a, b; output [3:0] q; assign q = a + b; endmodule"))
	// createGo(C.CString("module cnt_unit(CK, RES, EN, Q, CA);input CK, RES, EN;output Q, CA;always @(posedge CK or posedge RES)begin	if(RES == 1'b1)		Q <= 1'b0;	else if(EN == 1'b1)		Q <= ~Q;end assign CA = EN & Q;endmodule"))
	// createGo(C.CString("module saikoro(ck, reset, enable, lamp);input ck, reset, enable;output [6:0] lamp;reg [2:0] cnt;always@(posedge ck or posedge reset) begin	if(reset == 1'b1)		cnt <= 3'h1;	else if(enable == 1'b1)		if(cnt == 3'h6)			cnt <= 3'h1;		else			cnt <= cnt + 3'h1;end function [6:0] dec;input [2:0] din;	case(din)		3'h1: dec = 7'b0001000;		3'h2: dec = 7'b1000001;		3'h3: dec = 7'b0011100;		3'h4: dec = 7'b1010101;		3'h5: dec = 7'b1011101;		3'h6: dec = 7'b1110111;		default: dec = 7'bxxxxxxx;	endcase endfunction assign lamp = dec(cnt);endmodule"))
	sourceFile := os.Args[1]
	// bytes, err := ioutil.ReadFile(sourceFile)
	// if err != nil {
	// 	panic(err)
	// }
	preprocess.Run(sourceFile)
}
