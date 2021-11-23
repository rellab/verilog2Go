`include "defs.v"

module alu(a, b, f, s);

 input [15:0] a, b;
 input [4:0] 	f;
 output [15:0] s;
 wire [15:0] x, y;

 assign x = a + 16'h8000;
 assign y = b + 16'h8000;

 always @(a or b or x or y or f)
   case(f)
     `ADD : s = b + a;
     `SUB : s = b - a;
     `MUL : s = b * a;
     `SHL : s = b << a;
     `SHR : s = b >> a;
     `BAND: s = b & a;
     `BOR : s = b | a;
     `BXOR: s = b ^ a;
     `AND : s = b && a;
     `OR  : s = b || a;
     `EQ  : s = b == a;
     `NE  : s = b != a;
     `GE  : s = y >= x;
     `LE  : s = y <= x;
     `GT  : s = y > x;
     `LT  : s = y < x;
     `NEG : s = -a;
     `BNOT : s = ~a;
     `NOT : s = !a;
     default : s = 16'hxxxx;
   endcase

endmodule