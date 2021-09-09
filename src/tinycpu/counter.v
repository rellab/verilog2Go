module counter(clk,reset,load,inc,d,q);
  parameter N = 16;
   
  input clk,reset,load,inc;
  input [N-1:0] d;
  output [N-1:0] q;
  reg [N-1:0] 	 q;

  always @(posedge clk or negedge reset)
    if(!reset) q <= 0;
    else if(load) q <= d;
    else if(inc) q <= q + 1;

endmodule