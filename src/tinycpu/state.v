`include "defs.v"

module state(clk,reset,run,cont,halt,cs);
  
  input clk, reset, run, cont, halt;
  output [2:0] cs;
  reg [2:0]cs;

  always @(posedge clk or negedge reset)
    if(!reset) cs <= `IDLE;
    else
      case(cs)
        `IDLE: if(run) cs <= `FETCHA;
        `FETCHA: cs <= `FETCHB;
        `FETCHB: cs <= `EXECA;
        `EXECA: if(halt) cs <= `IDLE;
                else if(cont) cs <= `EXECB;
                else cs <= `FETCHA;
        `EXECB: cs <= `FETCHA;
        default: cs <= 3'bxxx;
      endcase

endmodule